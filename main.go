package main

import (
	"first_app/controller"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ping(c *gin.Context) error {
	c.JSON(200, gin.H{
		"message": "pong",
	})

	return nil
}

func main() {
	r := gin.Default()

	r.NoMethod(HandleNotFound)
	r.NoRoute(HandleNotFound)

	api := r.Group("/api")

	v1 := api.Group("v1")

	v1.GET("/ping", wrapper(ping))

	v1.GET("/accounts", func(c *gin.Context) {
		a := controller.GetAccounts()
		c.JSON(200, a)
	})

	v1.POST("/accounts", wrapper(func(c *gin.Context) error {
		var account controller.Account

		if c.BindJSON(&account) != nil {
			return UnknownError("")
		}

		a, err := controller.CreateAccount(account)

		if err != nil {
			return ParameterError(err.Error())
		}

		c.JSON(200, a)

		return nil
	}))

	v1.PATCH("/accounts/:account_id", func(c *gin.Context) {
		var account controller.Account

		if err := c.BindJSON(&account); err != nil {
			c.String(400, err.Error())
		}

		account_id, err := strconv.Atoi(c.Param("account_id"))

		if err != nil {
			c.String(400, "")
		}

		a, err := controller.UpdateAccount(account, account_id)

		if err != nil {
			c.String(400, err.Error())
		}

		c.JSON(200, a)

		return
	})

	v1.DELETE("/accounts/:account_id", func(c *gin.Context) {
		account_id, err := strconv.Atoi(c.Param("account_id"))

		if err != nil {
			c.String(400, "")
		}

		err = controller.DeleteAccount(account_id)

		if err != nil {
			c.String(400, err.Error())
		}

		c.String(204, "")

		return
	})
	r.Run() // 預設監聽本地8080埠，如果需要更改可以使用 r.Run(":9000")
}
