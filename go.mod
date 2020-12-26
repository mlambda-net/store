module github.com/mlambda-net/store

go 1.16

require (
	cloud.google.com/go v0.81.0
	github.com/AsynkronIT/protoactor-go v0.0.0-20201117200131-5ae73aa8b6fe
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/aws/aws-sdk-go v1.37.15
	github.com/elastic/go-elasticsearch/v8 v8.0.0-20210311100734-5d6b0c808457
	github.com/etherlabsio/healthcheck v0.0.0-20191224061800-dd3d2fd8c3f6
	github.com/gin-gonic/gin v1.6.3
	github.com/go-openapi/spec v0.20.3 // indirect
	github.com/go-pg/pg/v10 v10.8.0
	github.com/go-redis/redis/v8 v8.7.1
	github.com/gogo/protobuf v1.3.2
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/google/uuid v1.2.0
	github.com/gorilla/mux v1.8.0
	github.com/graphql-go/graphql v0.7.9
	github.com/graphql-go/handler v0.2.3
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mlambda-net/monads v1.0.4
	github.com/mlambda-net/net v1.0.45
	github.com/sirupsen/logrus v1.7.0
	github.com/stretchr/testify v1.7.0
	github.com/swaggo/files v0.0.0-20190704085106-630677cd5c14
	github.com/swaggo/gin-swagger v1.3.0
	github.com/swaggo/swag v1.7.0
	golang.org/x/net v0.0.0-20210423184538-5f58ad60dda6 // indirect
	golang.org/x/sys v0.0.0-20210423185535-09eb48e85fd7 // indirect
	golang.org/x/text v0.3.6
	golang.org/x/tools v0.1.0 // indirect
	google.golang.org/api v0.45.0
	google.golang.org/genproto v0.0.0-20210423144448-3a41ef94ed2b // indirect
)

replace github.com/swaggo/gin-swagger v1.3.0 => github.com/mlambda-net/gin-swagger v1.3.2
