package handler

import (
	Models "campaign/Model/campaign"
	"campaign/helper"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type campaignHandler struct {
	campaignService Models.Service
	//authService auth.Service
}

func NewCampaignHandler(campaignService Models.Service) *campaignHandler {
	return &campaignHandler{campaignService}
}

func (h *campaignHandler) FindCampaigns (c *gin.Context) {
	var input Models.Campaign
	err := c.Bind(&input.UserID)
	if err != nil {
		var errors []string

		for _, e := range err.(validator.ValidationErrors) {
			errors = append(errors, e.Error())

		}

		errorsMessage := gin.H{"errors": errors}

		response := helper.APIResponse("failed to get data", http.StatusUnprocessableEntity, "error", errorsMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return

	}
	
	FindCampaigns, err := h.campaignService.FindCampaigns(int(input.UserID))
	if err != nil {
		errorsMessage := gin.H{"errors": err}
		c.JSON(http.StatusBadRequest, helper.APIResponse("fail", http.StatusBadRequest, "not updated", errorsMessage))
		return
	}
	
	formatCampaign := Models.FormatCompaigns(FindCampaigns)
	
	response := helper.APIResponse("the data", http.StatusOK, "updated", formatCampaign)
	c.JSON(http.StatusOK, response)
}
