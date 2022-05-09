module go-boiler-plate

go 1.17

require (
	github.com/golang-jwt/jwt/v4 v4.4.1
	github.com/golang-migrate/migrate/v4 v4.15.2
	github.com/golang/mock v1.6.0
	github.com/jmoiron/sqlx v1.3.5
	github.com/joho/godotenv v1.4.0
	github.com/labstack/echo/v4 v4.7.2
	github.com/lib/pq v1.10.5
	github.com/onsi/ginkgo/v2 v2.1.3
	github.com/onsi/gomega v1.19.0
	github.com/spf13/viper v1.11.0
	golang.org/x/crypto v0.0.0-20220427172511-eb4f295cb31f
	gopkg.in/go-playground/validator.v9 v9.31.0
	repo.pegadaian.co.id/ms-pds/modules/pgdlogger v1.0.0
	repo.pegadaian.co.id/ms-pds/modules/pgdutil v0.0.2
)

replace (
	repo.pegadaian.co.id/ms-pds/modules/pgdlogger => repo.pegadaian.co.id/ms-pds/modules/pgdlogger.git v1.0.0
	repo.pegadaian.co.id/ms-pds/modules/pgdutil => repo.pegadaian.co.id/ms-pds/modules/pgdutil.git v0.0.2
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/fsnotify/fsnotify v1.5.1 // indirect
	github.com/go-playground/locales v0.14.0 // indirect
	github.com/go-playground/universal-translator v0.18.0 // indirect
	github.com/golang-jwt/jwt v3.2.2+incompatible // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/labstack/gommon v0.3.1 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/magiconair/properties v1.8.6 // indirect
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/mitchellh/mapstructure v1.4.3 // indirect
	github.com/pelletier/go-toml v1.9.4 // indirect
	github.com/pelletier/go-toml/v2 v2.0.0-beta.8 // indirect
	github.com/sirupsen/logrus v1.8.1 // indirect
	github.com/spf13/afero v1.8.2 // indirect
	github.com/spf13/cast v1.4.1 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/subosito/gotenv v1.2.0 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.1 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	golang.org/x/net v0.0.0-20220425223048-2871e0cb64e4 // indirect
	golang.org/x/sys v0.0.0-20220502124256-b6088ccd6cba // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/time v0.0.0-20220224211638-0e9765cccd65 // indirect
	gopkg.in/ini.v1 v1.66.4 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)
