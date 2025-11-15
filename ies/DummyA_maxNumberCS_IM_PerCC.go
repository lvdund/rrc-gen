package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	DummyA_maxNumberCS_IM_PerCC_Enum_n1  uper.Enumerated = 0
	DummyA_maxNumberCS_IM_PerCC_Enum_n2  uper.Enumerated = 1
	DummyA_maxNumberCS_IM_PerCC_Enum_n4  uper.Enumerated = 2
	DummyA_maxNumberCS_IM_PerCC_Enum_n8  uper.Enumerated = 3
	DummyA_maxNumberCS_IM_PerCC_Enum_n16 uper.Enumerated = 4
	DummyA_maxNumberCS_IM_PerCC_Enum_n32 uper.Enumerated = 5
)

type DummyA_maxNumberCS_IM_PerCC struct {
	Value uper.Enumerated `lb:0,ub:5,madatory`
}

func (ie *DummyA_maxNumberCS_IM_PerCC) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("Encode DummyA_maxNumberCS_IM_PerCC", err)
	}
	return nil
}

func (ie *DummyA_maxNumberCS_IM_PerCC) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("Decode DummyA_maxNumberCS_IM_PerCC", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
