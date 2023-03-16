package validator

import (
	"regexp"
	"testing"

	assert "github.com/go-playground/assert/v2"
)

func TestValidator(t *testing.T) {
	tests := []struct {
		title string
		args  string
		want  bool
	}{
		{
			title: "all lowercase",
			args:  `kurniawan`,
			want:  false,
		},
		{
			title: `all uppercase`,
			args:  `KURNIAWAN`,
			want:  false,
		},
		{
			title: `right way`,
			args:  `Kurniawan99??`,
			want:  true,
		},
		{
			title: `match string`,
			args:  `Abcd1234!`,
			want:  true,
		},
	}
	pattern := `^(?=.*?[A-Z])(?=.*?[a-z])(?=.*?[0-9])(?=.*?[#?!@$ %^&*-]).{8,}$`

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			assert.Equal(
				t,
				matchRegexString(
					pattern,
					test.args,
				),
				test.want,
			)
		})
	}
}

func matchRegexString(pattern, text string) bool {
	re, err := regexp.Compile(pattern)
	if err != nil {
		return false
	}
	return re.MatchString(text)
}
