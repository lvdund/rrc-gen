package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasResultFreqListFailMRDC struct {
	Value []MeasResult2EUTRA `lb:1,ub:maxFreq,madatory`
}

func (ie *MeasResultFreqListFailMRDC) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*MeasResult2EUTRA]([]*MeasResult2EUTRA{}, uper.Constraint{Lb: 1, Ub: maxFreq}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode MeasResultFreqListFailMRDC", err)
	}
	return nil
}

func (ie *MeasResultFreqListFailMRDC) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*MeasResult2EUTRA]([]*MeasResult2EUTRA{}, uper.Constraint{Lb: 1, Ub: maxFreq}, false)
	fn := func() *MeasResult2EUTRA {
		return new(MeasResult2EUTRA)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode MeasResultFreqListFailMRDC", err)
	}
	ie.Value = []MeasResult2EUTRA{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
