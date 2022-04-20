package docs

import "fmt"

func SimplePrinter(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}
