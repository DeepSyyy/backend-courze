package web

import (
	"courze-backend-app/model/domain"
)

type AdminRequest struct {
	Course domain.Course `json:"course" validate:"required"`
}
