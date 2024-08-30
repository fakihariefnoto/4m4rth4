package http

import (
	restMethod "github.com/fakihariefnoto/4m4rth4/module/presenter/rest"
)

func (h *httpHandler) AssignRoutes(httpPresenter restMethod.IPresenter) {

	// Create routes group.
	route := h.app.Group("/api/v1")

	// Routes for GET method:
	route.Get("/IsDelinquent/:customer_id", httpPresenter.IsDelinquent)
	route.Get("/GetOutstanding/:customer_id", httpPresenter.GetOutstanding)
	route.Get("/ListLoan/:customer_id", httpPresenter.GetOutstanding)
	route.Get("/GetLoanDetails/:loan_id", httpPresenter.GetOutstanding)
	route.Get("/GetCustomerData/:customer_id", httpPresenter.GetCustomer)

	// Routes for POST method:
	route.Post("/user/sign/up", httpPresenter.MakePayment)
	route.Post("/user/sign/in", httpPresenter.CreateLoan)
}
