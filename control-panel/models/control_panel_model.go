package models

type Credentials struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

type DataYT struct {
    ID          string `json:"id"`
    Title       string `json:"title"`
    Description string `json:"description"`
    URL         string `json:"url"`
}

type Live struct {
    ID    string `json:"id"`
    Title string `json:"title"`
    URL   string `json:"url"`
    Live  bool   `json:"live"`
}

type MonitorLog struct {
    ID        string      `json:"id"`
    Fecha     string      `json:"fecha"`
    Hora      string      `json:"hora"`
    Endpoint  string      `json:"endpoint"`
    Usuario   string      `json:"usuario"`
    Respuesta interface{} `json:"respuesta"`
}

type ErrorLog struct {
    ID       string `json:"id"`
    Fecha    string `json:"fecha"`
    Hora     string `json:"hora"`
    Endpoint string `json:"endpoint"`
    Usuario  string `json:"usuario"`
    Error    string `json:"error"`
}
