package configuration

import (
	"os"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type AppConfigServer struct {
	logger      *zap.Logger
	gormDB      *gorm.DB
	Environment string
	server      *gin.Engine
}

func (a *AppConfigServer) CreateGormPostgresqlDB() error {
	a.logger.Info("[AppConfigServer-CreateGormMysqlDB]", zap.String("Environment", a.Environment))
	dsn := "postgresql://neondb_owner:Z4zf7CFvglQK@ep-shrill-sound-a5g0r3g2.us-east-2.aws.neon.tech/dstock?sslmode=require"
	a.logger.Info("[AppConfigServer-CreateGormMysqlDB]", zap.String("EnvironmentMysqlAddr", dsn))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	a.gormDB = db
	return nil
}

func (a *AppConfigServer) SetLogger(logger *zap.Logger) {
	a.logger = logger
}

func InitAppConfigServer() {
	zapLogger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	appServerConfig := &AppConfigServer{
		Environment: os.Getenv("ENVIRONMENT"), // from docker-compose
	}
	appServerConfig.SetLogger(zapLogger)
	appServerConfig.server = gin.Default()
	apiGroup := appServerConfig.server.Group("/api")
	stockapi := apiGroup.Group("stock")
	//userGroup := apiGroup.Group("/users/:id")
	// transactionGroup := userGroup.Group("/transactions")
	tySInitStockRouter(appServerConfig.logger, stockapi, appServerConfig)
	appServerConfig.server.Run(":8080")
}
