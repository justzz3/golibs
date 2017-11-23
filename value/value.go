package value

import (
	"fmt"
	"strconv"
	"strings"
)

type Raw struct {
	value interface{}
}

func New(v interface{}) *Raw {
	return &Raw{v}
}

func (v *Raw) Get(path string) (ret interface{}, err error) {
	parts := strings.Split(path, ".")
	ret = v.value
	for _, k := range parts {
		switch vv := ret.(type) {
		case map[interface{}]interface{}:
			return nil, fmt.Errorf("unsupport %T, ret: %v", vv, ret)
		case map[string]interface{}:
			ret = vv[k]
		case []interface{}:
			i, err := strconv.Atoi(k)
			if err != nil {
				return nil, fmt.Errorf("%s is not a number", k)
			}
			ret = vv[i]
		default:
			return nil, fmt.Errorf("unsupport %T, ret: %v", vv, ret)
		}
	}
	return
}

func (v *Raw) Bool(path string) bool {
	ret, err := v.Get(path)
	if err != nil {
		return false
	}
	switch vv := ret.(type) {
	case bool:
		return vv
	default:
		fmt.Printf("unsupport %T, ret: %v\n", vv, ret)
		return false
	}
}
func (v *Raw) Int(path string) int {
	ret, err := v.Get(path)
	if err != nil {
		return 0
	}
	switch vv := ret.(type) {
	case int:
		return vv
	case int64:
		return int(vv)
	case float64:
		return int(vv)
	default:
		fmt.Printf("unsupport %T, ret: %v\n", vv, ret)
		return 0
	}
}

func (v *Raw) Int64(path string) int64 {
	ret, err := v.Get(path)
	if err != nil {
		return 0
	}
	switch vv := ret.(type) {
	case int:
		return int64(vv)
	case int64:
		return vv
	case float64:
		return int64(vv)
	default:
		fmt.Printf("unsupport %T, ret: %v\n", vv, ret)
		return 0
	}
}

func (v *Raw) String(path string) string {
	ret, err := v.Get(path)
	if err != nil {
		return ""
	}
	switch vv := ret.(type) {
	case string:
		return vv
	case Stringer:
		return vv.String()
	default:
		fmt.Printf("unsupport %T, ret: %v\n", vv, ret)
		return ""
	}
}

type Stringer interface {
	String() string
}
