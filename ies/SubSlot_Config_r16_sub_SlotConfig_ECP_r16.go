package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SubSlot_Config_r16_sub_SlotConfig_ECP_r16_Enum_n4 uper.Enumerated = 0
	SubSlot_Config_r16_sub_SlotConfig_ECP_r16_Enum_n5 uper.Enumerated = 1
	SubSlot_Config_r16_sub_SlotConfig_ECP_r16_Enum_n6 uper.Enumerated = 2
)

type SubSlot_Config_r16_sub_SlotConfig_ECP_r16 struct {
	Value uper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *SubSlot_Config_r16_sub_SlotConfig_ECP_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode SubSlot_Config_r16_sub_SlotConfig_ECP_r16", err)
	}
	return nil
}

func (ie *SubSlot_Config_r16_sub_SlotConfig_ECP_r16) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode SubSlot_Config_r16_sub_SlotConfig_ECP_r16", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
