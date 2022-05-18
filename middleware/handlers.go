package middleware

import (
	"database/sql"
	"encoding/json" // package to encode and decode the json into struct and vice versa
	"fmt"           // models package where schema is defined
	"log"
	"net/http" // used to access the request and response object of the api

	"github.com/meredsa01/go-mockbuster/models"

	"github.com/gorilla/mux" // used to get the params from the route
	_ "github.com/lib/pq"    // postgres golang driver
)

// response format
type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

const (
	host     = "localhost"
	port     = 5555
	user     = "postgres"
	password = "postgres"
	dbname   = "dvdrental"
	//DB_DSN = "postgres://postgres:postgres@localhost:5432/dvdrental?sslmode=disable"
)

// create connection with postgres db
func createConnection() *sql.DB {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Open the connection
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	//defer db.Close()

	// check the connection
	err = db.Ping()
	CheckError(err)

	fmt.Println("Successfully connected!")
	// return the connection
	return db
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

// GetAllFilms will return all the films
func GetAllFilms(w http.ResponseWriter, r *http.Request) {

	// get all the films in the db
	films, err := getAllFilms()

	if err != nil {
		log.Fatalf("Unable to get all films. %v", err)
	}

	// send all the users as response
	json.NewEncoder(w).Encode(films)
}

// GetFilmsByTitle will return films with given partial titles
func GetFilmsByTitle(w http.ResponseWriter, r *http.Request) {
	// get the title from the request params, key is "title"
	params := mux.Vars(r)

	title := params["title"]

	// call the getUser function with user id to retrieve a single user
	films, err := getFilmsByTitle(string(title))

	if err != nil {
		log.Fatalf("Unable to get user. %v", err)
	}

	// send the response
	json.NewEncoder(w).Encode(films)
}

// // UpdateUser update user's detail in the postgres db
// func UpdateUser(w http.ResponseWriter, r *http.Request) {

// 	// get the userid from the request params, key is "id"
// 	params := mux.Vars(r)

// 	// convert the id type from string to int
// 	id, err := strconv.Atoi(params["id"])

// 	if err != nil {
// 		log.Fatalf("Unable to convert the string into int.  %v", err)
// 	}

// 	// create an empty user of type models.User
// 	var user models.User

// 	// decode the json request to user
// 	err = json.NewDecoder(r.Body).Decode(&user)

// 	if err != nil {
// 		log.Fatalf("Unable to decode the request body.  %v", err)
// 	}

// 	// call update user to update the user
// 	updatedRows := updateUser(int64(id), user)

// 	// format the message string
// 	msg := fmt.Sprintf("User updated successfully. Total rows/record affected %v", updatedRows)

// 	// format the response message
// 	res := response{
// 		ID:      int64(id),
// 		Message: msg,
// 	}

// 	// send the response
// 	json.NewEncoder(w).Encode(res)
// }

// // DeleteUser delete user's detail in the postgres db
// func DeleteUser(w http.ResponseWriter, r *http.Request) {

// 	// get the userid from the request params, key is "id"
// 	params := mux.Vars(r)

// 	// convert the id in string to int
// 	id, err := strconv.Atoi(params["id"])

// 	if err != nil {
// 		log.Fatalf("Unable to convert the string into int.  %v", err)
// 	}

// 	// call the deleteUser, convert the int to int64
// 	deletedRows := deleteUser(int64(id))

// 	// format the message string
// 	msg := fmt.Sprintf("User updated successfully. Total rows/record affected %v", deletedRows)

// 	// format the reponse message
// 	res := response{
// 		ID:      int64(id),
// 		Message: msg,
// 	}

// 	// send the response
// 	json.NewEncoder(w).Encode(res)
// }

// //------------------------- handler functions ----------------
// // insert one user in the DB
// func insertUser(user models.User) int64 {

// 	// create the postgres db connection
// 	db := createConnection()

// 	// close the db connection
// 	defer db.Close()

// 	// create the insert sql query
// 	// returning userid will return the id of the inserted user
// 	sqlStatement := `INSERT INTO users (name, location, age) VALUES ($1, $2, $3) RETURNING userid`

// 	// the inserted id will store in this id
// 	var id int64

// 	// execute the sql statement
// 	// Scan function will save the insert id in the id
// 	err := db.QueryRow(sqlStatement, user.Name, user.Location, user.Age).Scan(&id)

// 	if err != nil {
// 		log.Fatalf("Unable to execute the query. %v", err)
// 	}

// 	fmt.Printf("Inserted a single record %v", id)

// 	// return the inserted id
// 	return id
// }

// get all films from the DB
func getAllFilms() ([]models.Film, error) {
	// create the postgres db connection
	db := createConnection()

	// close the db connection
	//defer db.Close()

	// create an array of models.Film type
	var films []models.Film

	// create the select sql query
	sqlStatement := `SELECT * FROM "film"`

	// execute the sql statement
	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// close the statement
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		// create a film of models.Film type
		var film models.Film

		// unmarshal the row object to film
		err = rows.Scan(&film.Film_id, &film.Title, &film.Description, &film.Release_year,
			&film.Language_id, &film.Rental_duration, &film.Rental_rate, &film.Length,
			&film.Replacement_cost, &film.Rating, &film.Last_update, &film.Special_features,
			&film.Fulltext)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		// append the film in the films slice
		films = append(films, film)

	}

	// return empty film on error
	return films, err
}

// get films from the DB by title
func getFilmsByTitle(title string) ([]models.Film, error) {
	// create the postgres db connection
	db := createConnection()

	// close the db connection
	//defer db.Close()

	// create an array of models.Film type
	var films []models.Film

	// create the select sql query
	sqlStatement := `SELECT * FROM "film" WHERE title LIKE '%$1%'`

	// execute the sql statement
	rows, err := db.Query(sqlStatement, title)

	// iterate over the rows
	for rows.Next() {
		// create a film of models.Film type
		var film models.Film
		// unmarshal the row object to film
		err = rows.Scan(&film.Film_id, &film.Title, &film.Description, &film.Release_year,
			&film.Language_id, &film.Rental_duration, &film.Rental_rate, &film.Length,
			&film.Replacement_cost, &film.Rating, &film.Last_update, &film.Special_features,
			&film.Fulltext)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		// append the film in the films slice
		films = append(films, film)
	}

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return films, nil
	case nil:
		return films, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}

	return films, err
}

// // update user in the DB
// func updateUser(id int64, user models.User) int64 {

// 	// create the postgres db connection
// 	db := createConnection()

// 	// close the db connection
// 	defer db.Close()

// 	// create the update sql query
// 	sqlStatement := `UPDATE users SET name=$2, location=$3, age=$4 WHERE userid=$1`

// 	// execute the sql statement
// 	res, err := db.Exec(sqlStatement, id, user.Name, user.Location, user.Age)

// 	if err != nil {
// 		log.Fatalf("Unable to execute the query. %v", err)
// 	}

// 	// check how many rows affected
// 	rowsAffected, err := res.RowsAffected()

// 	if err != nil {
// 		log.Fatalf("Error while checking the affected rows. %v", err)
// 	}

// 	fmt.Printf("Total rows/record affected %v", rowsAffected)

// 	return rowsAffected
// }

// // delete user in the DB
// func deleteUser(id int64) int64 {

// 	// create the postgres db connection
// 	db := createConnection()

// 	// close the db connection
// 	defer db.Close()

// 	// create the delete sql query
// 	sqlStatement := `DELETE FROM users WHERE userid=$1`

// 	// execute the sql statement
// 	res, err := db.Exec(sqlStatement, id)

// 	if err != nil {
// 		log.Fatalf("Unable to execute the query. %v", err)
// 	}

// 	// check how many rows affected
// 	rowsAffected, err := res.RowsAffected()

// 	if err != nil {
// 		log.Fatalf("Error while checking the affected rows. %v", err)
// 	}

// 	fmt.Printf("Total rows/record affected %v", rowsAffected)

// 	return rowsAffected
// }
