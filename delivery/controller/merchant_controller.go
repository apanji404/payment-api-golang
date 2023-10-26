package controller

import (
	"mnc/model"
	"mnc/usecase"
	"mnc/utils/common"

	"github.com/gin-gonic/gin"
)

type MerchantController struct {
	merchantUC usecase.MerchantUseCase
	router     *gin.Engine
}

func (b *MerchantController) createHandler(c *gin.Context) {
	var merchant model.Merchant
	if err := c.ShouldBindJSON(&merchant); err != nil {
		c.JSON(400, gin.H{"err": err.Error()})
		return
	}
	merchant.ID = common.GenerateID()
	if err := b.merchantUC.RegisterNewMerchant(merchant); err != nil {
		c.JSON(500, gin.H{"err": err.Error()})
		return
	}
	c.JSON(201, merchant)
}

func (b *MerchantController) listHandler(c *gin.Context) {
	merchants, err := b.merchantUC.FindAllMerchant()
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
		"data":   merchants,
	})
}

func (b *MerchantController) getHandler(c *gin.Context) {
	id := c.Param("id")
	merchant, err := b.merchantUC.FindByMerchantId(id)
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
		"data":   merchant,
	})
}

func (b *MerchantController) updateHandler(c *gin.Context) {
	var merchant model.Merchant
	if err := c.ShouldBindJSON(&merchant); err != nil {
		c.JSON(400, gin.H{"err": err.Error()})
		return
	}
	if err := b.merchantUC.UpdateMerchant(merchant); err != nil {
		c.JSON(500, gin.H{"err": err.Error()})
		return
	}
	c.JSON(200, merchant)
}

func (b *MerchantController) deleteHandler(c *gin.Context) {
	id := c.Param("id")
	if err := b.merchantUC.DeleteMerchant(id); err != nil {
		c.JSON(500, gin.H{"err": err.Error()})
		return
	}
	c.String(204, "")
}

func NewMerchantController(usecase usecase.MerchantUseCase, r *gin.Engine) *MerchantController {
	controller := MerchantController{
		router:     r,
		merchantUC: usecase,
	}

	rg := r.Group("/api")
	rg.POST("/merchant", controller.createHandler)
	rg.GET("/merchant", controller.listHandler)
	rg.GET("/merchant/:id", controller.getHandler)
	rg.PUT("/merchant", controller.updateHandler)
	rg.DELETE("/merchant/:id", controller.deleteHandler)
	return &controller
}
