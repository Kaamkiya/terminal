package animate

import (
	"fmt"
	"io"
	"time"
)

func TypeWriter(w io.Writer, waitTime time.Duration, text string) {
	for _, character := range text {
		fmt.Fprint(w, string(character))
		time.Sleep(waitTime)
	}
}
