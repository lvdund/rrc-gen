package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MCC struct {
	Value []MCC_MNC_Digit `lb:3,ub:3,madatory`
}

func (ie *MCC) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*MCC_MNC_Digit]([]*MCC_MNC_Digit{}, uper.Constraint{Lb: 3, Ub: 3}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode MCC", err)
	}
	return nil
}

func (ie *MCC) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*MCC_MNC_Digit]([]*MCC_MNC_Digit{}, uper.Constraint{Lb: 3, Ub: 3}, false)
	fn := func() *MCC_MNC_Digit {
		return new(MCC_MNC_Digit)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode MCC", err)
	}
	ie.Value = []MCC_MNC_Digit{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
