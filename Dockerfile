FROM golang:1.18-alpine

WORKDIR /home

RUN apk add build-base

COPY evmos.sh /home/evmos.sh

RUN chmod 777 /home/evmos.sh

ENTRYPOINT ["/home/evmos.sh"]