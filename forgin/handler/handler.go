package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var userData = make(map[string]Signupdata)

type Signupdata struct {
	ConfirmPassword string
	Email           string
	PhoneNumber     string
	Name            string
	Password        string
}

func HomePage(c *gin.Context) {
	c.Header("Cache-Control", "no-cache,no-store,must-revalidate")
	c.Header("Expires", "0")
	cookie, err := c.Cookie("Cookie")
	if err == nil && cookie != "" {
		c.HTML(http.StatusOK, "homepage.html", nil)
		return
	} else {
		c.Redirect(303, "/login")
	}

}
func SignupPage(c *gin.Context) {
	c.Header("Cache-Control", "no-cache,no-store,must-revalidate")
	c.Header("Expires", "0")
	c.HTML(http.StatusFound, "signupPage.html", nil)

}

func SignupPost(c *gin.Context) {
	c.Header("Cache-Control", "no-cache,no-store,must-revalidate")
	c.Header("Expires", "0")
	firstname := c.Request.FormValue("firstname")
	password := c.Request.FormValue("password")
	phonenumber := c.Request.FormValue("phonenumber")
	confirmpassword := c.Request.FormValue("confirmpassword")
	email := c.Request.FormValue("email")
	if firstname == "" {

		c.HTML(http.StatusUnauthorized, "signupPage.html", "invalid entry")
		return
	}

	if email == "" {

		c.HTML(http.StatusUnauthorized, "signupPage.html", "invalid entry")

		return
	}
	if password == "" {

		c.HTML(http.StatusUnauthorized, "signupPage.html", "invalid entry")

		return
	}
	if phonenumber == "" {

		c.HTML(http.StatusUnauthorized, "signupPage.html", "invalid entry")

		return
	}
	if confirmpassword != password {

		c.HTML(http.StatusUnauthorized, "signupPage.html", "invalid entry")

		return
	}
	userData[email] = Signupdata{
		Email:       email,
		Password:    password,
		Name:        firstname,
		PhoneNumber: phonenumber,
	}
	fmt.Print(userData)

	c.Redirect(303, "/login")

}
func LoginPage(c *gin.Context) {
	c.Header("Cache-Control", "no-cache,no-store,must-revalidate")
	c.Header("Expires", "0")
	cookie, err := c.Cookie("Cookie")
	if err == nil && cookie != "" {
		c.HTML(http.StatusOK, "homepage.html", nil)
		return
	}
	c.HTML(http.StatusOK, "loginPage.html", nil)
}
func Postmethod(c *gin.Context) {
	c.Header("Cache-Control", "no-cache,no-store,must-revalidate")
	c.Header("Expires", "0")
	email := c.Request.FormValue("emailLogin")
	password := c.Request.FormValue("passwordLogin")
	SignupData, ok := userData[email]
	if email == "" {
		c.HTML(http.StatusUnauthorized, "loginPage.html", "invalid entry")
		fmt.Println("email is not given")
		return
	} else if password == "" {
		c.HTML(http.StatusUnauthorized, "loginPage.html", "invalid entry")
		fmt.Println("password is not given")
		return
	}
	if ok && password == SignupData.Password {
		c.SetSameSite(http.SameSiteLaxMode)
		c.SetCookie("Cookie", email, 3600, "", "", false, true)
		c.Redirect(302, "/")

	} else {
		c.HTML(http.StatusOK, "loginPage.html", nil)
		fmt.Println("invalid credentials")
		return
	}
	c.HTML(http.StatusOK, "homepage.html", nil)
}
func Logout(c *gin.Context) {
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Cookie", "", 0, "", "", true, true)

	c.Redirect(303, "/")

}
