package user

import (
	"net/http"
	"strings"
	"vue-fiber/config"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type Handler struct {
	Service
	log zerolog.Logger
}

func NewHandler(s Service, l zerolog.Logger) *Handler {
	return &Handler{
		Service: s,
		log:     l,
	}
}

func (h *Handler) CreateUser(c *fiber.Ctx) error {
	u := new(CreateUserReq)

	if err := c.BodyParser(u); err != nil {
		h.log.Error().Err(err).Msg("Failed to parse body")
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid JSON body",
		})
	}

	res, err := h.Service.CreateUser(c.Context(), u)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create user",
		})
	}

	if err := c.JSON(res); err != nil {
		h.log.Error().Err(err).Msg("Failed to decode response")
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to encode response",
		})
	}

	return nil
}

func (h *Handler) Login(c *fiber.Ctx) error {
	user := new(LoginUserReq)

	if err := c.BodyParser(user); err != nil {
		h.log.Error().Err(err).Msg("Failed to decode json body")
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid login attempt",
		})
	}

	u, err := h.Service.Login(c.Context(), user)
	if err != nil {
		h.log.Error().Err(err).Msg("Failed to login")
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid login attempt",
		})
	}

	jwtCookie := fiber.Cookie{
		Name:     "accessToken",
		Value:    u.AccessToken,
		MaxAge:   3600,
		Path:     "/",
		Domain:   config.Cfg.Server.Domain,
		Secure:   false,
		HTTPOnly: true,
	}

	c.Cookie(&jwtCookie)

	res := &LoginUserRes{
		ID:           u.ID,
		Role:         u.Role,
		AccessToken:  u.AccessToken,
		RefreshToken: u.RefreshToken,
	}

	return c.Status(http.StatusOK).JSON(res)
}

func (h *Handler) Logout(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	if token == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Authorization header is missing",
		})
	}

	token = strings.TrimPrefix(token, "Bearer ")

	if err := h.Service.Logout(c.Context(), token); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	c.ClearCookie("accessToken")

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "logout successful",
	})
}
