package server

import (
	"context"
	"os"

	"github.com/BON4/gosubs/internal/domain"
	tokengen "github.com/BON4/gosubs/pkg/tokenGen"
	"github.com/BON4/timedQ/pkg/ttlstore"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// TODO: secret should NOT be stored in code or config
var SECRET = "secret"

func setUpLogger(fileName string) (*logrus.Logger, error) {
	// instantiation
	logger := logrus.New()

	if len(fileName) == 0 {
		logger.Out = os.Stdout
	} else {
		//Write to file
		src, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		if err != nil {
			return nil, err
		}
		//Set output
		logger.Out = src
	}

	//Set log level
	logger.SetLevel(logrus.DebugLevel)

	//Format log
	logger.SetFormatter(&logrus.TextFormatter{})
	return logger, nil
}

type Server struct {
	g      *gin.Engine
	Logger *logrus.Logger
	Cfg    ServerConfig
	Token  tokengen.Generator
	Store  *ttlstore.MapStore[string, *domain.Session]
}

func NewServer(configPath string) (*Server, error) {
	g := gin.Default()

	cfg, err := LoadServerConfig(configPath)
	if err != nil {
		return nil, err
	}

	log, err := setUpLogger(cfg.AppConfig.LogFile)
	if err != nil {
		return nil, err
	}

	token, err := tokengen.NewJWTGenerator(SECRET)
	if err != nil {
		return nil, err
	}

	log.Infof("Loaded config: %+v", cfg)

	return &Server{
		g:      g,
		Logger: log,
		Cfg:    cfg,
		Token:  token,
		Store:  ttlstore.NewMapStore[string, *domain.Session](context.Background(), cfg.TTLStoreConfig),
	}, nil
}
