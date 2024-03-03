package account

import (
	"api/app/lib"
	"api/app/middleware"
	"api/app/model"
	"api/app/services"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func PostLogin(c *fiber.Ctx) error {
	db := services.DB.WithContext(c.UserContext())
	redis := services.REDIS
	ctx := redis.Context()
	lib.GetXUserID()

	api := new(model.LoginAPI)
	if err := lib.BodyParser(c, api); nil != err {
		return lib.ErrorBadRequest(c, err)
	}

	user := model.User{}
	if row := db.Where(`username = ? OR email = ?`, api.Username, api.Username).Take(&user); row.RowsAffected < 1 {
		return lib.ErrorNotFound(c)
	}

	if ok := lib.PasswordCompare(lib.RevStr(user.Password), lib.RevStr(api.Password), viper.GetString("SALT"), viper.GetString("KEY")); !ok {
		return lib.ErrorUnauthorized(c)
	}

	res, err := middleware.GenerateAccessToken(&user)
	if nil != err {
		return lib.ErrorUnauthorized(c, err.Error())
	}

	if err := redis.Set(ctx, lib.RevStr(res.AccessToken), user.ID.String(), time.Duration(*res.ExpiresIn)).Err(); err != nil {
		return lib.ErrorInternal(c)
	}

	return lib.OK(c, &res)

}
