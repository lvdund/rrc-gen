package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CandidateCellCPC_r17 struct {
	ssbFrequency_r17      ARFCN_ValueNR `madatory`
	candidateCellList_r17 []PhysCellId  `lb:1,ub:maxNrofCondCells_r16,madatory`
}

func (ie *CandidateCellCPC_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.ssbFrequency_r17.Encode(w); err != nil {
		return utils.WrapError("Encode ssbFrequency_r17", err)
	}
	tmp_candidateCellList_r17 := utils.NewSequence[*PhysCellId]([]*PhysCellId{}, uper.Constraint{Lb: 1, Ub: maxNrofCondCells_r16}, false)
	for _, i := range ie.candidateCellList_r17 {
		tmp_candidateCellList_r17.Value = append(tmp_candidateCellList_r17.Value, &i)
	}
	if err = tmp_candidateCellList_r17.Encode(w); err != nil {
		return utils.WrapError("Encode candidateCellList_r17", err)
	}
	return nil
}

func (ie *CandidateCellCPC_r17) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.ssbFrequency_r17.Decode(r); err != nil {
		return utils.WrapError("Decode ssbFrequency_r17", err)
	}
	tmp_candidateCellList_r17 := utils.NewSequence[*PhysCellId]([]*PhysCellId{}, uper.Constraint{Lb: 1, Ub: maxNrofCondCells_r16}, false)
	fn_candidateCellList_r17 := func() *PhysCellId {
		return new(PhysCellId)
	}
	if err = tmp_candidateCellList_r17.Decode(r, fn_candidateCellList_r17); err != nil {
		return utils.WrapError("Decode candidateCellList_r17", err)
	}
	ie.candidateCellList_r17 = []PhysCellId{}
	for _, i := range tmp_candidateCellList_r17.Value {
		ie.candidateCellList_r17 = append(ie.candidateCellList_r17, *i)
	}
	return nil
}
