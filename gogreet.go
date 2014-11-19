package main
 
 
import (
  "log"
  "net/http"
  "html/template"
)
 
 
func main() {
  http.HandleFunc("/", root)
  http.HandleFunc("/greet", greeter)
  log.Println("Listening...")
  err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
  if err != nil {
    log.Fatal("ListenAndServe: ", err)
   }
}
 
 
func root(w http.ResponseWriter, r *http.Request) {
  t, _ := template.ParseFiles("public/index.html")
  t.Execute(w, nil)
}
 
 
func greeter(w http.ResponseWriter, r *http.Request) {
  username := r.FormValue("username")
  t, _ := template.ParseFiles("public/greeter.html")
  err := t.Execute(w, username)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}