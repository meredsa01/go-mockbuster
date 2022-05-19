package middleware

import (
	"database/sql"
	"encoding/json" // package to encode and decode the json into struct and vice versa
	"fmt"           // models package where schema is defined
	"net/http"      // used to access the request and response object of the api

	// used for string manipulation
	"strings"

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
		fmt.Printf("Unable to get all films. %v", err)
	}

	// send all the users as response
	json.NewEncoder(w).Encode(films)
}

// GetFilmsByTitle will return films with given full or partial titles
func GetFilmsByTitle(w http.ResponseWriter, r *http.Request) {
	// get the title from the request params, key is "title"
	params := mux.Vars(r)

	title := params["title"]

	// call the getUser function with user id to retrieve a single user
	films, err := getFilmsByTitle(string(title))

	if err != nil {
		fmt.Printf("Unable to get films. %v", err)
	}

	// send the response
	json.NewEncoder(w).Encode(films)
}

// GetFilmsByRating will return films with given rating
func GetFilmsByRating(w http.ResponseWriter, r *http.Request) {
	// get the rating from the request params, key is "rating"
	params := mux.Vars(r)

	rating := params["rating"]

	// call the getFilmsByRating function with rating to retrieve films
	films, err := getFilmsByRating(string(rating))

	if err != nil {
		fmt.Printf("Unable to get films. %v", err)
	}

	// send the response
	json.NewEncoder(w).Encode(films)
}

// GetFilmsByCategoryID will return films with given category id
func GetFilmsByCategoryID(w http.ResponseWriter, r *http.Request) {
	// get the category id from the request params, key is "id"
	params := mux.Vars(r)

	id := params["id"]

	// call the getFilmsByCategoryID function with id to retrieve films
	films, err := getFilmsByCategoryID(string(id))

	if err != nil {
		fmt.Printf("Unable to get films. %v", err)
	}

	// send the response
	json.NewEncoder(w).Encode(films)
}

// GetFilmsByCategory will return films with given category
func GetFilmsByCategory(w http.ResponseWriter, r *http.Request) {
	// get the category from the request params, key is "category"
	params := mux.Vars(r)

	category := params["category"]

	// call the getFilmsByCategory function with category to retrieve films
	films, err := getFilmsByCategory(string(category))

	if err != nil {
		fmt.Printf("Unable to get films. %v", err)
	}

	// send the response
	json.NewEncoder(w).Encode(films)
}

// GetFilmDetails will return a film's details
func GetFilmDetails(w http.ResponseWriter, r *http.Request) {
	// get the title from the request params, key is "title"
	params := mux.Vars(r)

	id := params["id"]

	// call the getFilmDetails function with film id to retrieve a single film's details
	filmdetails, err := getFilmDetails(string(id))

	if err != nil {
		fmt.Printf("Unable to get film details. %v", err)
	}

	// send the response
	json.NewEncoder(w).Encode(filmdetails)
}

// Create customer film comment
func InsertComment(w http.ResponseWriter, r *http.Request) {
	// get the comment from the request params
	params := mux.Vars(r)
	commentjson := params["comment"]

	var err error
	var comment models.Film_comment

	err = json.Unmarshal([]byte(commentjson), &comment)
	if err != nil {
		fmt.Printf("Unable to get comment. %v", err)
	}

	// call the insertComment function with comment from parameters
	comment.Comment_id, err = insertComment(comment)

	if err != nil {
		fmt.Printf("Error inserting comment. %v", err)
	}

	// format a response object
	res := response{
		ID:      comment.Comment_id,
		Message: "Comment inserted successfully",
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}

// GetCommentsByFilmID will return comment with given film id
func GetCommentsByFilmID(w http.ResponseWriter, r *http.Request) {
	// get the film id from the request params, key is "id"
	params := mux.Vars(r)

	id := params["film_id"]

	// call the getCommentsByFilmID function with id to retrieve comments
	films, err := getCommentsByFilmID(string(id))

	if err != nil {
		fmt.Printf("Unable to get comments. %v", err)
	}

	// send the response
	json.NewEncoder(w).Encode(films)
}

// get all films from the DB
func getAllFilms() ([]models.Film, error) {
	// create the postgres db connection
	db := createConnection()

	// close the db connection
	defer db.Close()

	// create an array of models.Film type
	var films []models.Film

	// create the select sql query
	sqlStatement := `SELECT * FROM "film"`

	// execute the sql statement
	rows, err := db.Query(sqlStatement)

	if err != nil {
		fmt.Printf("Unable to execute the query. %v", err)
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
			fmt.Printf("Unable to scan the row. %v", err)
		}

		// append the film in the films slice
		films = append(films, film)

	}

	// return empty film on error
	return films, err
}

// get films from the DB by partial title
func getFilmsByTitle(title string) ([]models.Film, error) {
	// create the postgres db connection
	db := createConnection()

	// close the db connection
	defer db.Close()

	// create an array of models.Film type
	var films []models.Film

	title = strings.ToLower(title)

	// create the select sql query
	sqlStatement := `SELECT * FROM "film" WHERE LOWER(title) LIKE '%' || $1 || '%'`

	// execute the sql statement
	rows, err := db.Query(sqlStatement, title)

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
			fmt.Printf("Unable to scan the row. %v", err)
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
		fmt.Printf("Unable to scan the row. %v", err)
	}

	return films, err
}

// get films from the DB by rating
func getFilmsByRating(rating string) ([]models.Film, error) {
	// create the postgres db connection
	db := createConnection()

	// close the db connection
	defer db.Close()

	// create an array of models.Film type
	var films []models.Film

	// ratings are stored as uppercase
	rating = strings.ToUpper(rating)

	// create the select sql query
	sqlStatement := `SELECT * FROM "film" f WHERE rating = $1`

	// execute the sql statement
	rows, err := db.Query(sqlStatement, rating)

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
			fmt.Printf("Unable to scan the row. %v", err)
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
		fmt.Printf("Unable to scan the row. %v", err)
	}

	return films, err
}

// get films from the DB by category id
func getFilmsByCategoryID(id string) ([]models.Film, error) {
	// create the postgres db connection
	db := createConnection()

	// close the db connection
	defer db.Close()

	// create an array of models.Film type
	var films []models.Film

	// create the select sql query
	sqlStatement := `SELECT f.* FROM "film_category" fc LEFT JOIN "film" f ON fc.film_id = f.film_id WHERE fc.category_id = $1`

	// execute the sql statement
	rows, err := db.Query(sqlStatement, id)

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
			fmt.Printf("Unable to scan the row. %v", err)
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
		fmt.Printf("Unable to scan the row. %v", err)
	}

	return films, err
}

// get films from the DB by category id
func getFilmsByCategory(category string) ([]models.Film, error) {
	// create the postgres db connection
	db := createConnection()

	// close the db connection
	defer db.Close()

	// create an array of models.Film type
	var films []models.Film

	// remove case for category
	category = strings.ToLower(category)

	// create the select sql query
	sqlStatement := `SELECT f.* FROM "film_category" fc LEFT JOIN "category" c ON fc.category_id = c.category_id LEFT JOIN "film" f ON fc.film_id = f.film_id WHERE LOWER(c.name) = $1`

	// execute the sql statement
	rows, err := db.Query(sqlStatement, category)

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
			fmt.Printf("Unable to scan the row. %v", err)
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
		fmt.Printf("Unable to scan the row. %v", err)
	}

	return films, err
}

// get film details from the DB by id
func getFilmDetails(id string) (models.Film_details, error) {
	// create the postgres db connection
	db := createConnection()

	// close the db connection
	defer db.Close()

	// create a variable of models.FilmDetails type
	var filmdetails models.Film_details

	// create the select sql query
	sqlStatement := `SELECT * FROM "film" WHERE film_id = $1`

	// execute the sql statement
	row := db.QueryRow(sqlStatement, id)

	// create a film of models.Film type
	var film models.Film
	// unmarshal the row object to film
	err := row.Scan(&film.Film_id, &film.Title, &film.Description, &film.Release_year,
		&film.Language_id, &film.Rental_duration, &film.Rental_rate, &film.Length,
		&film.Replacement_cost, &film.Rating, &film.Last_update, &film.Special_features,
		&film.Fulltext)

	if err != nil {
		fmt.Printf("Unable to scan the row. %v", err)
	}

	// set the film base
	filmdetails.FilmBase = film

	// create the select sql query
	sqlStatement = `SELECT name language FROM "language" WHERE language_id = $1`

	// execute the sql statement
	row = db.QueryRow(sqlStatement, film.Language_id)

	// create a language var
	var language string
	// unmarshal the row object to film
	err = row.Scan(&language)

	if err != nil {
		fmt.Printf("Unable to scan the row. %v", err)
	}

	// set the film base
	filmdetails.Language = language

	// create the select sql query
	sqlStatement = `SELECT c.category_id, c.name FROM "film_category" f LEFT JOIN "category" c ON f.category_id = c.category_id WHERE f.film_id = $1`

	// execute the sql statement
	rows, err := db.Query(sqlStatement, id)

	if err != nil {
		fmt.Printf("Unable to execute the query. %v", err)
	}

	// close the statement
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var category models.Film_category

		// unmarshal the row object to user
		err = rows.Scan(&category.Category_id, &category.Name)

		if err != nil {
			fmt.Printf("Unable to scan the row. %v", err)
		}

		// append the category in the film details slice
		filmdetails.Categories = append(filmdetails.Categories, category)
	}

	// create the select sql query
	sqlStatement = `SELECT a.actor_id, a.first_name, a.last_name FROM "film_actor" f LEFT JOIN "actor" a ON f.actor_id = a.actor_id WHERE f.film_id = $1`

	// execute the sql statement
	rows, err = db.Query(sqlStatement, id)

	if err != nil {
		fmt.Printf("Unable to execute the query. %v", err)
	}

	// close the statement
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var actor models.Film_actor

		// unmarshal the row object to user
		err = rows.Scan(&actor.Actor_id, &actor.First_name, &actor.Last_name)

		if err != nil {
			fmt.Printf("Unable to scan the row. %v", err)
		}

		// append the category in the film details slice
		filmdetails.Actors = append(filmdetails.Actors, actor)
	}

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return filmdetails, nil
	case nil:
		return filmdetails, nil
	default:
		fmt.Printf("Unable to scan the row. %v", err)
	}

	return filmdetails, err
}

func createCommentTable() {
	// create the postgres db connection
	db := createConnection()

	// close the db connection
	defer db.Close()

	createtable := `CREATE TABLE IF NOT EXISTS film_comment (
		comment_id serial PRIMARY KEY NOT NULL ,
		film_id integer NOT NULL ,
		customer_id integer NOT NULL ,
		comment varchar(1000) NOT NULL ,
		last_update timestamp with time zone DEFAULT CURRENT_TIMESTAMP)`

	stmt, err := db.Prepare(createtable)

	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		fmt.Println(err.Error())
	}
	//db.Close()
}

// insert a customer comment to a film
func insertComment(comment models.Film_comment) (int64, error) {

	createCommentTable()

	// create the postgres db connection
	db := createConnection()

	// close the db connection
	defer db.Close()

	// create the insert sql query
	// returning userid will return the id of the inserted user
	sqlStatement := `INSERT INTO film_comment (film_id, customer_id, comment) VALUES ($1, $2, $3) RETURNING comment_id`

	// the inserted id will store in this id
	var id int64

	// execute the sql statement
	// Scan function will save the insert id in the id
	err := db.QueryRow(sqlStatement, comment.Film_id, comment.Customer_id, comment.Comment).Scan(&id)

	if err != nil {
		fmt.Printf("Unable to execute the query. %v", err)
	}

	fmt.Printf("Inserted a single record %v", id)

	// return the inserted id
	return id, err
}

// get comments from the DB by film id
func getCommentsByFilmID(id string) ([]models.Film_comment, error) {

	createCommentTable()

	// create the postgres db connection
	db := createConnection()

	// close the db connection
	defer db.Close()

	// create an array of models.Film_comment type
	var comments []models.Film_comment

	// create the select sql query
	sqlStatement := `SELECT * FROM film_comment WHERE film_id = $1`

	// execute the sql statement
	rows, err := db.Query(sqlStatement, id)

	if err != nil {
		fmt.Printf("Unable to get comments. %v", err)
	}

	// close the statement
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		// create a comment of models.Film_comment type
		var comment models.Film_comment
		// unmarshal the row object to film
		err = rows.Scan(&comment.Comment_id, &comment.Film_id, &comment.Customer_id,
			&comment.Comment, &comment.Last_update)

		if err != nil {
			fmt.Printf("Unable to scan the row. %v", err)
		}

		// append the comment in the comments slice
		comments = append(comments, comment)
	}

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return comments, nil
	case nil:
		return comments, nil
	default:
		fmt.Printf("Unable to scan the row. %v", err)
	}

	return comments, err
}

