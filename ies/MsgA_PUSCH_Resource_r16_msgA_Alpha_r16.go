package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	MsgA_PUSCH_Resource_r16_msgA_Alpha_r16_Enum_alpha0  uper.Enumerated = 0
	MsgA_PUSCH_Resource_r16_msgA_Alpha_r16_Enum_alpha04 uper.Enumerated = 1
	MsgA_PUSCH_Resource_r16_msgA_Alpha_r16_Enum_alpha05 uper.Enumerated = 2
	MsgA_PUSCH_Resource_r16_msgA_Alpha_r16_Enum_alpha06 uper.Enumerated = 3
	MsgA_PUSCH_Resource_r16_msgA_Alpha_r16_Enum_alpha07 uper.Enumerated = 4
	MsgA_PUSCH_Resource_r16_msgA_Alpha_r16_Enum_alpha08 uper.Enumerated = 5
	MsgA_PUSCH_Resource_r16_msgA_Alpha_r16_Enum_alpha09 uper.Enumerated = 6
	MsgA_PUSCH_Resource_r16_msgA_Alpha_r16_Enum_alpha1  uper.Enumerated = 7
)

type MsgA_PUSCH_Resource_r16_msgA_Alpha_r16 struct {
	Value uper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *MsgA_PUSCH_Resource_r16_msgA_Alpha_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode MsgA_PUSCH_Resource_r16_msgA_Alpha_r16", err)
	}
	return nil
}

func (ie *MsgA_PUSCH_Resource_r16_msgA_Alpha_r16) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode MsgA_PUSCH_Resource_r16_msgA_Alpha_r16", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
