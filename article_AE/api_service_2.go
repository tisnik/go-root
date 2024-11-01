package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// Datova struktura predstavujici uzivatele
type User struct {
	Name    string
	Surname string
}

// Rozhrani s predpisem metod pro manipulace s databazi uzivatelu
type UserStorage interface {
	ReadListOfUsers() []User
	ReadUser(ID int) (User, bool)
	DeleteUser(ID int)
}

// Implementace databaze uzivatelu v operacni pameti
type MemoryStorage struct {
	users map[int]User
}

// Inicializace databaze uzivatelu
func NewMemoryStorage() MemoryStorage {
	users := make(map[int]User)
	users[1] = User{
		Name:    "Linus",
		Surname: "Torvalds",
	}
	users[2] = User{
		Name:    "Rob",
		Surname: "Pike",
	}
	users[3] = User{
		Name:    "Ken",
		Surname: "Iverson",
	}

	return MemoryStorage{
		users: users,
	}
}

// Implementace vsech metod predepsanych rozhranim UserStorage

func (s MemoryStorage) ReadListOfUsers() []User {
	users := make([]User, 0, len(s.users))
	for _, user := range s.users {
		users = append(users, user)
	}
	return users
}

func (s MemoryStorage) ReadUser(ID int) (User, bool) {
	user, found := s.users[ID]
	return user, found
}

func (s MemoryStorage) DeleteUser(ID int) {
	delete(s.users, ID)
}

// Rozhrani predepisujici metody serveru
type Server interface {
	Serve(port uint)
}

// Implementace HTTPServeru
type ServerImpl struct {
	storage UserStorage
}

// Inicialiace serveru, predani kontextu
func NewServer(storage UserStorage) Server {
	return ServerImpl{
		storage: storage,
	}
}

func (s ServerImpl) Serve(port uint) {
	log.Printf("Starting server on port 8080")

	// REST API endpoints
	http.HandleFunc("/users", s.returnListOfUsers)
	http.HandleFunc("/user", s.returnOneUser)
	http.HandleFunc("/delete-user", s.deleteOneUser)

	// start the server
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

// REST API handlery

func (s ServerImpl) returnListOfUsers(writer http.ResponseWriter, r *http.Request) {
	users := s.storage.ReadListOfUsers()
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(users)
}

func (s ServerImpl) returnOneUser(writer http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	IDs, found := params["ID"]
	if !found {
		writer.Header().Set("Content-Type", "text/plain")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	ID, err := strconv.Atoi(IDs[0])
	if err != nil {
		writer.Header().Set("Content-Type", "text/plain")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	user, found := s.storage.ReadUser(ID)
	if !found {
		writer.Header().Set("Content-Type", "text/plain")
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(user)
}

func (s ServerImpl) deleteOneUser(writer http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	IDs, found := params["ID"]
	if !found {
		writer.Header().Set("Content-Type", "text/plain")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	ID, err := strconv.Atoi(IDs[0])
	if err != nil {
		writer.Header().Set("Content-Type", "text/plain")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	s.storage.DeleteUser(ID)

	writer.Header().Set("Content-Type", "application/json")
	status := struct {
		Status string
		ID     int
	}{"deleted", ID}
	json.NewEncoder(writer).Encode(status)
}

func main() {
	storage := NewMemoryStorage()
	server := NewServer(storage)
	server.Serve(8080)
}
