package controller

import "github.com/Bubotka/Microservices/user/domain/models"

type CreateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ListResponse struct {
	Users []models.User `json:"users"`
}
