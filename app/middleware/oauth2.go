package middleware

import (
	"api/app/lib"
	"api/app/model"
	"bytes"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
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
func GenerateAccessToken(t *model.LoginAPI) (model.ResponseToken, error) {
	result := model.ResponseToken{}
	url := os.Getenv("HOST_OAUTH2_SERVER") + "/token"
	method := "POST"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("grant_type", "password")
	_ = writer.WriteField("client_id", os.Getenv("OAUTH2_CLIENT_ID"))
	_ = writer.WriteField("client_secret", os.Getenv("OAUTH2_CLIENT_SECRET"))
	_ = writer.WriteField("username", *t.Username)
	_ = writer.WriteField("password", *t.Password)
	err := writer.Close()
	if err != nil {
		lib.Logs.Println(err)
		return result, err
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		lib.Logs.Println(err)
		return result, err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		lib.Logs.Println(err)
		return result, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return result, fmt.Errorf("Unauthorized")
	}
	var jsonData map[string]interface{}
	json.NewDecoder(res.Body).Decode(&jsonData)
	jsonByte := []byte(string(lib.ConvertJsonToStr(jsonData)))
	err = json.Unmarshal(jsonByte, &result)
	if err != nil {
		lib.Logs.Println(err)
		return result, err
	}

	return result, nil
}
