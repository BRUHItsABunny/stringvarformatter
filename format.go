package stringvarformatter

import "strings"

type Formatter struct {
	FMT  string
	Keys []string
	null string
}

func NewFormatter(fmt string, keys ...string) *Formatter {
	return &Formatter{FMT: fmt, Keys: keys, null: "null"}
}

func (f *Formatter) GetVarMap() map[string]string {
	result := map[string]string{}
	for _, varKey := range f.Keys {
		result[varKey] = f.null
	}
	return result
}

func (f *Formatter) SetNull(null string) {
	f.null = null
}

// Format TODO: Optimize
func (f *Formatter) Format(depth int, values map[string]string) string {
	result := f.FMT
	for _, varName := range f.Keys {
		for i := 0; i < depth; i++ {
			varNameStart := strings.Index(result, varName)
			if varNameStart == -1 {
				break
			}
			value, ok := values[varName]
			if !ok {
				value = f.null
			}
			result = strings.Replace(result, varName, value, 1)
		}
	}
	return result
}
