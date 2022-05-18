# go-mockbuster

API endpoints examples:

List of all films
http://localhost:8081/api/films


List of films searched by (partial or full) title
http://localhost:8081/api/films/title/run
localhost:8081/api/films/title/airplane Sierra 
which converts to 
http://localhost:8081/api/films/title/airplane%20Sierra


List of films filtered by rating
http://localhost:8081/api/films/rating/pg-13
http://localhost:8081/api/films/rating/R


List of films filtered by category_id
http://localhost:8081/api/films/categoryid/3
or by category
http://localhost:8081/api/films/category/animation
http://localhost:8081/api/films/category/Music


Film details for a given film_id
http://localhost:8081/api/filmdetails/300
