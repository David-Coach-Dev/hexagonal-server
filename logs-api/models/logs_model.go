package models

type ErrorLog struct {
    ID       string `json:"id"`
    Fecha    string `json:"fecha"`
    Hora     string `json:"hora"`
    Endpoint string `json:"endpoint"`
    Usuario  string `json:"usuario"`
    Error    string `json:"error"`
}
