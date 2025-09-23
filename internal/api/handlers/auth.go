package handlers

import (
	"context"
	"errors"
	"github.com/Dmitrii-Gor/notification-bot/pkg/logger"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"

	"github.com/Dmitrii-Gor/notification-bot/internal/domain"
)

type AuthHandler struct {
	users      domain.UserRepo
	jwtSecret  []byte
	accessTTL  time.Duration
	refreshTTL time.Duration
	dbTimeout  time.Duration
}

func NewAuthHandler(
	users domain.UserRepo,
	jwtSecret []byte,
	accessTTL, refreshTTL time.Duration,
	dbTimeout time.Duration,
) *AuthHandler {
	return &AuthHandler{
		users:      users,
		jwtSecret:  jwtSecret,
		accessTTL:  accessTTL,
		refreshTTL: refreshTTL,
		dbTimeout:  dbTimeout,
	}
}

type emailRegisterRequest struct {
	Email    string `json:"email"    binding:"required,email"`
	Password string `json:"password" binding:"required,min=8,max=42"`
}

type tokenPair struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

func (h *AuthHandler) RegisterWithEmail(c *gin.Context) {
	var req emailRegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid_payload"})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), h.dbTimeout)
	defer cancel()

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		logger.Error("hash password", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal_error"})
		return
	}

	user, err := h.users.CreateWithEmail(ctx, req.Email, string(hash))
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrEmailAlreadyExists):
			c.JSON(http.StatusConflict, gin.H{"error": "email_in_use"})
		default:
			logger.Error("create user", zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal_error"})
		}
		return
	}

	access, err := h.buildToken(user, h.accessTTL)
	if err != nil {
		logger.Error("access token", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal_error"})
		return
	}

	refresh, err := h.buildToken(user, h.refreshTTL)
	if err != nil {
		logger.Error("refresh token", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal_error"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"user": gin.H{
			"id":    user.ID,
			"email": user.Email,
		},
		"tokens": tokenPair{
			AccessToken:  access,
			RefreshToken: refresh,
		},
	})
}

func (h *AuthHandler) buildToken(u *domain.User, ttl time.Duration) (string, error) {
	now := time.Now().UTC()
	claims := jwt.MapClaims{
		"sub":   u.ID,
		"email": u.Email,
		"iat":   now.Unix(),
		"exp":   now.Add(ttl).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(h.jwtSecret)
}
