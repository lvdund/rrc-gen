package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MNC struct {
	Value []MCC_MNC_Digit `lb:2,ub:3,madatory`
}

func (ie *MNC) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*MCC_MNC_Digit]([]*MCC_MNC_Digit{}, uper.Constraint{Lb: 2, Ub: 3}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode MNC", err)
	}
	return nil
}

func (ie *MNC) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*MCC_MNC_Digit]([]*MCC_MNC_Digit{}, uper.Constraint{Lb: 2, Ub: 3}, false)
	fn := func() *MCC_MNC_Digit {
		return new(MCC_MNC_Digit)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode MNC", err)
	}
	ie.Value = []MCC_MNC_Digit{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
