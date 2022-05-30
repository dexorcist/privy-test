package infra

import (
	"database/sql"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
	"privy-test/docs"
	"runtime"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/zerolog"
)

type Infra interface {
	Config() MergeConfig
	Logger() zerolog.Logger
	SQLDB() *sql.DB
}

type infraCtx struct {
	App          AppConfig
	CustomConfig AppCustomConfig
}

func (c infraCtx) Config() MergeConfig {
	return MergeConfig{
		c.App,
		c.CustomConfig,
	}
}

// New construct new infrastructure object manager
func New(namespace, buildVersion string) Infra {
	viper.AutomaticEnv()
	viper.SetConfigFile(".env")
	// Find and read the config file
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(fmt.Sprintf("Error Load Config %s", err.Error()))
	}
	var App AppConfig
	if err := viper.Unmarshal(&App); err != nil {
		log.Fatal(fmt.Sprintf("Error Unmarshal Config %s", err.Error()))
	}

	viper.SetConfigFile(App.CustomConfig)
	err = viper.ReadInConfig()
	if err != nil {
		log.Fatal(fmt.Sprintf("Error Load Config %s", err.Error()))
	}
	var CustomConfig AppCustomConfig
	if err := viper.Unmarshal(&CustomConfig); err != nil {
		log.Fatal(fmt.Sprintf("Error Unmarshal Config %s", err.Error()))
	}

	App.Namespace = namespace
	App.GoVersion = runtime.Version()
	App.BuildVersion = buildVersion
	docs.SwaggerInfo.Version = buildVersion
	docs.SwaggerInfo.Host = CustomConfig.Swagger.SwaggerTemplate.GoTemplate.Host
	docs.SwaggerInfo.BasePath = CustomConfig.Swagger.SwaggerTemplate.GoTemplate.BasePath
	docs.SwaggerInfo.Schemes = []string{CustomConfig.Swagger.SwaggerTemplate.GoTemplate.Schemes}
	docs.SwaggerInfo.Title = CustomConfig.Swagger.SwaggerTemplate.GoTemplate.Title
	docs.SwaggerInfo.Description = CustomConfig.Swagger.SwaggerTemplate.GoTemplate.Description

	return &infraCtx{
		App:          App,
		CustomConfig: CustomConfig,
	}
}

var (
	zerologgerOnce sync.Once
	zerologger     zerolog.Logger
)

// Logger singleton to get logger object
func (c *infraCtx) Logger() zerolog.Logger {
	zerologgerOnce.Do(func() {
		mainCfg := c.Config()

		var lvl zerolog.Level
		switch mainCfg.LoggerLevel {
		case "debug":
			lvl = zerolog.DebugLevel
		case "info":
			lvl = zerolog.InfoLevel
		case "warn":
			lvl = zerolog.WarnLevel
		case "error":
			lvl = zerolog.ErrorLevel
		case "fatal":
			lvl = zerolog.FatalLevel
		case "panic":
			lvl = zerolog.PanicLevel
		default:
			lvl = zerolog.InfoLevel
		}

		zerolog.SetGlobalLevel(lvl)
		zerolog.TimestampFieldName = "timestamp"
		zerolog.TimeFieldFormat = time.RFC3339Nano
		logger := zerolog.New(zerolog.MultiLevelWriter(os.Stdout)).With().Logger()

		if mainCfg.LoggerWithTimestamp {
			logger = logger.With().Timestamp().Logger()
		}

		zerologger = logger
	})

	return zerologger
}

// SQLDB singleton to get sql database object
func (c *infraCtx) SQLDB() *sql.DB {
	db, err := sql.Open("mysql", c.CustomConfig.Database.Host)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	return db
}
