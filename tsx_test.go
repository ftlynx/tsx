package tsx

import (
	"testing"
)

func TestMd5(t *testing.T) {
	t.Log(Md5([]byte("abc")))
	t.Log(Md5([]byte("abc")))

}
