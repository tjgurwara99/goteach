package interesting

import (
	"testing"
	"time"
)

// TestExportedAPI I have a Love-Hate relationship with this pattern.
// I don't like the mutability of functions but at the same time, it makes
// the services way to decouple systems with dependencies.
func TestExportedAPI(t *testing.T) {
	oldFunc := someDependency
	defer func() { someDependency = oldFunc }()
	someDependency = func() int {
		<-time.After(1 * time.Second)
		return 1
	}
	str := ExportedAPI()
	if str != "Hello there. It took 1 to complete" {
		t.Errorf("Test failed: expected %d, got %d", 1, 20)
	}
}
