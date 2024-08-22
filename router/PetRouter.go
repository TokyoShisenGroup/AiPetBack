package router

import (
	"AiPetBack/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func initPetRoutes(r *gin.Engine) {
	petCRUD := db.PetCRUD{}

	// 路由组 /pets
	pets := r.Group("/pets")
	{
		pets.POST("/", func(c *gin.Context) {
			var pet db.Pet
			if err := c.ShouldBindJSON(&pet); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			if err := petCRUD.CreateByObject(&pet); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, pet)
		})

		pets.GET("/:name", func(c *gin.Context) {
			name := c.Param("name")
			pet, err := petCRUD.GetPetByName(name)
			if err != nil {
				c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, pet)
		})

		pets.PUT("/:name", func(c *gin.Context) {
			name := c.Param("name")
			var pet db.Pet
			if err := c.ShouldBindJSON(&pet); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			existingPet, err := petCRUD.GetPetByName(name)
			if err != nil {
				c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
				return
			}
			pet.ID = existingPet.ID
			if err := petCRUD.UpdateByObject(&pet); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, pet)
		})

		pets.DELETE("/:name", func(c *gin.Context) {
			name := c.Param("name")
			if err := petCRUD.DeletePetbyName(name); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{"message": "Pet deleted"})
		})
	}
}
