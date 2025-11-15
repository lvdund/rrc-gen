package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasTimingList struct {
	Value []MeasTiming `lb:1,ub:maxMeasFreqsMN,madatory`
}

func (ie *MeasTimingList) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*MeasTiming]([]*MeasTiming{}, uper.Constraint{Lb: 1, Ub: maxMeasFreqsMN}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode MeasTimingList", err)
	}
	return nil
}

func (ie *MeasTimingList) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*MeasTiming]([]*MeasTiming{}, uper.Constraint{Lb: 1, Ub: maxMeasFreqsMN}, false)
	fn := func() *MeasTiming {
		return new(MeasTiming)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode MeasTimingList", err)
	}
	ie.Value = []MeasTiming{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
