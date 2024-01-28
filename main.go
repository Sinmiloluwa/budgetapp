package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Sinmiloluwa/budgetapp/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load()

	portString := os.Getenv("PORT")

	if portString == "" {
		log.Fatal("Port is not found in the environment")
	}

	dbURL := os.Getenv("DB_URL")

	if dbURL == "" {
		log.Fatal("Db url is not found in the environment")
	}

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Cannot connect to database")
	}

	apiCfg := apiConfig{
		DB: database.New(conn),
	}



	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: false,
		MaxAge: 300,
	}))

	v1Router := chi.NewRouter()
	v1Router.HandleFunc("/health", handleReadiness)
	v1Router.Get("/err", handleError)
	v1Router.Post("/create-user", apiCfg.handleCreateUser)
	v1Router.Get("/get-user-by-key", apiCfg.handleGetUser)

	router.Mount("/v1", v1Router)

	srv := &http.Server {
		Handler: router,
		Addr: ":" + portString,
	}

	log.Printf("Server starting on %v", portString)

	err = srv.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Port:", portString)

}

func homePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	const page = `
		<html>
		<body>
		<p>Another one bites the dust</p>
		</body>
		</html>
		`
	w.WriteHeader(200)
	w.Write([]byte(page))
}
