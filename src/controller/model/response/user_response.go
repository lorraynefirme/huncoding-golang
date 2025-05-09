package response

type UserResponse struct {
	ID    string `json:"id" binding:"required"`
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
	Age   int8   `json:"age" binding:"required"`
}