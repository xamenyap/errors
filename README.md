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

