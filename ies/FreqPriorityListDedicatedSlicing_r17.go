package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type FreqPriorityListDedicatedSlicing_r17 struct {
	Value []FreqPriorityDedicatedSlicing_r17 `lb:1,ub:maxFreq,madatory`
}

func (ie *FreqPriorityListDedicatedSlicing_r17) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*FreqPriorityDedicatedSlicing_r17]([]*FreqPriorityDedicatedSlicing_r17{}, uper.Constraint{Lb: 1, Ub: maxFreq}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode FreqPriorityListDedicatedSlicing_r17", err)
	}
	return nil
}

func (ie *FreqPriorityListDedicatedSlicing_r17) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*FreqPriorityDedicatedSlicing_r17]([]*FreqPriorityDedicatedSlicing_r17{}, uper.Constraint{Lb: 1, Ub: maxFreq}, false)
	fn := func() *FreqPriorityDedicatedSlicing_r17 {
		return new(FreqPriorityDedicatedSlicing_r17)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode FreqPriorityListDedicatedSlicing_r17", err)
	}
	ie.Value = []FreqPriorityDedicatedSlicing_r17{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
