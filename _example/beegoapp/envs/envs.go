package envs

import (
	"os"
	"strconv"

	"github.com/astaxie/beego"
)

const (
	ProdProdEnv         = "prod"
	BetaProdEnv         = "beta"
	AlphaProdEnv        = "alpha"
	DevProdEnv          = "dev"
	TestProdEnv         = "test"
	defaultServicePort  = "8080"
	UserTokenHeaderName = "X-USER-AUTH-TOKEN"
	secretKey           = "supersecret##jjgo"
	TestUsername        = "admin"
	TestUserPassword    = "jjgo"
)

//환경변수들
var (
	ProdEnv            = os.Getenv("PRODENV") //prod, beta, alpha, dev
	ServiceHost        = os.Getenv("SERVICE_HOST")
	ServicePort        = os.Getenv("SERVICE_PORT")
	DBHost             = os.Getenv("DB_HOST")
	DBPort             = os.Getenv("DB_PORT")
	DBName             = os.Getenv("DB_NAME")
	DBUser             = os.Getenv("DB_USER")
	DBPassword         = os.Getenv("DB_PASSWORD")
	SecretKey          = os.Getenv("SECRET_KEY")
	SystemToken        = os.Getenv("SYSTEM_TOKEN")
	SystemAdmin        = os.Getenv("SYSTEM_ADMIN")
	SystemPassword     = os.Getenv("SYSTEM_PASSWORD")
	EsURL              = os.Getenv("ES_URL")
	EsIndexPrefix      = os.Getenv("ES_INDEX_PREFIX")
	GoogleClientID     = os.Getenv("GOOGLE_CLIENT_ID")
	GoogleClientSecret = os.Getenv("GOOGLE_CLIENT_SECRET")
	SyncICSURL         = os.Getenv("SYNC_ICS_URL")
	SyncTrelloKey      = os.Getenv("SYNC_TRELLO_KEY")
	SyncTrelloToken    = os.Getenv("SYNC_TRELLO_TOKEN")
	SyncTrelloUser     = os.Getenv("SYNC_TRELLO_USER")
	AWSRegion          = os.Getenv("AWS_REGION")
	AWSClientID        = os.Getenv("AWS_CLIENT_ID")
	AWSClientSecret    = os.Getenv("AWS_CLIENT_SECRET")
	AWSUserPoolID      = os.Getenv("AWS_USER_POOL_ID")
)

func GetServicePort() int {
	port, err := strconv.ParseInt(ServicePort, 10, 64)
	if err != nil {
		panic("Service port가 없거나 잘못된 이름입니다.")
	}
	return int(port)
}

func InitEnvs() {
	if DBHost == "" || DBName == "" || DBUser == "" || DBPassword == "" {
		panic("DB 정보가 없습니다.")
	}
	if ProdEnv == "" {
		ProdEnv = DevProdEnv
	} else if ProdEnv == "prod" {
		ProdEnv = ProdProdEnv

	}

	if ServicePort == "" {
		ServicePort = defaultServicePort
		beego.BConfig.Listen.HTTPPort = GetServicePort()
	}
	if DBPort == "" {
		DBPort = "3306"
	}
	if SystemAdmin == "" {
		SystemAdmin = "admin@circle.land"
	}
	if SystemPassword == "" {
		SystemPassword = "1"
	}
	if SecretKey == "" {
		SecretKey = secretKey
	}

	if ProdEnv == ProdProdEnv {
		beego.BConfig.RunMode = "prod"
	} else {
		beego.BConfig.RunMode = "dev"
	}

	beego.BConfig.CopyRequestBody = true
}
