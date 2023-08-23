package main

import (

	routes "api/api/routes"
    
	"github.com/gin-gonic/gin"
)

func main() {
    app:= gin.Default()

    routes.AppRoutes(app)
    
    app.Run("localhost:3001")
}












// func main() {
//     fmt.Println("Hello, World!")
// }
// import(
// 	"fmt"
// 	"net/http"
// )

// func mainHandler(w http.ResponseWriter, r *http.Request){
// 	fmt.Println(w, "Bem-vindo à API em Go!")
// }

// func main(){
// 	http.HadleFunc("/", mainHandler)

// 	fmt.Println("Servidor rodando em http://localhost:8080")
// 	http.ListenAndServe(":8080", nill)
// }