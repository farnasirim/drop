FROM ubuntu:18.04

COPY drop /opt/drop
RUN chmod +x /opt/drop

CMD /opt/drop serve --addr=0.0.0.0:12345 --redis-addr=redis:6379
