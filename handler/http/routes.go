package http

import (
	restMethod "billingapp/internal/presenter/rest"
)

func (h *httpHandler) AssignRoutes(httpPresenter restMethod.IPresenter) {

	// Create routes group.
	route := h.app.Group("/api/v1")

	// Routes for GET method:
	route.Get("/IsDelinquent/:customer_id", httpPresenter.IsDelinquent)
	route.Get("/GetOutstanding/:loan_id", httpPresenter.GetOutstanding)
	route.Get("/ListLoan/:customer_id", httpPresenter.ListLoan)
	route.Get("/GetLoanDetails/:loan_id", httpPresenter.GetLoanDetails)
	route.Get("/GetCustomer/:customer_id", httpPresenter.GetCustomer)

	// Routes for POST method:
	route.Post("/MakePayment", httpPresenter.MakePayment)
	route.Post("/CreateLoan", httpPresenter.CreateLoan)
	route.Post("/CreateCustomer", httpPresenter.CreateCustomer)
}
