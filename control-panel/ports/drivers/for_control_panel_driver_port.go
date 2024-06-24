package drivers

import (
    "context"
    "control-panel/models"
)

type ControlPanelDriverPort interface {
    // Funciones de autenticación
    Authenticate(ctx context.Context, credentials models.Credentials) (string, error)
    Authorize(ctx context.Context, token string, resource string, action string) (bool, error)

    // Funciones de YouTube
    GetYouTubeSearching(ctx context.Context, searching string) ([]models.DataYT, error)
    GetYouTubeRandomVideo(ctx context.Context) (models.DataYT, error)
    GetYouTubePlaylists(ctx context.Context) ([]models.DataYT, error)
    GetYouTubePlaylistItems(ctx context.Context, playlistID string) ([]models.DataYT, error)
    GetYouTubeLive(ctx context.Context) (models.Live, error)

    // Funciones de monitorización
    SaveMonitorLog(ctx context.Context, log models.MonitorLog) error
    ListMonitorLogs(ctx context.Context) ([]string, error)
    ReadMonitorLog(ctx context.Context, filename string) (*models.MonitorLog, error)
    DeleteMonitorLog(ctx context.Context, filename string) error
    DeleteAllMonitorLogs(ctx context.Context) error

    // Funciones de logs
    SaveErrorLog(ctx context.Context, log models.ErrorLog) error
    ListLogs(ctx context.Context) ([]string, error)
    ReadLog(ctx context.Context, filename string) (*models.ErrorLog, error)
    DeleteLog(ctx context.Context, filename string) error
    DeleteAllLogs(ctx context.Context) error
}
