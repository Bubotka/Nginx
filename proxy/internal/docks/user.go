package docks

import (
	"github.com/Bubotka/Microservices/proxy/internal/modules/user/controller"
	"github.com/Bubotka/Microservices/user/domain/models"
)

// swagger:route GET /api/user/list user ListRequest
// Все пользователи.
// security:
// 	- Bearer: []
// Responses:
//   200: ListResponseDock

// swagger:response ListResponseDock
type ListResponseDock struct {
	// in: body
	// Response ответ содержащий пользователей и их количество в бд.
	Response controller.ListResponse
}

// swagger:route GET /api/user/profile/{username} user ProfileRequest
// Инфа о пользователе.
// security:
// 	- Bearer: []
// Responses:
//   200: ProfileResponse

// swagger:parameters ProfileRequest
type ProfileRequest struct {
	// Username - имя пользователя
	// in: path
	// required: true
	Username string `json:"Username"`
}

// swagger:response ProfileResponse
type ProfileResponse struct {
	// in: body
	// Response ответ содержащий пользователя.
	Response models.User
}
