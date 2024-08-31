package payment

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	loanModel "billingapp/internal/model/loan"
	paymentModel "billingapp/internal/model/payment"
)

func (p *payment) MakePayment(paymentReq PaymentReq) error {
	// get loan details by loan id
	resp, err := p.loanRepo.GetLoanDetailsByID(context.Background(), paymentReq.LoanDetailsID)
	if err != nil {
		return err
	}

	// check if amount is same or not, return error
	if fmt.Sprintf("%.2f", resp.Amount) != fmt.Sprintf("%.2f", paymentReq.Amount) {
		return errors.New("Invalid amount")
	}

	summary := fmt.Sprintf("Payment for loan details %s with amount %f", paymentReq.LoanDetailsID, paymentReq.Amount)
	details := fmt.Sprintf(`
		User : %v
		Loan Details ID : %v
		Amount : %f
		Date : %v
	`, paymentReq.CustomerID, paymentReq.LoanDetailsID, paymentReq.Amount, time.Now())

	// if amount valid, make payment
	paymentID, err := p.paymentRepo.InsertPaymentHistory(context.Background(), paymentModel.Payment{
		CustomerID: paymentReq.CustomerID,
		Details:    details,
		Summary:    summary,
		Amount:     paymentReq.Amount,
		Status:     paymentModel.PaymentStatusReceived,
	})
	if err != nil {
		log.Println("Error when inser payment ", err)
		return err
	}

	// if payment success, update loan detail status
	err = p.loanRepo.UpdateLoanDetailsStatus(context.Background(), paymentReq.LoanDetailsID, paymentID, loanModel.LoanDetailStatusPaid)
	if err != nil {
		// WE CAN NOT ROLL BACK PAYMENT
		// BECAUSE OF DIFF. TX, WE SHOULD TRY TO RETRY, WE CAN USE MSG QUEUE FOR THIS
		// LIMITED TIME, WE JUST IGNORE FOR NOW, CRON DO SWEEPING
		log.Println("Error when update loan detail status ", err)
	}

	return nil
}
