package middleware

import (
	"api/app/lib"
	"api/app/model"
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
)

// Oauth2Authentication authenticating oauth2 before calling next request
func Oauth2Authentication(c *fiber.Ctx) error {
	result := model.ResponseAuthenticate{}
	var url string
	if string(c.Request().Header.Peek("Host")) == "example.com" {
		result = model.ResponseAuthenticate{
			Access: lib.Strptr(viper.GetString("ACCESS_TOKEN")),
			UserID: lib.Strptr(viper.GetString("USER_ID")),
		}
	} else {
		url = os.Getenv("HOST_OAUTH2_SERVER") + "/authenticate"
		method := "POST"

		client := &http.Client{}
		req, err := http.NewRequest(method, url, nil)
		if err != nil {
			lib.Logs.Println(err)
			return err
		}
		req.Header.Add("Authorization", string(c.Request().Header.Peek("Authorization")))

		res, err := client.Do(req)
		if err != nil {
			lib.Logs.Println(err)
			return err
		}
		defer res.Body.Close()
		if res.StatusCode != 200 {
			return lib.ErrorUnauthorized(c, "Unauthorized")
		}

		var jsonData map[string]interface{}
		json.NewDecoder(res.Body).Decode(&jsonData)
		jsonByte := []byte(string(lib.ConvertJsonToStr(jsonData)))
		err = json.Unmarshal(jsonByte, &result)
		if err != nil {
			lib.Logs.Println(err)
			return err
		}
	}

	// Set the authenticated user data in the context for later use
	c.Locals("auth", result)

	return c.Next()
}

// GenerateAccessToken func
func GenerateAccessToken(user *model.User) (model.ResponseToken, error) {
	result := model.ResponseToken{}
	exp := viper.GetInt("TOKEN_EXPIRE_IN")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"user_id": user.ID.String(),
			"exp":     time.Now().Add(time.Second * time.Duration(int64(exp))).Unix(),
		})

	tokenString, err := token.SignedString(viper.GetString("TOKEN_KEY"))
	if err != nil {
		return result, err
	}

	return model.ResponseToken{
		AccessToken: &tokenString,
		ExpiresIn:   &exp,
		TokenType:   lib.Strptr("Bearer"),
	}, nil

}
