FROM golang:1.19.0-bullseye

WORKDIR /usr/src/api

COPY ./ ./

RUN go get -t api

CMD ["/bin/bash"]