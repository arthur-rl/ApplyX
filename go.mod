module github.com/arthur-rl/applyx

go 1.17

require (
	github.com/gofiber/fiber/v2 v2.23.0
	golang.org/x/crypto v0.0.0-20211215153901-e495a2d5b3d3
	gorm.io/datatypes v1.0.5
	gorm.io/driver/postgres v1.2.3
	gorm.io/gorm v1.22.4
)

require (
	github.com/andybalholm/brotli v1.0.4 // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/jackc/chunkreader/v2 v2.0.1 // indirect
	github.com/jackc/pgconn v1.10.1 // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgproto3/v2 v2.2.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20200714003250-2b9c44734f2b // indirect
	github.com/jackc/pgtype v1.9.1 // indirect
	github.com/jackc/pgx/v4 v4.14.1 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.4 // indirect
	github.com/klauspost/compress v1.13.6 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasthttp v1.31.0 // indirect
	github.com/valyala/tcplisten v1.0.0 // indirect
	golang.org/x/sys v0.0.0-20211216021012-1d35b9e2eb4e // indirect
	golang.org/x/text v0.3.7 // indirect
	gorm.io/driver/mysql v1.2.2 // indirect
	gorm.io/driver/sqlite v1.2.6 // indirect
	gorm.io/driver/sqlserver v1.2.1 // indirect
)

replace (
	github.com/arthur-rl/applyx/database => ./database
	github.com/arthur-rl/applyx/routes/api => ./routes/api
	github.com/arthur-rl/applyx/routes/auth => ./routes/auth
	github.com/arthur-rl/applyx/util => ./util
)
