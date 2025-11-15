package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type FreqPriorityListNR struct {
	Value []FreqPriorityNR `lb:1,ub:maxFreq,madatory`
}

func (ie *FreqPriorityListNR) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*FreqPriorityNR]([]*FreqPriorityNR{}, uper.Constraint{Lb: 1, Ub: maxFreq}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode FreqPriorityListNR", err)
	}
	return nil
}

func (ie *FreqPriorityListNR) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*FreqPriorityNR]([]*FreqPriorityNR{}, uper.Constraint{Lb: 1, Ub: maxFreq}, false)
	fn := func() *FreqPriorityNR {
		return new(FreqPriorityNR)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode FreqPriorityListNR", err)
	}
	ie.Value = []FreqPriorityNR{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
