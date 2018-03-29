# Google gRPC Client with Envoy Proxy

使用`Envoy Proxy`作为`Greet Service`的代理, 提供`tracing`功能.

## 运行指南

    make up               // 启动服务
    make down             // 关闭服务
    make build_cmd        // 生成本机可执行文件(主要针对MacOS用户)
    ./bin/greet --addr <addr>  // 调用Greet Service服务
    
## 详细介绍

### 1. 沙盒架构初览

`docker-compose.yaml`

    ...
    services:
      greet:
        ...
        networks:
          envoymesh:
            ...
      zipkin:
        ...
        networks:
          envoymesh:
            ...
    networks:
      envoymesh: {}

通过`docker-compose`构建一个沙盒, 方便读者部署现场环境.

先创建一个网络(`envoymesh`), 然后把容器都放置到该网络里面.

得到两个容器, GreetService(`greet`)和Zipkin Tracing Service(`zipkin`).

### 2. Envoy Proxy配置初览

接下来, 我们看看`Envoy Proxy`的配置文件.

`config/greet-grpc-envoy.yaml`

    # 配置文件分成三部分
    static_resources:
      # 代理主要配置
      ...
    tracing:
      # 调用跟踪配置
      ...
    admin:
      # Envoy Proxy管理配置
      ...

主要分析`static_resources`和`tracing`两个配置项, `admin`是`Envoy Proxy`的管理接口.

关于管理配置项可以参考这篇[文档](https://www.envoyproxy.io/docs/envoy/latest/api-v2/config/bootstrap/v2/bootstrap.proto#envoy-api-msg-config-bootstrap-v2-admin).

关于管理接口可以参考这篇[文档](https://www.envoyproxy.io/docs/envoy/latest/operations/admin).

### 3. static_resources配置初览

我们先分析代理相关的配置项.

`config/greet-grpc-envoy.yaml`

    static_resources:
      clusters:
        ...
      listeners:
        ...
    ...

在本项目的配置里面, `static_resources`有两个参数.

1. `clusters`定义`Envoy Proxy`的集群(cluster), 主要是提供给监听器指定流量流向使用的.
1. `listeners`定义`Envoy Proxy`的监听器(listener), 主要声明监听的地址和流量的处理器(filter).

### 4. clusters配置详解

下面分析`clusters`配置项.

`config/greet-grpc-envoy.yaml`

    static_resources:
      clusters:
      - name: local_service_grpc
        hosts:
        - socket_address:                 // 这里指定了Greet Service的服务地址
            address: 127.0.0.1
            port_value: 5556
        http2_protocol_options: {}        // gRPC的服务需要指定这个参数
        ...
      - name: zipkin
        hosts:
        - socket_address:                 // 这里指定Zipkin Tracing Service的服务地址
            address: zipkin
            port_value: 9411
        ...

上面分别定义了`local_service_grpc`和`zipkin`两个集群.

1. `local_service_grpc`是`Envoy Proxy`转发`Greet Service`流量的目的集群.
1. `zipkin`是`Envoy Proxy`转发跟踪流量的集群.

关于`clusters`配置项还可以参考这个[文档](https://www.envoyproxy.io/docs/envoy/latest/api-v2/api/v2/cds.proto#envoy-api-msg-cluster).

### 5. listeners配置初览

下面分析`listeners`配置项.

`config/greet-grpc-envoy.yaml`

    static_resources:
      ...
      listeners:
      - address:
          ...
        filter_chains:
          ...

`listener`由两个参数组成, `address`和`filter_chains`.

1. `address`声明了代理监听的地址. 具体参数可以参考[文档](https://www.envoyproxy.io/docs/envoy/latest/api-v2/api/v2/core/address.proto#envoy-api-msg-core-address).
1. `filter_chains`声明了代理接受到监听流量之后的行为.

### 6. listeners配置详解

`config/greet-grpc-envoy.yaml`

    static_resources:
      ...
      listeners:
      - filter_chains:
        - filters:
          - name: envoy.http_connection_manager     // 指定filter的驱动
            config:
              generate_request_id: true             // 自动生成request_id供调用跟踪使用.
              tracing:
                operation_name: ingress             // 作用在流量流入时
              route_config:
                name: local_route
                virtual_hosts:
                - name: greet_service               // 定义一个虚拟主机
                  routes:
                  - match:
                      prefix: "/"                   // 指定流量匹配规则
                      headers:                      // 指定http头部匹配规则
                      - name: content-type
                        value: application/grpc
                    route:
                      cluster: local_service_grpc   // 匹配到的流量导流到local_service_grpc
                    ...
                  ...
              http_filters:
              - name: envoy.router                  // 声明这个过滤器是流量转发的
                config: {}
              ...
        ...

`Envoy Proxy`的核心配置部分, 详细可以参考[文档](https://www.envoyproxy.io/docs/envoy/latest/api-v2/api/v2/lds.proto).

## 备注

如果启动服务失败的话, 有可能是在下载go依赖的时候出错. 请配置系统代理服务, 确保你的系统能够正常访问世界的互联网.
