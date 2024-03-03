package middleware

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
)

func OauthAuthentication(c *fiber.Ctx) error {
	db := services.DB
	redis := services.REDIS
	ctx := redis.Context()

	auth := c.Get("Authorization")

	return c.Next()
}

func bnc(c *fiber.Ctx) error {
	db := services.DB
	redis := services.REDIS
	ctx := redis.Context()

	clientKey := c.Get("X-CLIENT-KEY")
	timestamp, _ := time.Parse(time.RFC3339Nano, c.Get("X-TIMESTAMP"))
	signature := c.Get("X-SIGNATURE")
	externalID := c.Get("X-EXTERNAL-ID")

	serv := model.ServiceCredential{}
	db.Where(`integration_partner_code = ? AND callback_token = ?`, "bnc", clientKey).Take(&serv)

	conf := model.BncCredential{}

	if len(clientKey) == 0 {
		return lib.ErrorBadRequest(c, "missing X-CLIENT-KEY")
	}

	if err := validateSignature(signature); err != nil {
		return lib.ErrorUnauthorized(c, err.Error())
	}

	if timestamp.AddDate(0, 0, 1).Before(time.Now()) {
		log.Println(c.Get("X-TIMESTAMP"), timestamp, "+", time.Now())
		return lib.ErrorBadRequest(c, "X-TIMESTAMP is not valid")
	}

	if val, _ := redis.Get(ctx, externalID).Result(); len(val) > 0 {
		log.Println(val)
		return lib.ErrorUnauthorized(c, "duplicate X-EXTERNAL-ID")
	}
	if err := redis.Set(ctx, externalID, "true", time.Duration(24)*time.Hour).Err(); err != nil {
		return lib.ErrorInternal(c)
	}

	if clientKey != lib.RevStr(conf.CallbackToken) {
		return lib.ErrorUnauthorized(c) // TODO : response to partner with standart BNC response unauthorized
	}

	token, _ := createToken(conf, serv, clientKey)

	res := bncmodel.BncResponseToken{
		AccessToken:     lib.Strptr(token),
		ExpiresIn:       lib.Strptr(fmt.Sprint(viper.GetInt64("TOKEN_EXPIRE_IN"))),
		ResponseCode:    lib.Strptr(fmt.Sprint(2007300)),
		ResponseMessage: lib.Strptr("Successful"),
		TokenType:       lib.Strptr("Bearer"),
	}

	return lib.OK(c, res)
}

func createToken(conf model.BncCredential, crd model.ServiceCredential, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username":     conf.AppID,
			"exp":          time.Now().Add(time.Second * time.Duration(viper.GetInt64("TOKEN_EXPIRE_IN"))).Unix(),
			"partner_code": crd.IntegrationPartnerCode,
			"service_name": crd.ServiceName,
			"service_id":   crd.ID,
		})

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		log.Println(err)
		return "", err
	}

	return tokenString, nil
}

func validateSignature(sign string) error {
	return nil // TODO : validate signature
}
