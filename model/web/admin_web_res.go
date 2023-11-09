package web

import (
	"courze-backend-app/model/domain"
)

type AdminResponse struct {
	Course domain.Course `json:"course"`
}
