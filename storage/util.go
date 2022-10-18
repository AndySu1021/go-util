package storage

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func getFileName(filename string) string {
	rand.Seed(time.Now().UnixNano())
	tmpName := strings.Split(filename, ".")
	ext := tmpName[len(tmpName)-1]
	return fmt.Sprintf("%d.%s", time.Now().UnixMilli()*1000+int64(rand.Intn(999)+1), ext)
}

func getUrl(baseUrl, path, filename string) (string, string) {
	// url with domain
	builder := strings.Builder{}
	builder.WriteString(strings.TrimRight(baseUrl, "/"))
	builder.WriteString("/")
	builder.WriteString(strings.Trim(path, "/"))
	builder.WriteString("/")
	builder.WriteString(filename)
	urlWithDomain := builder.String()

	// url without domain
	builder = strings.Builder{}
	builder.WriteString(strings.Trim(path, "/"))
	builder.WriteString("/")
	builder.WriteString(filename)

	return urlWithDomain, builder.String()
}
