package main

import (
    "context"
    authdrivens "hexagonal-server/authentication-api/adapters/drivens"
    authdrivers "hexagonal-server/authentication-api/ports/drivers"
    authzdrivens "hexagonal-server/authorization-api/adapters/drivens"
    authzdrivers "hexagonal-server/authorization-api/ports/drivers"
    controlpaneldrivens "hexagonal-server/control-panel/adapters/drivens"
    controlpaneldrivers "hexagonal-server/control-panel/adapters/drivers"
    monitordrivens "hexagonal-server/monitor-api/adapters/drivens"
    monitordrivers "hexagonal-server/monitor-api/ports/drivers"
    logsdrivens "hexagonal-server/logs-api/adapters/drivens"
    logsdrivers "hexagonal-server/logs-api/ports/drivers"
    youtubedrivens "hexagonal-server/youtube-api/adapters/drivens"
    youtubedrivers "hexagonal-server/youtube-api/ports/drivers"
)

func Compose() *controlpaneldrivers.ControlPanelDriverAdapter {
    ctx := context.Background()

    // Create adapters for each service
    authAdapter := authdrivens.NewAuthenticationAdapter()
    authzAdapter := authzdrivens.NewAuthorizationAdapter()
    youtubeAdapter, err := youtubedrivens.NewYouTubeAdapter()
    if err != nil {
        panic(err)
    }
    monitorAdapter := monitordrivens.NewMonitorAdapter()
    logsAdapter := logsdrivens.NewLogsAdapter()

    // Create driven ports for each service
    authDrivenPort := authdrivers.NewAuthenticationDriverAdapter(authAdapter)
    authzDrivenPort := authzdrivers.NewAuthorizationDriverAdapter(authzAdapter)
    youtubeDrivenPort := youtubedrivers.NewYouTubeDriverAdapter(youtubeAdapter)
    monitorDrivenPort := monitordrivers.NewMonitorDriverAdapter(monitorAdapter)
    logsDrivenPort := logsdrivers.NewLogsDriverAdapter(logsAdapter)

    // Create control panel driven adapter
    controlPanelDrivenAdapter := controlpaneldrivens.NewControlPanelAdapter(
        authDrivenPort,
        authzDrivenPort,
        youtubeDrivenPort,
        monitorDrivenPort,
        logsDrivenPort,
    )

    // Create control panel driver adapter
    controlPanelDriverAdapter := controlpaneldrivers.NewControlPanelDriverAdapter(controlPanelDrivenAdapter)

    return controlPanelDriverAdapter
}
