package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PUSCH_PowerControl struct {
	tpc_Accumulation                 *PUSCH_PowerControl_tpc_Accumulation             `optional`
	msg3_Alpha                       *Alpha                                           `optional`
	p0_NominalWithoutGrant           *int64                                           `lb:-202,ub:24,optional`
	p0_AlphaSets                     []P0_PUSCH_AlphaSet                              `lb:1,ub:maxNrofP0_PUSCH_AlphaSets,optional`
	pathlossReferenceRSToAddModList  []PUSCH_PathlossReferenceRS                      `lb:1,ub:maxNrofPUSCH_PathlossReferenceRSs,optional`
	pathlossReferenceRSToReleaseList []PUSCH_PathlossReferenceRS_Id                   `lb:1,ub:maxNrofPUSCH_PathlossReferenceRSs,optional`
	twoPUSCH_PC_AdjustmentStates     *PUSCH_PowerControl_twoPUSCH_PC_AdjustmentStates `optional`
	deltaMCS                         *PUSCH_PowerControl_deltaMCS                     `optional`
	sri_PUSCH_MappingToAddModList    []SRI_PUSCH_PowerControl                         `lb:1,ub:maxNrofSRI_PUSCH_Mappings,optional`
	sri_PUSCH_MappingToReleaseList   []SRI_PUSCH_PowerControlId                       `lb:1,ub:maxNrofSRI_PUSCH_Mappings,optional`
}

func (ie *PUSCH_PowerControl) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.tpc_Accumulation != nil, ie.msg3_Alpha != nil, ie.p0_NominalWithoutGrant != nil, len(ie.p0_AlphaSets) > 0, len(ie.pathlossReferenceRSToAddModList) > 0, len(ie.pathlossReferenceRSToReleaseList) > 0, ie.twoPUSCH_PC_AdjustmentStates != nil, ie.deltaMCS != nil, len(ie.sri_PUSCH_MappingToAddModList) > 0, len(ie.sri_PUSCH_MappingToReleaseList) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.tpc_Accumulation != nil {
		if err = ie.tpc_Accumulation.Encode(w); err != nil {
			return utils.WrapError("Encode tpc_Accumulation", err)
		}
	}
	if ie.msg3_Alpha != nil {
		if err = ie.msg3_Alpha.Encode(w); err != nil {
			return utils.WrapError("Encode msg3_Alpha", err)
		}
	}
	if ie.p0_NominalWithoutGrant != nil {
		if err = w.WriteInteger(*ie.p0_NominalWithoutGrant, &uper.Constraint{Lb: -202, Ub: 24}, false); err != nil {
			return utils.WrapError("Encode p0_NominalWithoutGrant", err)
		}
	}
	if len(ie.p0_AlphaSets) > 0 {
		tmp_p0_AlphaSets := utils.NewSequence[*P0_PUSCH_AlphaSet]([]*P0_PUSCH_AlphaSet{}, uper.Constraint{Lb: 1, Ub: maxNrofP0_PUSCH_AlphaSets}, false)
		for _, i := range ie.p0_AlphaSets {
			tmp_p0_AlphaSets.Value = append(tmp_p0_AlphaSets.Value, &i)
		}
		if err = tmp_p0_AlphaSets.Encode(w); err != nil {
			return utils.WrapError("Encode p0_AlphaSets", err)
		}
	}
	if len(ie.pathlossReferenceRSToAddModList) > 0 {
		tmp_pathlossReferenceRSToAddModList := utils.NewSequence[*PUSCH_PathlossReferenceRS]([]*PUSCH_PathlossReferenceRS{}, uper.Constraint{Lb: 1, Ub: maxNrofPUSCH_PathlossReferenceRSs}, false)
		for _, i := range ie.pathlossReferenceRSToAddModList {
			tmp_pathlossReferenceRSToAddModList.Value = append(tmp_pathlossReferenceRSToAddModList.Value, &i)
		}
		if err = tmp_pathlossReferenceRSToAddModList.Encode(w); err != nil {
			return utils.WrapError("Encode pathlossReferenceRSToAddModList", err)
		}
	}
	if len(ie.pathlossReferenceRSToReleaseList) > 0 {
		tmp_pathlossReferenceRSToReleaseList := utils.NewSequence[*PUSCH_PathlossReferenceRS_Id]([]*PUSCH_PathlossReferenceRS_Id{}, uper.Constraint{Lb: 1, Ub: maxNrofPUSCH_PathlossReferenceRSs}, false)
		for _, i := range ie.pathlossReferenceRSToReleaseList {
			tmp_pathlossReferenceRSToReleaseList.Value = append(tmp_pathlossReferenceRSToReleaseList.Value, &i)
		}
		if err = tmp_pathlossReferenceRSToReleaseList.Encode(w); err != nil {
			return utils.WrapError("Encode pathlossReferenceRSToReleaseList", err)
		}
	}
	if ie.twoPUSCH_PC_AdjustmentStates != nil {
		if err = ie.twoPUSCH_PC_AdjustmentStates.Encode(w); err != nil {
			return utils.WrapError("Encode twoPUSCH_PC_AdjustmentStates", err)
		}
	}
	if ie.deltaMCS != nil {
		if err = ie.deltaMCS.Encode(w); err != nil {
			return utils.WrapError("Encode deltaMCS", err)
		}
	}
	if len(ie.sri_PUSCH_MappingToAddModList) > 0 {
		tmp_sri_PUSCH_MappingToAddModList := utils.NewSequence[*SRI_PUSCH_PowerControl]([]*SRI_PUSCH_PowerControl{}, uper.Constraint{Lb: 1, Ub: maxNrofSRI_PUSCH_Mappings}, false)
		for _, i := range ie.sri_PUSCH_MappingToAddModList {
			tmp_sri_PUSCH_MappingToAddModList.Value = append(tmp_sri_PUSCH_MappingToAddModList.Value, &i)
		}
		if err = tmp_sri_PUSCH_MappingToAddModList.Encode(w); err != nil {
			return utils.WrapError("Encode sri_PUSCH_MappingToAddModList", err)
		}
	}
	if len(ie.sri_PUSCH_MappingToReleaseList) > 0 {
		tmp_sri_PUSCH_MappingToReleaseList := utils.NewSequence[*SRI_PUSCH_PowerControlId]([]*SRI_PUSCH_PowerControlId{}, uper.Constraint{Lb: 1, Ub: maxNrofSRI_PUSCH_Mappings}, false)
		for _, i := range ie.sri_PUSCH_MappingToReleaseList {
			tmp_sri_PUSCH_MappingToReleaseList.Value = append(tmp_sri_PUSCH_MappingToReleaseList.Value, &i)
		}
		if err = tmp_sri_PUSCH_MappingToReleaseList.Encode(w); err != nil {
			return utils.WrapError("Encode sri_PUSCH_MappingToReleaseList", err)
		}
	}
	return nil
}

func (ie *PUSCH_PowerControl) Decode(r *uper.UperReader) error {
	var err error
	var tpc_AccumulationPresent bool
	if tpc_AccumulationPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var msg3_AlphaPresent bool
	if msg3_AlphaPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var p0_NominalWithoutGrantPresent bool
	if p0_NominalWithoutGrantPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var p0_AlphaSetsPresent bool
	if p0_AlphaSetsPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var pathlossReferenceRSToAddModListPresent bool
	if pathlossReferenceRSToAddModListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var pathlossReferenceRSToReleaseListPresent bool
	if pathlossReferenceRSToReleaseListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var twoPUSCH_PC_AdjustmentStatesPresent bool
	if twoPUSCH_PC_AdjustmentStatesPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var deltaMCSPresent bool
	if deltaMCSPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var sri_PUSCH_MappingToAddModListPresent bool
	if sri_PUSCH_MappingToAddModListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var sri_PUSCH_MappingToReleaseListPresent bool
	if sri_PUSCH_MappingToReleaseListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if tpc_AccumulationPresent {
		ie.tpc_Accumulation = new(PUSCH_PowerControl_tpc_Accumulation)
		if err = ie.tpc_Accumulation.Decode(r); err != nil {
			return utils.WrapError("Decode tpc_Accumulation", err)
		}
	}
	if msg3_AlphaPresent {
		ie.msg3_Alpha = new(Alpha)
		if err = ie.msg3_Alpha.Decode(r); err != nil {
			return utils.WrapError("Decode msg3_Alpha", err)
		}
	}
	if p0_NominalWithoutGrantPresent {
		var tmp_int_p0_NominalWithoutGrant int64
		if tmp_int_p0_NominalWithoutGrant, err = r.ReadInteger(&uper.Constraint{Lb: -202, Ub: 24}, false); err != nil {
			return utils.WrapError("Decode p0_NominalWithoutGrant", err)
		}
		ie.p0_NominalWithoutGrant = &tmp_int_p0_NominalWithoutGrant
	}
	if p0_AlphaSetsPresent {
		tmp_p0_AlphaSets := utils.NewSequence[*P0_PUSCH_AlphaSet]([]*P0_PUSCH_AlphaSet{}, uper.Constraint{Lb: 1, Ub: maxNrofP0_PUSCH_AlphaSets}, false)
		fn_p0_AlphaSets := func() *P0_PUSCH_AlphaSet {
			return new(P0_PUSCH_AlphaSet)
		}
		if err = tmp_p0_AlphaSets.Decode(r, fn_p0_AlphaSets); err != nil {
			return utils.WrapError("Decode p0_AlphaSets", err)
		}
		ie.p0_AlphaSets = []P0_PUSCH_AlphaSet{}
		for _, i := range tmp_p0_AlphaSets.Value {
			ie.p0_AlphaSets = append(ie.p0_AlphaSets, *i)
		}
	}
	if pathlossReferenceRSToAddModListPresent {
		tmp_pathlossReferenceRSToAddModList := utils.NewSequence[*PUSCH_PathlossReferenceRS]([]*PUSCH_PathlossReferenceRS{}, uper.Constraint{Lb: 1, Ub: maxNrofPUSCH_PathlossReferenceRSs}, false)
		fn_pathlossReferenceRSToAddModList := func() *PUSCH_PathlossReferenceRS {
			return new(PUSCH_PathlossReferenceRS)
		}
		if err = tmp_pathlossReferenceRSToAddModList.Decode(r, fn_pathlossReferenceRSToAddModList); err != nil {
			return utils.WrapError("Decode pathlossReferenceRSToAddModList", err)
		}
		ie.pathlossReferenceRSToAddModList = []PUSCH_PathlossReferenceRS{}
		for _, i := range tmp_pathlossReferenceRSToAddModList.Value {
			ie.pathlossReferenceRSToAddModList = append(ie.pathlossReferenceRSToAddModList, *i)
		}
	}
	if pathlossReferenceRSToReleaseListPresent {
		tmp_pathlossReferenceRSToReleaseList := utils.NewSequence[*PUSCH_PathlossReferenceRS_Id]([]*PUSCH_PathlossReferenceRS_Id{}, uper.Constraint{Lb: 1, Ub: maxNrofPUSCH_PathlossReferenceRSs}, false)
		fn_pathlossReferenceRSToReleaseList := func() *PUSCH_PathlossReferenceRS_Id {
			return new(PUSCH_PathlossReferenceRS_Id)
		}
		if err = tmp_pathlossReferenceRSToReleaseList.Decode(r, fn_pathlossReferenceRSToReleaseList); err != nil {
			return utils.WrapError("Decode pathlossReferenceRSToReleaseList", err)
		}
		ie.pathlossReferenceRSToReleaseList = []PUSCH_PathlossReferenceRS_Id{}
		for _, i := range tmp_pathlossReferenceRSToReleaseList.Value {
			ie.pathlossReferenceRSToReleaseList = append(ie.pathlossReferenceRSToReleaseList, *i)
		}
	}
	if twoPUSCH_PC_AdjustmentStatesPresent {
		ie.twoPUSCH_PC_AdjustmentStates = new(PUSCH_PowerControl_twoPUSCH_PC_AdjustmentStates)
		if err = ie.twoPUSCH_PC_AdjustmentStates.Decode(r); err != nil {
			return utils.WrapError("Decode twoPUSCH_PC_AdjustmentStates", err)
		}
	}
	if deltaMCSPresent {
		ie.deltaMCS = new(PUSCH_PowerControl_deltaMCS)
		if err = ie.deltaMCS.Decode(r); err != nil {
			return utils.WrapError("Decode deltaMCS", err)
		}
	}
	if sri_PUSCH_MappingToAddModListPresent {
		tmp_sri_PUSCH_MappingToAddModList := utils.NewSequence[*SRI_PUSCH_PowerControl]([]*SRI_PUSCH_PowerControl{}, uper.Constraint{Lb: 1, Ub: maxNrofSRI_PUSCH_Mappings}, false)
		fn_sri_PUSCH_MappingToAddModList := func() *SRI_PUSCH_PowerControl {
			return new(SRI_PUSCH_PowerControl)
		}
		if err = tmp_sri_PUSCH_MappingToAddModList.Decode(r, fn_sri_PUSCH_MappingToAddModList); err != nil {
			return utils.WrapError("Decode sri_PUSCH_MappingToAddModList", err)
		}
		ie.sri_PUSCH_MappingToAddModList = []SRI_PUSCH_PowerControl{}
		for _, i := range tmp_sri_PUSCH_MappingToAddModList.Value {
			ie.sri_PUSCH_MappingToAddModList = append(ie.sri_PUSCH_MappingToAddModList, *i)
		}
	}
	if sri_PUSCH_MappingToReleaseListPresent {
		tmp_sri_PUSCH_MappingToReleaseList := utils.NewSequence[*SRI_PUSCH_PowerControlId]([]*SRI_PUSCH_PowerControlId{}, uper.Constraint{Lb: 1, Ub: maxNrofSRI_PUSCH_Mappings}, false)
		fn_sri_PUSCH_MappingToReleaseList := func() *SRI_PUSCH_PowerControlId {
			return new(SRI_PUSCH_PowerControlId)
		}
		if err = tmp_sri_PUSCH_MappingToReleaseList.Decode(r, fn_sri_PUSCH_MappingToReleaseList); err != nil {
			return utils.WrapError("Decode sri_PUSCH_MappingToReleaseList", err)
		}
		ie.sri_PUSCH_MappingToReleaseList = []SRI_PUSCH_PowerControlId{}
		for _, i := range tmp_sri_PUSCH_MappingToReleaseList.Value {
			ie.sri_PUSCH_MappingToReleaseList = append(ie.sri_PUSCH_MappingToReleaseList, *i)
		}
	}
	return nil
}
