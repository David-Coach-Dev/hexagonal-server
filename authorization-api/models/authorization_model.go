package models

type AuthorizationRequest struct {
    Token    string `json:"token"`
    Resource string `json:"resource"`
    Action   string `json:"action"`
}
