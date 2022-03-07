# tetris_battle_royale

## Service architecture
![Service architecture diagram](docs/service_architecture.png)
## Software architecture
![Software architecture diagram](docs/software_architecture.png)
## Package structure
```
cmd/
├── game_service/
├── gateway/
├── matchmaking_service/
├── statistics_service/
└── user_service/
internal/
├── core/
│   ├── driven_adapters/
│   │   ├── game_adapter/
│   │   ├── ipc/
│   │   │   └── grpc/
│   │   └── repository/
│   │       ├── postgres/
│   │       └── redis/
│   ├── driven_ports/
│   │   ├── ipc/
│   │   └── repository/
│   ├── driving_ports/
│   ├── protofiles/
│   │   ├── game_service/
│   │   ├── statistics_service/
│   │   └── user_service/
│   ├── services/
│   │   ├── game_service/
│   │   ├── matchmaking_service/
│   │   ├── statistics_service/
│   │   └── user_service/
│   └── types/
└── driving_adapters/
    ├── rest/
    └── websocket/
```
