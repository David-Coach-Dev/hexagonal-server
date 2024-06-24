package drivens

import (
    "context"
    "errors"
    "authorization-api/models"
    "encoding/json"
    "io/ioutil"
    "authorization-api/utils"
)

type AuthorizationRule struct {
    Rol      string `json:"rol"`
    Resource string `json:"resource"`
    Action   string `json:"action"`
}

type AuthorizationAdapter struct {
    rules  []AuthorizationRule
    jwtUtil *utils.JWTUtil
}

func NewAuthorizationAdapter() *AuthorizationAdapter {
    // Cargar reglas de autorización desde un archivo mock
    rules, err := loadMockRules("authorization_rules_mock.json")
    if err != nil {
        panic(err)
    }

    jwtUtil := utils.NewJWTUtil("your_secret_key")

    return &AuthorizationAdapter{
        rules:  rules,
        jwtUtil: jwtUtil,
    }
}

func (aa *AuthorizationAdapter) Authorize(ctx context.Context, token string, resource string, action string) (bool, error) {
    // Decodificar el token para obtener el rol del usuario
    claims, err := aa.jwtUtil.ValidateJWT(token)
    if err != nil {
        return false, err
    }

    userRol := claims["rol"].(string)

    // Verificar las reglas de autorización
    for _, rule := range aa.rules {
        if rule.Rol == userRol && rule.Resource == resource && rule.Action == action {
            return true, nil
        }
    }

    return false, errors.New("not authorized")
}

func loadMockRules(filename string) ([]AuthorizationRule, error) {
    data, err := ioutil.ReadFile(filename)
    if err != nil {
        return nil, err
    }

    var rules []AuthorizationRule
    if err := json.Unmarshal(data, &rules); err != nil {
        return nil, err
    }

    return rules, nil
}
