package main

import(
	"fmt"
	mrand "math/rand"
	"time"
	"errors"
	"strings"
	"os"
	"path"
	"path/filepath"


)
var (
	rnd = mrand.NewSource(time.Now().UnixNano())
	// ErrNotImplemented used to return errors for functions not implemented
	ErrNotImplemented = errors.New("NOT YET IMPLEMENTED")
)


const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

// RandomString returns a random string
func RandomString(n int) string {
	b := make([]byte, n)

	for i, cache, remain := n-1, rnd.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = rnd.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}



// RemoveQuotes removes outer quotes from a string if necessary
func RemoveQuotes(str string) string {
	if str == "" {
		return str
	}
	if (strings.HasPrefix(str, "'") && strings.HasSuffix(str, "'")) ||
		(strings.HasPrefix(str, "\"") && strings.HasSuffix(str, "\"")) {
		str = str[1 : len(str)-1]
	}
	return str
}


// MakeFileAbs makes 'file' absolute relative to 'dir' if not already absolute
func MakeFileAbs(file, dir string) (string, error) {
	if file == "" {
		return "", nil
	}
	if filepath.IsAbs(file) {
		return file, nil
	}
	path, err := filepath.Abs(filepath.Join(dir, file))
	if err != nil {
		str := fmt.Sprint("Failed making %s absolute based on %s", file, dir)
		strArray := []string{err.Error(),str}
		errs := strings.Join(strArray, " ")
		return "", errors.New(errs)
	}
	return path, nil
}


func main() {
	fmt.Println(RandomString(50))

	fmt.Println(("'hello world'"))
	fmt.Println(("\"hello world\""))
	fmt.Println(RemoveQuotes("'hello world'"))

	dirpath := path.Join(os.Getenv("HOME"), ".fabric-ca-client", "fabric-ca-server-config.yaml")
	fmt.Println(dirpath)

	arrayStr := []string{"hello","world"}
	onlyStr := strings.Join(arrayStr," ")
	fmt.Println(onlyStr)


	fmt.Println(os.Getenv("PWD"))
	s ,_:= MakeFileAbs("main",os.Getenv("PWD"))
	fmt.Println(s)
}

