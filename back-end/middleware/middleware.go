package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/AbdelkarimBENGRINE/todoapp/models"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/sideshow/apns2/payload"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

func init() {
	loadTheEnv()
	createDBInstance()
}

func loadTheEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error while loading the .env file")
	}
}

func createDBInstance() {
	connectionString := os.Getenv("DB_URI")
	dbName := os.Getenv("DB_NAME")
	collectionName := os.Getenv("DB_COLLECTION_NAME")
	clientOptions := options.Client().ApplyURI(connectionString) // TODO FIX URL
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil) 

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("connected to mongodb")
	client.Database(dbName).Collection(collectionName)
	fmt.Println("collection instance created")
}

func GetAllTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	payload := getAllTasks()
	json.NewEncoder(w).Encode(payload)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var task models.TodoList
	json.NewDecoder(r.Body).Decode(&task)
	insetOneTask(task)
	json.NewEncoder(w).Encode(task)
}

func TaskComplete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)
	TaskComplete(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func UndoTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)
	UndoTask(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)
	DeleteTask(params["id"])
}

func DeleteAllTasks(w http.ResponseWriter, r *http.Request) {

}

func getAllTasks(w http.ResponseWriter, r *http.Request) {

}

func TaskComplete(task string) {

}