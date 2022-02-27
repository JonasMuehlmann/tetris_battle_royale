FROM  golang:latest

# RUN mkdir -p /go/src/microservice
RUN mkdir -p /go/src/github.com/JonasMuehlmann/tetris_battle_royale/microservice
# RUN mkdir /go/
# COPY ../../../../../../../ /go

# WORKDIR /go/src/microservice
WORKDIR /go/src/github.com/JonasMuehlmann/tetris_battle_royale/microservice

# COPY . /go/src/microservice
COPY . /go/src/github.com/JonasMuehlmann/tetris_battle_royale/microservice

RUN go build ./cmd/user_service
CMD ./user_service

EXPOSE 8080
EXPOSE 8081
