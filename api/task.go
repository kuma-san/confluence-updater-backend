package api

import (

	//"github.com/Sirupsen/logrus"
	"fmt"

	"github.com/Sirupsen/logrus"
	"github.com/hashicorp/go-uuid"
	"github.com/kuma-san/confluence-updater-backend/db"
	"github.com/labstack/echo"
	"github.com/valyala/fasthttp"
	"gopkg.in/go-playground/validator.v8"
)

type (
	req struct {
		Ancestor     string `json:"ancestor" validate:"required"`
		SpaceKey     string `json:"space_key" validate:"required"`
		Title        string `json:"title"`
		Template     string `json:"template" validate:"required"`
		Weekday      string `json:"weekday" validate:"required"`
		Hour         string `json:"hour" validate:"required"`
		SlackURL     string `json:"slack_url"`
		SlackChannel string `json:"slack_channel"`
		Email        string `json:"email"`
	}
)

func PostTask() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		r := new(req)
		if err := c.Bind(r); err != nil {
			logrus.Error(err)
			return c.NoContent(fasthttp.StatusBadRequest)
		}

		config := &validator.Config{TagName: "validate"}
		validate := validator.New(config)
		err = validate.Struct(r)
		if err != nil {
			fmt.Println(r)
			logrus.Error(err)
			return c.NoContent(fasthttp.StatusBadRequest)
		}

		uuid, _ := uuid.GenerateUUID()
		r.Template = "made by kumabot" + r.Template

		a := map[string]string{
			"task_id":       uuid,
			"ancestor":      r.Ancestor,
			"space_key":     r.SpaceKey,
			"title":         r.Title,
			"template":      r.Template,
			"weekday":       r.Weekday,
			"hour":          r.Hour,
			"slack_url":     r.SlackURL,
			"slack_channel": r.SlackChannel,
			"email":         r.Email,
		}

		redis := db.RedisInit()
		err = redis.HMSet("confluence-task:"+uuid, a).Err()
		if err != nil {
			logrus.Error(err)
			return c.NoContent(fasthttp.StatusInternalServerError)
		}
		err = redis.SAdd("confluence-queue:"+r.Weekday+":"+r.Hour, uuid).Err()
		if err != nil {
			logrus.Error(err)
			return c.NoContent(fasthttp.StatusInternalServerError)
		}
		err = redis.SAdd("confluence-task-list", uuid).Err()
		if err != nil {
			logrus.Error(err)
			return c.NoContent(fasthttp.StatusInternalServerError)
		}

		return c.NoContent(fasthttp.StatusOK)

	}
}
