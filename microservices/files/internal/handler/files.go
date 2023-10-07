package handler

import (
	"fmt"
	"github.com/blazee5/cloud-drive/microservices/files/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"os"
	"strings"
)

func (h *Handler) Upload(c *fiber.Ctx) error {
	var fileInput models.File

	file, err := c.FormFile("file")

	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	uniqueId := uuid.New()
	fileName := strings.Replace(uniqueId.String(), "-", "", -1)
	fileExt := strings.Split(file.Filename, ".")[1]
	image := fmt.Sprintf("%s.%s", fileName, fileExt)

	if _, err := os.Stat("public"); os.IsNotExist(err) {
		err = os.Mkdir("public", os.ModePerm)
	}

	if err = c.SaveFile(file, fmt.Sprintf("./public/%s", image)); err != nil {
		h.log.Infof("error while save file: %s", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	fileInput.Name = image

	_, err = h.service.Upload(c.Context(), fileInput)

	if err != nil {
		h.log.Infof("error while upload file: %s", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"filename": fileName,
	})
}

func (h *Handler) Download(c *fiber.Ctx) error {
	fileName := c.Params("filename")

	err := h.service.File.AddCount(c.Context(), fileName)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "file not found",
		})
	}

	return c.Download("./public/"+fileName, fileName)
}
