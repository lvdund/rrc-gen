package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type AvailabilityIndicator_r16 struct {
	ai_RNTI_r16                    AI_RNTI_r16                                `madatory`
	dci_PayloadSizeAI_r16          int64                                      `lb:1,ub:maxAI_DCI_PayloadSize_r16,madatory`
	availableCombToAddModList_r16  []AvailabilityCombinationsPerCell_r16      `lb:1,ub:maxNrofDUCells_r16,optional`
	availableCombToReleaseList_r16 []AvailabilityCombinationsPerCellIndex_r16 `lb:1,ub:maxNrofDUCells_r16,optional`
}

func (ie *AvailabilityIndicator_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{len(ie.availableCombToAddModList_r16) > 0, len(ie.availableCombToReleaseList_r16) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.ai_RNTI_r16.Encode(w); err != nil {
		return utils.WrapError("Encode ai_RNTI_r16", err)
	}
	if err = w.WriteInteger(ie.dci_PayloadSizeAI_r16, &uper.Constraint{Lb: 1, Ub: maxAI_DCI_PayloadSize_r16}, false); err != nil {
		return utils.WrapError("WriteInteger dci_PayloadSizeAI_r16", err)
	}
	if len(ie.availableCombToAddModList_r16) > 0 {
		tmp_availableCombToAddModList_r16 := utils.NewSequence[*AvailabilityCombinationsPerCell_r16]([]*AvailabilityCombinationsPerCell_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofDUCells_r16}, false)
		for _, i := range ie.availableCombToAddModList_r16 {
			tmp_availableCombToAddModList_r16.Value = append(tmp_availableCombToAddModList_r16.Value, &i)
		}
		if err = tmp_availableCombToAddModList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode availableCombToAddModList_r16", err)
		}
	}
	if len(ie.availableCombToReleaseList_r16) > 0 {
		tmp_availableCombToReleaseList_r16 := utils.NewSequence[*AvailabilityCombinationsPerCellIndex_r16]([]*AvailabilityCombinationsPerCellIndex_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofDUCells_r16}, false)
		for _, i := range ie.availableCombToReleaseList_r16 {
			tmp_availableCombToReleaseList_r16.Value = append(tmp_availableCombToReleaseList_r16.Value, &i)
		}
		if err = tmp_availableCombToReleaseList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode availableCombToReleaseList_r16", err)
		}
	}
	return nil
}

func (ie *AvailabilityIndicator_r16) Decode(r *uper.UperReader) error {
	var err error
	var availableCombToAddModList_r16Present bool
	if availableCombToAddModList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var availableCombToReleaseList_r16Present bool
	if availableCombToReleaseList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.ai_RNTI_r16.Decode(r); err != nil {
		return utils.WrapError("Decode ai_RNTI_r16", err)
	}
	var tmp_int_dci_PayloadSizeAI_r16 int64
	if tmp_int_dci_PayloadSizeAI_r16, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: maxAI_DCI_PayloadSize_r16}, false); err != nil {
		return utils.WrapError("ReadInteger dci_PayloadSizeAI_r16", err)
	}
	ie.dci_PayloadSizeAI_r16 = tmp_int_dci_PayloadSizeAI_r16
	if availableCombToAddModList_r16Present {
		tmp_availableCombToAddModList_r16 := utils.NewSequence[*AvailabilityCombinationsPerCell_r16]([]*AvailabilityCombinationsPerCell_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofDUCells_r16}, false)
		fn_availableCombToAddModList_r16 := func() *AvailabilityCombinationsPerCell_r16 {
			return new(AvailabilityCombinationsPerCell_r16)
		}
		if err = tmp_availableCombToAddModList_r16.Decode(r, fn_availableCombToAddModList_r16); err != nil {
			return utils.WrapError("Decode availableCombToAddModList_r16", err)
		}
		ie.availableCombToAddModList_r16 = []AvailabilityCombinationsPerCell_r16{}
		for _, i := range tmp_availableCombToAddModList_r16.Value {
			ie.availableCombToAddModList_r16 = append(ie.availableCombToAddModList_r16, *i)
		}
	}
	if availableCombToReleaseList_r16Present {
		tmp_availableCombToReleaseList_r16 := utils.NewSequence[*AvailabilityCombinationsPerCellIndex_r16]([]*AvailabilityCombinationsPerCellIndex_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofDUCells_r16}, false)
		fn_availableCombToReleaseList_r16 := func() *AvailabilityCombinationsPerCellIndex_r16 {
			return new(AvailabilityCombinationsPerCellIndex_r16)
		}
		if err = tmp_availableCombToReleaseList_r16.Decode(r, fn_availableCombToReleaseList_r16); err != nil {
			return utils.WrapError("Decode availableCombToReleaseList_r16", err)
		}
		ie.availableCombToReleaseList_r16 = []AvailabilityCombinationsPerCellIndex_r16{}
		for _, i := range tmp_availableCombToReleaseList_r16.Value {
			ie.availableCombToReleaseList_r16 = append(ie.availableCombToReleaseList_r16, *i)
		}
	}
	return nil
}
