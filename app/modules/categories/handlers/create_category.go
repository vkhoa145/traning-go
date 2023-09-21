package handlers

import (
	"fmt"
	"net/http"
	"reflect"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/vkhoa145/go-training/app/mails"
	"github.com/vkhoa145/go-training/app/models"
	"github.com/vkhoa145/go-training/config"
)

func (h *CategoryHandlers) CreateCategory() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		payload := models.CreateCategoryInput{}

		if err := ctx.BodyParser(&payload); err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(&fiber.Map{"status": http.StatusBadRequest, "error": err.Error()})
		}

		validate := validator.New()
		if err := validate.Struct(payload); err != nil {
			errors := map[string]string{}
			for _, err := range err.(validator.ValidationErrors) {
				A := err.Field()
				field, _ := reflect.TypeOf(payload).FieldByName(A)
				jsonTag := field.Tag.Get("json")
				errors[jsonTag] = err.Tag()
			}

			ctx.Status(http.StatusUnprocessableEntity)
			return ctx.JSON(&fiber.Map{"status": http.StatusUnprocessableEntity, "message": "Unprocessable Content", "errors": errors})
		}

		userId := ctx.Get("User_id")
		userIdFloat, err := strconv.ParseFloat(userId, 64)
		if err != nil {
			return ctx.JSON(&fiber.Map{"status": http.StatusBadRequest, "error": err.Error()})
		}

		payload.UserId = uint(userIdFloat)

		createdCategory, err := h.categoryUseCase.CreateCategory(ctx, &payload)
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(&fiber.Map{"status": http.StatusBadRequest, "error": err.Error()})
		}

		sender := mails.NewGmailSender(config.LoadConfig().EmailSenderName, config.LoadConfig().EmailSenderAddress, config.LoadConfig().EmailSenderPassword)
		subject := "a test email"
		content := `
		<h1>Hello World</h1>
		<p>This is a test message from <a href="http://google.com">Link</a></p>
		`
		to := []string{"khoavodang1451997@gmail.com"}
		err1 := sender.SendMail(subject, content, to, nil, nil, nil)
		if err1 != nil {
			fmt.Println("mail err", err1)
			return ctx.JSON(&fiber.Map{"status": http.StatusBadRequest, "error": err1.Error()})
		}

		ctx.Status(http.StatusCreated)
		return ctx.JSON(&fiber.Map{"status": http.StatusCreated, "data": createdCategory, "error": nil})
	}
}
