# Notes

## CH1 - Foundations

 `go run` compiles your code, creates an executable binary in your `/tmp`
directory.

```go

http.NewServeMux() //  function to initialize a new servemux

```


```go

// Add a snippetView handler function
func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific snippet ..."))
}

mux.HandleFunc("/snippet/view", snippetView) // Registers the Handler to  route

```

Servemux sanitize urls

The r.PathValue() method always returns a string value IT IS NOT SANITIZED !

Be carefull of route precedence , the most specific wins, make your route no there is no overlapping routes.

there is no right or wrong way to name your handlers in Go

It’s only possible to call w.WriteHeader() once per response, and after the status code
has been written it can’t be changed.


---------
##### Folder Structure 

cmd : Contains application specific code 

internal : ancillary non-application-specific code used in the project. Reusable code like validation helpers and the SQL database models for the project

ui: user-interface assets used by the application.



## CH2 - Configuration and error handling


For flags defined with flag.Bool(), omitting a value when starting the application is the
same as writing -flag=true. The following two commands are equivalent:

##### Dependency Injection 

What we really want to answer is: how can we make any dependency available to our handlers? 
=> Inject dependencies into your handlers. It makes your code more explicit, less error prone and easier to unit test than if you use global variables.



## CH3 - Database-driven response


 `go mod verify` => verifies that the checksums of the downloaded packages on the machine match the entries in go.sum.

 `go mod download` => this will get an error if there is any mismatch between the packages they are downloading and the checksums in the file.

 To upgrade to the latest available minor or patch release of a package, you can simply run the go get with the -u flag like so : 

 ```
 go get -u github.com/foo/bar
 ```

 For specific version 
 
 ```
 go get -u github.com/foo/bar@v2.0.0
 ```

 Removing a package 


 run 
 
 ```
 go get github.com/foo/bar@none
 ```

 or if all refs are gone from project 

 ``` 
 go mod tidy
 ```

============

The `sql.Open()` function returns a `sql.DB object`. This isn’t a database connection — it’s a pool of many connections. This is an important difference to understand. Go manages the connections in this pool as needed, automatically opening and closing connections
to the database via the driver.


The connection pool is intended to be long-lived. In a web application it’s normal to initialize the connection pool in your `main()` function and then pass the pool to your handlers. You shouldn’t call `sql.Open()` in a short-lived HTTP handler itself — it would be a waste of memory and network resources.



## TODO : 

Understand middlewares.
Understand closure.
Understand pattern for dependence injection