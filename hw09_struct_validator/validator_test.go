package hw09structvalidator

import (
	"encoding/json"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

type UserRole string

// Test the function on different structures and other types.
type (
	User struct {
		ID     string `json:"id" validate:"len:36"`
		Name   string
		Age    int      `validate:"min:18|max:50"`
		Email  string   `validate:"regexp:^\\w+@\\w+\\.\\w+$"`
		Role   UserRole `validate:"in:admin,stuff"`
		Phones []string `validate:"len:11"`
		Limits []int    `validate:"min:3|max:7"`
		meta   json.RawMessage
	}

	App struct {
		Version string `validate:"len:5"`
	}

	Token struct {
		Header    []byte
		Payload   []byte
		Signature []byte
	}

	Response struct {
		Code int    `validate:"in:200,404,500"`
		Body string `json:"omitempty"`
	}
)

func TestValidate(t *testing.T) {
	sec := `{"rawMessage"}`
	byte1, _ := json.Marshal(sec)
	rawMessage := json.RawMessage(byte1)
	user := User{
		"1",
		"John",
		12,
		"parovozzzik@yandex.ru",
		"admin",
		[]string{"9312701311", "89214098329"},
		[]int{3, 4, 55},
		rawMessage,
	}
	userErrors := []ValidationError{
		{"ID", ErrStrLen},
		{"Age", ErrIntMin},
		{"Phones", ErrStrLen},
		{"Limits", ErrIntMax},
	}

	app := App{"1234"}
	appErrors := []ValidationError{{"Version", ErrStrLen}}
	app2 := App{"12345"}

	response := Response{404, "NotFound"}
	response2 := Response{504, "TimeOut"}
	response2Errors := []ValidationError{{"Code", ErrIntIn}}

	tests := []struct {
		in               interface{}
		validationErrors ValidationErrors
	}{
		{in: user, validationErrors: userErrors},
		{in: app, validationErrors: appErrors},
		{in: app2, validationErrors: nil},
		{in: response, validationErrors: nil},
		{in: response2, validationErrors: response2Errors},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			tt := tt
			t.Parallel()

			errs := Validate(tt.in)
			require.Equal(t, len(errs), len(tt.validationErrors))

			for i := 0; i < len(tt.validationErrors); i++ {
				require.True(t, errors.Is(tt.validationErrors[i].Err, errs[i].Err))
			}

			require.Equal(t, errs, tt.validationErrors)

			_ = tt
		})
	}
}
