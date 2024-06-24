# Structure Project Hexagonal Server

```bash
hexagonal-server/
├── control-panel/
│   ├── adapters/
│   │   ├── drivens/
│   │   │   └── control_panel_driven_adapter.go
│   │   ├── drivers/
│   │   │   └── control_panel_driver_adapter.go
│   ├── handler/
│   │   ├── error_handler.go
│   │   └── response_handler.go
│   ├── models/
│   │   └── control_panel_model.go
│   ├── ports/
│   │   ├── drivens/
│   │   │   └── for_control_panel_driven_port.go
│   │   ├── drivers/
│   │   │   └── for_control_panel_driver_port.go
│   ├── .env
│   └── control_panel.go
├── youtube-api/
│   ├── adapters/
│   │   ├── drivens/
│   │   │   └── youtube_driven_adapter.go
│   │   ├── drivers/
│   │   │   └── youtube_driver_adapter.go
│   ├── models/
│   │   └── youtube_model.go
│   ├── ports/
│   │   ├── drivens/
│   │   │   └── for_youtube_driven_port.go
│   │   ├── drivers/
│   │   │   └── for_youtube_driver_port.go
│   ├── .env
│   └── youtube.go
├── monitor-api/
│   ├── adapters/
│   │   ├── drivens/
│   │   │   └── monitor_driven_adapter.go
│   │   ├── drivers/
│   │   │   └── monitor_driver_adapter.go
│   ├── models/
│   │   └── monitor_model.go
│   ├── ports/
│   │   ├── drivens/
│   │   │   └── for_monitor_driven_port.go
│   │   ├── drivers/
│   │   │   └── for_monitor_driver_port.go
│   ├── .env
│   └── monitor.go
├── logs-api/
│   ├── adapters/
│   │   ├── drivens/
│   │   │   └── logs_driven_adapter.go
│   │   ├── drivers/
│   │   │   └── logs_driver_adapter.go
│   ├── models/
│   │   └── logs_model.go
│   ├── ports/
│   │   ├── drivens/
│   │   │   └── for_logs_driven_port.go
│   │   ├── drivers/
│   │   │   └── for_logs_driver_port.go
│   ├── .env
│   └── logs.go
├── authentication-api/
│   ├── adapters/
│   │   ├── drivens/
│   │   │   └── authentication_driven_adapter.go
│   │   ├── drivers/
│   │   │   └── authentication_driver_adapter.go
│   ├── models/
│   │   └── authentication_model.go
│   ├── ports/
│   │   ├── drivens/
│   │   │   └── for_authentication_driven_port.go
│   │   ├── drivers/
│   │   │   └── for_authentication_driver_port.go
│   ├── .env
│   └── authentication.go
├── authorization-api/
│   ├── adapters/
│   │   ├── drivens/
│   │   │   └── authorization_driven_adapter.go
│   │   ├── drivers/
│   │   │   └── authorization_driver_adapter.go
│   ├── models/
│   │   └── authorization_model.go
│   ├── ports/
│   │   ├── drivens/
│   │   │   └── for_authorization_driven_port.go
│   │   ├── drivers/
│   │   │   └── for_authorization_driver_port.go
│   ├── .env
│   └── authorization.go
├── composer.go
├── main.go
└── routing.go
```
