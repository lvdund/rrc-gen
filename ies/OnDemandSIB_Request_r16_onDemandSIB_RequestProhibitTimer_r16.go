package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	OnDemandSIB_Request_r16_onDemandSIB_RequestProhibitTimer_r16_Enum_s0     uper.Enumerated = 0
	OnDemandSIB_Request_r16_onDemandSIB_RequestProhibitTimer_r16_Enum_s0dot5 uper.Enumerated = 1
	OnDemandSIB_Request_r16_onDemandSIB_RequestProhibitTimer_r16_Enum_s1     uper.Enumerated = 2
	OnDemandSIB_Request_r16_onDemandSIB_RequestProhibitTimer_r16_Enum_s2     uper.Enumerated = 3
	OnDemandSIB_Request_r16_onDemandSIB_RequestProhibitTimer_r16_Enum_s5     uper.Enumerated = 4
	OnDemandSIB_Request_r16_onDemandSIB_RequestProhibitTimer_r16_Enum_s10    uper.Enumerated = 5
	OnDemandSIB_Request_r16_onDemandSIB_RequestProhibitTimer_r16_Enum_s20    uper.Enumerated = 6
	OnDemandSIB_Request_r16_onDemandSIB_RequestProhibitTimer_r16_Enum_s30    uper.Enumerated = 7
)

type OnDemandSIB_Request_r16_onDemandSIB_RequestProhibitTimer_r16 struct {
	Value uper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *OnDemandSIB_Request_r16_onDemandSIB_RequestProhibitTimer_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode OnDemandSIB_Request_r16_onDemandSIB_RequestProhibitTimer_r16", err)
	}
	return nil
}

func (ie *OnDemandSIB_Request_r16_onDemandSIB_RequestProhibitTimer_r16) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode OnDemandSIB_Request_r16_onDemandSIB_RequestProhibitTimer_r16", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
