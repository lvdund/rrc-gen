package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CG_CandidateList_r17_IEs struct {
	cg_CandidateToAddModList_r17  []CG_CandidateInfo_r17   `lb:1,ub:maxNrofCondCells_r16,optional`
	cg_CandidateToReleaseList_r17 []CG_CandidateInfoId_r17 `lb:1,ub:maxNrofCondCells_r16,optional`
	nonCriticalExtension          interface{}              `optional`
}

func (ie *CG_CandidateList_r17_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{len(ie.cg_CandidateToAddModList_r17) > 0, len(ie.cg_CandidateToReleaseList_r17) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.cg_CandidateToAddModList_r17) > 0 {
		tmp_cg_CandidateToAddModList_r17 := utils.NewSequence[*CG_CandidateInfo_r17]([]*CG_CandidateInfo_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofCondCells_r16}, false)
		for _, i := range ie.cg_CandidateToAddModList_r17 {
			tmp_cg_CandidateToAddModList_r17.Value = append(tmp_cg_CandidateToAddModList_r17.Value, &i)
		}
		if err = tmp_cg_CandidateToAddModList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode cg_CandidateToAddModList_r17", err)
		}
	}
	if len(ie.cg_CandidateToReleaseList_r17) > 0 {
		tmp_cg_CandidateToReleaseList_r17 := utils.NewSequence[*CG_CandidateInfoId_r17]([]*CG_CandidateInfoId_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofCondCells_r16}, false)
		for _, i := range ie.cg_CandidateToReleaseList_r17 {
			tmp_cg_CandidateToReleaseList_r17.Value = append(tmp_cg_CandidateToReleaseList_r17.Value, &i)
		}
		if err = tmp_cg_CandidateToReleaseList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode cg_CandidateToReleaseList_r17", err)
		}
	}
	return nil
}

func (ie *CG_CandidateList_r17_IEs) Decode(r *uper.UperReader) error {
	var err error
	var cg_CandidateToAddModList_r17Present bool
	if cg_CandidateToAddModList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var cg_CandidateToReleaseList_r17Present bool
	if cg_CandidateToReleaseList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if cg_CandidateToAddModList_r17Present {
		tmp_cg_CandidateToAddModList_r17 := utils.NewSequence[*CG_CandidateInfo_r17]([]*CG_CandidateInfo_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofCondCells_r16}, false)
		fn_cg_CandidateToAddModList_r17 := func() *CG_CandidateInfo_r17 {
			return new(CG_CandidateInfo_r17)
		}
		if err = tmp_cg_CandidateToAddModList_r17.Decode(r, fn_cg_CandidateToAddModList_r17); err != nil {
			return utils.WrapError("Decode cg_CandidateToAddModList_r17", err)
		}
		ie.cg_CandidateToAddModList_r17 = []CG_CandidateInfo_r17{}
		for _, i := range tmp_cg_CandidateToAddModList_r17.Value {
			ie.cg_CandidateToAddModList_r17 = append(ie.cg_CandidateToAddModList_r17, *i)
		}
	}
	if cg_CandidateToReleaseList_r17Present {
		tmp_cg_CandidateToReleaseList_r17 := utils.NewSequence[*CG_CandidateInfoId_r17]([]*CG_CandidateInfoId_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofCondCells_r16}, false)
		fn_cg_CandidateToReleaseList_r17 := func() *CG_CandidateInfoId_r17 {
			return new(CG_CandidateInfoId_r17)
		}
		if err = tmp_cg_CandidateToReleaseList_r17.Decode(r, fn_cg_CandidateToReleaseList_r17); err != nil {
			return utils.WrapError("Decode cg_CandidateToReleaseList_r17", err)
		}
		ie.cg_CandidateToReleaseList_r17 = []CG_CandidateInfoId_r17{}
		for _, i := range tmp_cg_CandidateToReleaseList_r17.Value {
			ie.cg_CandidateToReleaseList_r17 = append(ie.cg_CandidateToReleaseList_r17, *i)
		}
	}
	return nil
}
