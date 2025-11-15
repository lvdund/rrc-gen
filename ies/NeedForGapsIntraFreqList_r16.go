package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type NeedForGapsIntraFreqList_r16 struct {
	Value []NeedForGapsIntraFreq_r16 `lb:1,ub:maxNrofServingCells,madatory`
}

func (ie *NeedForGapsIntraFreqList_r16) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*NeedForGapsIntraFreq_r16]([]*NeedForGapsIntraFreq_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofServingCells}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode NeedForGapsIntraFreqList_r16", err)
	}
	return nil
}

func (ie *NeedForGapsIntraFreqList_r16) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*NeedForGapsIntraFreq_r16]([]*NeedForGapsIntraFreq_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofServingCells}, false)
	fn := func() *NeedForGapsIntraFreq_r16 {
		return new(NeedForGapsIntraFreq_r16)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode NeedForGapsIntraFreqList_r16", err)
	}
	ie.Value = []NeedForGapsIntraFreq_r16{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
