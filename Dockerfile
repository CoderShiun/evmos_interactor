FROM golang:1.18-alpine

WORKDIR /home

COPY evmos.sh /home/tc_command.sh

RUN chmod 777 /home/evmos.sh

ENTRYPOINT ["/home/evmos.sh"]