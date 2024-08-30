package main

import (
	configPkg "github.com/fakihariefnoto/4m4rth4/pkg/config"
	databasePkg "github.com/fakihariefnoto/4m4rth4/pkg/database"

	customerRepository "github.com/fakihariefnoto/4m4rth4/module/repository/customer"
	loanRepository "github.com/fakihariefnoto/4m4rth4/module/repository/loan"
	paymentRepository "github.com/fakihariefnoto/4m4rth4/module/repository/payment"

	customerUsecase "github.com/fakihariefnoto/4m4rth4/module/usecase/customer"
	loanUsecase "github.com/fakihariefnoto/4m4rth4/module/usecase/loan"
	paymentUsecase "github.com/fakihariefnoto/4m4rth4/module/usecase/payment"

	presenterHttp "github.com/fakihariefnoto/4m4rth4/module/presenter/rest"

	httpHandler "github.com/fakihariefnoto/4m4rth4/handler/http"

	"log"
)

func main() {
	err := configPkg.Init()
	chekErrMain("config", err)

	config := configPkg.Get()

	checkDB(config)

	databasePkg.Init(config)

	err = databasePkg.AddConnection(databasePkg.CustomerDB)
	chekErrMain("connect customer db", err)
	err = databasePkg.AddConnection(databasePkg.LoanDB)
	chekErrMain("connect loan db", err)
	err = databasePkg.AddConnection(databasePkg.PaymentDB)
	chekErrMain("connect payment db", err)

	customerDBConn, err := databasePkg.GetDBConn(databasePkg.CustomerDB)
	chekErrMain("get customer db conn", err)
	loanDBConn, err := databasePkg.GetDBConn(databasePkg.LoanDB)
	chekErrMain("get loan db conn", err)
	paymentDBConn, err := databasePkg.GetDBConn(databasePkg.PaymentDB)
	chekErrMain("get payment db conn", err)

	customerRepo := customerRepository.New(customerDBConn)
	loanRepo := loanRepository.New(loanDBConn)
	paymentRepo := paymentRepository.New(paymentDBConn)

	loanUse := loanUsecase.New(loanRepo)
	customerUse := customerUsecase.New(customerRepo)
	paymentUse := paymentUsecase.New(paymentRepo, loanRepo)

	httpPresenter := presenterHttp.New(loanUse, paymentUse, customerUse)

	httpServer := httpHandler.New()
	httpServer.AssignRoutes(httpPresenter)
	httpServer.Start(config.Port)

}

func chekErrMain(label string, err error) {
	if err != nil {
		log.Fatalln("Init process : ", label, " error -> ", err)
	}
}

func checkDB(cfg configPkg.Config) {
	for _, db := range cfg.DB {
		databasePkg.CreateDB(db.Name + ".sql")
	}
}
