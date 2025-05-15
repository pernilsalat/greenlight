package validator

import "regexp"

var (
	EmailRX = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
)

// TODO: change name to T ¿?
type Validator struct {
	Errors map[string]string
}

func New() *Validator {
	return &Validator{Errors: make(map[string]string)}
}

func (v *Validator) Valid() bool {
	return len(v.Errors) == 0
}

func (v *Validator) AddError(key, message string) {
	if _, has := v.Errors[key]; !has {
		v.Errors[key] = message
	}
}

func (v *Validator) Check(predicate bool, key, message string) {
	if !predicate {
		v.AddError(key, message)
	}
}

func In[T comparable](value T, slice ...T) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func Matches(value string, regexp *regexp.Regexp) bool {
	return regexp.MatchString(value)
}

func Unique[T comparable](slice []T) bool {
	seen := make(map[T]bool)

	for _, v := range slice {
		seen[v] = true
	}

	return len(seen) == len(slice)
}

// TASK: add all conditions
func NotBlank(value any) bool {
	return value != nil && value != 0 && value != ""
}

func GTE(value, min int) bool {
	return value >= min
}

func STE(value, max int) bool {
	return value <= max
}

func Between(value, min, max int) bool {
	return value >= min && value <= max
}
