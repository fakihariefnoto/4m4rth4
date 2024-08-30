package main

import (
	configPkg "billingapp/pkg/config"
	databasePkg "billingapp/service/database"

	customerRepository "billingapp/internal/repository/customer"
	loanRepository "billingapp/internal/repository/loan"
	paymentRepository "billingapp/internal/repository/payment"

	customerUsecase "billingapp/internal/usecase/customer"
	loanUsecase "billingapp/internal/usecase/loan"
	paymentUsecase "billingapp/internal/usecase/payment"

	presenterHttp "billingapp/internal/presenter/rest"

	httpHandler "billingapp/handler/http"

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

	customerRepo := customerRepository.New(customerDBConn.DB)
	loanRepo := loanRepository.New(loanDBConn.DB)
	paymentRepo := paymentRepository.New(paymentDBConn.DB)

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
