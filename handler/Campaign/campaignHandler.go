package handler

import (
	Models "campaign/Model/campaign"
	"campaign/helper"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"

)

type campaignHandler struct {
	campaignService Models.Service
	//authService auth.Service
}

func NewCampaignHandler(campaignService Models.Service) *campaignHandler {
	return &campaignHandler{campaignService}
}

func (h *campaignHandler) FindCampaignsByuserId (c *gin.Context) {
	UserID, _ :=strconv.Atoi(c.Query("userid"))

	
	FindCampaigns, err := h.campaignService.FindCampaignsByuserId(UserID)
	if err != nil {
		errorsMessage := gin.H{"errors": err}
		c.JSON(http.StatusBadRequest, helper.APIResponse("fail", http.StatusBadRequest, "not updated", errorsMessage))
		return
	}
	
	formatCampaign := Models.FormatCompaigns(FindCampaigns)
	
	response := helper.APIResponse("the data", http.StatusOK, "updated", formatCampaign)
	c.JSON(http.StatusOK, response)
}


func (h *campaignHandler) FindCampaignByid (c *gin.Context) {
	var input Models.GetCampaignByID
	err := c.ShouldBindUri(&input)
	if err != nil {
		errorsMessage := gin.H{"errors": err}
		c.JSON(http.StatusBadRequest, helper.APIResponse("fail", http.StatusBadRequest, "There something wrong ", errorsMessage))
		return
	}

	FindCampaigns, err := h.campaignService.FindCampaignByid(input.ID)
	if err != nil {
		errorsMessage := gin.H{"errors": err}
		c.JSON(http.StatusBadRequest, helper.APIResponse("fail", http.StatusBadRequest, "not updated", errorsMessage))
		return
	}
	
	formatCampaign := Models.FormatCompaignDetail(FindCampaigns)
	
	response := helper.APIResponse("the data", http.StatusOK, "updated", formatCampaign)
	c.JSON(http.StatusOK, response)

	
}
