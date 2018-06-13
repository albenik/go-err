package errx

import (
	"path/filepath"
	"runtime"
	"strings"
)

func getCallerInfo(skip int) (string, int) {
	pc, file, line, _ := runtime.Caller(skip)
	fn := runtime.FuncForPC(pc).Name()
	if dot := strings.LastIndex(fn, "."); dot > -1 {
		return fn[:dot] + "/" + filepath.Base(file), line
	}
	return fn, line
}
