package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Document struct {
	ID    primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name  string             `json:"name" bson:"name"`
	Value string             `json:"value" bson:"value"`
}

type Server struct {
	collection *mongo.Collection
	router     *http.ServeMux
}

type Handlers struct {
	DocumentHandler *DocumentHandler
}

type DocumentHandler struct {
	collection *mongo.Collection
}

func main() {
	// MongoDB connection
	uri := "mongodb://localhost:27017"
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	if err = client.Ping(ctx, nil); err != nil {
		log.Fatal("Failed to connect:", err)
	}

	db := client.Database("helloayush")
	collection := db.Collection("documents")
	fmt.Println("Connected to MongoDB!")

	// Initialize server
	server := &Server{
		collection: collection,
		router:     http.NewServeMux(),
	}

	// Initialize handlers
	handlers := Handlers{
		DocumentHandler: &DocumentHandler{collection: collection},
	}

	// Initialize routes
	server.InitRoutes(handlers)

	fmt.Println("Server starting on port 8080")
	log.Fatal(http.ListenAndServe(":8080", server.router))
}

// Routes to handle the requests
func (s *Server) InitRoutes(h Handlers) {
	router := s.router

	router.HandleFunc("GET /ayush", h.DocumentHandler.GetHiAyush)
	router.HandleFunc("GET /allProducts", h.DocumentHandler.GetAll)
	router.HandleFunc("POST /create", h.DocumentHandler.Create)
}

// DB Operations to get data and create data in DB
func (h *DocumentHandler) GetHiAyush(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Hi Ayush, Welcome to the Rapido"})
}

func (h *DocumentHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := h.collection.Find(ctx, bson.M{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cursor.Close(ctx)

	var docs []Document
	if err := cursor.All(ctx, &docs); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return empty array instead of null if no documents
	if docs == nil {
		docs = []Document{}
	}
	json.NewEncoder(w).Encode(docs)
}

func (h *DocumentHandler) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var doc Document
	json.NewDecoder(r.Body).Decode(&doc)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	doc.ID = primitive.NewObjectID()
	h.collection.InsertOne(ctx, doc)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(doc)
}
