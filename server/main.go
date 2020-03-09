package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"gopkg.in/natefinch/lumberjack.v2"
)

type pizza struct {
	ID    string `json:"ID"`
	Name  string `json:"Name"`
	Price uint32 `json:"Price"`
}

type topping struct {
	ID   string `json:"ID"`
	Name string `json:"Name"`
}

type allPizzas []pizza
type allToppings []topping

var pizzas = allPizzas{
	{
		ID:    "1",
		Name:  "Margherita",
		Price: 9000,
	},
}

var toppings = allToppings{
	{
		ID:   "1",
		Name: "Tomato sauce",
	},
	{
		ID:   "2",
		Name: "Mozzarella",
	},
}

func handler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	name := query.Get("name")
	if name == "" {
		name = "Guest"
	}
	log.Printf("Received request for %s\n", name)
	w.Write([]byte(fmt.Sprintf("Hello, %s\n", name)))
}

/*
Pizza Route Handlers
*/
func getOnePizza(w http.ResponseWriter, r *http.Request) {
	pizzaID := mux.Vars(r)["id"]

	for _, singlePizza := range pizzas {
		if singlePizza.ID == pizzaID {
			json.NewEncoder(w).Encode(singlePizza)
		}
	}
}

func getAllPizzas(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(pizzas)
}

/*
Topping Route Handler
*/
func getAllToppings(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(toppings)
}

func main() {
	// Create Server and Route Handlers
	r := mux.NewRouter().StrictSlash(true)

	// Routes
	r.HandleFunc("/", handler)
	r.HandleFunc("/pizzas", getAllPizzas).Methods("GET")
	r.HandleFunc("/pizzas/{id}", getOnePizza).Methods("GET")

	srv := &http.Server{
		Handler:      r,
		Addr:         ":8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Configure Logging
	logFileLocation := os.Getenv("LOG_FILE_LOCATION")
	if logFileLocation != "" {
		log.SetOutput(&lumberjack.Logger{
			Filename:   logFileLocation,
			MaxSize:    500, // megabytes
			MaxBackups: 3,
			MaxAge:     28,   //days
			Compress:   true, // disabled by default
		})
	}

	// Start Server
	go func() {
		log.Println("Starting Server")
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	// Graceful Shutdown
	waitForShutdown(srv)
}

func waitForShutdown(srv *http.Server) {
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Block until we receive our signal.
	<-interruptChan

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	srv.Shutdown(ctx)

	log.Println("Shutting down")
	os.Exit(0)
}
