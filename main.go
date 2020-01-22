package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Cookie struct {
	ID string `bson:"_id" json:"id"`
	Name string `json:"Name" bson:"Name`
	Quantity int `json:"Quantity" bson:"Quantity"`
	LastBaked time.Time `json:"Last Baked" bson:"Last Baked"`
	Expiry time.Time `json:"Expiry" bson:"Expiry"`
	Price float64 `json:"Price per cookie" bson:"Price per cookie"`
	Description string `json:"Description" bson:"Description"`
}

func getCookies() []Cookie {
	session, err := mgo.Dial("mongodb://127.0.0.1:27017/bakery")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	c := session.DB("bakery").C("myCookies")
	cookies := []Cookie{}
	err = c.Find(nil).All(&cookies)
	if err != nil {
		log.Fatal(err)
	}
	return cookies
}

func removeCookie(id string) Cookie {
	session, err := mgo.Dial("mongodb://127.0.0.1:27017/bakery")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	c := session.DB("bakery").C("myCookies")
	cookie := Cookie{}
	err = c.Remove(bson.M{"id": id})
	if err != nil {
		log.Fatal(err)
	}
	return cookie

}

func createCookie(e Cookie) error {
	session, err := mgo.Dial("mongodb://127.0.0.1:27017/bakery")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	c := session.DB("bakery").C("cookies")
	return c.Insert(e)
}

func getAllCookies(res http.ResponseWriter, req *http.Request) {
	cookies := getCookies()
	json.NewEncoder(res).Encode(cookies)
}

func postCookie(res http.ResponseWriter, req *http.Request) {
	var item Cookie
	json.NewDecoder(req.Body).Decode(&item)
	createCookie(item)
}

func deleteCookie(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	removeCookie(params["id"])
	res.Write([]byte("OK"))
}

func handleRequest() {
	router := mux.NewRouter()
	router.HandleFunc("/api/listCookies", getAllCookies).Methods("GET")
	router.HandleFunc("/api/insertCookie", postCookie).Methods("POST")
	router.HandleFunc("/api/deleteCookie/{id}", deleteCookie).Methods("DELETE")
	http.ListenAndServe(":8080", router)
}

func main() {
	handleRequest()
}