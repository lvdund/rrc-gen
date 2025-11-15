package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SL_InterUE_CoordinationScheme1_r17_sl_IUC_Explicit_r17_Enum_enabled  uper.Enumerated = 0
	SL_InterUE_CoordinationScheme1_r17_sl_IUC_Explicit_r17_Enum_disabled uper.Enumerated = 1
)

type SL_InterUE_CoordinationScheme1_r17_sl_IUC_Explicit_r17 struct {
	Value uper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *SL_InterUE_CoordinationScheme1_r17_sl_IUC_Explicit_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode SL_InterUE_CoordinationScheme1_r17_sl_IUC_Explicit_r17", err)
	}
	return nil
}

func (ie *SL_InterUE_CoordinationScheme1_r17_sl_IUC_Explicit_r17) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode SL_InterUE_CoordinationScheme1_r17_sl_IUC_Explicit_r17", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
