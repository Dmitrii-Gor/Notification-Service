# Notification-Service

## Configuration

The service reads its configuration from environment variables:

- `APP_ENV` &mdash; application environment name (defaults to `development`).
- `DATABASE_URL` &mdash; connection string for the primary database (**required**).
- `DB_TIMEOUT` &mdash; database operation timeout parsed via `time.ParseDuration` (defaults to `5s`).
- `JWT_SECRET` &mdash; secret key used to sign JWT tokens (**required**).
- `JWT_ACCESS_TTL` &mdash; access token lifetime parsed via `time.ParseDuration` (defaults to `15m`).
- `JWT_REFRESH_TTL` &mdash; refresh token lifetime parsed via `time.ParseDuration` (defaults to `720h`).

Set these variables in your environment or in a `.env` file before starting the service.
