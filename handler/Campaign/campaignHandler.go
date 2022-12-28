package handler

import (
	Models "campaign/Model/campaign"
	models_user "campaign/Model/user"
	"fmt"

	"campaign/helper"

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

func (h *campaignHandler) FindCampaignsByUserId (c *gin.Context) {
	var input Models.GetCampaignByID


	err :=c.ShouldBindUri(&input)
	if err !=nil{
		errorsMessage := gin.H{"errors": err}
		c.JSON(http.StatusBadRequest, helper.APIResponse("fail", http.StatusBadRequest, "not updated",errorsMessage ))
		return
	}
	

	FindCampaigns, err := h.campaignService.FindCampaignsByUserId(input.ID)
	if err != nil {
		errorsMessage := gin.H{"errors": err}
		c.JSON(http.StatusBadRequest, helper.APIResponse("fail", http.StatusBadRequest, "not updated", errorsMessage))
		return
	}
	
	formatCampaign := Models.FormatCompaignsDetail(FindCampaigns)
	
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

		c.JSON(http.StatusBadRequest, helper.APIResponse("fail", http.StatusBadRequest, "not updated", nil))
		return
	}

	response := helper.APIResponse("campaign created", http.StatusOK, "updated", Models.FormatCompaign(inputdata))
	c.JSON(http.StatusOK, response)
}





func (h *campaignHandler) Updatecampaign(c *gin.Context) {
	var modelUpdate Models.CreateCampaign

	err := c.Bind(&modelUpdate)
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


	var input Models.GetCampaignByID

	err =c.ShouldBindUri(&input)
	fmt.Println(input.ID)
	currentUser := c.MustGet("Current_User").(models_user.User)
	if err !=nil{
		errorsMessage := gin.H{"errors": err}
		c.JSON(http.StatusBadRequest, helper.APIResponse("fail", http.StatusBadRequest, "not updated",errorsMessage ))
		return
	}
	

	checkID, _ := h.campaignService.FindById(input)
	if checkID.UserID != int(currentUser.ID){
		fmt.Println(checkID.ID , currentUser.ID)
		c.JSON(http.StatusBadRequest, helper.APIResponse("unauthorize", http.StatusBadRequest, "token dan id berbeda", nil))
		return
	}

	update_data, err := h.campaignService.Updatecampaign(input, modelUpdate)
	if err!= nil{
		errorsMessage := gin.H{"errors": err}
		c.JSON(http.StatusBadRequest, helper.APIResponse("fail", http.StatusBadRequest, "not updated",errorsMessage ))
		return
	}
	formatCampaign := Models.FormatCompaignDetail(update_data)
	c.JSON(http.StatusAccepted, helper.APIResponse("ok", http.StatusAccepted, "ok", formatCampaign))
	return


	// response := helper.APIResponse("campaign updated", http.StatusOK, "updated", Models.FormatCompaign(updatedata))
	// c.JSON(http.StatusOK, response)
}