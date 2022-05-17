package middleware

import (
	"database/sql"
	"encoding/json" // package to encode and decode the json into struct and vice versa
	"fmt"           // models package where schema is defined
	"github.com/meredsa01/go-mockbuster/models"
	"log"
	"net/http" // used to access the request and response object of the api

	// package used to covert string into int type
	// used to get the params from the route

	_ "github.com/lib/pq" // postgres golang driver
)

// response format
type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

const (
	host     = "localhost"
	port     = 5555 //5432
	user     = "postgres"
	password = "postgres"
	dbname   = "dvdrental"
)

// create connection with postgres db
func createConnection() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=enable",
		host, port, user, password, dbname)

	// Open the connection
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// check the connection
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	// return the connection
	return db
}

// GetAllFilms will return all the films
func GetFilms(w http.ResponseWriter, r *http.Request) {

	// get all the films in the db
	films, err := getAllFilms()

	if err != nil {
		log.Fatalf("Unable to get all films. %v", err)
	}

	// send all the users as response
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

// // get one user from the DB by its userid
// func getUser(id int64) (models.User, error) {
// 	// create the postgres db connection
// 	db := createConnection()

// 	// close the db connection
// 	defer db.Close()

// 	// create a user of models.User type
// 	var user models.User

// 	// create the select sql query
// 	sqlStatement := `SELECT * FROM users WHERE userid=$1`

// 	// execute the sql statement
// 	row := db.QueryRow(sqlStatement, id)

// 	// unmarshal the row object to user
// 	err := row.Scan(&user.ID, &user.Name, &user.Age, &user.Location)

// 	switch err {
// 	case sql.ErrNoRows:
// 		fmt.Println("No rows were returned!")
// 		return user, nil
// 	case nil:
// 		return user, nil
// 	default:
// 		log.Fatalf("Unable to scan the row. %v", err)
// 	}

// 	// return empty user on error
// 	return user, err
// }

// get all films from the DB
func getAllFilms() ([]models.film, error) {
	// create the postgres db connection
	db := createConnection()

	// close the db connection
	defer db.Close()

	var films []models.film

	// create the select sql query
	sqlStatement := `SELECT * FROM film`

	// execute the sql statement
	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// close the statement
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var film models.film

		// unmarshal the row object to film
		err = rows.Scan(&film.film_id, &film.title, &film.description, &film.release_year)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		// append the film in the films slice
		films = append(films, film)

	}

	// return empty film on error
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
