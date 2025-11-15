package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CandidateCellInfo_r17 struct {
	ssbFrequency_r17  ARFCN_ValueNR       `madatory`
	candidateList_r17 []CandidateCell_r17 `lb:1,ub:maxNrofCondCells_r16,madatory`
}

func (ie *CandidateCellInfo_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.ssbFrequency_r17.Encode(w); err != nil {
		return utils.WrapError("Encode ssbFrequency_r17", err)
	}
	tmp_candidateList_r17 := utils.NewSequence[*CandidateCell_r17]([]*CandidateCell_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofCondCells_r16}, false)
	for _, i := range ie.candidateList_r17 {
		tmp_candidateList_r17.Value = append(tmp_candidateList_r17.Value, &i)
	}
	if err = tmp_candidateList_r17.Encode(w); err != nil {
		return utils.WrapError("Encode candidateList_r17", err)
	}
	return nil
}

func (ie *CandidateCellInfo_r17) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.ssbFrequency_r17.Decode(r); err != nil {
		return utils.WrapError("Decode ssbFrequency_r17", err)
	}
	tmp_candidateList_r17 := utils.NewSequence[*CandidateCell_r17]([]*CandidateCell_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofCondCells_r16}, false)
	fn_candidateList_r17 := func() *CandidateCell_r17 {
		return new(CandidateCell_r17)
	}
	if err = tmp_candidateList_r17.Decode(r, fn_candidateList_r17); err != nil {
		return utils.WrapError("Decode candidateList_r17", err)
	}
	ie.candidateList_r17 = []CandidateCell_r17{}
	for _, i := range tmp_candidateList_r17.Value {
		ie.candidateList_r17 = append(ie.candidateList_r17, *i)
	}
	return nil
}
