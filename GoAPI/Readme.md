#simple APi to  manage books



#API :

Request : get all book
Methode : GET
Url     : http://localhost:8000/api/books


Request : get a book
Method  : GET
Url     : http://localhost:8000/api/books/{id}


Request : get a book
Method  : GET
Url     : http://localhost:8000/api/books/{id}


Request : create a book
Method  : POST
Url     : http://localhost:8000/api/book
example body : 
{
  "isbn":"33332",
   "title":"tiga teman",
   "author":{"firstname":"Untung",  "lastname":"Saja"}
}



Request : update a book
Method  : PUT
Url     : http://localhost:8000/api/books/{id}
example body : 
{
  "isbn":"33332",
   "title":"empat teman",
   "author":{"firstname":"Untung",  "lastname":"Saja"}
}


Request : delete a book
Method  : DELETE
Url     : http://localhost:8000/api/books/{id}


