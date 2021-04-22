package forms

type CreateRequest struct {
	TierId uint64 `json:"tierId" binding:"required"`
}
