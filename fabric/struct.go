package main

import (
	"reflect"
	"bytes"
	"fmt"
	"regexp"
	"strings"
)

// URLRegex is the regular expression to check if a value is an URL
var URLRegex = regexp.MustCompile("(ldap|http)s*://(\\S+):(\\S+)@")


const SecretTag = "mask"

// GetMaskedURL returns masked URL. It masks username and password from the URL
// if present
func GetMaskedURL(url string) string {
	matches := URLRegex.FindStringSubmatch(url)

	// If there is a match, there should be four entries: 1 for
	// the match and 3 for submatches
	if len(matches) == 4 {
		matchIdxs := URLRegex.FindStringSubmatchIndex(url)
		matchStr := url[matchIdxs[0]:matchIdxs[1]]
		for idx := 2; idx < len(matches); idx++ {
			if matches[idx] != "" {
				matchStr = strings.Replace(matchStr, matches[idx], "****", 1)
			}
		}
		url = url[:matchIdxs[0]] + matchStr + url[matchIdxs[1]:len(url)]
	}
	return url
}

// StructToString converts a struct to a string. If a field
// has a 'secret' tag, it is masked in the returned string
func StructToString(si interface{}) string {
	rval := reflect.ValueOf(si).Elem()
	tipe := rval.Type()

	fmt.Println(rval,tipe,rval.NumField())

	var buffer bytes.Buffer
	buffer.WriteString("{ ")
	for i := 0; i < rval.NumField(); i++ {
		tf := tipe.Field(i)
		if !rval.FieldByName(tf.Name).CanSet() {
			continue // skip unexported fields
		}
		var fStr string
		tagv := tf.Tag.Get(SecretTag)
		if tagv == "password" || tagv == "username" {
			fStr = fmt.Sprintf("%s:**** ", tf.Name)
		} else if tagv == "url" {
			val, ok := rval.Field(i).Interface().(string)
			if ok {
				val = GetMaskedURL(val)
				fStr = fmt.Sprintf("%s:%v ", tf.Name, val)
			} else {
				fStr = fmt.Sprintf("%s:%v ", tf.Name, rval.Field(i).Interface())
			}
		} else {
			fStr = fmt.Sprintf("%s:%v ", tf.Name, rval.Field(i).Interface())
		}
		buffer.WriteString(fStr)
	}
	buffer.WriteString(" }")
	return buffer.String()
}

type configID struct {
	Name string `mask:"username"`
	Addr string `json:"address"`
	Pass string `mask:"password"`
	URL  string `mask:"url"`
	ID   int    `mask:"url"`
}

func main() {
	a := configID{"lhy","beijing","passwd","http://www.baidu.com", 32}
	fmt.Println(StructToString(&a))
}

//struct 转为string