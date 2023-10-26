package delivery

import (
	"fmt"
	"mnc/config"
	"mnc/delivery/controller"
	"mnc/delivery/middleware"
	"mnc/manager"
	exceptions "mnc/utils/exception"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Server struct {
	useCaseManager manager.UsecaseManager
	engine         *gin.Engine
	host           string
	log            *logrus.Logger
}

func (s *Server) Run() {
	s.initController()
	err := s.engine.Run(s.host)
	if err != nil {
		panic(err)
	}
}

func (s *Server) initController() {
	s.engine.Use(middleware.LogRequestMiddleware(s.log))
	// semua controller disini
	controller.NewBankController(s.useCaseManager.BankUsecase(), s.engine)
	controller.NewCustomerController(s.useCaseManager.CustomerUsecase(), s.engine)
	controller.NewMerchantController(s.useCaseManager.MerchantUsecase(), s.engine)
	controller.NewTransactionController(s.useCaseManager.TransactionUsecase(), s.engine)
}

func NewServer() *Server {
	cfg, err := config.NewConfig()
	exceptions.CheckError(err)
	infraManager, _ := manager.NewInfraManager(cfg)
	repoManager := manager.NewRepoManager(infraManager)
	useCaseManager := manager.NewUseCaseManager(repoManager)

	engine := gin.Default()
	host := fmt.Sprintf("%s:%s", cfg.ApiHost, cfg.ApiPort)
	return &Server{
		useCaseManager: useCaseManager,
		engine:         engine,
		host:           host,
		log:            logrus.New(),
	}
}
