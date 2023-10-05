# errors
Go provides a few ways to create error instances, such as these standard functions `errors.New` and `fmt.Errorf`.
Simple as they are, all standard error types lack these capabilities:
- No contextual data can be added. For instance, it's not possible with standard error types to monitor how many errors originate from the database,
  how many errors come from an external dependency.
- No way to know which part of the error is safe for consumption by an external client, usually the ugly parts of the error (why the database malfunctions, for example) are often mixed in with the useful
  message that a system may want to return to its client.

This package aims to solve the problems above with a new type `Error` 
that is fully compatible with the standard `errors` package.

## Examples
In a low level package, we can wrap our error like this
```go
package mydatabase

import "github.com/xamenyap/errors"

func Exec() error {
  // run the sql command
  err := db.Exec("INSERT INTO Products VALUES ...")
  
  // the error message can contain low level details that we want to hide with a friendly message
  return errors.Wrap(err, "cannot create new product", 
	  errors.ContextualOption("package", "mydatabase"), 
	  errors.ContextualOption("query", "INSERT INTO Products VALUES ..."), 
	  )
}
```

Then in your communication layer, log the contextual data of the error, and return only the friendly message part.
```go
package myserver

import (
  "github.com/xamenyap/errors"
  "net/http"
)

func createProduct() http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    if err := mydatabase.Exec(); err != nil {
      contextualErr, ok := err.(*errors.Error)
	  if ok {
		// use your favorite logger to log contextualErr and its contextual data, 
		// then return the friendly message to your client  
		http.Error(w, contextualErr.FriendlyMessage, http.StatusInternalServerError)
        return
      }
	  
	  http.Error(w, err.Error(), http.StatusInternalServerError)
	  return 
    }

    w.WriteHeader(http.StatusCreated)
  }
}
```

