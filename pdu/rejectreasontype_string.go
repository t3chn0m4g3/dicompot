// Code generated by "stringer -type RejectReasonType"; DO NOT EDIT

package pdu

import "fmt"

const (
	_RejectReasonType_name_0 = "RejectReasonNoneRejectReasonApplicationContextNameNotSupportedRejectReasonCallingAETitleNotRecognized"
	_RejectReasonType_name_1 = "RejectReasonCalledAETitleNotRecognized"
)

var (
	_RejectReasonType_index_0 = [...]uint8{0, 16, 62, 101}
	_RejectReasonType_index_1 = [...]uint8{0, 38}
)

func (i RejectReasonType) String() string {
	switch {
	case 1 <= i && i <= 3:
		i--
		return _RejectReasonType_name_0[_RejectReasonType_index_0[i]:_RejectReasonType_index_0[i+1]]
	case i == 7:
		return _RejectReasonType_name_1
	default:
		return fmt.Sprintf("RejectReasonType(%d)", i)
	}
}
