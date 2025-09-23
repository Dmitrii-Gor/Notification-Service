package domain

import (
	"context"
	"errors"
	"time"
)

var (
	ErrEmailAlreadyExists = errors.New("email already exists")
	UserNotFound          = errors.New("user not found")
)

type User struct {
	ID             int64     `db:"id"`
	Email          string    `db:"email"`
	PasswordHash   string    `db:"password_hash"`
	TelegramChatID *int64    `db:"telegram_chat_id"`
	CreatedAt      time.Time `db:"created_at"`
	UpdatedAt      time.Time `db:"updated_at"`
}

type UserRepo interface {
	CreateWithEmail(ctx context.Context, email, passwordHash string) (*User, error)
	AttachEmail(ctx context.Context, userID int64, email, passwordHash string) (*User, error)
	CreateWithTelegram(ctx context.Context, chatID int64) (*User, error)
	LinkTelegram(ctx context.Context, userID, chatID int64) (*User, error)
	Delete(ctx context.Context, user *User) error
}
