package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type FreqPriorityListSlicing_r17 struct {
	Value []FreqPrioritySlicing_r17 `lb:1,ub:maxFreqPlus1,madatory`
}

func (ie *FreqPriorityListSlicing_r17) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*FreqPrioritySlicing_r17]([]*FreqPrioritySlicing_r17{}, uper.Constraint{Lb: 1, Ub: maxFreqPlus1}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode FreqPriorityListSlicing_r17", err)
	}
	return nil
}

func (ie *FreqPriorityListSlicing_r17) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*FreqPrioritySlicing_r17]([]*FreqPrioritySlicing_r17{}, uper.Constraint{Lb: 1, Ub: maxFreqPlus1}, false)
	fn := func() *FreqPrioritySlicing_r17 {
		return new(FreqPrioritySlicing_r17)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode FreqPriorityListSlicing_r17", err)
	}
	ie.Value = []FreqPrioritySlicing_r17{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
