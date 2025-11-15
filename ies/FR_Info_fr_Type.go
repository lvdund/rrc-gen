package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	FR_Info_fr_Type_Enum_fr1 uper.Enumerated = 0
	FR_Info_fr_Type_Enum_fr2 uper.Enumerated = 1
)

type FR_Info_fr_Type struct {
	Value uper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *FR_Info_fr_Type) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode FR_Info_fr_Type", err)
	}
	return nil
}

func (ie *FR_Info_fr_Type) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode FR_Info_fr_Type", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
