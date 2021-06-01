package convert

import "github.com/pkgng/copit"

// 转换A至B
func Copy(src interface{}, dest interface{}) error {
	return copit.Copy(dest, src)
}
