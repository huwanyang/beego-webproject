package bean

/*
	User 结构体
*/
type User struct {
	id      int    `form:"-"`
	Name    string `form:"username"`
	Age     int    `form:"age"`
	Email   string
	address string
}
