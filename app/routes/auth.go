package routes

import (
	"goauth/models"
	"strings"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func NewMiddleware() fiber.Handler {
	return AuthMiddleware
}

func AuthMiddleware(c *fiber.Ctx) error {
	sess, err := store.Get(c)

	if strings.Split(c.Path(), "/")[1] == "auth" {
		return c.Next()
	}

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "not authorized",
		})
	}

	if sess.Get(AUTH_KEY) == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	return c.Next()
}

type user struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

func Register(c *fiber.Ctx) error {
	c.Accepts("application/json")
	var data user

	err := c.BodyParser(&data)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "error parsing body" + err.Error(),
		})
	}

	password, bcErr := bcrypt.GenerateFromPassword([]byte(data.Password), 14)
	if bcErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error hashing password" + bcErr.Error(),
		})
	}

	user := models.User{
		Name:     data.Name,
		Email:    data.Email,
		Password: string(password),
	}

	err = user.CreateUser()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error creating user" + err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "user created",
	})
}

func Login(c *fiber.Ctx) error {
	c.Accepts("application/json")
	var data user

	err := c.BodyParser(&data)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "error parsing body" + err.Error(),
		})
	}

	user := models.User{
		Email: data.Email,
	}

	if !user.CheckEmail() {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "invalid credentials",
		})
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "invalid credentials",
		})
	}

	sess, err := store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error getting session" + err.Error(),
		})
	}

	sess.Set(AUTH_KEY, true)
	sess.Set(USER_ID, user.ID)

	err = sess.Save()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error saving session" + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
	})

}

func Logout(c *fiber.Ctx) error {
	sess, err := store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Loged out (no session)" + err.Error(),
		})
	}

	err = sess.Destroy()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error destroying session" + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
	})
}

func HealthCheck(c *fiber.Ctx) error {
	sess, err := store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	if sess.Get(AUTH_KEY) == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
	})
}

func GetUser(c *fiber.Ctx) error {
	sess, err := store.Get(c)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	if sess.Get(AUTH_KEY) == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	userId := sess.Get(USER_ID)

	user := models.User{
		ID: userId.(uint64),
	}

	err = user.GetUser()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error getting user" + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"data": fiber.Map{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
		},
	})
}
