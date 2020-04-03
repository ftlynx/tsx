package tsx

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

//创建文件父目录
func CreateParentDir(dstfile string) error {
	dirname, err := filepath.Abs(filepath.Dir(dstfile))
	if err != nil {
		return err
	}
	return os.MkdirAll(dirname, 0755)
}

func Errx(s string, e error) error{
	return fmt.Errorf("%s %w", s, e)
}

func Millisecond() int64 {
	return time.Now().UnixNano() / (1000 * 1000)
}