package tsx

import (
	"fmt"
	"os"
	"path/filepath"
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