package finalizer

import (
	"fmt"
	"syscall"
)

func Write(typ string, f string) {
	s := fmt.Sprintf("%s:%s\n", typ, f)
	syscall.Write(9, []byte(s))
}
