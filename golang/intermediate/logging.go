package main

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"log/slog"
)

func main() {

	log.Println("standard logger") // Output: 2023/08/22 10:45:16 standard logger

	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.Println("with micro") // Output: 2023/08/22 10:45:16.904141 with micro

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("with file/line") // Output: 2023/08/22 10:45:16 logging.go:40: with file/line

	mylog := log.New(os.Stdout, "my:", log.LstdFlags)
	mylog.Println("from mylog") // Output: my:2023/08/22 10:45:16 from mylog

	mylog.SetPrefix("ohmy:")
	mylog.Println("from mylog") // Output: ohmy:2023/08/22 10:45:16 from mylog

	var buf bytes.Buffer
	buflog := log.New(&buf, "buf:", log.LstdFlags)
	buflog.Println("hello")

	fmt.Print("from buflog:", buf.String()) // Output: from buflog:buf:2023/08/22 10:45:16 hello

	jsonHandler := slog.NewJSONHandler(os.Stderr, nil)
	myslog := slog.New(jsonHandler)
	myslog.Info("hi there")

	myslog.Info("hello again", "key", "val", "age", 25)
	// Output:
	// {"time":"2023-08-22T10:45:16.904166391-07:00", "level":"INFO", "msg":"hi there"}
	// {"time":"2023-08-22T10:45:16.904178985-07:00", "level":"INFO", "msg":"hello again", "key":"val", "age":25}
}
