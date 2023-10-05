package controllers

import (
	"api/api/entities"
	"api/database"
	"net/http"

	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	users []entities.User
}

func NewUserController() *UserController {
	// Initialize an empty UserController
	userController := &UserController{
		users: []entities.User{},
	}

	// Fetch all users from the database and store them in the userController
	allUsers, err := database.GetAllUsers() // Implement this function in the database package
	if err == nil {
		userController.users = allUsers
	}

	return userController
}

func CreateUser(c *gin.Context, userController *UserController) {
	var user entities.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create a new User instance and save it to the database
	newUser := entities.NewUser(user.Name, user.Surname, user.Endereco, user.NumeroCelular, nil)

	if err := database.CreateUser(newUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	userController.users = append(userController.users, *newUser)

	c.JSON(http.StatusCreated, newUser)
}

func GetUser(c *gin.Context) {
	userID := c.Param("id")
	if _, err := uuid.Parse(userID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Retrieve the User from the database by ID
	user, err := database.GetUser(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func UpdateUser(c *gin.Context, userController *UserController) {
	userID := c.Param("id")
	if _, err := uuid.Parse(userID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var user entities.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call the UpdateUser function with a pointer to the user variable
	updatedUser, err := database.UpdateUser(userID, &user)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Update the UserController's users slice with the updated user
	for i, u := range userController.users {
		if u.ID == updatedUser.ID {
			userController.users[i] = *updatedUser
			break
		}
	}

	c.JSON(http.StatusOK, updatedUser)
}

func DeleteUser(c *gin.Context, userController *UserController) {
	userID := c.Param("id")
	if _, err := uuid.Parse(userID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Delete the User from the database
	if err := database.DeleteUser(userID); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Remove the deleted user from the UserController's users slice
	for i, user := range userController.users {
		if user.ID == userID {
			userController.users = append(userController.users[:i], userController.users[i+1:]...)
			break
		}
	}

	c.Status(http.StatusNoContent)
}

func (u *UserController) FindAll(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, u.users)
}
