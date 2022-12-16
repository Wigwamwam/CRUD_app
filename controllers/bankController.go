package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/wigwamwam/CRUD_app/initializers"
	"github.com/wigwamwam/CRUD_app/models"
)

func BanksCreate(c *gin.Context) {

	var body struct {
		Name string
		IBAN string
	}

	c.Bind(&body)

	bank := models.Bank{Name: body.Name, IBAN: body.IBAN}
	result := initializers.DB.Create(&bank)
	// what happens to this - why are we not passing this into the json?

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"bank": bank,
	})
}

func BanksIndex(c *gin.Context) {
	var banks []models.Bank
	initializers.DB.Find(&banks)

	c.JSON(200, gin.H{
		"banks": banks,
	})
}

func BanksShow(c *gin.Context) {
	id := c.Param("id")

	var bank []models.Bank
	initializers.DB.Find(&bank, id)

	c.JSON(200, gin.H{
		"bank": bank,
	})
}

func BanksUpdate(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Name string
		IBAN string
	}

	c.Bind(&body)

	var bank []models.Bank
	initializers.DB.Find(&bank, id)

	initializers.DB.Model(&bank).Updates(models.Bank{
		Name: body.Name,
		IBAN: body.IBAN,
	})

	c.JSON(200, gin.H{
		"bank": bank,
	})

}

func BanksDelete(c *gin.Context) {
	id := c.Param("id")

	initializers.DB.Delete(&models.Bank{}, id)
	// why is the delete and find method written differently?

	c.Status(200)

}
