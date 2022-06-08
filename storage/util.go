package storage

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func getFileName(filename string) string {
	tmpName := strings.Split(filename, ".")
	ext := tmpName[len(tmpName)-1]
	return fmt.Sprintf("%d%d.%s", time.Now().Unix(), rand.Intn(999)+1, ext)
}

func getUrl(baseUrl string, filename string) string {
	builder := strings.Builder{}
	builder.WriteString(strings.TrimRight(baseUrl, "/"))
	builder.WriteString("/")
	builder.WriteString(filename)
	return builder.String()
}
