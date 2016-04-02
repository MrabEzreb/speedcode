package main

import (
	"fmt"
	"net/http"
	"html"
	"strings"
	"io/ioutil"
	"encoding/json"
	"log"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if strings.Index(html.EscapeString(r.URL.Path), "/g/") == 0 {
			name := strings.Replace(strings.Replace(strings.Split(html.EscapeString(r.URL.Path), "g/")[1], "\"", "", -1), "/", " ", -1)
			fmt.Fprintf(w, "Goodbye %s!", name)
		} else if strings.Index(html.EscapeString(r.URL.Path), "/h/") == 0 {
			name := strings.Replace(strings.Replace(strings.Split(html.EscapeString(r.URL.Path), "h/")[1], "\"", "", -1), "/", " ", -1)
			fmt.Fprintf(w, "Hello %s!", name)
		} else if strings.Index(html.EscapeString(r.URL.Path), "/u/h") == 0 {
			nameC, _ := r.Cookie("UName")
			name := nameC.Value
			fmt.Fprintf(w, "Hello %s!", name)
		} else if strings.Index(html.EscapeString(r.URL.Path), "/u/g") == 0 {
			nameC, _ := r.Cookie("UName")
			name := nameC.Value
			fmt.Fprintf(w, "Goodbye %s!", name)
		} else if strings.Index(html.EscapeString(r.URL.Path), "/u/o") == 0 {
			cookie := http.Cookie{Name: "UName", Value: "", Path: "/"}
			http.SetCookie(w, &cookie)
			http.Redirect(w, r, "/", http.StatusFound)
		} else if strings.Index(html.EscapeString(r.URL.Path), "/u") == 0 {
			fmt.Fprintf(w, "<html><body><h2>Welcome to SpeedCodeSite User Section!</h2><p>Say Goodbye: /u/g</p><p>Say Hello: /u/h</p><p>Logout: /u/o</p></body></html>")
		} else if strings.Index(html.EscapeString(r.URL.Path), "/s/post") == 0 {
			username := r.FormValue("username")
			_, err := LoadUser(username)
			if err == nil {
				fmt.Fprintf(w, "<a href=\"/l/\">Error</a>")
			} else {
				U := User{Username: r.FormValue("username"), Password: r.FormValue("password"), Name: r.FormValue("name")}
				SaveUser(U)
				http.Redirect(w, r, "/l/", http.StatusFound)
			}
		} else if strings.Index(html.EscapeString(r.URL.Path), "/s") == 0 {
			fmt.Fprintf(w, "<html><body><h2>Signup</h2><form method=\"post\" action=\"/s/post\">Username <input type=\"text\" name=\"username\"><br>Name <input type=\"text\" name=\"name\"><br>Password <input type=\"password\" name=\"password\"><br><input type=\"submit\" value=\"Submit\"></form></body></html>")
		} else if strings.Index(html.EscapeString(r.URL.Path), "/l/post") == 0 {
			username := r.FormValue("username")
			User, err := LoadUser(username)
			if err != nil {
				fmt.Fprintf(w, "<a href=\"/l/\">Error</a>")
			} else {
				cookie := http.Cookie{Name: "UName", Value: User.Name, Path: "/"}
				http.SetCookie(w, &cookie)
				http.Redirect(w, r, "/u/", http.StatusFound)
			}
		} else if strings.Index(html.EscapeString(r.URL.Path), "/l") == 0 {
			fmt.Fprintf(w, "<html><body><h2>Login</h2><form method=\"post\" action=\"/l/post\">Username <input type=\"text\" name=\"username\"><br>Password <input type=\"password\" name=\"password\"><br><input type=\"submit\" value=\"Submit\"></form></body></html>")
		} else if strings.Index(html.EscapeString(r.URL.Path), "/k") == 0 {
			os.Exit(0)
		} else {
			fmt.Fprintf(w, "<html><body><h2>Welcome to SpeedCodeSite!</h2><p>Say Goodbye: /g/_name_</p><p>Say Hello: /h/_name_</p><p>Login: /l</p><p>Signup: /s</p></body></html>")
		}
	})
	fmt.Println("test")
	log.Fatal(http.ListenAndServe(":8080", nil))
	fmt.Println("test2")
}

type User struct {
	Username	string
	Password	string
	Name		string
}

func LoadUser(Username string) (User, error) {
	filename := Username + ".user"
	raw, err0 := ioutil.ReadFile(filename)
	if err0 != nil {
		return User{}, err0
	}
	var u User
	err1 := json.Unmarshal(raw, &u)
	if err1 != nil {
		return User{}, err1
	}
	return u, nil
}

func SaveUser(User User) error {
	raw, err0 := json.Marshal(User)
	if err0 != nil {
		return err0
	}
	filename := User.Username + ".user"
	err1 := ioutil.WriteFile(filename, raw, 0600)
	if err1 != nil {
		return err1
	}
	return nil
}