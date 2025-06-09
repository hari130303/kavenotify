package controller

import (
	"fmt"
	dbfunction "kavenotify/db_function"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {

	var RequestData struct {
		UserName string `json:"userid,omitempty"`
		Password string `json:"password,omitempty"`
	}
	err := ReadJson(c, &RequestData)
	if err != nil {
		ErrorJson(c, http.StatusBadRequest, err)
		return
	}

	returnData, err := dbfunction.UserCheck(RequestData.UserName, RequestData.Password)
	if err != nil {
		ErrorJson(c, http.StatusBadRequest, err)
		return
	}
	dbfunction.UserLogger(RequestData.UserName, "LOGIN")

	WriteJson(c, http.StatusOK, "Login Successful", returnData)
}

func RegisterApi(c *gin.Context) {

	var RequestData struct {
		GroupName     string  `json:"group_name"`
		GroupDesc     *string `json:"group_desc"`
		ContactNumber string  `json:"contact_number"`
		MailId        string  `json:"mail_id"`
		UserName      string  `json:"user_name"`
		Password      string  `json:"password"`
		CustomerType  string  `json:"customer_type"`
	}
	err := ReadJson(c, &RequestData)
	if err != nil {
		ErrorJson(c, http.StatusBadRequest, err)
		return
	}

	returnData, err := dbfunction.RegisterApi(RequestData.GroupName, RequestData.GroupDesc, RequestData.ContactNumber,
		RequestData.MailId, RequestData.UserName, RequestData.Password, RequestData.CustomerType)
	if err != nil {
		ErrorJson(c, http.StatusBadRequest, fmt.Errorf("internal server error"))
		return
	}
	dbfunction.UserLogger(RequestData.UserName, "REGISTER")

	WriteJson(c, http.StatusOK, "Register Successful", returnData)
}

func GetCustomerType(c *gin.Context) {

	var RequestData struct {
		SessionString string `json:"session_string"`
	}
	err := ReadJson(c, &RequestData)
	if err != nil {
		ErrorJson(c, http.StatusBadRequest, fmt.Errorf("internal server error"))
		return
	}

	UserData, err := dbfunction.SessionCheck(RequestData.SessionString)
	if err != nil {
		ErrorJson(c, http.StatusBadRequest, fmt.Errorf("internal server error"))
		return
	}

	returnData, err := dbfunction.GetCustomerType(UserData.GroupId)
	if err != nil {
		ErrorJson(c, http.StatusBadRequest, fmt.Errorf("internal server error"))
		return
	}

	WriteJson(c, http.StatusOK, "Customer types reuturned Successful", returnData)
}

func SendSingleMessage(c *gin.Context) {

	var RequestData struct {
		SessionString  string `json:"session_id"`
		PhoneNumber    string `json:"phone_number"`
		MessageContent string `json:"message_content"`
		UserName       string `json:"user_name"`
		CustomerId     int    `json:"customer_type_id"`
	}
	err := ReadJson(c, &RequestData)
	if err != nil {
		// ErrorJson(c, http.StatusBadRequest, fmt.Errorf("internal server error"))
		ErrorJson(c, http.StatusBadRequest, err)
		return
	}

	_, err = dbfunction.SessionCheck(RequestData.SessionString)
	if err != nil {
		// ErrorJson(c, http.StatusBadRequest, fmt.Errorf("internal server error"))
		fmt.Println("errro during check session")
		ErrorJson(c, http.StatusBadRequest, err)
		return
	}

	returnData, err := dbfunction.MessageSendAndSave(RequestData.PhoneNumber, RequestData.MessageContent, RequestData.CustomerId, RequestData.UserName)
	if err != nil {
		fmt.Println("errro during message send and save function")
		ErrorJson(c, http.StatusBadRequest, err)
		// ErrorJson(c, http.StatusBadRequest, fmt.Errorf("internal server error"))
		return
	}

	WriteJson(c, http.StatusOK, "Message send Successfully", returnData)
}

func GetCustomerTypeTable(c *gin.Context) {

	var RequestData struct {
		SessionString string `json:"session_string"`
		SearchTerm    string `json:"search_term"`
		Start         int    `json:"start"`
		Limit         int    `json:"limit"`
		OrderBy       int    `json:"order_by"`
	}
	err := ReadJson(c, &RequestData)
	if err != nil {
		ErrorJson(c, http.StatusBadRequest, fmt.Errorf("internal server error"))
		return
	}

	// fmt.Println("request payload : ", RequestData)

	UserData, err := dbfunction.SessionCheck(RequestData.SessionString)
	if err != nil {
		ErrorJson(c, http.StatusBadRequest, fmt.Errorf("internal server error"))
		return
	}

	returnData, err := dbfunction.GetCustomerTypeTable(UserData.GroupId, RequestData.SearchTerm, RequestData.Start, RequestData.Limit, RequestData.OrderBy)
	if err != nil {
		ErrorJson(c, http.StatusBadRequest, fmt.Errorf("internal server error : %s", err))
		return
	}

	WriteJson(c, http.StatusOK, "Customer type details reuturned Successful", returnData)
}

func NumberManageTable(c *gin.Context) {

	var RequestData struct {
		SessionString string `json:"session_string"`
		SearchTerm    string `json:"search_term"`
		Start         int    `json:"start"`
		Limit         int    `json:"limit"`
		OrderBy       int    `json:"order_by"`
	}
	err := ReadJson(c, &RequestData)
	if err != nil {
		ErrorJson(c, http.StatusBadRequest, fmt.Errorf("internal server error"))
		return
	}

	UserData, err := dbfunction.SessionCheck(RequestData.SessionString)
	if err != nil {
		ErrorJson(c, http.StatusBadRequest, fmt.Errorf("internal server error"))
		return
	}

	returnData, err := dbfunction.NumberManageTable(UserData.GroupId, RequestData.SearchTerm, RequestData.Start, RequestData.Limit, RequestData.OrderBy)
	if err != nil {
		ErrorJson(c, http.StatusBadRequest, fmt.Errorf("internal server error : %s", err))
		return
	}

	WriteJson(c, http.StatusOK, "Nunber and details reuturned Successfully", returnData)
}

func GetUserLoginLog(c *gin.Context) {

	var RequestData struct {
		SessionString string `json:"session_string"`
		SearchTerm    string `json:"search_term"`
		Start         int    `json:"start"`
		Limit         int    `json:"limit"`
		OrderBy       int    `json:"order_by"`
	}
	err := ReadJson(c, &RequestData)
	if err != nil {
		ErrorJson(c, http.StatusBadRequest, fmt.Errorf("internal server error"))
		return
	}

	UserData, err := dbfunction.SessionCheck(RequestData.SessionString)
	if err != nil {
		ErrorJson(c, http.StatusBadRequest, fmt.Errorf("internal server error"))
		return
	}

	returnData, err := dbfunction.NumberManageTable(UserData.GroupId, RequestData.SearchTerm, RequestData.Start, RequestData.Limit, RequestData.OrderBy)
	if err != nil {
		ErrorJson(c, http.StatusBadRequest, fmt.Errorf("internal server error : %s", err))
		return
	}

	WriteJson(c, http.StatusOK, "Login logs reuturned Successfully", returnData)
}
