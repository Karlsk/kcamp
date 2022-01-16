FROM ubuntu

RUN \
apt-get update -y && \
apt-get install vim -y && \
apt-get install curl -y && \
apt-get install ca-certificates -y

ENV Version="V1.0"

COPY ./httpserver httpserver

CMD ["./httpserver"]