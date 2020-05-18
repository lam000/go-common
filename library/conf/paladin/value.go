package paladin

import (
	"encoding"
	"encoding/json"
	"errors"
	"github.com/BurntSushi/toml"
	"gopkg.in/yaml.v2"
	"time"
)

var (
	ErrNotExist      = errors.New("paladin: value key not exist")
	ErrTypeAssertion = errors.New("paladin: value type assertion not match")
)

type Value struct {
	val   interface{}
	slice interface{}
	raw   string
}

func (v *Value) Bool() (bool, error) {
	if v.val == nil {
		return false, ErrNotExist
	}
	b, ok := v.val.(bool)
	if !ok {
		return false, ErrTypeAssertion
	}

	return b, nil
}

func (v *Value) Int() (int, error) {
	i, err := v.Int64()
	return int(i), err
}

func (v *Value) Int32() (int32, error) {
	i, err := v.Int64()
	return int32(i), err
}

func (v *Value) Int64() (int64, error) {
	if v.val == nil {
		return 0, ErrNotExist
	}

	i, ok := v.val.(int64)
	if !ok {
		return 0, ErrTypeAssertion
	}

	return i, nil
}

func (v *Value) Float64() (float64, error) {
	if v.val == nil {
		return 0.0, ErrNotExist
	}

	f, ok := v.val.(float64)
	if !ok {
		return 0.0, ErrTypeAssertion
	}

	return f, nil
}

func (v *Value) Float32() (float32, error) {
	f, err := v.Float64()
	if err != nil {
		return 0.0, err
	}

	return float32(f), nil
}

func (v *Value) String() (string, error) {
	if v.val == nil {
		return "", ErrNotExist
	}

	s, ok := v.val.(string)
	if !ok {
		return "", ErrTypeAssertion
	}

	return s, nil
}

func (v *Value) Duration() (time.Duration, error) {
	s, err := v.String()
	if err != nil {
		return time.Duration(0), err
	}

	return time.ParseDuration(s)
}

func (v *Value) Raw() (string, error) {
	if v.val == nil {
		return "", ErrNotExist
	}

	return v.raw, nil
}

func (v *Value) Unmarshal(un encoding.TextUnmarshaler) error {
	text, err := v.Raw()
	if err != nil {
		return err
	}

	return un.UnmarshalText([]byte(text))
}

func (v *Value) UnmarshalTOML(dst interface{}) error {
	text, err := v.Raw()
	if err != nil {
		return err
	}

	return toml.Unmarshal([]byte(text), dst)
}

func (v *Value) UnmarshalJSON(dst interface{}) error {
	text, err := v.Raw()
	if err != nil {
		return err
	}

	return json.Unmarshal([]byte(text), dst)
}

func (v *Value) UnmarshalYAML(dst interface{}) error {
	text, err := v.Raw()
	if err != nil {
		return err
	}

	return yaml.Unmarshal([]byte(text), dst)
}
