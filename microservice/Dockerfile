FROM  golang:latest

RUN mkdir -p /go/src/microservice

WORKDIR /go/src/microservice

COPY . /go/src/microservice

# RUN go install microservicea/cmd/server

# CMD /go/bin/server

RUN go build ./cmd/server
CMD ./server

EXPOSE 8080
