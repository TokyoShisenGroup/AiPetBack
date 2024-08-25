package router

import (
	"net/http"
	"strconv"

	"AiPetBack/db"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var userCRUD = db.UserCRUD{}

func RegisterUserRoutes(r *gin.Engine) {
	r.POST("/users", createUser)
	r.GET("/users/:name", getUserByName)
	r.GET("/users", getAllUsers)
	r.PUT("/users/:name", updateUser)
	r.DELETE("/users/:name", deleteUserByName)
	r.GET("/users/location", getUsersByLocation)
	r.POST("/users/register", Register)
	r.POST("/users/login", Login)
}

func createUser(c *gin.Context) {
	var user db.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := userCRUD.CreateByObject(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func getUserByName(c *gin.Context) {
	name := c.Param("name")

	user, err := userCRUD.GetUserByName(name)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, user)
}

func getAllUsers(c *gin.Context) {
	users, err := userCRUD.GetAllUser()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

func updateUser(c *gin.Context) {
	name := c.Param("name")

	var user db.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.UserName = name
	if err := userCRUD.UpdateByObject(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func deleteUserByName(c *gin.Context) {
	name := c.Param("name")

	if err := userCRUD.DeleteUserbyName(name); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

func getUsersByLocation(c *gin.Context) {
	lat, err := strconv.ParseFloat(c.Query("lat"), 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid latitude"})
		return
	}

	long, err := strconv.ParseFloat(c.Query("long"), 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid longitude"})
		return
	}

	radius, err := strconv.ParseFloat(c.Query("radius"), 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid radius"})
		return
	}

	users, err := userCRUD.GetUserByLocation(lat, long, radius)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

func Register(c *gin.Context) {
	crud := &db.UserCRUD{}
	var Register db.UserLogin
	if err := c.ShouldBindJSON(&Register); err != nil {
		c.JSON(500, gin.H{"err": err.Error()})
		return
	}
	newuser := db.User{
		UserName: Register.UserName,
		PassWord: Register.Password,
	}

	err := crud.CreateByObject(&newuser)

	if err != nil {
		c.JSON(500, gin.H{"error": "Register failed!"})
		return
	}

	c.JSON(200, gin.H{"message": "Registration successful"})
}

func Login(c *gin.Context) {
	crud := &db.UserCRUD{}
	var Login db.UserLogin
	if err := c.ShouldBindJSON(&Login); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	user, err := crud.GetUserByName(Login.UserName)
	if err != nil {
		c.JSON(404, gin.H{"error": "Invalid username or password"})
		return
	}

	if user.PassWord != Login.Password {
		c.JSON(404, gin.H{"error": "Invalid username or password"})
		return
	}

	err = crud.UpdateByObject(user)
	if err != nil {
		c.JSON(500, gin.H{"error": "Database error"})
	}

	c.JSON(200, gin.H{"message": "Login successful"})
}
