package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type IntraBandCC_Combination_r17 struct {
	Value []CC_State_r17 `lb:1,ub:maxNrofServingCells,madatory`
}

func (ie *IntraBandCC_Combination_r17) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*CC_State_r17]([]*CC_State_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofServingCells}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode IntraBandCC_Combination_r17", err)
	}
	return nil
}

func (ie *IntraBandCC_Combination_r17) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*CC_State_r17]([]*CC_State_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofServingCells}, false)
	fn := func() *CC_State_r17 {
		return new(CC_State_r17)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode IntraBandCC_Combination_r17", err)
	}
	ie.Value = []CC_State_r17{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
