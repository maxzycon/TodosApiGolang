package create

type CreateCategoryInput struct {
	Name string `form:"name" json:"name" validate:"required"`
}