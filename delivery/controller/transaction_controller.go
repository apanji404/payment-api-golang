package controller

import (
	"mnc/model"
	"mnc/usecase"
	"mnc/utils/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransactionController struct {
	txUC   usecase.TransactionUseCase
	router *gin.Engine
}

func (t *TransactionController) createHandler(c *gin.Context) {
	var transaction model.Transaction
	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(400, gin.H{"err": err.Error()})
		return
	}
	transaction.ID = common.GenerateID()
	if err := t.txUC.CreateTransaction(transaction); err != nil {
		c.JSON(500, gin.H{"err": err.Error()})
		return
	}
	userResponse := map[string]any{
		"id":             transaction.ID,
		"customer_id":    transaction.Customer_ID,
		"merchant_id":    transaction.Merchant_ID,
		"bank_id":        transaction.Bank_ID,
		"payment_amount": transaction.Payment_Amount,
		"payment_time":   transaction.Payment_Time,
	}

	c.JSON(http.StatusOK, userResponse)
}

func (t *TransactionController) listHandler(c *gin.Context) {
	transactions, err := t.txUC.FindAllTransaction()
	if err != nil {
		c.JSON(500, gin.H{"err": err.Error()})
		return
	}
	status := map[string]any{
		"code":        200,
		"description": "Get All Data Successfully",
	}
	c.JSON(200, gin.H{
		"status": status,
		"data":   transactions,
	})
}

func (t *TransactionController) getHandler(c *gin.Context) {
	id := c.Param("id")
	transaction, err := t.txUC.FindByTransactionId(id)
	if err != nil {
		c.JSON(500, gin.H{"err": err.Error()})
		return
	}
	status := map[string]any{
		"code":        200,
		"description": "Get By Id Data Successfully",
	}
	c.JSON(200, gin.H{
		"status": status,
		"data":   transaction,
	})
}

func NewTransactionController(usecase usecase.TransactionUseCase, r *gin.Engine) *TransactionController {
	controller := TransactionController{
		router: r,
		txUC:   usecase,
	}
	rg := r.Group("/api")
	rg.POST("/transaction", controller.createHandler)
	rg.GET("/transaction", controller.listHandler)
	rg.GET("/transaction/:id", controller.getHandler)
	return &controller
}
