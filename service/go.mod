module github.com/kzpolicy/user/service

go 1.14

require (
	github.com/PuerkitoBio/goquery v1.5.1
	github.com/friendsofgo/errors v0.9.2 // indirect
	github.com/go-sql-driver/mysql v1.5.0
	github.com/gofrs/uuid v3.3.0+incompatible // indirect
	github.com/saintfish/chardet v0.0.0-20120816061221-3af4cd4741ca
	github.com/spf13/cast v1.3.1 // indirect
	github.com/stretchr/testify v1.2.2
	github.com/volatiletech/inflect v0.0.1 // indirect
	github.com/volatiletech/null v8.0.0+incompatible // indirect
	github.com/volatiletech/sqlboiler v3.7.1+incompatible
	golang.org/x/crypto v0.0.0-20200820211705-5c72a883971a
	golang.org/x/net v0.0.0-20200202094626-16171245cfb2
	local.packages/models v0.0.0-00010101000000-000000000000

)

replace local.packages/models => ./../models
