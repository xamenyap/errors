# errors
Go provides a few ways to create error instances, such as these standard functions `errors.New` and `fmt.Errorf`. 
However, none of them allows decorating the error with extra contextual data. Why is adding contextual data important? 
While simple errors are easy to handle, in large applications, 
we may also want to programmatically monitor our errors: how many errors originate from the database, 
how many errors come from an external dependency, etc. 
Furthermore, when working with an error created by `errors.New` and `fmt.Errorf`, 
usually the ugly parts of the error (why the database malfunctions, for example) are often mixed in with the useful
message that a system may want to return to its client.

This package aims to solve the problems above with a new type `Error`, 
allowing adding contextual data in the form of key/value pairs, 
and forming a friendly message that hides the ugly details of an error when necessary. 
The `Error` type is fully compatible with the standard `errors` package.

## Examples

