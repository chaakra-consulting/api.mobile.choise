package requests

type FindEmailRequest struct {
	Email string `json:"email" binding:"required"`
}
