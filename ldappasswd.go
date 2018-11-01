package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	ldap "gopkg.in/ldap.v2"
)

var (
	LDAP_SERVER = os.Getenv("LDAP_SERVER")
	LDAP_PORT   = os.Getenv("LDAP_PORT")
)

type Account struct {
	DN        string `form:"dn" json:"dn" binding:"required"`
	OldPasswd string `form:"oldpasswd" json:"oldpasswd" binding:"required"`
	NewPasswd string `form:"newpasswd" json:"newpasswd" binding:"required"`
}

func main() {
	router := gin.Default()
	router.LoadHTMLFiles("index.html")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
	router.POST("/modifypwd", modifypwd())
	router.Run(":8389")
}

func modifypwd() gin.HandlerFunc {
	return func(c *gin.Context) {
		var account Account
		if err := c.ShouldBind(&account); err != nil {
			c.HTML(http.StatusOK, "index.html", gin.H{"message": err.Error()})
			return
		}
		err := account.passwordModify()
		if err != nil {
			c.HTML(http.StatusOK, "index.html", gin.H{"message": err.Error()})
			return
		}
		c.HTML(http.StatusOK, "index.html", gin.H{"message": "Password has been modified."})
	}
}

func (a *Account) passwordModify() error {
	l, err := ldap.Dial("tcp", fmt.Sprintf("%v:%v", LDAP_SERVER, LDAP_PORT))
	if err != nil {
		return err
	}
	defer l.Close()

	err = l.Bind(a.DN, a.OldPasswd)
	if err != nil {
		return err
	}

	passwordModifyRequest := ldap.NewPasswordModifyRequest("", a.OldPasswd, a.NewPasswd)
	_, err = l.PasswordModify(passwordModifyRequest)

	if err != nil {
		return err
	}
	return nil
}
