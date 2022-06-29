package main 

import(
	"log"
	"net/http"
	"github.com/gorilla/mux"
	// "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/aviuday/Globalfair-project/pkg/routes"
)

func main(){
	r := mux.NewRouter()
	routes.RegisterRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:3000", r))
}