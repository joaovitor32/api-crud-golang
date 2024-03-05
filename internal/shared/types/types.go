package types

type RequestBody struct {
	Name string `json:"name" binding:"omitempty"`
	Age int `json:"age" binding:"omitempty"`
}
