package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type TDD_UL_DL_ConfigDedicated struct {
	slotSpecificConfigurationsToAddModList  []TDD_UL_DL_SlotConfig `lb:1,ub:maxNrofSlots,optional`
	slotSpecificConfigurationsToReleaseList []TDD_UL_DL_SlotIndex  `lb:1,ub:maxNrofSlots,optional`
}

func (ie *TDD_UL_DL_ConfigDedicated) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{len(ie.slotSpecificConfigurationsToAddModList) > 0, len(ie.slotSpecificConfigurationsToReleaseList) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.slotSpecificConfigurationsToAddModList) > 0 {
		tmp_slotSpecificConfigurationsToAddModList := utils.NewSequence[*TDD_UL_DL_SlotConfig]([]*TDD_UL_DL_SlotConfig{}, uper.Constraint{Lb: 1, Ub: maxNrofSlots}, false)
		for _, i := range ie.slotSpecificConfigurationsToAddModList {
			tmp_slotSpecificConfigurationsToAddModList.Value = append(tmp_slotSpecificConfigurationsToAddModList.Value, &i)
		}
		if err = tmp_slotSpecificConfigurationsToAddModList.Encode(w); err != nil {
			return utils.WrapError("Encode slotSpecificConfigurationsToAddModList", err)
		}
	}
	if len(ie.slotSpecificConfigurationsToReleaseList) > 0 {
		tmp_slotSpecificConfigurationsToReleaseList := utils.NewSequence[*TDD_UL_DL_SlotIndex]([]*TDD_UL_DL_SlotIndex{}, uper.Constraint{Lb: 1, Ub: maxNrofSlots}, false)
		for _, i := range ie.slotSpecificConfigurationsToReleaseList {
			tmp_slotSpecificConfigurationsToReleaseList.Value = append(tmp_slotSpecificConfigurationsToReleaseList.Value, &i)
		}
		if err = tmp_slotSpecificConfigurationsToReleaseList.Encode(w); err != nil {
			return utils.WrapError("Encode slotSpecificConfigurationsToReleaseList", err)
		}
	}
	return nil
}

func (ie *TDD_UL_DL_ConfigDedicated) Decode(r *uper.UperReader) error {
	var err error
	var slotSpecificConfigurationsToAddModListPresent bool
	if slotSpecificConfigurationsToAddModListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var slotSpecificConfigurationsToReleaseListPresent bool
	if slotSpecificConfigurationsToReleaseListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if slotSpecificConfigurationsToAddModListPresent {
		tmp_slotSpecificConfigurationsToAddModList := utils.NewSequence[*TDD_UL_DL_SlotConfig]([]*TDD_UL_DL_SlotConfig{}, uper.Constraint{Lb: 1, Ub: maxNrofSlots}, false)
		fn_slotSpecificConfigurationsToAddModList := func() *TDD_UL_DL_SlotConfig {
			return new(TDD_UL_DL_SlotConfig)
		}
		if err = tmp_slotSpecificConfigurationsToAddModList.Decode(r, fn_slotSpecificConfigurationsToAddModList); err != nil {
			return utils.WrapError("Decode slotSpecificConfigurationsToAddModList", err)
		}
		ie.slotSpecificConfigurationsToAddModList = []TDD_UL_DL_SlotConfig{}
		for _, i := range tmp_slotSpecificConfigurationsToAddModList.Value {
			ie.slotSpecificConfigurationsToAddModList = append(ie.slotSpecificConfigurationsToAddModList, *i)
		}
	}
	if slotSpecificConfigurationsToReleaseListPresent {
		tmp_slotSpecificConfigurationsToReleaseList := utils.NewSequence[*TDD_UL_DL_SlotIndex]([]*TDD_UL_DL_SlotIndex{}, uper.Constraint{Lb: 1, Ub: maxNrofSlots}, false)
		fn_slotSpecificConfigurationsToReleaseList := func() *TDD_UL_DL_SlotIndex {
			return new(TDD_UL_DL_SlotIndex)
		}
		if err = tmp_slotSpecificConfigurationsToReleaseList.Decode(r, fn_slotSpecificConfigurationsToReleaseList); err != nil {
			return utils.WrapError("Decode slotSpecificConfigurationsToReleaseList", err)
		}
		ie.slotSpecificConfigurationsToReleaseList = []TDD_UL_DL_SlotIndex{}
		for _, i := range tmp_slotSpecificConfigurationsToReleaseList.Value {
			ie.slotSpecificConfigurationsToReleaseList = append(ie.slotSpecificConfigurationsToReleaseList, *i)
		}
	}
	return nil
}
