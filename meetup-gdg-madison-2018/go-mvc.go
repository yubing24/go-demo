package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

// Contact contains the ID, Name, JobTitle, PhoneNumber of my contact. json tag specifies the json
// field name when this struct is converted to JSON format
type Contact struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	JobTitle    string `json:"title"`
	PhoneNumber string `json:"phone"`
}

// ContactRepository is a repository of Contact
type ContactRepository struct {
	DatabaseConnection *sql.DB
}

// A global ContactRepository
var repo ContactRepository

// GetAllContact retrieve all my contact from the repository
func (repo ContactRepository) GetAllContact() ([]Contact, error) {
	rows, err := repo.DatabaseConnection.Query("SELECT * FROM CONTACT")
	results := make([]Contact, 0)
	for rows.Next() {
		each := Contact{}
		rows.Scan(&each.ID, &each.Name, &each.JobTitle, &each.PhoneNumber)
		results = append(results, each)
	}
	rows.Close()
	return results, err
}

// AddContact adds a new contact to the repository
func (repo ContactRepository) AddContact(contact *Contact) error {
	tx, err := repo.DatabaseConnection.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare("INSERT INTO CONTACT (NAME, JOBTITLE, PHONENUMBER) values(?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec(contact.Name, contact.JobTitle, contact.PhoneNumber)
	tx.Commit()
	return err
}

func createSchema(db *sql.DB) {
	db.Exec(`CREATE TABLE CONTACT (
		ID INTEGER PRIMARY KEY AUTOINCREMENT,
		NAME VARCHAR(64) NULL,
		JOBTITLE VARCHAR(64) NULL,
		PHONENUMBER VARCHAR(64) NULL );`)
}

func getContactHandler(w http.ResponseWriter, r *http.Request) {
	contacts, _ := repo.GetAllContact()
	output, _ := json.Marshal(contacts)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(output)
}

func main() {
	os.Remove("./demo.db") // remove existing db to start fresh

	db, err := sql.Open("sqlite3", "./demo.db") // open a new sqlite3 database connection
	defer db.Close()                            // do not close connection until main() returns
	createSchema(db)                            // create a fresh contact schema

	if err != nil {
		log.Fatal("cannot find demo database")
	}

	repo = ContactRepository{DatabaseConnection: db} // initialize the contact repository

	// an instance of Contact
	contact1 := Contact{
		Name:        "Bob",
		JobTitle:    "Business Analyst",
		PhoneNumber: "123-456-7890",
	}

	// anothe instance
	contact2 := Contact{
		Name:        "Alice",
		JobTitle:    "Chief Executive Officer",
		PhoneNumber: "098-765-4321",
	}

	repo.AddContact(&contact1) // add a contact to the repo
	repo.AddContact(&contact2) // add another contact

	http.HandleFunc("/contacts", getContactHandler) // set route
	http.ListenAndServe(":8000", nil)               // start the server
}
