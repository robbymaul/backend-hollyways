package handlers

import (
	dtoResult "hollyways/dto/result"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

func (h *transactionHandler) GetPayment(c *gin.Context) {
	var err error

	transactionId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dtoResult.ErrorResult{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	transaction, err := h.TransactionRepository.GetTransaction(transactionId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtoResult.ErrorResult{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	var s = snap.Client{}
	s.New(os.Getenv("SERVER_KEY"), midtrans.Sandbox)

	request := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  strconv.Itoa(transactionId),
			GrossAmt: int64(transaction.Donation),
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: transaction.User.FullName,
			Email: transaction.User.Email,
		},
	}

	snapResponse, err := s.CreateTransaction(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtoResult.ErrorResult{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dtoResult.SuccessResult{
		Status: http.StatusOK,
		Data:   snapResponse,
	})
}

func (h *transactionHandler) Notification(c *gin.Context) {
	var err error
	var notificationPayload map[string]interface{}

	if err := c.ShouldBind(&notificationPayload); err != nil {
		c.JSON(http.StatusBadRequest, dtoResult.ErrorResult{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	status := notificationPayload["transaction_status"].(string)
	fraud := notificationPayload["fraud_status"].(string)
	orderId, err := strconv.Atoi(notificationPayload["order_id"].(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, dtoResult.ErrorResult{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	transaction, _ := h.TransactionRepository.GetTransaction(orderId)

	if status == "capture" {
		if fraud == "challange" {
			h.TransactionRepository.UpdateTransaction("pending", orderId)
		} else if status == "accept" {
			sendMail("success", transaction)
			h.TransactionRepository.UpdateTransaction("success", orderId)
		}
	} else if status == "settlement" {
		sendMail("success", transaction)
		h.TransactionRepository.UpdateTransaction("success", orderId)
	} else if status == "deny" {
		h.TransactionRepository.UpdateTransaction("failed", orderId)
	} else if status == "cancel" || status == "expire" {
		h.TransactionRepository.UpdateTransaction("failed", orderId)
	} else if status == "pending" {
		h.TransactionRepository.UpdateTransaction("pending", orderId)
	}

	c.JSON(http.StatusOK, dtoResult.SuccessResult{
		Status: http.StatusOK,
		Data:   notificationPayload,
	})
}
