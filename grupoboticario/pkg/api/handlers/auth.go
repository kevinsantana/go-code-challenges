package handlers

import (
	log "github.com/sirupsen/logrus"

	"github.com/kevinsantana/go-code-challenges/grupoboticario/pkg/app/models"
	"github.com/kevinsantana/go-code-challenges/grupoboticario/pkg/app/config"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// UserSignUp method to create a new user.
// @Description Create a new user.
// @Summary create a new user
// @Tags User
// @Accept json
// @Produce json
// @Param email body string true "Email"
// @Param password body string true "Password"
// @Param user_role body string true "User role"
// @Success 200 {object} models.User
// @Router /v1/user/sign/up [post]
func UserSignUp(ctx *fiber.Ctx) error {

	// checking received data from JSON body.
	var httpRequest models.SignUp
	if err := ctx.BodyParser(&httpRequest); err != nil {
		return Error(ctx, err)
	}

	// validate sign up fields.
	if err := validator.New().Struct(httpRequest); err != nil {
		return Error(ctx, err)
	}

	user, err := config.AuthService.SignerUp(ctx.UserContext(), models.SignUp{
		Name:     httpRequest.Name,
		CPF:      httpRequest.CPF,
		Email:    httpRequest.Email,
		Password: httpRequest.Password,
	})

	if err != nil {
		log.WithError(err).Error("Error to sign up user")

		return Error(ctx, err)
	}

	return Success(ctx, user)
}
