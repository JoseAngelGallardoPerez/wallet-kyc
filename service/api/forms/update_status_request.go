package forms

type UpdateStatusRequest struct {
	Status string `json:"status" binding:"required,requestStatus"`
}
