package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasResultListLogging2NR_r16 struct {
	Value []MeasResultLogging2NR_r16 `lb:1,ub:maxFreq,madatory`
}

func (ie *MeasResultListLogging2NR_r16) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*MeasResultLogging2NR_r16]([]*MeasResultLogging2NR_r16{}, uper.Constraint{Lb: 1, Ub: maxFreq}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode MeasResultListLogging2NR_r16", err)
	}
	return nil
}

func (ie *MeasResultListLogging2NR_r16) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*MeasResultLogging2NR_r16]([]*MeasResultLogging2NR_r16{}, uper.Constraint{Lb: 1, Ub: maxFreq}, false)
	fn := func() *MeasResultLogging2NR_r16 {
		return new(MeasResultLogging2NR_r16)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode MeasResultListLogging2NR_r16", err)
	}
	ie.Value = []MeasResultLogging2NR_r16{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
