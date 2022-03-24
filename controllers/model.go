package controllers

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

type Transaction struct {
	ID        int `json:"id"`
	UserID    int `json:"user id"`
	ProductID int `json:"product id"`
	Quantity  int `json:"quantity"`
}

type DetailTransaction struct {
	ID          int     `json:"id"`
	DataUser    User    `json:"user id"`
	DataProduct Product `json:"product id"`
	Quantity    int     `json:"quantity"`
}

type UserResponse struct {
	Response GeneralResponse `json:"response"`
	Data     User            `json:"data"`
}

type UsersResponse struct {
	Response GeneralResponse `json:"response"`
	Data     []User          `json:"data"`
}

type ProductResponse struct {
	Response GeneralResponse `json:"response"`
	Data     Product         `json:"data"`
}

type ProductsResponse struct {
	Response GeneralResponse `json:"response"`
	Data     []Product       `json:"data"`
}

type TransactionResponse struct {
	Response GeneralResponse `json:"response"`
	Data     Transaction     `json:"data"`
}

type TransactionsResponse struct {
	Response GeneralResponse `json:"response"`
	Data     []Transaction   `json:"data"`
}

type DetailTransactionsResponse struct {
	Response GeneralResponse     `json:"response"`
	Data     []DetailTransaction `json:"data"`
}

type GeneralResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
