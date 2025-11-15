package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CandidateServingFreqListEUTRA struct {
	Value []ARFCN_ValueEUTRA `lb:1,ub:maxFreqIDC_MRDC,madatory`
}

func (ie *CandidateServingFreqListEUTRA) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*ARFCN_ValueEUTRA]([]*ARFCN_ValueEUTRA{}, uper.Constraint{Lb: 1, Ub: maxFreqIDC_MRDC}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode CandidateServingFreqListEUTRA", err)
	}
	return nil
}

func (ie *CandidateServingFreqListEUTRA) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*ARFCN_ValueEUTRA]([]*ARFCN_ValueEUTRA{}, uper.Constraint{Lb: 1, Ub: maxFreqIDC_MRDC}, false)
	fn := func() *ARFCN_ValueEUTRA {
		return new(ARFCN_ValueEUTRA)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode CandidateServingFreqListEUTRA", err)
	}
	ie.Value = []ARFCN_ValueEUTRA{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
