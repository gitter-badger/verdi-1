// generated by stringer -type=Status; DO NOT EDIT

package model

import "fmt"

const _Status_name = "DownUpUnknown"

var _Status_index = [...]uint8{0, 4, 6, 13}

func (i Status) String() string {
	if i < 0 || i+1 >= Status(len(_Status_index)) {
		return fmt.Sprintf("Status(%d)", i)
	}
	return _Status_name[_Status_index[i]:_Status_index[i+1]]
}
