package paladin

import (
	"bytes"
	"github.com/BurntSushi/toml"
	"github.com/pkg/errors"
	"reflect"
	"strconv"
)

type TOML = Map

func (m *TOML) Set(text string) error {
}

func (m *TOML) UnmarshalText(text []byte) error {
	raws := map[string]interface{}{}
	if err := toml.Unmarshal(text, &raws); err != nil {
		return err
	}
	values := map[string]*Value{}
	for k, v := range raws {
		k = keyNamed(k)
		rv := reflect.ValueOf(v)
		switch rv.Kind() {
		case reflect.Map:
			buf := bytes.NewBuffer(nil)
			err := toml.NewEncoder(buf).Encode(v)
			if err != nil {
				return err
			}
			values[k] = &Value{
				val: v,
				raw: buf.String(),
			}
		case reflect.Bool:
			b := v.(bool)
			values[k] = &Value{
				val: v.(bool),
				raw: strconv.FormatBool(b),
			}
		case reflect.Int64:
			i := v.(int64)
			values[k] = &Value{
				val: i,
				raw: strconv.FormatInt(i, 10),
			}
		case reflect.Float64:
			f := v.(float64)
			values[k] = &Value{
				val: f,
				raw: strconv.FormatFloat(f, 'f', -1, 64),
			}
		case reflect.String:
			s := v.(string)
			values[k] = &Value{
				val: s,
				raw: s,
			}
		default:
			return errors.Errorf("UnmarshalTOML: unknown kind(%v)", rv.Kind())
		}
	}
	m.Store(values)
	return nil
}
