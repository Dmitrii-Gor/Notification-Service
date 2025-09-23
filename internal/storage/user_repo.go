package storage

import (
	"context"
	"database/sql"
	"errors"
	"github.com/Dmitrii-Gor/notification-bot/internal/domain"
	"github.com/lib/pq"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) Delete(ctx context.Context, user *domain.User) error {
	result, err := r.db.ExecContext(ctx, `DELETE FROM users WHERE id=$1`, user.ID)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return domain.UserNotFound
	}
	return nil
}

func (r *UserRepo) CreateWithEmail(ctx context.Context, email string, passwordHash string) (*domain.User, error) {
	var user domain.User
	err := r.db.QueryRowContext(ctx,
		`INSERT INTO users (email, password_hash)
		VALUES ($1, $2)
		RETURNING id, email, password_hash, telegram_chat_id, created_at, updated_at`,
		email, passwordHash,
	).Scan(&user.ID, &user.Email, &user.PasswordHash, &user.TelegramChatID, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		var pgErr *pq.Error
		if errors.As(err, &pgErr) && pgErr.Code == pq.ErrorCode("23505") {
			return nil, domain.ErrEmailAlreadyExists
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepo) AttachEmail(ctx context.Context, userID int64, email, passwordHash string) (*domain.User, error) {
	var user domain.User
	err := r.db.QueryRowContext(
		ctx,
		`UPDATE users
         SET email = $1, password_hash = $2, updated_at = now()
         WHERE id = $3
         RETURNING id, email, password_hash, telegram_chat_id, created_at, updated_at`,
		email, passwordHash, userID,
	).Scan(&user.ID, &user.Email, &user.PasswordHash, &user.TelegramChatID, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepo) CreateWithTelegram(ctx context.Context, chatID int64) (*domain.User, error) {
	var user domain.User
	err := r.db.QueryRowContext(
		ctx,
		`INSERT INTO users (telegram_chat_id)
         VALUES ($1)
         RETURNING id, email, password_hash, telegram_chat_id, created_at, updated_at`,
		chatID,
	).Scan(&user.ID, &user.Email, &user.PasswordHash, &user.TelegramChatID, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepo) LinkTelegram(ctx context.Context, userID int64, chatID int64) (*domain.User, error) {
	var user domain.User
	err := r.db.QueryRowContext(
		ctx,
		`UPDATE users
         SET telegram_chat_id = $1, updated_at = now()
         WHERE id = $2
         RETURNING id, email, password_hash, telegram_chat_id, created_at, updated_at`,
		chatID, userID,
	).Scan(&user.ID, &user.Email, &user.PasswordHash, &user.TelegramChatID, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
