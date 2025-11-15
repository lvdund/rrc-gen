package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type NeedForGapsBandListNR_r16 struct {
	Value []NeedForGapsNR_r16 `lb:1,ub:maxBands,madatory`
}

func (ie *NeedForGapsBandListNR_r16) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*NeedForGapsNR_r16]([]*NeedForGapsNR_r16{}, uper.Constraint{Lb: 1, Ub: maxBands}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode NeedForGapsBandListNR_r16", err)
	}
	return nil
}

func (ie *NeedForGapsBandListNR_r16) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*NeedForGapsNR_r16]([]*NeedForGapsNR_r16{}, uper.Constraint{Lb: 1, Ub: maxBands}, false)
	fn := func() *NeedForGapsNR_r16 {
		return new(NeedForGapsNR_r16)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode NeedForGapsBandListNR_r16", err)
	}
	ie.Value = []NeedForGapsNR_r16{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
