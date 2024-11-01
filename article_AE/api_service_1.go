package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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

	// start the server
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

// REST API handlery

func (s ServerImpl) returnListOfUsers(writer http.ResponseWriter, r *http.Request) {
	users := s.storage.ReadListOfUsers()
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(users)
}

func main() {
	storage := NewMemoryStorage()
	server := NewServer(storage)
	server.Serve(8080)
}
