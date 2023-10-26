package controller

import (
	"mnc/model"
	"mnc/usecase"
	"mnc/utils/common"

	"github.com/gin-gonic/gin"
)

type BankController struct {
	bankUC usecase.BankUseCase
	router *gin.Engine
}

func (b *BankController) createHandler(c *gin.Context) {
	var bank model.Bank
	if err := c.ShouldBindJSON(&bank); err != nil {
		c.JSON(400, gin.H{"err": err.Error()})
		return
	}
	bank.ID = common.GenerateID()
	if err := b.bankUC.RegisterNewBank(bank); err != nil {
		c.JSON(500, gin.H{"err": err.Error()})
		return
	}
	c.JSON(201, bank)
}

func (b *BankController) listHandler(c *gin.Context) {
	banks, err := b.bankUC.FindAllBank()
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
		"data":   banks,
	})
}

func (b *BankController) getHandler(c *gin.Context) {
	id := c.Param("id")
	bank, err := b.bankUC.FindByBankId(id)
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
		"data":   bank,
	})
}

func (b *BankController) updateHandler(c *gin.Context) {
	var bank model.Bank
	if err := c.ShouldBindJSON(&bank); err != nil {
		c.JSON(400, gin.H{"err": err.Error()})
		return
	}
	if err := b.bankUC.UpdateBank(bank); err != nil {
		c.JSON(500, gin.H{"err": err.Error()})
		return
	}
	c.JSON(200, bank)
}

func (b *BankController) deleteHandler(c *gin.Context) {
	id := c.Param("id")
	if err := b.bankUC.DeleteBank(id); err != nil {
		c.JSON(500, gin.H{"err": err.Error()})
		return
	}
	c.String(204, "")
}

func NewBankController(usecase usecase.BankUseCase, r *gin.Engine) *BankController {
	controller := BankController{
		router: r,
		bankUC: usecase,
	}
	rg := r.Group("/api")
	rg.POST("/bank", controller.createHandler)
	rg.GET("/bank", controller.listHandler)
	rg.GET("/bank/:id", controller.getHandler)
	rg.PUT("/bank", controller.updateHandler)
	rg.DELETE("/bank/:id", controller.deleteHandler)
	return &controller
}
