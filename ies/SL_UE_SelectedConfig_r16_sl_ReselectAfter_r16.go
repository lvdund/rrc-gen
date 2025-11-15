package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SL_UE_SelectedConfig_r16_sl_ReselectAfter_r16_Enum_n1 uper.Enumerated = 0
	SL_UE_SelectedConfig_r16_sl_ReselectAfter_r16_Enum_n2 uper.Enumerated = 1
	SL_UE_SelectedConfig_r16_sl_ReselectAfter_r16_Enum_n3 uper.Enumerated = 2
	SL_UE_SelectedConfig_r16_sl_ReselectAfter_r16_Enum_n4 uper.Enumerated = 3
	SL_UE_SelectedConfig_r16_sl_ReselectAfter_r16_Enum_n5 uper.Enumerated = 4
	SL_UE_SelectedConfig_r16_sl_ReselectAfter_r16_Enum_n6 uper.Enumerated = 5
	SL_UE_SelectedConfig_r16_sl_ReselectAfter_r16_Enum_n7 uper.Enumerated = 6
	SL_UE_SelectedConfig_r16_sl_ReselectAfter_r16_Enum_n8 uper.Enumerated = 7
	SL_UE_SelectedConfig_r16_sl_ReselectAfter_r16_Enum_n9 uper.Enumerated = 8
)

type SL_UE_SelectedConfig_r16_sl_ReselectAfter_r16 struct {
	Value uper.Enumerated `lb:0,ub:8,madatory`
}

func (ie *SL_UE_SelectedConfig_r16_sl_ReselectAfter_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 8}, false); err != nil {
		return utils.WrapError("Encode SL_UE_SelectedConfig_r16_sl_ReselectAfter_r16", err)
	}
	return nil
}

func (ie *SL_UE_SelectedConfig_r16_sl_ReselectAfter_r16) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 8}, false); err != nil {
		return utils.WrapError("Decode SL_UE_SelectedConfig_r16_sl_ReselectAfter_r16", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
