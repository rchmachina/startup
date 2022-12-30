package handler

import (
	Models "campaign/Model/campaign"
	models_user "campaign/Model/user"
	"fmt"
	"strconv"

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

func (h *campaignHandler) FindCampaignsAll (c *gin.Context) {

	FindCampaigns, err := h.campaignService.FindCampaignsByUserId(0)
	if err != nil {
		errorsMessage := gin.H{"errors": err}
		c.JSON(http.StatusBadRequest, helper.APIResponse("fail", http.StatusBadRequest, "not updated", errorsMessage))
		return
	}
	
	formatCampaign := Models.FormatCompaignsDetail(FindCampaigns)
	
	response := helper.APIResponse("the data", http.StatusOK, "updated", formatCampaign)
	c.JSON(http.StatusOK, response)



}
func (h *campaignHandler) DeleteCampaign (c *gin.Context) {

	id := c.Params.ByName("id")
	idcampaign,err := strconv.Atoi(id)
	if err != nil{
		c.JSON(http.StatusBadRequest, helper.APIResponse("fail", http.StatusBadRequest, "wrong input", nil))
		return
	}

	currentUser := c.MustGet("Current_User").(models_user.User)


	checkID, _ := h.campaignService.FindById(idcampaign)
	if checkID.UserID != int(currentUser.ID){
		fmt.Println(checkID.ID , currentUser.ID)
		c.JSON(http.StatusBadRequest, helper.APIResponse("unauthorize", http.StatusBadRequest, "token dan id berbeda", nil))
		return
	}



	checkdeletederr := h.campaignService.DeleteCampaign(idcampaign)
	if checkdeletederr != nil{		errorsMessage := gin.H{"errors": err}
	c.JSON(http.StatusBadRequest, helper.APIResponse("fail", http.StatusBadRequest, "data tidak ada/sudah tidak ada", errorsMessage))
	return
	}

	response := helper.APIResponse("the data already deleted", http.StatusOK, "deleted", id)
	c.JSON(http.StatusOK, response)
}

func (h *campaignHandler) FindCampaignsByUserId (c *gin.Context) {

	id := c.Params.ByName("id")
	idcampaign,err := strconv.Atoi(id)
	if err != nil{
		c.JSON(http.StatusBadRequest, helper.APIResponse("fail", http.StatusBadRequest, "wrong input", nil))
	}
	FindCampaigns,_ := h.campaignService.FindCampaignsByUserId(idcampaign)
	if FindCampaigns != nil {
		errorsMessage := gin.H{"errors": err}
		c.JSON(http.StatusBadRequest, helper.APIResponse("fail", http.StatusBadRequest, "User tidak ditemukan", errorsMessage))
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

	currentUser := c.MustGet("Current_User").(models_user.User)
	id := c.Params.ByName("id")
	iduser,err := strconv.Atoi(id)
	if err != nil{
		c.JSON(http.StatusBadRequest, helper.APIResponse("fail", http.StatusBadRequest, "wrong input", nil))
	}

	checkID, _ := h.campaignService.FindById(iduser)
	if checkID.UserID != int(currentUser.ID){
		fmt.Println(checkID.ID , currentUser.ID)
		c.JSON(http.StatusBadRequest, helper.APIResponse("unauthorize", http.StatusBadRequest, "token dan id berbeda", nil))
		return
	}

	update_data, err := h.campaignService.Updatecampaign(iduser, modelUpdate)
	if err!= nil{
		errorsMessage := gin.H{"errors": err}
		c.JSON(http.StatusBadRequest, helper.APIResponse("fail", http.StatusBadRequest, "not updated",errorsMessage ))
		return
	}
	formatCampaign := Models.FormatCompaignDetail(update_data)
	c.JSON(http.StatusAccepted, helper.APIResponse("ok", http.StatusAccepted, "ok", formatCampaign))


}