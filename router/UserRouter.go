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