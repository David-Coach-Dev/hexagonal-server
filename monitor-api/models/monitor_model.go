package models

type MonitorLog struct {
    ID        string      `json:"id"`
    Fecha     string      `json:"fecha"`
    Hora      string      `json:"hora"`
    Endpoint  string      `json:"endpoint"`
    Usuario   string      `json:"usuario"`
    Respuesta interface{} `json:"respuesta"`
}
