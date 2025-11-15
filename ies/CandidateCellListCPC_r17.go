package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CandidateCellListCPC_r17 struct {
	Value []CandidateCellCPC_r17 `lb:1,ub:maxFreq,madatory`
}

func (ie *CandidateCellListCPC_r17) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*CandidateCellCPC_r17]([]*CandidateCellCPC_r17{}, uper.Constraint{Lb: 1, Ub: maxFreq}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode CandidateCellListCPC_r17", err)
	}
	return nil
}

func (ie *CandidateCellListCPC_r17) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*CandidateCellCPC_r17]([]*CandidateCellCPC_r17{}, uper.Constraint{Lb: 1, Ub: maxFreq}, false)
	fn := func() *CandidateCellCPC_r17 {
		return new(CandidateCellCPC_r17)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode CandidateCellListCPC_r17", err)
	}
	ie.Value = []CandidateCellCPC_r17{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
