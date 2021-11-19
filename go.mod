module gitlab.com/procsy/attorney/api

go 1.16

require (
	github.com/go-playground/validator/v10 v10.8.0
	github.com/gofrs/uuid v4.0.0+incompatible
	github.com/golang-jwt/jwt v3.2.1+incompatible
	github.com/labstack/echo v3.3.10+incompatible
	github.com/labstack/gommon v0.3.0 // indirect
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/viper v1.8.1
	go.mongodb.org/mongo-driver v1.5.4
	golang.org/x/crypto v0.0.0-20210817164053-32db794688a5 // indirect
	golang.org/x/net v0.0.0-20210428140749-89ef3d95e781 // indirect
	golang.org/x/sys v0.0.0-20210823070655-63515b42dcdf // indirect
)

// replace go.cypherpunks.ru/gogost/v5 => ../gogost-5.6.0
