package server

import (
	"context"
	"database/sql"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/BON4/gosubs/config"
	"github.com/BON4/gosubs/internal/domain"
	"github.com/BON4/gosubs/internal/middleware"
	tguser_handle "github.com/BON4/gosubs/internal/tguser/delivery/http"
	user_usecase "github.com/BON4/gosubs/internal/tguser/usecase/boil"

	account_handle "github.com/BON4/gosubs/internal/account/delivery/http"
	account_usecase "github.com/BON4/gosubs/internal/account/usecase/boil"

	sub_handle "github.com/BON4/gosubs/internal/subscription/delivery/http"
	sub_usecase "github.com/BON4/gosubs/internal/subscription/usecase/boil"

	tokengen "github.com/BON4/gosubs/pkg/tokenGen"
	"github.com/BON4/timedQ/pkg/ttlstore"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	_ "github.com/lib/pq"

	_ "github.com/BON4/gosubs/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// TODO: secret should NOT be stored in code or config
var SECRET = "8TKBySYHx9C7p0I2Pz3JlemhYkl10go24"

func SetupHandlers(s *Server) {
	s.g.Use(s.MidWar.CORS())

	uuc := user_usecase.NewBoilTgUserUsecase(s.DB, s.Logger.WithField("user usecase", struct{}{}))
	users_group := s.g.Group("/user", s.MidWar.AuthMiddleware())

	auc := account_usecase.NewBoilAccountUsecase(s.DB, s.Logger.WithField("account usecase", struct{}{}))
	acc_group := s.g.Group("/account", s.MidWar.AuthMiddleware())
	auth_group := s.g.Group("")

	suc := sub_usecase.NewBoilSubscriptionUsecase(s.DB, s.Logger.WithField("subscription usecase", struct{}{}))
	sub_group := s.g.Group("/sub", s.MidWar.AuthMiddleware())

	tguser_handle.NewTgUserHandler(users_group, uuc, s.MidWar, s.Cfg, s.Logger.WithField("users", ""))
	account_handle.NewAccountHandler(acc_group, auc, uuc, s.MidWar, s.Cfg, s.Logger.WithField("accounts", ""))
	account_handle.NewAuthHandler(auth_group, auc, s.MidWar, s.Cfg, s.Token, s.Store, s.Logger.WithField("auth", ""))
	sub_handle.NewSubscriptionHandler(sub_group, suc, uuc, s.MidWar, s.Cfg, s.Logger.WithField("subscription", ""))
}

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
	Cfg    config.ServerConfig
	Token  tokengen.Generator
	Store  *ttlstore.MapStore[string, *domain.Session]
	MidWar *middleware.ServerMiddleware
	DB     *sql.DB
}

func NewServer(configPath string) (*Server, error) {
	g := gin.Default()

	cfg, err := config.LoadServerConfig(configPath)
	if err != nil {
		return nil, err
	}

	log, err := setUpLogger(cfg.LogFile)
	if err != nil {
		return nil, err
	}

	log.Info(cfg)

	token, err := tokengen.NewJWTGenerator(SECRET)
	if err != nil {
		return nil, err
	}

	midW := middleware.NewServerMiddleware(token, cfg, log.WithField("middleware", ""))

	log.Infof("Loaded config: %+v", cfg)

	//TODO: postgres db connection string parse from config
	db, err := sql.Open("postgres", cfg.DBconn)
	if err != nil {
		panic(err)
	}

	store := ttlstore.NewMapStore[string, *domain.Session](context.Background(), cfg.Store)

	log.Infof("Creating db file in: %s", cfg.Store.SavePath)

	if err := store.Load(); err != nil {
		panic(err)
	}

	if err := store.Run(); err != nil {
		panic(err)
	}

	return &Server{
		g:      g,
		Logger: log,
		DB:     db,
		Cfg:    cfg,
		Token:  token,
		MidWar: midW,
		Store:  store,
	}, nil
}

func (s *Server) Run() error {
	srv := &http.Server{
		Handler: s.g,
		Addr:    s.Cfg.Port,
	}

	s.Logger.Infof("Running on: %s", s.Cfg.Port)

	//Swagger
	s.g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	SetupHandlers(s)

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.Logger.Errorf("listen: %s\n", err)
			return
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// Stop every store
	if err := s.Store.Close(); err != nil {
		return err
	}

	s.Logger.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		s.Logger.Error("Server Shutdown Err:", err)
		return err
	}

	select {
	case <-ctx.Done():
		s.Logger.Info("timeout of 5 seconds.")
	}
	s.Logger.Info("Server exiting")
	return nil
}
