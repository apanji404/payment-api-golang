package controller

import (
	"mnc/model"
	"mnc/usecase"
	"mnc/utils/common"

	"github.com/gin-gonic/gin"
)

type CustomerController struct {
	customerUC usecase.CustomerUseCase
	router     *gin.Engine
}

func (b *CustomerController) createHandler(c *gin.Context) {
	var customer model.Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(400, gin.H{"err": err.Error()})
		return
	}
	customer.ID = common.GenerateID()
	if err := b.customerUC.RegisterNewCustomer(customer); err != nil {
		c.JSON(500, gin.H{"err": err.Error()})
		return
	}
	c.JSON(201, customer)
}

func (b *CustomerController) listHandler(c *gin.Context) {
	customers, err := b.customerUC.FindAllCustomer()
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
		"data":   customers,
	})
}

func (b *CustomerController) getHandler(c *gin.Context) {
	id := c.Param("id")
	customer, err := b.customerUC.FindByCustomerId(id)
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
		"data":   customer,
	})
}

func (b *CustomerController) updateHandler(c *gin.Context) {
	var customer model.Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(400, gin.H{"err": err.Error()})
		return
	}
	if err := b.customerUC.UpdateCustomer(customer); err != nil {
		c.JSON(500, gin.H{"err": err.Error()})
		return
	}
	c.JSON(200, customer)
}

func (b *CustomerController) deleteHandler(c *gin.Context) {
	id := c.Param("id")
	if err := b.customerUC.DeleteCustomer(id); err != nil {
		c.JSON(500, gin.H{"err": err.Error()})
		return
	}
	c.String(204, "")
}

func NewCustomerController(usecase usecase.CustomerUseCase, r *gin.Engine) *CustomerController {
	controller := CustomerController{
		router:     r,
		customerUC: usecase,
	}
	rg := r.Group("/api")
	rg.POST("/customer", controller.createHandler)
	rg.GET("/customer", controller.listHandler)
	rg.GET("/customer/:id", controller.getHandler)
	rg.PUT("/customer", controller.updateHandler)
	rg.DELETE("/customer/:id", controller.deleteHandler)
	return &controller
}
