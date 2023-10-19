package controllers

import (
	"fmt"
	"github.com/blynn-dn/go_examples/crud/initializer"
	"github.com/blynn-dn/go_examples/crud/models"
	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	tempTable := models.TestTable{Name: "blynn"}

	result := initializer.DB.Create(&tempTable)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"message": fmt.Sprintf("error adding item to test_table: %s", result.Error),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": tempTable,
	})
}
