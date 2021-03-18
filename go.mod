module nippo

go 1.15

require (
	github.com/go-sql-driver/mysql v1.5.0 // indirect
	github.com/skratchdot/open-golang v0.0.0-20200116055534-eef842397966
	package.local/database v0.0.0-00010101000000-000000000000
)

replace package.local/database => ./database
