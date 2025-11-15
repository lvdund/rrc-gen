package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UAC_BarringInfoSetList struct {
	Value []UAC_BarringInfoSet `lb:1,ub:maxBarringInfoSet,madatory`
}

func (ie *UAC_BarringInfoSetList) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*UAC_BarringInfoSet]([]*UAC_BarringInfoSet{}, uper.Constraint{Lb: 1, Ub: maxBarringInfoSet}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode UAC_BarringInfoSetList", err)
	}
	return nil
}

func (ie *UAC_BarringInfoSetList) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*UAC_BarringInfoSet]([]*UAC_BarringInfoSet{}, uper.Constraint{Lb: 1, Ub: maxBarringInfoSet}, false)
	fn := func() *UAC_BarringInfoSet {
		return new(UAC_BarringInfoSet)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode UAC_BarringInfoSetList", err)
	}
	ie.Value = []UAC_BarringInfoSet{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
