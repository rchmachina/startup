package handler

import (
	Models "campaign/Model/campaign"
	models_user "campaign/Model/user"

	"campaign/helper"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)



type campaignHandler struct {
	campaignService Models.Service

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


func (h *campaignHandler) Createcampaign(c *gin.Context) {
	var input Models.CreateCampaign
	err := c.Bind(&input)
	if err != nil {
		var errors []string

		for _, e := range err.(validator.ValidationErrors) {
			errors = append(errors, e.Error())

		}

		errorsMessage := gin.H{"errors": errors}

		response := helper.APIResponse("create campaign failed", http.StatusUnprocessableEntity, "error", errorsMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return 

	}

	currentUser := c.MustGet("Current_User").(models_user.User)
	input.User = currentUser
	inputdata, err := h.campaignService.CreateCampaign(input)
	if err != nil {
		//errorsMessage := gin.H{"errors": err.Error()}
		//
		c.JSON(http.StatusBadRequest, helper.APIResponse("fail", http.StatusBadRequest, "not updated", nil))
		return
	}

	response := helper.APIResponse("campaign created", http.StatusOK, "updated", Models.FormatCompaign(inputdata))
	c.JSON(http.StatusOK, response)
}