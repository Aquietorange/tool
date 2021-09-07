package tuuid

import (
	"fmt"
	"testing"
)

func Test_NewUUID(t *testing.T) {
	uuid := NewUUID()
	//uuid, _ = ParseString(uuid.String())
	fmt.Println(uuid.String())
}
