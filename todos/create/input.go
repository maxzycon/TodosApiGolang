package create

type CreateTodoInput struct {
	Title string `binding:"required"`
	Description string `binding:"required"`
	CategoryID int `binding:"required" json:"category_id"`
}