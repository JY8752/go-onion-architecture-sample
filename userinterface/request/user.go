package request

type CreateUserRequest struct {
	Name string `json:"name" form:"name" query:"name"`
}
