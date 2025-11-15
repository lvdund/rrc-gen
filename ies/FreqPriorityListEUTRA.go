package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type FreqPriorityListEUTRA struct {
	Value []FreqPriorityEUTRA `lb:1,ub:maxFreq,madatory`
}

func (ie *FreqPriorityListEUTRA) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*FreqPriorityEUTRA]([]*FreqPriorityEUTRA{}, uper.Constraint{Lb: 1, Ub: maxFreq}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode FreqPriorityListEUTRA", err)
	}
	return nil
}

func (ie *FreqPriorityListEUTRA) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*FreqPriorityEUTRA]([]*FreqPriorityEUTRA{}, uper.Constraint{Lb: 1, Ub: maxFreq}, false)
	fn := func() *FreqPriorityEUTRA {
		return new(FreqPriorityEUTRA)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode FreqPriorityListEUTRA", err)
	}
	ie.Value = []FreqPriorityEUTRA{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
