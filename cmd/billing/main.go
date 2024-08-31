package main

import (
	configPkg "billingapp/pkg/config"
	databasePkg "billingapp/service/database"
	"io"
	"os"

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
	startInit("log")
	LOG_FILE := "log/app.log"
	logFile, err := os.OpenFile(LOG_FILE, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Panic(err)
	}
	defer logFile.Close()

	log.SetFlags(log.Lshortfile | log.LstdFlags)
	mw := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(mw)
	finishInit("log")

	startInit("config")
	err = configPkg.Init()
	chekErrMain("config", err)
	finishInit("config")

	config := configPkg.Get()

	checkDB(config)

	startInit("db")
	databasePkg.Init(config)
	finishInit("db")

	startInit("customerDB conn")
	err = databasePkg.AddConnection(databasePkg.CustomerDB)
	chekErrMain("connect customer db", err)
	finishInit("customerDB conn")

	startInit("loanDB conn")
	err = databasePkg.AddConnection(databasePkg.LoanDB)
	chekErrMain("connect loan db", err)
	finishInit("loanDB conn")

	startInit("paymentDB conn")
	err = databasePkg.AddConnection(databasePkg.PaymentDB)
	chekErrMain("connect payment db", err)
	finishInit("paymentDB conn")

	startInit("get customerDB conn")
	customerDBConn, err := databasePkg.GetDBConn(databasePkg.CustomerDB)
	chekErrMain("get customer db conn", err)
	finishInit("get customerDB conn")

	startInit("get loanDB conn")
	loanDBConn, err := databasePkg.GetDBConn(databasePkg.LoanDB)
	chekErrMain("get loan db conn", err)
	finishInit("get loanDB conn")

	startInit("get paymentDB conn")
	paymentDBConn, err := databasePkg.GetDBConn(databasePkg.PaymentDB)
	chekErrMain("get payment db conn", err)
	finishInit("get paymentDB conn")

	startInit("repository")
	customerRepo := customerRepository.New(customerDBConn.DB)
	loanRepo := loanRepository.New(loanDBConn.DB)
	paymentRepo := paymentRepository.New(paymentDBConn.DB)
	finishInit("repository")

	startInit("usecase")
	loanUse := loanUsecase.New(loanRepo)
	customerUse := customerUsecase.New(customerRepo, loanRepo)
	paymentUse := paymentUsecase.New(paymentRepo, loanRepo)
	finishInit("usecase")

	startInit("presenter")
	httpPresenter := presenterHttp.New(loanUse, paymentUse, customerUse)
	finishInit("presenter")

	httpServer := httpHandler.New()
	httpServer.AssignRoutes(httpPresenter)
	httpServer.Start(config.Port)

}

func chekErrMain(label string, err error) {
	if err != nil {
		log.Fatalln("Init process : ", label, " error -> ", err)
	}
}

func startInit(label string) {
	log.Println("Initialize " + label + "....")
}

func finishInit(label string) {
	log.Println(label + " initialized sucessfully.")
}
