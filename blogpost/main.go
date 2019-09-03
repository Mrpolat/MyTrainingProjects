package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"text/template"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //Parse url parameters passed, then parse the response packet for the POST body (request body)
	// attention: If you do not call ParseForm method, the following data can not be obtained form
	fmt.Println(r.Form) // print information on server side.
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	    for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello Polat!") // write data to response

}

type LoginStatus struct {
	Status string
}

type BlogPostData struct {
	Title    string
	BlogPost string
}

func Login(w http.ResponseWriter, r *http.Request) {
	var loginStatus LoginStatus

	fmt.Println("method:", r.Method) //get request method
	  if r.Method == "GET" {
		t, _ := template.ParseFiles("login.html")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		t, _ := template.ParseFiles("login.html")
		// we logged into the port
		if r.Form["username"][0] == "jack" && r.Form["password"][0] == "daniel" {
			//"if" checked
			t, _ = template.ParseFiles("blog_post.html")
			t.Execute(w, nil)

			//worked if correct and went to blog_post
		} else {
			loginStatus.Status = "Hatalı Giriş!"
			t, _ = template.ParseFiles("login.html")
			t.Execute(w, loginStatus)
			fmt.Println("giriş yanlış")
			
		}

	}
}

func BlogPost(w http.ResponseWriter, r *http.Request) {
	var blogPost BlogPostData
	    r.ParseForm()
	    if r.Method == "POST" {

		blogPost.Title = r.Form["title"][0]
		blogPost.BlogPost = r.Form["message"][0]

		fmt.Println(blogPost.Title)
		fmt.Println(blogPost.BlogPost)
		t, _ := template.ParseFiles("blog_page.html")
		t.Execute(w, blogPost)
	}

}

func main() {
	fmt.Println("Access URL: localhost:9090/login")
	http.HandleFunc("/", sayhelloName) // setting router rule
	http.HandleFunc("/login", Login)
	http.HandleFunc("/blog-post", BlogPost)
	err := http.ListenAndServe(":9090", nil) // setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}