package conf

import (
	"encoding/json"
	"fmt"
)

// Data is the core type that holds the JSON data. Config encapsulates
// a single Data type.
type Data map[string]string

// String fulfills the Stringer interface my marshaling the Data into
// compressed (no indents) JSON. Returns JSON containing just ERROR if
// an error occurred during marshalling.
func (d Data) String() string {
	byt, err := json.Marshal(d)
	if err != nil {
		return fmt.Sprintf("{\"ERROR\":\"%v\"}", err)
	}
	return string(byt)
}
