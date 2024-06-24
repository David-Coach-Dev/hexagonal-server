package models

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
