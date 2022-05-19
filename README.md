# go-mockbuster

API endpoint examples:

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


Add a customer comment to a film
localhost:8081/api/films/comment/{"comment_id":null,"film_id":1,"customer_id":1,"comment":"\"This\" isn't great!"}
which converts to
http://localhost:8081/api/films/comment/%7B%22comment_id%22:null,%22film_id%22:1,%22customer_id%22:1,%22comment%22:%22This%20isn't%20great!%22%7D
