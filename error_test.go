package tsx

import (
	"fmt"
	"testing"
)

func TestErrxErrort(t *testing.T) {
	e := fmt.Errorf("test")
	fmt.Println(ErrxError(e))
}