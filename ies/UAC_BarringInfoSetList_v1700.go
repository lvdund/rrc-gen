package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UAC_BarringInfoSetList_v1700 struct {
	Value []UAC_BarringInfoSet_v1700 `lb:1,ub:maxBarringInfoSet,madatory`
}

func (ie *UAC_BarringInfoSetList_v1700) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*UAC_BarringInfoSet_v1700]([]*UAC_BarringInfoSet_v1700{}, uper.Constraint{Lb: 1, Ub: maxBarringInfoSet}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode UAC_BarringInfoSetList_v1700", err)
	}
	return nil
}

func (ie *UAC_BarringInfoSetList_v1700) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*UAC_BarringInfoSet_v1700]([]*UAC_BarringInfoSet_v1700{}, uper.Constraint{Lb: 1, Ub: maxBarringInfoSet}, false)
	fn := func() *UAC_BarringInfoSet_v1700 {
		return new(UAC_BarringInfoSet_v1700)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode UAC_BarringInfoSetList_v1700", err)
	}
	ie.Value = []UAC_BarringInfoSet_v1700{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
