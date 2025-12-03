package handlers

import (
	"time"

	"go-backend/database"
	"go-backend/models"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// GetSubscription returns the subscription status of the authenticated user
func GetSubscription(c *fiber.Ctx) error {
	// Get token from header (middleware should have handled validation, but we check again or extract claims)
	// Assuming AuthMiddleware sets the user ID or we extract it here like in Me handler
	tokenString := c.Get("Authorization")
	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:]
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthenticated",
		})
	}

	claims := token.Claims.(jwt.MapClaims)
	var user models.User
	if err := database.DB.Where("id = ?", claims["sub"]).First(&user).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	return c.JSON(fiber.Map{
		"subscription_status":     user.SubscriptionStatus,
		"subscription_plan":       user.SubscriptionPlan,
		"subscription_expires_at": user.SubscriptionExpiresAt,
	})
}

// Subscribe handles the subscription process (Mock implementation)
func Subscribe(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	plan := data["plan"] // "monthly" or "yearly"
	if plan != "monthly" && plan != "yearly" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid plan",
		})
	}

	tokenString := c.Get("Authorization")
	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:]
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthenticated",
		})
	}

	claims := token.Claims.(jwt.MapClaims)
	var user models.User
	if err := database.DB.Where("id = ?", claims["sub"]).First(&user).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	// Mock Payment Logic: Immediately activate subscription
	expiresAt := time.Now()
	if plan == "monthly" {
		expiresAt = expiresAt.AddDate(0, 1, 0)
	} else {
		expiresAt = expiresAt.AddDate(1, 0, 0)
	}

	user.SubscriptionStatus = "active"
	user.SubscriptionPlan = plan
	user.SubscriptionExpiresAt = &expiresAt

	database.DB.Save(&user)

	return c.JSON(fiber.Map{
		"message":                 "Subscription successful",
		"subscription_status":     user.SubscriptionStatus,
		"subscription_plan":       user.SubscriptionPlan,
		"subscription_expires_at": user.SubscriptionExpiresAt,
	})
}
