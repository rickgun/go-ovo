package ovo

type Env int8

const (
	_ Env = iota

	// Staging : represent sandbox environment
	Staging

	// Production : represent production environment
	Production
)

var typeString = map[Env]string{
	Staging:    "https://api.byte-stack.net",
	Production: "https://api.ovo.id",
}

// implement stringer
func (e Env) String() string {
	for k, v := range typeString {
		if k == e {
			return v
		}
	}
	return "undefined"
}
