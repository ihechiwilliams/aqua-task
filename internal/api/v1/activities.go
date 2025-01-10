package v1

import (
	"aqua-backend/internal/repositories/activities"
)

type ActivitiesHandler struct {
	activitiesRepo activities.Repository
}

func NewActivitiesHandler(activitiesRepo activities.Repository) *ActivitiesHandler {
	return &ActivitiesHandler{
		activitiesRepo: activitiesRepo,
	}
}
