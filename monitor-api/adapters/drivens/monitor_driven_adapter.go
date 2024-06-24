package drivens

import (
    "context"
    "encoding/json"
    "errors"
    "io/ioutil"
    "monitor-api/models"
    "os"
    "path/filepath"
    "time"
)

type MonitorAdapter struct {
    logDir string
}

func NewMonitorAdapter() *MonitorAdapter {
    logDir := "monitor"
    if _, err := os.Stat(logDir); os.IsNotExist(err) {
        os.Mkdir(logDir, os.ModePerm)
    }
    return &MonitorAdapter{logDir: logDir}
}

func (ma *MonitorAdapter) SaveLog(ctx context.Context, log models.MonitorLog) error {
    filename := filepath.Join(ma.logDir, "monitor-"+time.Now().Format("2006-01-02-15-04-05")+".json")
    file, err := os.Create(filename)
    if err != nil {
        return err
    }
    defer file.Close()

    encoder := json.NewEncoder(file)
    if err := encoder.Encode(log); err != nil {
        return err
    }

    return nil
}

func (ma *MonitorAdapter) ListLogs(ctx context.Context) ([]string, error) {
    files, err := ioutil.ReadDir(ma.logDir)
    if err != nil {
        return nil, err
    }

    var filenames []string
    for _, file := range files {
        if !file.IsDir() {
            filenames = append(filenames, file.Name())
        }
    }
    return filenames, nil
}

func (ma *MonitorAdapter) ReadLog(ctx context.Context, filename string) (*models.MonitorLog, error) {
    filepath := filepath.Join(ma.logDir, filename)
    file, err := os.Open(filepath)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var log models.MonitorLog
    decoder := json.NewDecoder(file)
    if err := decoder.Decode(&log); err != nil {
        return nil, err
    }

    return &log, nil
}

func (ma *MonitorAdapter) DeleteLog(ctx context.Context, filename string) error {
    filepath := filepath.Join(ma.logDir, filename)
    if _, err := os.Stat(filepath); os.IsNotExist(err) {
        return errors.New("file not found")
    }

    return os.Remove(filepath)
}

func (ma *MonitorAdapter) DeleteAllLogs(ctx context.Context) error {
    files, err := ioutil.ReadDir(ma.logDir)
    if err != nil {
        return err
    }

    for _, file := range files {
        if !file.IsDir() {
            filepath := filepath.Join(ma.logDir, file.Name())
            if err := os.Remove(filepath); err != nil {
                return err
            }
        }
    }
    return nil
}
