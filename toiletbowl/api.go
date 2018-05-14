package toiletbowl

import (
	"net/http"

	"github.com/labstack/echo"
)

func APICreatePoo(c echo.Context) error {
	username := c.FormValue("username")
	content := c.FormValue("content")

	db := GetDB()
	var user User
	db.FirstOrCreate(&user, User{Username: username})
	poo := Poo{Content: content, UserId: user.ID}
	db.Save(&poo)
	var poos []Poo
	db.Model(&user).Association("Poos").Find(&poos)

	db.Save(&user)

	return c.JSON(http.StatusOK, map[string]interface{}{"data": map[string]interface{}{
		"poos":    poos,
		"content": content,
	}, "error": false})
}

func APIDeletePoo(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{"data": nil, "error": false})
}

func APIGetPoos(c echo.Context) error {
	db := GetDB()

	username := c.FormValue("username")
	var user User
	var poos []Poo
	c.Logger().Print(username)
	if username == "" {
		db.Find(&poos)
		c.Logger().Print("Finding all poos")

		return c.JSON(http.StatusOK, map[string]interface{}{"data": poos, "error": false})
	}

	if err := db.Preload("Poos").First(&user, User{Username: username}).Error; err != nil {
		return c.JSON(http.StatusOK, map[string]interface{}{"data": nil, "error": err})
	}
	c.Logger().Print("finding poos for ", user)

	return c.JSON(http.StatusOK, map[string]interface{}{"data": user.Poos, "error": false})
}

func APIUpdatePoo(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{"data": nil, "error": false})
}

func APICreateUser(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{"data": nil, "error": false})
}

func APIDeleteUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{"data": nil, "error": false})
}

func APIGetUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{"data": nil, "error": false})
}

func APIUpdateUser(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{"data": nil, "error": false})
}
