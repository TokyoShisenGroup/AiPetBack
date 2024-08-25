package router

import (
	"AiPetBack/db"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func initPetRoutes(r *gin.Engine) {
	petCRUD := db.PetCRUD{}

	// 路由组 /pets
	pets := r.Group("/pets")
	{
		pets.POST("/", func(c *gin.Context) {
			var pet db.RequestPet
			if err := c.ShouldBindJSON(&pet); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			// 创建宠物
			var CreatePet = db.Pet{
				Model:     gorm.Model{},
				PetName:   pet.PetName,
				Kind:      pet.Kind,
				Type:      pet.Type,
				Age:       pet.Age,
				Birthday:  pet.Birthday,
				Weight:    pet.Weight,
				OwnerName: pet.OwnerName,
				IsDeleted: false,
			}

			if err := petCRUD.CreateByObject(&CreatePet); err != nil {
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

		//get all pets of an owner
		pets.GET("/owner/:ownerName", func(c *gin.Context) {
			ownerName := c.Param("ownerName")
			pets, err := petCRUD.GetPetByOwner(ownerName)
			if err != nil {
				c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, pets)
		})
	}
}
