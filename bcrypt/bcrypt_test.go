package bcrypt_test

import (
	"fmt"
	"testing"

	"github.com/goark/gnkf/bcrypt"
)

func TestBCrypt(t *testing.T) {
	s := "password"
	h, err := bcrypt.Hash(s, bcrypt.DefaultCost)
	if err != nil {
		t.Errorf("Hash() error = \"%+v\", want nil.", err)
	}
	fmt.Println(h)
	err = bcrypt.Compare(h, s)
	if err != nil {
		t.Errorf("Compare() is \"%+v\", want nil.", err)
	}
}
