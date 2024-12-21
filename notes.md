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


