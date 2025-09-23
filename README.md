# Notification-Service

## Configuration

The service reads its configuration from environment variables:

- `APP_ENV` - application environment name (defaults to `development`).
- `DATABASE_URL` - connection string for the primary database (**required**).
- `DB_TIMEOUT` - database operation timeout parsed via `time.ParseDuration` (defaults to `5s`).
- `JWT_SECRET` - secret key used to sign JWT tokens (**required**).
- `JWT_ACCESS_TTL` - access token lifetime parsed via `time.ParseDuration` (defaults to `15m`).
- `JWT_REFRESH_TTL` - refresh token lifetime parsed via `time.ParseDuration` (defaults to `720h`).

Set these variables in your environment or in a `.env` file before starting the service.
