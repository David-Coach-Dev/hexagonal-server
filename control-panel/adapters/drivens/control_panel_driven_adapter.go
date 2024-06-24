package drivens

import (
    "context"
    "control-panel/models"
    "authentication-api/ports/drivers"
    "authorization-api/ports/drivers"
    "youtube-api/ports/drivers"
    "monitor-api/ports/drivers"
    "logs-api/ports/drivers"
)

type ControlPanelAdapter struct {
    authAdapter      drivers.AuthenticationDriverPort
    authzAdapter     drivers.AuthorizationDriverPort
    youtubeAdapter   drivers.YouTubeDriverPort
    monitorAdapter   drivers.MonitorDriverPort
    logsAdapter      drivers.LogsDriverPort
}

func NewControlPanelAdapter(
    authAdapter drivers.AuthenticationDriverPort,
    authzAdapter drivers.AuthorizationDriverPort,
    youtubeAdapter drivers.YouTubeDriverPort,
    monitorAdapter drivers.MonitorDriverPort,
    logsAdapter drivers.LogsDriverPort,
) *ControlPanelAdapter {
    return &ControlPanelAdapter{
        authAdapter:    authAdapter,
        authzAdapter:   authzAdapter,
        youtubeAdapter: youtubeAdapter,
        monitorAdapter: monitorAdapter,
        logsAdapter:    logsAdapter,
    }
}

// Funciones de autenticación
func (cpa *ControlPanelAdapter) Authenticate(ctx context.Context, credentials models.Credentials) (string, error) {
    return cpa.authAdapter.Authenticate(ctx, credentials)
}

func (cpa *ControlPanelAdapter) Authorize(ctx context.Context, token string, resource string, action string) (bool, error) {
    return cpa.authzAdapter.Authorize(ctx, token, resource, action)
}

// Funciones de YouTube
func (cpa *ControlPanelAdapter) GetYouTubeSearching(ctx context.Context, searching string) ([]models.DataYT, error) {
    return cpa.youtubeAdapter.GetSearching(ctx, searching)
}

func (cpa *ControlPanelAdapter) GetYouTubeRandomVideo(ctx context.Context) (models.DataYT, error) {
    return cpa.youtubeAdapter.GetRandomVideo(ctx)
}

func (cpa *ControlPanelAdapter) GetYouTubePlaylists(ctx context.Context) ([]models.DataYT, error) {
    return cpa.youtubeAdapter.GetSearchPlaylists(ctx)
}

func (cpa *ControlPanelAdapter) GetYouTubePlaylistItems(ctx context.Context, playlistID string) ([]models.DataYT, error) {
    return cpa.youtubeAdapter.GetSearchPlaylistsItems(ctx, playlistID)
}

func (cpa *ControlPanelAdapter) GetYouTubeLive(ctx context.Context) (models.Live, error) {
    return cpa.youtubeAdapter.GetCheckLive(ctx)
}

// Funciones de monitorización
func (cpa *ControlPanelAdapter) SaveMonitorLog(ctx context.Context, log models.MonitorLog) error {
    return cpa.monitorAdapter.SaveLog(ctx, log)
}

func (cpa *ControlPanelAdapter) ListMonitorLogs(ctx context.Context) ([]string, error) {
    return cpa.monitorAdapter.ListLogs(ctx)
}

func (cpa *ControlPanelAdapter) ReadMonitorLog(ctx context.Context, filename string) (*models.MonitorLog, error) {
    return cpa.monitorAdapter.ReadLog(ctx, filename)
}

func (cpa *ControlPanelAdapter) DeleteMonitorLog(ctx context.Context, filename string) error {
    return cpa.monitorAdapter.DeleteLog(ctx, filename)
}

func (cpa *ControlPanelAdapter) DeleteAllMonitorLogs(ctx context.Context) error {
    return cpa.monitorAdapter.DeleteAllLogs(ctx)
}

// Funciones de logs
func (cpa *ControlPanelAdapter) SaveErrorLog(ctx context.Context, log models.ErrorLog) error {
    return cpa.logsAdapter.SaveLog(ctx, log)
}

func (cpa *ControlPanelAdapter) ListLogs(ctx context.Context) ([]string, error) {
    return cpa.logsAdapter.ListLogs(ctx)
}

func (cpa *ControlPanelAdapter) ReadLog(ctx context.Context, filename string) (*models.ErrorLog, error) {
    return cpa.logsAdapter.ReadLog(ctx, filename)
}

func (cpa *ControlPanelAdapter) DeleteLog(ctx context.Context, filename string) error {
    return cpa.logsAdapter.DeleteLog(ctx, filename)
}

func (cpa *ControlPanelAdapter) DeleteAllLogs(ctx context.Context) error {
    return cpa.logsAdapter.DeleteAllLogs(ctx)
}
