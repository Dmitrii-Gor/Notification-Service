CREATE TABLE IF NOT EXISTS users (
                                     id BIGSERIAL PRIMARY KEY,
                                     email TEXT UNIQUE,               -- может быть NULL
                                     password_hash TEXT,              -- может быть NULL
                                     telegram_chat_id BIGINT UNIQUE,  -- может быть NULL
                                     created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
    );

-- индекс для быстрого поиска по email (только если он не NULL)
CREATE UNIQUE INDEX IF NOT EXISTS idx_users_email ON users(email) WHERE email IS NOT NULL;

-- индекс для быстрого поиска по telegram_chat_id (только если он не NULL)
CREATE UNIQUE INDEX IF NOT EXISTS idx_users_telegram ON users(telegram_chat_id) WHERE telegram_chat_id IS NOT NULL;
