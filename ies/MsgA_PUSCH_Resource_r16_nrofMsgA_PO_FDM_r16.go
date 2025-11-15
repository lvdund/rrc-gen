package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	MsgA_PUSCH_Resource_r16_nrofMsgA_PO_FDM_r16_Enum_one   uper.Enumerated = 0
	MsgA_PUSCH_Resource_r16_nrofMsgA_PO_FDM_r16_Enum_two   uper.Enumerated = 1
	MsgA_PUSCH_Resource_r16_nrofMsgA_PO_FDM_r16_Enum_four  uper.Enumerated = 2
	MsgA_PUSCH_Resource_r16_nrofMsgA_PO_FDM_r16_Enum_eight uper.Enumerated = 3
)

type MsgA_PUSCH_Resource_r16_nrofMsgA_PO_FDM_r16 struct {
	Value uper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *MsgA_PUSCH_Resource_r16_nrofMsgA_PO_FDM_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode MsgA_PUSCH_Resource_r16_nrofMsgA_PO_FDM_r16", err)
	}
	return nil
}

func (ie *MsgA_PUSCH_Resource_r16_nrofMsgA_PO_FDM_r16) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode MsgA_PUSCH_Resource_r16_nrofMsgA_PO_FDM_r16", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
