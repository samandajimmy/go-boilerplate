module go-boiler-plate

go 1.17

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-pg/pg/v9 v9.2.1
	github.com/joho/godotenv v1.4.0
	github.com/labstack/echo/v4 v4.7.2
	github.com/lib/pq v1.10.5
	golang.org/x/crypto v0.0.0-20220427172511-eb4f295cb31f
	repo.pegadaian.co.id/ms-pds/modules/pgdlogger v1.0.0
	repo.pegadaian.co.id/ms-pds/modules/pgdutil v0.0.2
)

replace (
	repo.pegadaian.co.id/ms-pds/modules/pgdlogger => repo.pegadaian.co.id/ms-pds/modules/pgdlogger.git v1.0.0
	repo.pegadaian.co.id/ms-pds/modules/pgdutil => repo.pegadaian.co.id/ms-pds/modules/pgdutil.git v0.0.2
)

require (
	github.com/codemodus/kace v0.5.1 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/go-pg/zerochecker v0.2.0 // indirect
	github.com/go-playground/locales v0.14.0 // indirect
	github.com/go-playground/universal-translator v0.18.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/labstack/gommon v0.3.1 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/onsi/ginkgo v1.16.4 // indirect
	github.com/segmentio/asm v1.2.0 // indirect
	github.com/segmentio/encoding v0.3.5 // indirect
	github.com/sirupsen/logrus v1.8.1 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.1 // indirect
	github.com/vmihailenco/bufpool v0.1.11 // indirect
	github.com/vmihailenco/msgpack/v4 v4.3.12 // indirect
	github.com/vmihailenco/tagparser v0.1.2 // indirect
	golang.org/x/net v0.0.0-20220425223048-2871e0cb64e4 // indirect
	golang.org/x/sys v0.0.0-20220502124256-b6088ccd6cba // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
	gopkg.in/go-playground/validator.v9 v9.31.0 // indirect
	mellium.im/sasl v0.2.1 // indirect
)
