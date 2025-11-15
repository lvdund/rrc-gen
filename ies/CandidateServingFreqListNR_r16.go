package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CandidateServingFreqListNR_r16 struct {
	Value []ARFCN_ValueNR `lb:1,ub:maxFreqIDC_r16,madatory`
}

func (ie *CandidateServingFreqListNR_r16) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*ARFCN_ValueNR]([]*ARFCN_ValueNR{}, uper.Constraint{Lb: 1, Ub: maxFreqIDC_r16}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode CandidateServingFreqListNR_r16", err)
	}
	return nil
}

func (ie *CandidateServingFreqListNR_r16) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*ARFCN_ValueNR]([]*ARFCN_ValueNR{}, uper.Constraint{Lb: 1, Ub: maxFreqIDC_r16}, false)
	fn := func() *ARFCN_ValueNR {
		return new(ARFCN_ValueNR)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode CandidateServingFreqListNR_r16", err)
	}
	ie.Value = []ARFCN_ValueNR{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
