package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {

	route := mux.NewRouter()

	route.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public/"))))

	route.HandleFunc("/", home).Methods("GET")
	route.HandleFunc("/home", home).Methods("GET")
	route.HandleFunc("/addProject", addProject).Methods("GET")
	route.HandleFunc("/addProject", addProjectPost).Methods("POST")
	route.HandleFunc("/contactMe", contactMe).Methods("GET")
	route.HandleFunc("/addContactMe", contactMePost).Methods("POST")
	route.HandleFunc("/projectDetail", projectDetail).Methods("GET")

	fmt.Println("server running on port 5000")
	http.ListenAndServe("localhost:5000", route)

}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "text/html; charset=utf8")
	var tmpl, err = template.ParseFiles("views/home.html")

	if err != nil {
		// w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("massage : " + err.Error()))
		return
	}

	// w.Write([]byte("home"))
	// w.WriteHeader(http.StatusOK)

	tmpl.Execute(w, nil)
}

func addProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "text/html; charset=utf8")
	var tmpl, err = template.ParseFiles("views/add-my-project.html")

	if err != nil {
		w.Write([]byte("massage : " + err.Error()))
		return
	}

	tmpl.Execute(w, nil)
}

func contactMe(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "text/html; charset=utf8")
	var tmpl, err = template.ParseFiles("views/contact.html")

	if err != nil {
		w.Write([]byte("massage : " + err.Error()))
		return
	}

	tmpl.Execute(w, nil)
}

func projectDetail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "text/html; charset=utf8")
	var tmpl, err = template.ParseFiles("views/my-project-detail.html")

	if err != nil {
		w.Write([]byte("massage : " + err.Error()))
		return
	}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	data := map[string]interface{}{
		"NameProject": "DUMBWAYS WEB APP",
		"Description": "Lorem ipsum dolor sit amet consectetur, adipisicing elit. Nulla optio pariatur quos doloremque neque vitae aliquam voluptate perferendis? Eaque enim quisquam ipsam unde, expedita saepe aliquid a praesentium est fuga.",
		"StartDate":   "12 Jan 2021",
		"EndDate":     "11 Feb 2021",
		"Duration":    "1 Month",
		"Image":       "/public/image/4625702.jpg",
		"Id":          id,
	}

	tmpl.Execute(w, data)
}

func addProjectPost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("projectName : " + r.PostForm.Get("input-nameProject"))
	fmt.Println("startDate : " + r.PostForm.Get("input-startDate"))
	fmt.Println("endDate : " + r.PostForm.Get("input-endDate"))
	fmt.Println("description : " + r.PostForm.Get("description"))
	fmt.Println("technologies : " + r.PostForm.Get("icon"))
	fmt.Println("image : " + r.PostForm.Get("input-image"))

	http.Redirect(w, r, "/home", http.StatusMovedPermanently)

}

func contactMePost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("name : " + r.PostForm.Get("input-name"))
	fmt.Println("email : " + r.PostForm.Get("input-email"))
	fmt.Println("phoneNumber : " + r.PostForm.Get("input-phonenumber"))
	fmt.Println("subject : " + r.PostForm.Get("input-subject"))
	fmt.Println("message : " + r.PostForm.Get("input-yourmessage"))

	http.Redirect(w, r, "/home", http.StatusMovedPermanently)

}

// go run main.go
