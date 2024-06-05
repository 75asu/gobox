package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/thedevsaddam/renderer"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var rnd *renderer.Render
var db *mgo.Database

const (
	hostName       string = "mongodb://localhost:27017"
	dbName         string = "todo1_db"
	collectionName string = "todo1_collection"
	port           string = ":9090"
)

// Two models will be there - json and bson
// BSON will interact with database
// JSON will interact with frontend
type (
	todo1Model struct {
		ID        bson.ObjectId `bson:"_id, omitempty"`
		Title     string        `bson:"title"`
		Completed bool          `bson:"completed"`
		CreatedAt time.Time     `bson:"createdAt"`
	}

	todo1 struct {
		ID        string    `json:"id"`
		Title     string    `json:"title"`
		Completed bool      `json:"completed"`
		CreatedAt time.Time `json:"created_at"`
	}
)

// Connect with database and start a session
func init() {
	rnd = renderer.New()
	log.Print("renderer starter")
	sess, err := mgo.Dial(hostName)
	log.Print("session started")
	checkErr(err)
	sess.SetMode(mgo.Monotonic, true)
	log.Print("mode set")
	db = sess.DB(dbName)
	log.Print("connection done")
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	err := rnd.Template(w, http.StatusOK, []string{"static/home.tpl"}, nil)
	checkErr(err)
}

func getTodo(w http.ResponseWriter, r *http.Request) {
	// Collect data from database in BSON format
	todos := []todo1Model{}
	if err := db.C(collectionName).Find(bson.M{}).All(&todos); err != nil {
		rnd.JSON(w, http.StatusProcessing, renderer.M{
			"message": "Failed to fetch todo",
			"error":   err,
		})
		return
	}

	// Convert the BSON data into JSON to show in frontend
	todoList := []todo1{}
	for _, t := range todos {
		todoList = append(todoList, todo1{
			ID:        t.ID.Hex(),
			Title:     t.Title,
			Completed: t.Completed,
			CreatedAt: t.CreatedAt,
		})
	}
	rnd.JSON(w, http.StatusOK, renderer.M{
		"data": todoList,
	})
}

func postTodo(w http.ResponseWriter, r *http.Request) {
	// 1. Decode the response
	// 2. Validate wether it has a title or not
	// 3. Create a TODO model
	// 4. Send the model to database
	// 5. Send a response back to the user
	var t todo1

	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		rnd.JSON(w, http.StatusProcessing, err)
		return
	}

	if t.Title == "" {
		rnd.JSON(w, http.StatusBadRequest, renderer.M{
			"message": "The title is required",
		})
	}

	tm := todo1Model{
		ID:        bson.NewObjectId(),
		Title:     t.Title,
		Completed: false,
		CreatedAt: time.Now(),
	}

	if err := db.C(collectionName).Insert(&tm); err != nil {
		rnd.JSON(w, http.StatusProcessing, renderer.M{
			"message": "Failed to save todo",
			"error":   err,
		})
		return
	}

	rnd.JSON(w, http.StatusCreated, renderer.M{
		"message": "TODO created successfully",
		"todo_id": tm.ID.Hex(),
	})
}

func putTodo(w http.ResponseWriter, r *http.Request) {
	// 1. Collect the object by ID
	// 2. Delete the object
	// 3. Create a new object with the body taken from request
	// 4. Update the new object with the existing object in database
	// 5. Update the user
	id := strings.TrimSpace(chi.URLParam(r, "id"))

	if !bson.IsObjectIdHex(id) {
		rnd.JSON(w, http.StatusBadRequest, renderer.M{
			"message": "The id is invalid",
		})
		return
	}

	var t todo1

	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		rnd.JSON(w, http.StatusProcessing, err)
		return
	}

	if t.Title == "" {
		rnd.JSON(w, http.StatusBadRequest, renderer.M{
			"message": "The title field id is required",
		})
		return
	}

	if err := db.C(collectionName).
		Update(
			bson.M{
				"_id": bson.ObjectIdHex(id),
			},
			bson.M{
				"title":     t.Title,
				"completed": t.Completed,
			},
		); err != nil {
		rnd.JSON(w, http.StatusProcessing, renderer.M{
			"messsage": "Failed to update TODO",
			"error":    err,
		})
		return
	}
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	// 1. Work with URL parameter
	// 2. Check wether the ID we sent is hex or not
	// 3. Remove the particular record with the ID from database
	// 4. Write back to frontend saying successfully deleted
	id := strings.TrimSpace(chi.URLParam(r, "id"))

	if !bson.IsObjectIdHex(id) {
		rnd.JSON(w, http.StatusBadRequest, renderer.M{
			"message": "The ID is invalid",
		})
		return
	}

	if err := db.C(collectionName).RemoveId(bson.ObjectIdHex(id)); err != nil {
		rnd.JSON(w, http.StatusProcessing, renderer.M{
			"message": "Failed to delete TODO",
			"error":   err,
		})
	}

	rnd.JSON(w, http.StatusOK, renderer.M{
		"message": "TODO deleted successfully",
	})
}

// handle all the routes starting with "/todo"
func todoHandlers() http.Handler {
	rg := chi.NewRouter()
	rg.Group(func(r chi.Router) {
		r.Get("/", getTodo)
		r.Post("/", postTodo)
		r.Put("/{id}", putTodo)
		r.Delete("/{id}", deleteTodo)
	})
	return rg
}

func main() {
	// Graceful exit of server
	stopChan := make(chan os.Signal)
	signal.Notify(stopChan, os.Interrupt)

	// Create routes
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", homeHandler)
	r.Mount("/todo", todoHandlers())

	// Server parameters
	srv := http.Server{
		Addr:         port,
		Handler:      r,
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Start and stop server
	go func() {
		log.Println("Listening on port: ", port)
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("listen: %s\n", err)
		}
	}()
	<-stopChan
	log.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	srv.Shutdown(ctx)
	defer cancel()
	log.Println("Server gracefully stopped")
}
