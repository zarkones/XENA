package sec

import (
	"fmt"
	"testing"
)

func TestFindSecrets(t *testing.T) {
	data := `asdaasdas;dasdasf;asfsdfsdfsdf;sdfdsfsdf;access_key="supersecretkey";asdasdasfasfasfasf`
	secrets, err := FindSecret(&data)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	if len(secrets) <= 0 {
		t.Log("secret not found")
		t.FailNow()
	}
	fmt.Println(secrets)
}
