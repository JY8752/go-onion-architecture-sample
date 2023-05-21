package request

type CreateTodoRequest struct {
	UserId      int64  `json:"userId" form:"userId" query:"userId" param:"userId"`
	Title       string `json:"title" form:"title" query:"title"`
	Description string `json:"description" form:"description" query:"description"`
}
