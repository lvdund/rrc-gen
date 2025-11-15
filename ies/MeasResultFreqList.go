package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasResultFreqList struct {
	Value []MeasResult2NR `lb:1,ub:maxFreq,madatory`
}

func (ie *MeasResultFreqList) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*MeasResult2NR]([]*MeasResult2NR{}, uper.Constraint{Lb: 1, Ub: maxFreq}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode MeasResultFreqList", err)
	}
	return nil
}

func (ie *MeasResultFreqList) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*MeasResult2NR]([]*MeasResult2NR{}, uper.Constraint{Lb: 1, Ub: maxFreq}, false)
	fn := func() *MeasResult2NR {
		return new(MeasResult2NR)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode MeasResultFreqList", err)
	}
	ie.Value = []MeasResult2NR{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
