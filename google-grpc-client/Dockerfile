FROM envoyproxy/envoy:latest

RUN mkdir /var/log/envoy/
COPY ./bin/greet /usr/local/bin/greet
COPY ./script/greet_start.sh /usr/local/bin/greet_start.sh
CMD /usr/local/bin/greet_start.sh
