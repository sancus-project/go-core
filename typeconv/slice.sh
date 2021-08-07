#!/bin/sh

set -eu

F="${0%.sh}.go"
trap "rm -f '$F~'" EXIT
exec > "$F~"

cat <<EOT
package typeconv

//go:generate $0 $*

// Code generated by $0 DO NOT EDIT
EOT

generate() {
	local n="$1"
	local t="$(echo "$n" | tr A-Z a-z)"

cat <<EOT

// As${n}Slice tries to convert data into a slice of $t
func As${n}Slice(v interface{}) ([]$t, bool) {
	switch w := v.(type) {
	case []$t:
		// ready
		return w, true
	case []interface{}:
		// slice of something else, convert them to $n
		out := make([]$t, 0, len(w))
		for _, o := range w {
			if n, ok := As$n(o); ok {
				out = append(out, n)
			} else {
				// failed
				return nil, false
			}
		}
		return out, true
	case interface{}:
		// promote single element to slice
		if n, ok := As$n(v); ok {
			return []$t{n}, true
		}
	}

	return nil, false
}
EOT
}

for x; do
	generate $x
done

if ! cmp -s "$F" "$F~" 2> /dev/null; then
	mv "$F~" "$F"
fi
