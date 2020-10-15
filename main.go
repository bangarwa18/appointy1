package main

import (
	"fmt"
	//"io/ioutil"
	"log"
	"net/http"
	"encoding/json"
)

type Article struct {
	Id      string   `json:"Id"`
	Title   string `json:"Title"`
	SubTitle    string `json:"SubTitle"`
	Content string `json:"content"`
}

// let's declare a global Articles array
// that we can then populate in our main function
// to simulate a database
var Articles []Article

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
    http.HandleFunc("/", homePage)
    // add our articles route and map it to our 
    // returnAllArticles function like so
    http.HandleFunc("/articles", returnAllArticles)
    http.HandleFunc("/articles1", createNewArticle)
    log.Fatal(http.ListenAndServe(":10000", nil))
}

func createNewArticle(w http.ResponseWriter, r *http.Request) {
    // get the body of our POST request
    // unmarshal this into a new Article struct
    // append this to our Articles array.    
    //reqBody, _ := ioutil.ReadAll(r.Body)
    //var article Article
    if r.Method == http.MethodPost  {
    article := Article{
    	Id:r.FormValue("Id"),
    	Title:r.FormValue("Title"),
    	SubTitle:r.FormValue("SubTitle"),
    	Content:r.FormValue("Content"),
    }
    //json.Unmarshal(reqBody, &article)
    // update our global Articles array to include
    // our new Article
    Articles = append(Articles, article)

    json.NewEncoder(w).Encode(article)
}
}

func returnAllArticles(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint Hit: returnAllArticles")
    fmt.Fprintf(w, "<h1>Editing</h1>"+
        "<form action=\"/articles1\" method=\"POST\">"+
        "<input type=\"text\" value=\"Id\"></input><br>"+
        "<input type=\"text\" value=\"Title\"></input><br>"+
        "<input type=\"text\" value=\"SubTitle\"></input><br>"+
        "<input type=\"text\" value=\"Content\"></input><br>"+
        "<input type=\"submit\" value=\"Save\">"+
        "</form>")
    json.NewEncoder(w).Encode(Articles)
}

func main() {
    Articles = []Article{
        Article{Id:"1",Title: "Hello", SubTitle: "Article Description", Content: "Article Content"},
        Article{Id:"2",Title: "Hello 2", SubTitle: "Article Description", Content: "Article Content"},
    }
    handleRequests()
}

