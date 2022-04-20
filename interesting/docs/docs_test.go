package docs_test

import "github.com/tjgurwara99/goteach/interesting/docs"

func ExampleSimplePrinter() {
	docs.SimplePrinter("Something %s", "here")
	// Output:
	// Something here
}
