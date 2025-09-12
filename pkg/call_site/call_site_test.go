package call_site

import (
	"fmt"
	"testing"
)

func TestGetCallSite(t *testing.T) {
	callSite := GetCallSite()
	for _, part := range callSite {
		fmt.Println(part)
	}
}
