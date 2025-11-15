package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	MsgA_PUSCH_Config_r16_msgA_TransformPrecoder_r16_Enum_enabled  uper.Enumerated = 0
	MsgA_PUSCH_Config_r16_msgA_TransformPrecoder_r16_Enum_disabled uper.Enumerated = 1
)

type MsgA_PUSCH_Config_r16_msgA_TransformPrecoder_r16 struct {
	Value uper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *MsgA_PUSCH_Config_r16_msgA_TransformPrecoder_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode MsgA_PUSCH_Config_r16_msgA_TransformPrecoder_r16", err)
	}
	return nil
}

func (ie *MsgA_PUSCH_Config_r16_msgA_TransformPrecoder_r16) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode MsgA_PUSCH_Config_r16_msgA_TransformPrecoder_r16", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
