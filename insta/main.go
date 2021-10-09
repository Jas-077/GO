package main

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Name     string             `json:"name"`
	Email    string             `json:"mail"`
	Password string             `json:"pass"`
}

type Post struct {
	Id        string `json:"usr-id"`
	Name      string `json:"name"`
	Caption   string `json:"mail"`
	URL       string `json:"pass"`
	TimeStamp string `json:"time"`
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Endpoint Hit: Home Page")
	fmt.Println("method:", r.Method)
}

func addUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: Add User")
	fmt.Println("method:", r.Method)
	if r.Method == http.MethodPost {
		fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", r.PostForm)
		var name string
		var email string
		var pass string
		//var new_art Article
		name = r.FormValue("name")
		email = r.FormValue("mail")
		pass = r.FormValue("pass")
		//var new_user User = User{name, email, pass}
		fmt.Fprintf(w, "name = %s\n", name)
		fmt.Fprintf(w, "email = %s\n", email)
		fmt.Fprintf(w, "password= %s\n", pass)

		clientOptions := options.Client().
			ApplyURI("mongodb+srv://jasjeet:123@cluster0.glzuy.mongodb.net/insta?retryWrites=true&w=majority")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		client, err := mongo.Connect(ctx, clientOptions)
		if err != nil {
			log.Fatal(err)
		}

		data := []byte(pass)
		fmt.Printf("%x", md5.Sum(data))

		new_pass := hex.EncodeToString(data[:])

		collection := client.Database("insta").Collection("user")

		user := bson.D{{Key: "Name", Value: name},
			{Key: "Email", Value: email}, {Key: "Password", Value: new_pass}}

		res, insertErr := collection.InsertOne(ctx, user)
		if insertErr != nil {
			log.Fatal(insertErr)
		}
		fmt.Println(res)

	} else {
		http.ServeFile(w, r, "userForm.html")
	}
}

func addPost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: Add Post")
	fmt.Println("method:", r.Method)
	if r.Method == http.MethodPost {
		fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", r.PostForm)
		var name string
		var cap string
		var url string
		//var new_art Article
		name = r.FormValue("name")
		cap = r.FormValue("cap")

		//var new_user User = User{name, email, pass}
		fmt.Fprintf(w, "name = %s\n", name)
		fmt.Fprintf(w, "caption = %s\n", cap)

		file, handler, err := r.FormFile("myFile")
		if err != nil {
			fmt.Println("Error Retrieving the File")
			fmt.Println(err)
			return
		}

		defer file.Close()

		// Create file
		dst, err := os.Create(handler.Filename)
		defer dst.Close()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Copy the uploaded file to the created file on the filesystem
		if _, err := io.Copy(dst, file); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "Successfully Uploaded File\n")

		pwd, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(pwd)

		var new_path string = pwd + handler.Filename
		url = new_path
		fmt.Printf("Uploaded File: %+v\n", handler.Filename)
		fmt.Printf("File Size: %+v\n", handler.Size)
		fmt.Printf("MIME Header: %+v\n", handler.Header)
		fmt.Println("New url: ", url)

		clientOptions := options.Client().
			ApplyURI("mongodb+srv://jasjeet:123@cluster0.glzuy.mongodb.net/insta?retryWrites=true&w=majority")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		client, err := mongo.Connect(ctx, clientOptions)
		if err != nil {
			log.Fatal(err)
		}

		var user User
		collection := client.Database("insta").Collection("user")
		if err = collection.FindOne(ctx, bson.M{"Name": name}).Decode(&user); err != nil {
			log.Fatal(err)
		}
		fmt.Println(user)

		collection2 := client.Database("insta").Collection("posts")
		t := time.Now()
		post := bson.D{{Key: "Name", Value: name},
			{Key: "Caption", Value: cap}, {Key: "URL", Value: url},
			{Key: "Timestamp", Value: t.String()},
			{Key: "User_ID", Value: user.ID}}

		res, insertErr := collection2.InsertOne(ctx, post)
		if insertErr != nil {
			log.Fatal(insertErr)
		}
		fmt.Println(res)
	} else {
		http.ServeFile(w, r, "Post.html")
	}
}

func getUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: Get User")
	fmt.Println("method:", r.Method)
	if r.Method == http.MethodGet {
		clientOptions := options.Client().
			ApplyURI("mongodb+srv://jasjeet:123@cluster0.glzuy.mongodb.net/insta?retryWrites=true&w=majority")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		client, err := mongo.Connect(ctx, clientOptions)
		if err != nil {
			log.Fatal(err)
		}

		keys, ok := r.URL.Query()["id"]
		if !ok || len(keys[0]) < 1 {
			log.Println("Url Param 'id' is missing")
			return
		}
		key := keys[0]

		objectId, err := primitive.ObjectIDFromHex(key)
		if err != nil {
			log.Println("Invalid id")
		}

		var user bson.M
		collection := client.Database("insta").Collection("user")
		if err = collection.FindOne(ctx, bson.M{"_id": objectId}).Decode(&user); err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(w, "User: %s", user)
	} else {
		fmt.Fprintf(w, "Only Get Request allowed")
	}
}

func getPost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: Get Post")
	fmt.Println("method:", r.Method)
	if r.Method == http.MethodGet {
		clientOptions := options.Client().
			ApplyURI("mongodb+srv://jasjeet:123@cluster0.glzuy.mongodb.net/insta?retryWrites=true&w=majority")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		client, err := mongo.Connect(ctx, clientOptions)
		if err != nil {
			log.Fatal(err)
		}
		keys, ok := r.URL.Query()["id"]
		if !ok || len(keys[0]) < 1 {
			log.Println("Url Param 'id' is missing")
			return
		}
		key := keys[0]

		objectId, err := primitive.ObjectIDFromHex(key)
		if err != nil {
			log.Println("Invalid id")
		}

		var post bson.M
		collection := client.Database("insta").Collection("posts")
		if err = collection.FindOne(ctx, bson.M{"_id": objectId}).Decode(&post); err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(w, "Post: %s", post)
	} else {
		fmt.Fprintf(w, "Only Get Request allowed")
	}
}

func getPosts_User(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: Get Posts from User")
	fmt.Println("method:", r.Method)
	if r.Method == http.MethodGet {
		clientOptions := options.Client().
			ApplyURI("mongodb+srv://jasjeet:123@cluster0.glzuy.mongodb.net/insta?retryWrites=true&w=majority")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		client, err := mongo.Connect(ctx, clientOptions)
		if err != nil {
			log.Fatal(err)
		}
		keys, ok := r.URL.Query()["id"]
		if !ok || len(keys[0]) < 1 {
			log.Println("Url Param 'id' is missing")
			return
		}
		//fmt.Println(keys)
		id := keys[0]

		objectId, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			log.Println("Invalid id")
		}

		collection := client.Database("insta").Collection("posts")
		var posts []bson.M
		filterCursor, err := collection.Find(ctx, bson.M{"User_ID": objectId})
		if err != nil {
			log.Fatal(err)
		}
		if err = filterCursor.All(ctx, &posts); err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(w, "All Posts by %s:\n", id)
		//fmt.Fprintf(w, "%s", posts)
		for i := 0; i < len(posts); i++ {
			fmt.Fprintf(w, "\n%d: %s\n", (i + 1), posts[i])
		}
	} else {
		fmt.Fprintf(w, "Only Get Request allowed")
	}
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/users", addUser)
	http.HandleFunc("/users/", getUser)
	http.HandleFunc("/posts", addPost)
	http.HandleFunc("/posts/", getPost)
	http.HandleFunc("/posts/users/", getPosts_User)
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}

func main() {
	handleRequests()
}
