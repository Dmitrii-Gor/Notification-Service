package domain

import "time"

type User struct {
	ID             int64     `db:"id"`
	Email          string    `db:"email"`
	PasswordHash   string    `db:"password_hash"`
	TelegramChatID *int64    `db:"telegram_chat_id"`
	CreatedAt      time.Time `db:"created_at"`
	UpdatedAt      time.Time `db:"updated_at"`
}
