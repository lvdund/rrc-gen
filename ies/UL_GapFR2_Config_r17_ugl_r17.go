package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	UL_GapFR2_Config_r17_ugl_r17_Enum_ms0dot125 uper.Enumerated = 0
	UL_GapFR2_Config_r17_ugl_r17_Enum_ms0dot25  uper.Enumerated = 1
	UL_GapFR2_Config_r17_ugl_r17_Enum_ms0dot5   uper.Enumerated = 2
	UL_GapFR2_Config_r17_ugl_r17_Enum_ms1       uper.Enumerated = 3
)

type UL_GapFR2_Config_r17_ugl_r17 struct {
	Value uper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *UL_GapFR2_Config_r17_ugl_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode UL_GapFR2_Config_r17_ugl_r17", err)
	}
	return nil
}

func (ie *UL_GapFR2_Config_r17_ugl_r17) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode UL_GapFR2_Config_r17_ugl_r17", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
