package main

import (
	"log"
	"net/http"

	"github.com/go-martini/martini"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/cors"

	controllers "latihan/controllers"
)

func main() {
	m := martini.Classic()
	// 0 = user biasa, 1 = admin

	m.Get("/products", controllers.Authenticate(0), controllers.GetAllProducts)
	m.Post("/products", controllers.Authenticate(1), controllers.InsertProduct)
	m.Put("/products", controllers.Authenticate(1), controllers.UpdateProduct)
	m.Delete("/products/:id", controllers.Authenticate(1), controllers.DeleteProduct)

	// m.Group("/products", func(r martini.Router) {
	// 	r.Get("/get", controllers.Authenticate(0), controllers.GetAllProducts)
	// 	r.Post("/add", controllers.Authenticate(1), controllers.InsertProduct)
	// 	r.Put("/update", controllers.Authenticate(1), controllers.UpdateProduct)
	// 	r.Delete("/:id", controllers.Authenticate(1), controllers.DeleteProduct)
	// })

	// m.Get("/", func() (int, string) {
	// 	return 418, "this is handler" // HTTP 418 : "this is handler"
	// })

	// m.Get("/hello/:name", func(params martini.Params) string {
	// 	return "Hello " + params["name"]
	// })

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowCredentials: true,
	})
	handler := corsHandler.Handler(m)

	m.Post("/login", controllers.Login)
	m.Get("/logout", controllers.Logout)

	http.Handle("/", m)
	m.RunOnAddr(":8080")

	log.Fatal(http.ListenAndServe(":8080", handler))
}
