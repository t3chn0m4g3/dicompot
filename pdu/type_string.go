// Code generated by "stringer -type Type"; DO NOT EDIT

package pdu

import "fmt"

const _Type_name = "TypeAAssociateRq"

var _Type_index = [...]uint8{0, 16}

func (i Type) String() string {
	i--
	if i >= Type(len(_Type_index)-1) {
		return fmt.Sprintf("Type(%d)", i+1)
	}
	return _Type_name[_Type_index[i]:_Type_index[i+1]]
}
