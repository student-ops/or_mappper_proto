module main

go 1.22.2

replace main/HawkORM => ./HawkORM

require main/HawkORM v0.0.0-00010101000000-000000000000

require (
	github.com/HarrisonEagle/HawkORM v0.0.0-20231004173236-bf8f87020e4f // indirect
	github.com/go-sql-driver/mysql v1.7.1 // indirect
)
