package handlers

import (
	dtoResult "hollyways/dto/result"
	dtoTransaction "hollyways/dto/transaction"
	"hollyways/models"
	"hollyways/packages/middleware"
	"hollyways/repositories"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

type transactionHandler struct {
	TransactionRepository repositories.TransactionRepository
}

func HandlerTransaction(TransactionRepository repositories.TransactionRepository) *transactionHandler {
	return &transactionHandler{TransactionRepository}
}

func (h *transactionHandler) CreateTransaction(c *gin.Context) {
	var err error
	var transactionIsMatch = false
	var transactionID int

	userLogin, _ := c.Get("userLogin")
	if userLogin == "" {
		c.JSON(http.StatusUnauthorized, dtoResult.ErrorResult{
			Status:  http.StatusUnauthorized,
			Message: "Unauthorized",
		})
		return
	}
	userId := int(userLogin.(jwt.MapClaims)["id"].(float64))

	request := new(dtoTransaction.TransactionRequestDTO)
	if err := c.ShouldBind(request); err != nil {
		c.JSON(http.StatusBadRequest, dtoResult.ErrorResult{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	validation := validator.New()
	if err := validation.Struct(request); err != nil {
		c.JSON(http.StatusBadRequest, dtoResult.ErrorResult{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	for !transactionIsMatch {
		transactionID = int(time.Now().Unix())
		transactionData, _ := h.TransactionRepository.GetTransaction(transactionID)
		if transactionData.ID == 0 {
			transactionIsMatch = true
		}
	}

	transaction := models.Transaction{
		Model:     gorm.Model{ID: uint(transactionID)},
		UserID:    userId,
		ProjectID: request.ProjectID,
		Donation:  request.Donation,
	}

	err = h.TransactionRepository.CreateTransaction(transaction)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtoResult.ErrorResult{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Transaction has been created, Please the payment promptly.",
	})
}

func (h *transactionHandler) FindTransaction(c *gin.Context) {
	var err error
	var transactionDTO []dtoTransaction.TransactionResponseDTO

	authorized := middleware.Authorized(c)
	if authorized != "authorize" {
		c.JSON(http.StatusUnauthorized, dtoResult.ErrorResult{
			Status:  http.StatusUnauthorized,
			Message: authorized,
		})
		return
	}

	transactions, err := h.TransactionRepository.FindTransaction()
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtoResult.ErrorResult{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	if len(transactions) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "data not found",
			"data":    transactions,
		})
		return
	}

	for _, transaction := range transactions {
		transactionDTO = append(transactionDTO, convertTransactionResponse(transaction))
	}

	c.JSON(http.StatusOK, dtoResult.SuccessResult{
		Status: http.StatusOK,
		Data:   transactionDTO,
	})

}

func (h *transactionHandler) GetTransaction(c *gin.Context) {
	var err error

	authorized := middleware.Authorized(c)
	if authorized != "authorize" {
		c.JSON(http.StatusUnauthorized, dtoResult.ErrorResult{
			Status:  http.StatusUnauthorized,
			Message: authorized,
		})
		return
	}

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
		if strings.Contains(err.Error(), "record not found") {
			c.JSON(http.StatusOK, dtoResult.SuccessResult{
				Status: http.StatusOK,
				Data:   transaction,
			})
			return
		}

		c.JSON(http.StatusInternalServerError, dtoResult.ErrorResult{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dtoResult.SuccessResult{
		Status: http.StatusOK,
		Data:   convertTransactionResponse(transaction),
	})
}

func convertTransactionResponse(t models.Transaction) dtoTransaction.TransactionResponseDTO {
	return dtoTransaction.TransactionResponseDTO{
		ID:       int(t.ID),
		User:     ConvertUserResponse(t.User),
		Project:  ConvertProjectResponse(t.Project),
		Donation: t.Donation,
		Status:   t.Status,
	}
}
