package forms

type OrderForm struct {
	Address string `json:"address" binding:"required"`
	Name    string `json:"name" binding:"required"`
	Mobile  string `json:"mobile" binding:"required"`
	Post    string `json:"post" binding:"required"` //留言
}
