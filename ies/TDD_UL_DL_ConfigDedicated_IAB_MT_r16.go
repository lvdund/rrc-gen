package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type TDD_UL_DL_ConfigDedicated_IAB_MT_r16 struct {
	slotSpecificConfigurationsToAddModList_IAB_MT_r16  []TDD_UL_DL_SlotConfig_IAB_MT_r16 `lb:1,ub:maxNrofSlots,optional`
	slotSpecificConfigurationsToReleaseList_IAB_MT_r16 []TDD_UL_DL_SlotIndex             `lb:1,ub:maxNrofSlots,optional`
}

func (ie *TDD_UL_DL_ConfigDedicated_IAB_MT_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{len(ie.slotSpecificConfigurationsToAddModList_IAB_MT_r16) > 0, len(ie.slotSpecificConfigurationsToReleaseList_IAB_MT_r16) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.slotSpecificConfigurationsToAddModList_IAB_MT_r16) > 0 {
		tmp_slotSpecificConfigurationsToAddModList_IAB_MT_r16 := utils.NewSequence[*TDD_UL_DL_SlotConfig_IAB_MT_r16]([]*TDD_UL_DL_SlotConfig_IAB_MT_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofSlots}, false)
		for _, i := range ie.slotSpecificConfigurationsToAddModList_IAB_MT_r16 {
			tmp_slotSpecificConfigurationsToAddModList_IAB_MT_r16.Value = append(tmp_slotSpecificConfigurationsToAddModList_IAB_MT_r16.Value, &i)
		}
		if err = tmp_slotSpecificConfigurationsToAddModList_IAB_MT_r16.Encode(w); err != nil {
			return utils.WrapError("Encode slotSpecificConfigurationsToAddModList_IAB_MT_r16", err)
		}
	}
	if len(ie.slotSpecificConfigurationsToReleaseList_IAB_MT_r16) > 0 {
		tmp_slotSpecificConfigurationsToReleaseList_IAB_MT_r16 := utils.NewSequence[*TDD_UL_DL_SlotIndex]([]*TDD_UL_DL_SlotIndex{}, uper.Constraint{Lb: 1, Ub: maxNrofSlots}, false)
		for _, i := range ie.slotSpecificConfigurationsToReleaseList_IAB_MT_r16 {
			tmp_slotSpecificConfigurationsToReleaseList_IAB_MT_r16.Value = append(tmp_slotSpecificConfigurationsToReleaseList_IAB_MT_r16.Value, &i)
		}
		if err = tmp_slotSpecificConfigurationsToReleaseList_IAB_MT_r16.Encode(w); err != nil {
			return utils.WrapError("Encode slotSpecificConfigurationsToReleaseList_IAB_MT_r16", err)
		}
	}
	return nil
}

func (ie *TDD_UL_DL_ConfigDedicated_IAB_MT_r16) Decode(r *uper.UperReader) error {
	var err error
	var slotSpecificConfigurationsToAddModList_IAB_MT_r16Present bool
	if slotSpecificConfigurationsToAddModList_IAB_MT_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var slotSpecificConfigurationsToReleaseList_IAB_MT_r16Present bool
	if slotSpecificConfigurationsToReleaseList_IAB_MT_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if slotSpecificConfigurationsToAddModList_IAB_MT_r16Present {
		tmp_slotSpecificConfigurationsToAddModList_IAB_MT_r16 := utils.NewSequence[*TDD_UL_DL_SlotConfig_IAB_MT_r16]([]*TDD_UL_DL_SlotConfig_IAB_MT_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofSlots}, false)
		fn_slotSpecificConfigurationsToAddModList_IAB_MT_r16 := func() *TDD_UL_DL_SlotConfig_IAB_MT_r16 {
			return new(TDD_UL_DL_SlotConfig_IAB_MT_r16)
		}
		if err = tmp_slotSpecificConfigurationsToAddModList_IAB_MT_r16.Decode(r, fn_slotSpecificConfigurationsToAddModList_IAB_MT_r16); err != nil {
			return utils.WrapError("Decode slotSpecificConfigurationsToAddModList_IAB_MT_r16", err)
		}
		ie.slotSpecificConfigurationsToAddModList_IAB_MT_r16 = []TDD_UL_DL_SlotConfig_IAB_MT_r16{}
		for _, i := range tmp_slotSpecificConfigurationsToAddModList_IAB_MT_r16.Value {
			ie.slotSpecificConfigurationsToAddModList_IAB_MT_r16 = append(ie.slotSpecificConfigurationsToAddModList_IAB_MT_r16, *i)
		}
	}
	if slotSpecificConfigurationsToReleaseList_IAB_MT_r16Present {
		tmp_slotSpecificConfigurationsToReleaseList_IAB_MT_r16 := utils.NewSequence[*TDD_UL_DL_SlotIndex]([]*TDD_UL_DL_SlotIndex{}, uper.Constraint{Lb: 1, Ub: maxNrofSlots}, false)
		fn_slotSpecificConfigurationsToReleaseList_IAB_MT_r16 := func() *TDD_UL_DL_SlotIndex {
			return new(TDD_UL_DL_SlotIndex)
		}
		if err = tmp_slotSpecificConfigurationsToReleaseList_IAB_MT_r16.Decode(r, fn_slotSpecificConfigurationsToReleaseList_IAB_MT_r16); err != nil {
			return utils.WrapError("Decode slotSpecificConfigurationsToReleaseList_IAB_MT_r16", err)
		}
		ie.slotSpecificConfigurationsToReleaseList_IAB_MT_r16 = []TDD_UL_DL_SlotIndex{}
		for _, i := range tmp_slotSpecificConfigurationsToReleaseList_IAB_MT_r16.Value {
			ie.slotSpecificConfigurationsToReleaseList_IAB_MT_r16 = append(ie.slotSpecificConfigurationsToReleaseList_IAB_MT_r16, *i)
		}
	}
	return nil
}
