package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MIMOParam_r17 struct {
	additionalPCI_ToAddModList_r17       []SSB_MTC_AdditionalPCI_r17             `lb:1,ub:maxNrofAdditionalPCI_r17,optional`
	additionalPCI_ToReleaseList_r17      []AdditionalPCIIndex_r17                `lb:1,ub:maxNrofAdditionalPCI_r17,optional`
	unifiedTCI_StateType_r17             *MIMOParam_r17_unifiedTCI_StateType_r17 `optional`
	uplink_PowerControlToAddModList_r17  []Uplink_powerControl_r17               `lb:1,ub:maxUL_TCI_r17,optional`
	uplink_PowerControlToReleaseList_r17 []Uplink_powerControlId_r17             `lb:1,ub:maxUL_TCI_r17,optional`
	sfnSchemePDCCH_r17                   *MIMOParam_r17_sfnSchemePDCCH_r17       `optional`
	sfnSchemePDSCH_r17                   *MIMOParam_r17_sfnSchemePDSCH_r17       `optional`
}

func (ie *MIMOParam_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{len(ie.additionalPCI_ToAddModList_r17) > 0, len(ie.additionalPCI_ToReleaseList_r17) > 0, ie.unifiedTCI_StateType_r17 != nil, len(ie.uplink_PowerControlToAddModList_r17) > 0, len(ie.uplink_PowerControlToReleaseList_r17) > 0, ie.sfnSchemePDCCH_r17 != nil, ie.sfnSchemePDSCH_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.additionalPCI_ToAddModList_r17) > 0 {
		tmp_additionalPCI_ToAddModList_r17 := utils.NewSequence[*SSB_MTC_AdditionalPCI_r17]([]*SSB_MTC_AdditionalPCI_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofAdditionalPCI_r17}, false)
		for _, i := range ie.additionalPCI_ToAddModList_r17 {
			tmp_additionalPCI_ToAddModList_r17.Value = append(tmp_additionalPCI_ToAddModList_r17.Value, &i)
		}
		if err = tmp_additionalPCI_ToAddModList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode additionalPCI_ToAddModList_r17", err)
		}
	}
	if len(ie.additionalPCI_ToReleaseList_r17) > 0 {
		tmp_additionalPCI_ToReleaseList_r17 := utils.NewSequence[*AdditionalPCIIndex_r17]([]*AdditionalPCIIndex_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofAdditionalPCI_r17}, false)
		for _, i := range ie.additionalPCI_ToReleaseList_r17 {
			tmp_additionalPCI_ToReleaseList_r17.Value = append(tmp_additionalPCI_ToReleaseList_r17.Value, &i)
		}
		if err = tmp_additionalPCI_ToReleaseList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode additionalPCI_ToReleaseList_r17", err)
		}
	}
	if ie.unifiedTCI_StateType_r17 != nil {
		if err = ie.unifiedTCI_StateType_r17.Encode(w); err != nil {
			return utils.WrapError("Encode unifiedTCI_StateType_r17", err)
		}
	}
	if len(ie.uplink_PowerControlToAddModList_r17) > 0 {
		tmp_uplink_PowerControlToAddModList_r17 := utils.NewSequence[*Uplink_powerControl_r17]([]*Uplink_powerControl_r17{}, uper.Constraint{Lb: 1, Ub: maxUL_TCI_r17}, false)
		for _, i := range ie.uplink_PowerControlToAddModList_r17 {
			tmp_uplink_PowerControlToAddModList_r17.Value = append(tmp_uplink_PowerControlToAddModList_r17.Value, &i)
		}
		if err = tmp_uplink_PowerControlToAddModList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode uplink_PowerControlToAddModList_r17", err)
		}
	}
	if len(ie.uplink_PowerControlToReleaseList_r17) > 0 {
		tmp_uplink_PowerControlToReleaseList_r17 := utils.NewSequence[*Uplink_powerControlId_r17]([]*Uplink_powerControlId_r17{}, uper.Constraint{Lb: 1, Ub: maxUL_TCI_r17}, false)
		for _, i := range ie.uplink_PowerControlToReleaseList_r17 {
			tmp_uplink_PowerControlToReleaseList_r17.Value = append(tmp_uplink_PowerControlToReleaseList_r17.Value, &i)
		}
		if err = tmp_uplink_PowerControlToReleaseList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode uplink_PowerControlToReleaseList_r17", err)
		}
	}
	if ie.sfnSchemePDCCH_r17 != nil {
		if err = ie.sfnSchemePDCCH_r17.Encode(w); err != nil {
			return utils.WrapError("Encode sfnSchemePDCCH_r17", err)
		}
	}
	if ie.sfnSchemePDSCH_r17 != nil {
		if err = ie.sfnSchemePDSCH_r17.Encode(w); err != nil {
			return utils.WrapError("Encode sfnSchemePDSCH_r17", err)
		}
	}
	return nil
}

func (ie *MIMOParam_r17) Decode(r *uper.UperReader) error {
	var err error
	var additionalPCI_ToAddModList_r17Present bool
	if additionalPCI_ToAddModList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var additionalPCI_ToReleaseList_r17Present bool
	if additionalPCI_ToReleaseList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var unifiedTCI_StateType_r17Present bool
	if unifiedTCI_StateType_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var uplink_PowerControlToAddModList_r17Present bool
	if uplink_PowerControlToAddModList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var uplink_PowerControlToReleaseList_r17Present bool
	if uplink_PowerControlToReleaseList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sfnSchemePDCCH_r17Present bool
	if sfnSchemePDCCH_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sfnSchemePDSCH_r17Present bool
	if sfnSchemePDSCH_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if additionalPCI_ToAddModList_r17Present {
		tmp_additionalPCI_ToAddModList_r17 := utils.NewSequence[*SSB_MTC_AdditionalPCI_r17]([]*SSB_MTC_AdditionalPCI_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofAdditionalPCI_r17}, false)
		fn_additionalPCI_ToAddModList_r17 := func() *SSB_MTC_AdditionalPCI_r17 {
			return new(SSB_MTC_AdditionalPCI_r17)
		}
		if err = tmp_additionalPCI_ToAddModList_r17.Decode(r, fn_additionalPCI_ToAddModList_r17); err != nil {
			return utils.WrapError("Decode additionalPCI_ToAddModList_r17", err)
		}
		ie.additionalPCI_ToAddModList_r17 = []SSB_MTC_AdditionalPCI_r17{}
		for _, i := range tmp_additionalPCI_ToAddModList_r17.Value {
			ie.additionalPCI_ToAddModList_r17 = append(ie.additionalPCI_ToAddModList_r17, *i)
		}
	}
	if additionalPCI_ToReleaseList_r17Present {
		tmp_additionalPCI_ToReleaseList_r17 := utils.NewSequence[*AdditionalPCIIndex_r17]([]*AdditionalPCIIndex_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofAdditionalPCI_r17}, false)
		fn_additionalPCI_ToReleaseList_r17 := func() *AdditionalPCIIndex_r17 {
			return new(AdditionalPCIIndex_r17)
		}
		if err = tmp_additionalPCI_ToReleaseList_r17.Decode(r, fn_additionalPCI_ToReleaseList_r17); err != nil {
			return utils.WrapError("Decode additionalPCI_ToReleaseList_r17", err)
		}
		ie.additionalPCI_ToReleaseList_r17 = []AdditionalPCIIndex_r17{}
		for _, i := range tmp_additionalPCI_ToReleaseList_r17.Value {
			ie.additionalPCI_ToReleaseList_r17 = append(ie.additionalPCI_ToReleaseList_r17, *i)
		}
	}
	if unifiedTCI_StateType_r17Present {
		ie.unifiedTCI_StateType_r17 = new(MIMOParam_r17_unifiedTCI_StateType_r17)
		if err = ie.unifiedTCI_StateType_r17.Decode(r); err != nil {
			return utils.WrapError("Decode unifiedTCI_StateType_r17", err)
		}
	}
	if uplink_PowerControlToAddModList_r17Present {
		tmp_uplink_PowerControlToAddModList_r17 := utils.NewSequence[*Uplink_powerControl_r17]([]*Uplink_powerControl_r17{}, uper.Constraint{Lb: 1, Ub: maxUL_TCI_r17}, false)
		fn_uplink_PowerControlToAddModList_r17 := func() *Uplink_powerControl_r17 {
			return new(Uplink_powerControl_r17)
		}
		if err = tmp_uplink_PowerControlToAddModList_r17.Decode(r, fn_uplink_PowerControlToAddModList_r17); err != nil {
			return utils.WrapError("Decode uplink_PowerControlToAddModList_r17", err)
		}
		ie.uplink_PowerControlToAddModList_r17 = []Uplink_powerControl_r17{}
		for _, i := range tmp_uplink_PowerControlToAddModList_r17.Value {
			ie.uplink_PowerControlToAddModList_r17 = append(ie.uplink_PowerControlToAddModList_r17, *i)
		}
	}
	if uplink_PowerControlToReleaseList_r17Present {
		tmp_uplink_PowerControlToReleaseList_r17 := utils.NewSequence[*Uplink_powerControlId_r17]([]*Uplink_powerControlId_r17{}, uper.Constraint{Lb: 1, Ub: maxUL_TCI_r17}, false)
		fn_uplink_PowerControlToReleaseList_r17 := func() *Uplink_powerControlId_r17 {
			return new(Uplink_powerControlId_r17)
		}
		if err = tmp_uplink_PowerControlToReleaseList_r17.Decode(r, fn_uplink_PowerControlToReleaseList_r17); err != nil {
			return utils.WrapError("Decode uplink_PowerControlToReleaseList_r17", err)
		}
		ie.uplink_PowerControlToReleaseList_r17 = []Uplink_powerControlId_r17{}
		for _, i := range tmp_uplink_PowerControlToReleaseList_r17.Value {
			ie.uplink_PowerControlToReleaseList_r17 = append(ie.uplink_PowerControlToReleaseList_r17, *i)
		}
	}
	if sfnSchemePDCCH_r17Present {
		ie.sfnSchemePDCCH_r17 = new(MIMOParam_r17_sfnSchemePDCCH_r17)
		if err = ie.sfnSchemePDCCH_r17.Decode(r); err != nil {
			return utils.WrapError("Decode sfnSchemePDCCH_r17", err)
		}
	}
	if sfnSchemePDSCH_r17Present {
		ie.sfnSchemePDSCH_r17 = new(MIMOParam_r17_sfnSchemePDSCH_r17)
		if err = ie.sfnSchemePDSCH_r17.Decode(r); err != nil {
			return utils.WrapError("Decode sfnSchemePDSCH_r17", err)
		}
	}
	return nil
}
