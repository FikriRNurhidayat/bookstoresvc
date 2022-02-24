# Config Package

This directory is being used to store the configuration object.
New function should be added if there will be new configuration value.

For example:

```go
func GetDatabaseConnectionString() string {
  return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
    Getenv("DATABASE_USER", "fain"),
    Getenv("DATABASE_PASSWORD", "awurenwae"),
    Getenv("DATABASE_HOST", "localhost"),
    Getenv("DATABASE_PORT", "5432"),
    Getenv("DATABASE_NAME", "bookstoresvc_development"),
    Getenv("DATABASE_SSL_MODE", "disable"),
  )
}
```

Also this package have `GetEnv` function,
it is being used to fetch environment variable of the process.
