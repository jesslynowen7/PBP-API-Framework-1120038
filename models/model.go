package models

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Address  string `json:"address"`
	Email    string `json:"email"`
	UserType int    `json:"userType"`
}

type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type UsersssResponse struct {
	Response GeneralResponse `json:"response"`
	Data     []User          `json:"data"`
}

type UserResponse struct {
	Response GeneralResponse `json:"response"`
	Data     User            `json:"data"`
}

type ProductResponse struct {
	Response GeneralResponse `json:"response"`
	Data     Product         `json:"data"`
}

type ProductsResponse struct {
	Response GeneralResponse `json:"response"`
	Data     []Product       `json:"data"`
}

type GeneralResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
