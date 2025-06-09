package controller

import (
	"embed"
	"html/template"
	"log"

	"github.com/gin-gonic/gin"
)

//go:embed templates
var Templates embed.FS

func LoginTemplate(c *gin.Context) {
	// Note the call to ParseFS instead of Parse
	t, err := template.ParseFS(Templates, "templates/login.html")
	if err != nil {
		log.Fatal(err)
	}
	// c.HTML(http.StatusOK, string(t), nil)
	t.Execute(c.Writer, nil)
}

func RegisterTemplate(c *gin.Context) {
	// Note the call to ParseFS instead of Parse
	t, err := template.ParseFS(Templates, "templates/register_form.html")
	if err != nil {
		log.Fatal(err)
	}
	// c.HTML(http.StatusOK, string(t), nil)
	t.Execute(c.Writer, nil)
}

func SingleMessage(c *gin.Context) {
	// Note the call to ParseFS instead of Parse
	t, err := template.ParseFS(Templates, "templates/dashboard_single_message.html")
	if err != nil {
		log.Fatal(err)
	}
	// c.HTML(http.StatusOK, string(t), nil)
	t.Execute(c.Writer, nil)
}

func BulkMessage(c *gin.Context) {
	// Note the call to ParseFS instead of Parse
	t, err := template.ParseFS(Templates, "templates/dashboard_bulk_message.html")
	if err != nil {
		log.Fatal(err)
	}
	// c.HTML(http.StatusOK, string(t), nil)
	t.Execute(c.Writer, nil)
}

func CustomerSetup(c *gin.Context) {
	// Note the call to ParseFS instead of Parse
	t, err := template.ParseFS(Templates, "templates/dashboard_cutomer_setup.html")
	if err != nil {
		log.Fatal(err)
	}
	// c.HTML(http.StatusOK, string(t), nil)
	t.Execute(c.Writer, nil)
}

func NumberManage(c *gin.Context) {
	// Note the call to ParseFS instead of Parse
	t, err := template.ParseFS(Templates, "templates/dashboard_number_manage.html")
	if err != nil {
		log.Fatal(err)
	}
	// c.HTML(http.StatusOK, string(t), nil)
	t.Execute(c.Writer, nil)
}

func UserLoginLog(c *gin.Context) {
	// Note the call to ParseFS instead of Parse
	t, err := template.ParseFS(Templates, "templates/dashboard_user_login_log.html")
	if err != nil {
		log.Fatal(err)
	}
	// c.HTML(http.StatusOK, string(t), nil)
	t.Execute(c.Writer, nil)
}
