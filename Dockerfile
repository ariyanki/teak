FROM alpine:latest

RUN apk -U add ca-certificates

EXPOSE 8080

ADD teak /app/teak
ADD cfg/temp.json /app/cfg/config.json

RUN mkdir -p /app/log

CMD cd /app && ./teak
