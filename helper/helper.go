package helper

import (
	"context"
	"crypto/md5"
	"fmt"
	"runtime"
)

func Recover(ctx context.Context) {
	if r := recover(); r != nil {
		var msg string
		for i := 0; ; i++ {
			_, file, line, ok := runtime.Caller(i)
			if !ok {
				break
			}
			msg += fmt.Sprintf("%s:%d\n", file, line)
		}
		logger.Logger.Errorf("%s\n↧↧↧↧↧↧ PANIC ↧↧↧↧↧↧\n%s↥↥↥↥↥↥ PANIC ↥↥↥↥↥↥", r, msg)
	}
}

func EncryptPassword(password, salt string) string {
	if salt == "" {
		salt = "my-salt"
	}
	return fmt.Sprintf("%x", md5.Sum([]byte(salt+password)))
}
