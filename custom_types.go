package owl

import "encoding/json"

type FlexibleInteger int

func (fi *FlexibleInteger) UnmarshalJSON(b []byte) error {
	if b[0] == '"' {
		*fi = FlexibleInteger(-1)
		return nil
	}
	return json.Unmarshal(b, (*int)(fi))
}
