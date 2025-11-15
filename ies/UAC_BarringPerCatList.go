package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UAC_BarringPerCatList struct {
	Value []UAC_BarringPerCat `lb:1,ub:maxAccessCat_1,madatory`
}

func (ie *UAC_BarringPerCatList) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*UAC_BarringPerCat]([]*UAC_BarringPerCat{}, uper.Constraint{Lb: 1, Ub: maxAccessCat_1}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode UAC_BarringPerCatList", err)
	}
	return nil
}

func (ie *UAC_BarringPerCatList) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*UAC_BarringPerCat]([]*UAC_BarringPerCat{}, uper.Constraint{Lb: 1, Ub: maxAccessCat_1}, false)
	fn := func() *UAC_BarringPerCat {
		return new(UAC_BarringPerCat)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode UAC_BarringPerCatList", err)
	}
	ie.Value = []UAC_BarringPerCat{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
