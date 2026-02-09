package validation

import (
	"errors"
	"fmt"
	"net/mail"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Validator func(string) error

func Required() Validator {
	return func(s string) error {
		if strings.TrimSpace(s) == "" {
			return errors.New("required")
		}
		return nil
	}
}

func Optional(v Validator) Validator {
	return func(s string) error {
		if strings.TrimSpace(s) == "" {
			return nil
		}
		return v(s)
	}
}

func IsInt() Validator {
	return func(s string) error {
		if _, err := strconv.Atoi(strings.TrimSpace(s)); err != nil {
			return fmt.Errorf("must be integer")
		}
		return nil
	}
}

func IsFloat() Validator {
	return func(s string) error {
		if _, err := strconv.ParseFloat(strings.TrimSpace(s), 64); err != nil {
			return fmt.Errorf("must be number")
		}
		return nil
	}
}

func MinLen(n int) Validator {
	return func(s string) error {
		if len(strings.TrimSpace(s)) < n {
			return fmt.Errorf("min length %d", n)
		}
		return nil
	}
}

func MaxLen(n int) Validator {
	return func(s string) error {
		if len(strings.TrimSpace(s)) > n {
			return fmt.Errorf("max length %d", n)
		}
		return nil
	}
}

func Percent() Validator {
	return func(s string) error {
		f, err := strconv.ParseFloat(strings.TrimSpace(s), 64)
		if err != nil {
			return fmt.Errorf("must be number")
		}
		if f < 0 || f > 100 {
			return fmt.Errorf("must be between 0 and 100")
		}
		return nil
	}
}

func IntRange(min, max int) Validator {
	return func(s string) error {
		i, err := strconv.Atoi(strings.TrimSpace(s))
		if err != nil {
			return fmt.Errorf("must be integer")
		}
		if i < min || i > max {
			return fmt.Errorf("must be between %d and %d", min, max)
		}
		return nil
	}
}

func FloatRange(min, max float64) Validator {
	return func(s string) error {
		f, err := strconv.ParseFloat(strings.TrimSpace(s), 64)
		if err != nil {
			return fmt.Errorf("must be number")
		}
		if f < min || f > max {
			return fmt.Errorf("must be between %.2f and %.2f", min, max)
		}
		return nil
	}
}

func OneOfStrict(allowed ...string) Validator {
	set := map[string]struct{}{}
	for _, a := range allowed {
		set[a] = struct{}{}
	}
	return func(s string) error {
		val := strings.TrimSpace(s)
		if val == "" {
			return errors.New("value required")
		}
		if _, ok := set[val]; !ok {
			return fmt.Errorf("invalid value %q", val)
		}
		return nil
	}
}

func Matches(rx *regexp.Regexp) Validator {
	return func(s string) error {
		if !rx.MatchString(strings.TrimSpace(s)) {
			return fmt.Errorf("invalid format")
		}
		return nil
	}
}

func IsEmail() Validator {
	return func(s string) error {
		if _, err := mail.ParseAddress(strings.TrimSpace(s)); err != nil {
			return fmt.Errorf("invalid email")
		}
		return nil
	}
}

func DateFlexible() Validator {
	formats := []string{time.RFC3339, "2006-01-02"}
	return func(s string) error {
		val := strings.TrimSpace(s)
		for _, f := range formats {
			if _, err := time.Parse(f, val); err == nil {
				return nil
			}
		}
		return fmt.Errorf("invalid date")
	}
}

// CSVList: comma-separated non-empty tokens
func CSVList() Validator {
	return func(s string) error {
		val := strings.TrimSpace(s)
		if val == "" {
			return errors.New("list cannot be empty")
		}
		parts := strings.Split(val, ",")
		for _, p := range parts {
			if strings.TrimSpace(p) == "" {
				return errors.New("empty value in list")
			}
		}
		return nil
	}
}

// CSVListOfRequired enforces a non-empty comma-separated list
// whose values must be from the allowed set (case-sensitive).
func CSVListOfRequired(allowed ...string) Validator {
	allowedSet := map[string]struct{}{}
	for _, a := range allowed {
		allowedSet[a] = struct{}{}
	}

	return func(s string) error {
		val := strings.TrimSpace(s)
		if val == "" {
			return errors.New("list cannot be empty")
		}

		parts := strings.Split(val, ",")
		for _, p := range parts {
			item := strings.TrimSpace(p)
			if item == "" {
				return errors.New("empty value in list")
			}
			if _, ok := allowedSet[item]; !ok {
				return fmt.Errorf("invalid value %q (allowed: %v)", item, allowed)
			}
		}
		return nil
	}
}

// CSVListOfOptional allows empty; if non-empty, applies CSVListOfRequired.
func CSVListOfOptional(allowed ...string) Validator {
	required := CSVListOfRequired(allowed...)
	return func(s string) error {
		if strings.TrimSpace(s) == "" {
			return nil
		}
		return required(s)
	}
}
