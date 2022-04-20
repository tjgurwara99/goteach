package interesting

import (
	"strconv"
	"time"
)

var someDependency = func() int {
	<-time.After(20 * time.Second)
	return 20
}

func ExportedAPI() string {
	return "Hello there. It took " + strconv.Itoa(someDependency()) + " to complete"
}
