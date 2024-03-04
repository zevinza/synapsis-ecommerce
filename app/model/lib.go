/*
this file contains local library for package model
all func should be limited for this package only
*/
package model

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/go-openapi/strfmt"
)

// boolInt func convert 0 and 1 any into false and true boolean
func boolInt(b interface{}) *bool {
	var t, f bool = true, false
	s := fmt.Sprint(b)
	if s == "1" {
		return &t
	}
	return &f
}

func now() *strfmt.DateTime {
	now := strfmt.DateTime(time.Now())
	return &now
}

// randomNumber function
func randomNumber(length int) *string {
	charset := "0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	res := string(result)
	return &res
}

func strptr(s string) *string { return &s }
