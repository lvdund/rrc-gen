package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	CodebookParameters_type1_multiPanel_modes_Enum_mode1 uper.Enumerated = 0
	CodebookParameters_type1_multiPanel_modes_Enum_mode2 uper.Enumerated = 1
	CodebookParameters_type1_multiPanel_modes_Enum_both  uper.Enumerated = 2
)

type CodebookParameters_type1_multiPanel_modes struct {
	Value uper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *CodebookParameters_type1_multiPanel_modes) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode CodebookParameters_type1_multiPanel_modes", err)
	}
	return nil
}

func (ie *CodebookParameters_type1_multiPanel_modes) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode CodebookParameters_type1_multiPanel_modes", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
