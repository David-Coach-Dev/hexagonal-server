package drivers

import (
    "context"
    "monitor-api/models"
    "monitor-api/ports/drivens"
)

type MonitorDriverAdapter struct {
    drivenAdapter drivens.MonitorDrivenPort
}

func NewMonitorDriverAdapter(drivenAdapter drivens.MonitorDrivenPort) *MonitorDriverAdapter {
    return &MonitorDriverAdapter{
        drivenAdapter: drivenAdapter,
    }
}

func (mda *MonitorDriverAdapter) SaveLog(ctx context.Context, log models.MonitorLog) error {
    return mda.drivenAdapter.SaveLog(ctx, log)
}

func (mda *MonitorDriverAdapter) ListLogs(ctx context.Context) ([]string, error) {
    return mda.drivenAdapter.ListLogs(ctx)
}

func (mda *MonitorDriverAdapter) ReadLog(ctx context.Context, filename string) (*models.MonitorLog, error) {
    return mda.drivenAdapter.ReadLog(ctx, filename)
}

func (mda *MonitorDriverAdapter) DeleteLog(ctx context.Context, filename string) error {
    return mda.drivenAdapter.DeleteLog(ctx, filename)
}

func (mda *MonitorDriverAdapter) DeleteAllLogs(ctx context.Context) error {
    return mda.drivenAdapter.DeleteAllLogs(ctx)
}
