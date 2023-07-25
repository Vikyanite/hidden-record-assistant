package util

import (
	md52 "crypto/md5"
	"fmt"
)

func MD5(name, dq string) string {
	return fmt.Sprintf("%x", md52.Sum([]byte("name="+name+"dq-"+dq)))
}
