package handlers

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"go-backend/database"
	"go-backend/models"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// Helper to get user ID from token
func getUserID(c *fiber.Ctx) (uint, error) {
	tokenString := c.Get("Authorization")
	if tokenString == "" {
		return 0, fmt.Errorf("unauthenticated")
	}

	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:]
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil // TODO: Use env var
	})

	if err != nil || !token.Valid {
		return 0, fmt.Errorf("unauthenticated")
	}

	claims := token.Claims.(jwt.MapClaims)
	userID := uint(claims["sub"].(float64))
	return userID, nil
}

func UploadSave(c *fiber.Ctx) error {
	userID, err := getUserID(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Unauthenticated"})
	}

	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "No file uploaded"})
	}

	gameName := c.FormValue("game_name")
	platform := c.FormValue("platform")

	// Create upload directory
	uploadDir := fmt.Sprintf("./uploads/%d", userID)
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Could not create upload directory"})
	}

	// Save file
	filename := filepath.Base(file.Filename)
	savePath := filepath.Join(uploadDir, filename)
	if err := c.SaveFile(file, savePath); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Could not save file"})
	}

	// Create database record
	gameSave := models.GameSave{
		UserID:   userID,
		GameName: gameName,
		Platform: platform,
		Filename: savePath,
		Size:     file.Size,
	}

	if err := database.DB.Create(&gameSave).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Could not save metadata"})
	}

	return c.JSON(gameSave)
}

func GetSaves(c *fiber.Ctx) error {
	userID, err := getUserID(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Unauthenticated"})
	}

	var saves []models.GameSave
	database.DB.Where("user_id = ?", userID).Find(&saves)

	return c.JSON(saves)
}

func DownloadSave(c *fiber.Ctx) error {
	userID, err := getUserID(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Unauthenticated"})
	}

	saveID := c.Params("id")
	var save models.GameSave
	if err := database.DB.First(&save, saveID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Save not found"})
	}

	if save.UserID != userID {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"message": "Unauthorized"})
	}

	return c.Download(save.Filename)
}
