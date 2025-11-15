package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PosSIB_MappingInfo_r16 struct {
	Value []PosSIB_Type_r16 `lb:1,ub:maxSIB,madatory`
}

func (ie *PosSIB_MappingInfo_r16) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*PosSIB_Type_r16]([]*PosSIB_Type_r16{}, uper.Constraint{Lb: 1, Ub: maxSIB}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode PosSIB_MappingInfo_r16", err)
	}
	return nil
}

func (ie *PosSIB_MappingInfo_r16) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*PosSIB_Type_r16]([]*PosSIB_Type_r16{}, uper.Constraint{Lb: 1, Ub: maxSIB}, false)
	fn := func() *PosSIB_Type_r16 {
		return new(PosSIB_Type_r16)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode PosSIB_MappingInfo_r16", err)
	}
	ie.Value = []PosSIB_Type_r16{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
