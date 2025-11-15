package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	RAT_Type_Enum_nr             uper.Enumerated = 0
	RAT_Type_Enum_eutra_nr       uper.Enumerated = 1
	RAT_Type_Enum_eutra          uper.Enumerated = 2
	RAT_Type_Enum_utra_fdd_v1610 uper.Enumerated = 3
)

type RAT_Type struct {
	Value uper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *RAT_Type) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode RAT_Type", err)
	}
	return nil
}

func (ie *RAT_Type) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode RAT_Type", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
