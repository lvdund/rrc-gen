package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	RRCResumeComplete_v1610_IEs_mobilityState_r16_Enum_normal uper.Enumerated = 0
	RRCResumeComplete_v1610_IEs_mobilityState_r16_Enum_medium uper.Enumerated = 1
	RRCResumeComplete_v1610_IEs_mobilityState_r16_Enum_high   uper.Enumerated = 2
	RRCResumeComplete_v1610_IEs_mobilityState_r16_Enum_spare  uper.Enumerated = 3
)

type RRCResumeComplete_v1610_IEs_mobilityState_r16 struct {
	Value uper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *RRCResumeComplete_v1610_IEs_mobilityState_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode RRCResumeComplete_v1610_IEs_mobilityState_r16", err)
	}
	return nil
}

func (ie *RRCResumeComplete_v1610_IEs_mobilityState_r16) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode RRCResumeComplete_v1610_IEs_mobilityState_r16", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
