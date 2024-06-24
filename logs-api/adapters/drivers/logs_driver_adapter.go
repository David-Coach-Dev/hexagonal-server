package drivers

import (
    "context"
    "logs-api/models"
    "logs-api/ports/drivens"
)

type LogsDriverAdapter struct {
    drivenAdapter drivens.LogsDrivenPort
}

func NewLogsDriverAdapter(drivenAdapter drivens.LogsDrivenPort) *LogsDriverAdapter {
    return &LogsDriverAdapter{
        drivenAdapter: drivenAdapter,
    }
}

func (lda *LogsDriverAdapter) SaveLog(ctx context.Context, log models.ErrorLog) error {
    return lda.drivenAdapter.SaveLog(ctx, log)
}

func (lda *LogsDriverAdapter) ListLogs(ctx context.Context) ([]string, error) {
    return lda.drivenAdapter.ListLogs(ctx)
}

func (lda *LogsDriverAdapter) ReadLog(ctx context.Context, filename string) (*models.ErrorLog, error) {
    return lda.drivenAdapter.ReadLog(ctx, filename)
}

func (lda *LogsDriverAdapter) DeleteLog(ctx context.Context, filename string) error {
    return lda.drivenAdapter.DeleteLog(ctx, filename)
}

func (lda *LogsDriverAdapter) DeleteAllLogs(ctx context.Context) error {
    return lda.drivenAdapter.DeleteAllLogs(ctx)
}
