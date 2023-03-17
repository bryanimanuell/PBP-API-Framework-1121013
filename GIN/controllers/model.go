package controllers

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Address  string `json:"address"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Response struct {
	Message string      `json:"error"`
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
}
