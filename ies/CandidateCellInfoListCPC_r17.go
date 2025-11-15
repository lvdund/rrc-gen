package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CandidateCellInfoListCPC_r17 struct {
	Value []CandidateCellInfo_r17 `lb:1,ub:maxFreq,madatory`
}

func (ie *CandidateCellInfoListCPC_r17) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*CandidateCellInfo_r17]([]*CandidateCellInfo_r17{}, uper.Constraint{Lb: 1, Ub: maxFreq}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode CandidateCellInfoListCPC_r17", err)
	}
	return nil
}

func (ie *CandidateCellInfoListCPC_r17) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*CandidateCellInfo_r17]([]*CandidateCellInfo_r17{}, uper.Constraint{Lb: 1, Ub: maxFreq}, false)
	fn := func() *CandidateCellInfo_r17 {
		return new(CandidateCellInfo_r17)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode CandidateCellInfoListCPC_r17", err)
	}
	ie.Value = []CandidateCellInfo_r17{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
