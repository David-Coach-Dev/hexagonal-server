package drivens

import (
    "context"
    "errors"
    "authentication-api/models"
    "authentication-api/utils"
    "encoding/json"
    "io/ioutil"
)

type User struct {
    Avatar    string `json:"avatar"`
    Name      string `json:"name"`
    Lastname  string `json:"lastname"`
    User      string `json:"user"`
    Password  string `json:"password"`
    Rol       string `json:"rol"`
}

type AuthenticationAdapter struct {
    users   map[string]User
    jwtUtil *utils.JWTUtil
}

func NewAuthenticationAdapter() *AuthenticationAdapter {
    // Cargar usuarios desde un archivo mock
    users, err := loadMockUsers("users_mock.json")
    if err != nil {
        panic(err)
    }

    jwtUtil := utils.NewJWTUtil("your_secret_key")

    return &AuthenticationAdapter{
        users:   users,
        jwtUtil: jwtUtil,
    }
}

func (aa *AuthenticationAdapter) Authenticate(ctx context.Context, credentials models.Credentials) (string, error) {
    // Verificar si el usuario y la contraseña son válidos
    if user, ok := aa.users[credentials.Username]; ok && user.Password == credentials.Password {
        // Generar un JWT
        token, err := aa.jwtUtil.GenerateJWT(credentials.Username, user.Rol)
        if err != nil {
            return "", err
        }
        return token, nil
    }
    return "", errors.New("invalid credentials")
}

func loadMockUsers(filename string) (map[string]User, error) {
    data, err := ioutil.ReadFile(filename)
    if err != nil {
        return nil, err
    }

    var users []User
    if err := json.Unmarshal(data, &users); err != nil {
        return nil, err
    }

    userMap := make(map[string]User)
    for _, user := range users {
        userMap[user.User] = user
    }

    return userMap, nil
}
