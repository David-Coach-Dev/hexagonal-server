package main

import (
    "log"
    "github.com/gin-gonic/gin"
    controlpanel "hexagonal-server/control-panel"
    "hexagonal-server/control-panel/handler"
)

func main() {
    // Create the main router
    router := gin.Default()

    // Compose all services
    controlPanelDriverAdapter := Compose()

    // Create response handler
    responseHandler := handler.NewResponseHandler(
        controlPanelDriverAdapter,
        controlPanelDriverAdapter,
    )

    // Create control panel service
    controlPanelService := controlpanel.NewControlPanelService(
        controlPanelDriverAdapter,
        responseHandler,
    )

    // Setup routes
    setupRoutes(router, controlPanelService)

    // Start the server
    if err := router.Run(":8080"); err != nil {
        log.Fatalf("Failed to run server: %v", err)
    }
}
