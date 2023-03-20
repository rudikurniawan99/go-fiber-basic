package validator

import (
	"regexp"
	"testing"
)

func TestValidator(t *testing.T) {
	// pattern := `^(?=.*\[A-Z])(?=.*\[a-z])(?=.*\d)(?=.*[#?!@\$%\^&\*\-])[A-Za-z\d#?!@\$%\^&\*\-]{8,}$`
	// fix bug => the go version doesn't support lookahead pattern
	pattern := `^[A-Za-z0-9#?!@\$%\^&\*\-]{8,}$`

	tests := []struct {
		title string
		args  string
		want  bool
	}{
		{
			title: `success password`,
			args:  `Kurniawan99??`,
			want:  true,
		},
		{
			title: `match string`,
			args:  `Abcd1234!`,
			want:  true,
		},
	}

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			got := matchRegexString(pattern, test.args)
			if got != test.want {
				t.Errorf("For %q, expected %v but got %v", test.args, test.want, got)
			}
		})
	}
}

type Identity struct {
	No int
}

func TestValueConverter(t *testing.T) {
	t.Run("test interface convert to struct", func(t *testing.T) {
		result := testValueConvert(89).(Identity)
		t.Log(result)
	})
}

func testValueConvert(param interface{}) interface{} {
	return Identity{
		No: param.(int),
	}
}

func matchRegexString(pattern, text string) bool {
	re, err := regexp.Compile(pattern)
	if err != nil {
		return false
	}
	return re.MatchString(text)
}
