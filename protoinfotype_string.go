// Code generated by "stringer -type=ProtoInfoType"; DO NOT EDIT.

package conntrack

import "strconv"

const _ProtoInfoType_name = "ctaProtoInfoUnspecctaProtoInfoTCPctaProtoInfoDCCPctaProtoInfoSCTP"

var _ProtoInfoType_index = [...]uint8{0, 18, 33, 49, 65}

func (i ProtoInfoType) String() string {
	if i >= ProtoInfoType(len(_ProtoInfoType_index)-1) {
		return "ProtoInfoType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ProtoInfoType_name[_ProtoInfoType_index[i]:_ProtoInfoType_index[i+1]]
}
