package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	Alpha_Enum_alpha0  uper.Enumerated = 0
	Alpha_Enum_alpha04 uper.Enumerated = 1
	Alpha_Enum_alpha05 uper.Enumerated = 2
	Alpha_Enum_alpha06 uper.Enumerated = 3
	Alpha_Enum_alpha07 uper.Enumerated = 4
	Alpha_Enum_alpha08 uper.Enumerated = 5
	Alpha_Enum_alpha09 uper.Enumerated = 6
	Alpha_Enum_alpha1  uper.Enumerated = 7
)

type Alpha struct {
	Value uper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *Alpha) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode Alpha", err)
	}
	return nil
}

func (ie *Alpha) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode Alpha", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
