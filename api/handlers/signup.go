package handlers

import (
    "context"
    "net/http"
    "go-mongo-api/config"
    "go-mongo-api/models"
    "golang.org/x/crypto/bcrypt"
    "github.com/gin-gonic/gin"
)

// Signup handles the signup endpoint
func Signup(c *gin.Context) {
    var newUser models.User

    // Bind request body to the User model
    if err := c.BindJSON(&newUser); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
        return
    }

    // Hash the user's password
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing password"})
        return
    }

    // Store the hashed password in the user model
    newUser.Password = string(hashedPassword)

    // Get MongoDB client
    mongoClient, err := config.GetMongoDBClient()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error connecting to MongoDB"})
        return
    }
    defer mongoClient.Disconnect(context.Background())

    // Get "users" collection
    usersCollection := config.GetUsersCollection(mongoClient)

    // Insert the new user into the "users" collection
    _, err = usersCollection.InsertOne(context.Background(), newUser)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error inserting user data into MongoDB"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}
