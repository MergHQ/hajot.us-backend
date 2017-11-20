package utils

import "../domain"

type ApiResponse struct {
	Message string
	Data *domain.Post
}

type ApiResponseArray struct {
	Message string
	Data []domain.Post
}