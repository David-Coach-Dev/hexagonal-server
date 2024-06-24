package drivers

import (
    "context"
    "control-panel/models"
    "control-panel/ports/drivens"
)

type ControlPanelDriverAdapter struct {
    drivenAdapter drivens.ControlPanelDrivenPort
}

func NewControlPanelDriverAdapter(drivenAdapter drivens.ControlPanelDrivenPort) *ControlPanelDriverAdapter {
    return &ControlPanelDriverAdapter{
        drivenAdapter: drivenAdapter,
    }
}

// Funciones de autenticación
func (cpda *ControlPanelDriverAdapter) Authenticate(ctx context.Context, credentials models.Credentials) (string, error) {
    return cpda.drivenAdapter.Authenticate(ctx, credentials)
}

func (cpda *ControlPanelDriverAdapter) Authorize(ctx context.Context, token string, resource string, action string) (bool, error) {
    return cpda.drivenAdapter.Authorize(ctx, token, resource, action)
}

// Funciones de YouTube
func (cpda *ControlPanelDriverAdapter) GetYouTubeSearching(ctx context.Context, searching string) ([]models.DataYT, error) {
    return cpda.drivenAdapter.GetYouTubeSearching(ctx, searching)
}

func (cpda *ControlPanelDriverAdapter) GetYouTubeRandomVideo(ctx context.Context) (models.DataYT, error) {
    return cpda.drivenAdapter.GetYouTubeRandomVideo(ctx)
}

func (cpda *ControlPanelDriverAdapter) GetYouTubePlaylists(ctx context.Context) ([]models.DataYT, error) {
    return cpda.drivenAdapter.GetYouTubePlaylists(ctx)
}

func (cpda *ControlPanelDriverAdapter) GetYouTubePlaylistItems(ctx context.Context, playlistID string) ([]models.DataYT, error) {
    return cpda.drivenAdapter.GetYouTubePlaylistItems(ctx, playlistID)
}

func (cpda *ControlPanelDriverAdapter) GetYouTubeLive(ctx context.Context) (models.Live, error) {
    return cpda.drivenAdapter.GetYouTubeLive(ctx)
}

// Funciones de monitorización
func (cpda *ControlPanelDriverAdapter) SaveMonitorLog(ctx context.Context, log models.MonitorLog) error {
    return cpda.drivenAdapter.SaveMonitorLog(ctx, log)
}

func (cpda *ControlPanelDriverAdapter) ListMonitorLogs(ctx context.Context) ([]string, error) {
    return cpda.drivenAdapter.ListMonitorLogs(ctx)
}

func (cpda *ControlPanelDriverAdapter) ReadMonitorLog(ctx context.Context, filename string) (*models.MonitorLog, error) {
    return cpda.drivenAdapter.ReadMonitorLog(ctx, filename)
}

func (cpda *ControlPanelDriverAdapter) DeleteMonitorLog(ctx context.Context, filename string) error {
    return cpda.drivenAdapter.DeleteMonitorLog(ctx, filename)
}

func (cpda *ControlPanelDriverAdapter) DeleteAllMonitorLogs(ctx context.Context) error {
    return cpda.drivenAdapter.DeleteAllMonitorLogs(ctx)
}

// Funciones de logs
func (cpda *ControlPanelDriverAdapter) SaveErrorLog(ctx context.Context, log models.ErrorLog) error {
    return cpda.drivenAdapter.SaveErrorLog(ctx, log)
}

func (cpda *ControlPanelDriverAdapter) ListLogs(ctx context.Context) ([]string, error) {
    return cpda.drivenAdapter.ListLogs(ctx)
}

func (cpda *ControlPanelDriverAdapter) ReadLog(ctx context.Context, filename string) (*models.ErrorLog, error) {
    return cpda.drivenAdapter.ReadLog(ctx, filename)
}

func (cpda *ControlPanelDriverAdapter) DeleteLog(ctx context.Context, filename string) error {
    return cpda.drivenAdapter.DeleteLog(ctx, filename)
}

func (cpda *ControlPanelDriverAdapter) DeleteAllLogs(ctx context.Context) error {
    return cpda.drivenAdapter.DeleteAllLogs(ctx)
}
