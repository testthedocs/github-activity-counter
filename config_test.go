package counter

import (
	"testing"
)

func TestDefaultConfigInitializer(t *testing.T) {
	err := defaultConfigInitializer()
	if err != nil {
		t.Error(err)
	}
}