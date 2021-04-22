package forms

type UpdateStatusRequirement struct {
	Status string `json:"status" binding:"required,requirementStatus"`
}
