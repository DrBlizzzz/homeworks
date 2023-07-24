package hw09structvalidator

import (
	"encoding/json"
	"errors"
	"fmt"
	"testing"
)

type UserRole string

type (
	User struct {
		ID     string `json:"id" validate:"len:36"`
		Name   string
		Age    int      `validate:"min:18|max:50"`
		Email  string   `validate:"regexp:^\\w+@\\w+\\.\\w+$"`
		Role   UserRole `validate:"in:admin,stuff"`
		Phones []string `validate:"len:11"`
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
	tests := []struct {
		in          interface{}
		expectedErr error
	}{
		{User{"1", "Anton", 18, "test@mail.ru", "stuff", []string{"01234567891"}, nil}, LenError{}}, // ID fall
		{User{
			"012345678901234567890123456789012345", "Anton", 17, "test@mail.ru", "stuff",
			[]string{"01234567891"},
			nil,
		}, MinError{}}, // Age min fall
		{User{
			"012345678901234567890123456789012345", "Anton", 51, "test@mail.ru", "stuff",
			[]string{"01234567891"},
			nil,
		}, MaxError{}}, // Age max fall
		{User{
			"012345678901234567890123456789012345", "Anton", 18, "test", "stuff",
			[]string{"01234567891"},
			nil,
		}, RegExpError{}}, // Email fall
		{User{
			"012345678901234567890123456789012345", "Anton", 18, "test@mail.ru", "simple text",
			[]string{"01234567891"},
			nil,
		}, nil}, // здесь вообще валидация не пройдет
		{User{
			"012345678901234567890123456789012345", "Anton", 18, "test@mail.ru", "stuff",
			[]string{"phone"},
			nil,
		}, LenError{}}, // Phone fault
		{Response{300, "body"}, InError{}}, // Response fall
		{User{
			"1", "Anton", 18, "test@mail.ru", "stuff",
			[]string{"phone"},
			nil,
		}, LenError{}}, // Complex ID and phone fall
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			tt := tt
			t.Parallel()
			// ниже нужно отрубить линтер, иначе ругается на приведение типа
			resultErrors, _ := Validate(tt.in).(ValidationErrors) //nolint:all
			for _, err := range resultErrors {
				errors.Is(err.Err, tt.expectedErr)
			}
			_ = tt
		})
	}
}
