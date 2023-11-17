package controllers

import "Scramble/app/backend/pkg/models"

type apiResponse struct {
	GameID       *string      `json:"gameID,omitempty"`
	GameResp     *models.Game `json:"gameState,omitempty"`
	ErrorMessage *string      `json:"message,omitempty"`
	Valid        bool         `json:"valid"`
}
