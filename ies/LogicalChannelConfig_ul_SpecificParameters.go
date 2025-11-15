package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type LogicalChannelConfig_ul_SpecificParameters struct {
	priority                           int64                                                                    `lb:1,ub:16,madatory`
	prioritisedBitRate                 LogicalChannelConfig_ul_SpecificParameters_prioritisedBitRate            `madatory`
	bucketSizeDuration                 LogicalChannelConfig_ul_SpecificParameters_bucketSizeDuration            `madatory`
	allowedServingCells                []ServCellIndex                                                          `lb:1,ub:maxNrofServingCells_1,optional`
	allowedSCS_List                    []SubcarrierSpacing                                                      `lb:1,ub:maxSCSs,optional`
	maxPUSCH_Duration                  *LogicalChannelConfig_ul_SpecificParameters_maxPUSCH_Duration            `optional`
	configuredGrantType1Allowed        *LogicalChannelConfig_ul_SpecificParameters_configuredGrantType1Allowed  `optional`
	logicalChannelGroup                *int64                                                                   `lb:0,ub:maxLCG_ID,optional`
	schedulingRequestID                *SchedulingRequestId                                                     `optional`
	logicalChannelSR_Mask              bool                                                                     `madatory`
	logicalChannelSR_DelayTimerApplied bool                                                                     `madatory`
	bitRateQueryProhibitTimer          *LogicalChannelConfig_ul_SpecificParameters_bitRateQueryProhibitTimer    `optional`
	allowedCG_List_r16                 []ConfiguredGrantConfigIndexMAC_r16                                      `lb:0,ub:maxNrofConfiguredGrantConfigMAC_1_r16,optional`
	allowedPHY_PriorityIndex_r16       *LogicalChannelConfig_ul_SpecificParameters_allowedPHY_PriorityIndex_r16 `optional`
	logicalChannelGroupIAB_Ext_r17     *int64                                                                   `lb:0,ub:maxLCG_ID_IAB_r17,optional`
	allowedHARQ_mode_r17               *LogicalChannelConfig_ul_SpecificParameters_allowedHARQ_mode_r17         `optional`
}

func (ie *LogicalChannelConfig_ul_SpecificParameters) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{len(ie.allowedServingCells) > 0, len(ie.allowedSCS_List) > 0, ie.maxPUSCH_Duration != nil, ie.configuredGrantType1Allowed != nil, ie.logicalChannelGroup != nil, ie.schedulingRequestID != nil, ie.bitRateQueryProhibitTimer != nil, len(ie.allowedCG_List_r16) > 0, ie.allowedPHY_PriorityIndex_r16 != nil, ie.logicalChannelGroupIAB_Ext_r17 != nil, ie.allowedHARQ_mode_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = w.WriteInteger(ie.priority, &uper.Constraint{Lb: 1, Ub: 16}, false); err != nil {
		return utils.WrapError("WriteInteger priority", err)
	}
	if err = ie.prioritisedBitRate.Encode(w); err != nil {
		return utils.WrapError("Encode prioritisedBitRate", err)
	}
	if err = ie.bucketSizeDuration.Encode(w); err != nil {
		return utils.WrapError("Encode bucketSizeDuration", err)
	}
	if len(ie.allowedServingCells) > 0 {
		tmp_allowedServingCells := utils.NewSequence[*ServCellIndex]([]*ServCellIndex{}, uper.Constraint{Lb: 1, Ub: maxNrofServingCells_1}, false)
		for _, i := range ie.allowedServingCells {
			tmp_allowedServingCells.Value = append(tmp_allowedServingCells.Value, &i)
		}
		if err = tmp_allowedServingCells.Encode(w); err != nil {
			return utils.WrapError("Encode allowedServingCells", err)
		}
	}
	if len(ie.allowedSCS_List) > 0 {
		tmp_allowedSCS_List := utils.NewSequence[*SubcarrierSpacing]([]*SubcarrierSpacing{}, uper.Constraint{Lb: 1, Ub: maxSCSs}, false)
		for _, i := range ie.allowedSCS_List {
			tmp_allowedSCS_List.Value = append(tmp_allowedSCS_List.Value, &i)
		}
		if err = tmp_allowedSCS_List.Encode(w); err != nil {
			return utils.WrapError("Encode allowedSCS_List", err)
		}
	}
	if ie.maxPUSCH_Duration != nil {
		if err = ie.maxPUSCH_Duration.Encode(w); err != nil {
			return utils.WrapError("Encode maxPUSCH_Duration", err)
		}
	}
	if ie.configuredGrantType1Allowed != nil {
		if err = ie.configuredGrantType1Allowed.Encode(w); err != nil {
			return utils.WrapError("Encode configuredGrantType1Allowed", err)
		}
	}
	if ie.logicalChannelGroup != nil {
		if err = w.WriteInteger(*ie.logicalChannelGroup, &uper.Constraint{Lb: 0, Ub: maxLCG_ID}, false); err != nil {
			return utils.WrapError("Encode logicalChannelGroup", err)
		}
	}
	if ie.schedulingRequestID != nil {
		if err = ie.schedulingRequestID.Encode(w); err != nil {
			return utils.WrapError("Encode schedulingRequestID", err)
		}
	}
	if err = w.WriteBoolean(ie.logicalChannelSR_Mask); err != nil {
		return utils.WrapError("WriteBoolean logicalChannelSR_Mask", err)
	}
	if err = w.WriteBoolean(ie.logicalChannelSR_DelayTimerApplied); err != nil {
		return utils.WrapError("WriteBoolean logicalChannelSR_DelayTimerApplied", err)
	}
	if ie.bitRateQueryProhibitTimer != nil {
		if err = ie.bitRateQueryProhibitTimer.Encode(w); err != nil {
			return utils.WrapError("Encode bitRateQueryProhibitTimer", err)
		}
	}
	if len(ie.allowedCG_List_r16) > 0 {
		tmp_allowedCG_List_r16 := utils.NewSequence[*ConfiguredGrantConfigIndexMAC_r16]([]*ConfiguredGrantConfigIndexMAC_r16{}, uper.Constraint{Lb: 0, Ub: maxNrofConfiguredGrantConfigMAC_1_r16}, false)
		for _, i := range ie.allowedCG_List_r16 {
			tmp_allowedCG_List_r16.Value = append(tmp_allowedCG_List_r16.Value, &i)
		}
		if err = tmp_allowedCG_List_r16.Encode(w); err != nil {
			return utils.WrapError("Encode allowedCG_List_r16", err)
		}
	}
	if ie.allowedPHY_PriorityIndex_r16 != nil {
		if err = ie.allowedPHY_PriorityIndex_r16.Encode(w); err != nil {
			return utils.WrapError("Encode allowedPHY_PriorityIndex_r16", err)
		}
	}
	if ie.logicalChannelGroupIAB_Ext_r17 != nil {
		if err = w.WriteInteger(*ie.logicalChannelGroupIAB_Ext_r17, &uper.Constraint{Lb: 0, Ub: maxLCG_ID_IAB_r17}, false); err != nil {
			return utils.WrapError("Encode logicalChannelGroupIAB_Ext_r17", err)
		}
	}
	if ie.allowedHARQ_mode_r17 != nil {
		if err = ie.allowedHARQ_mode_r17.Encode(w); err != nil {
			return utils.WrapError("Encode allowedHARQ_mode_r17", err)
		}
	}
	return nil
}

func (ie *LogicalChannelConfig_ul_SpecificParameters) Decode(r *uper.UperReader) error {
	var err error
	var allowedServingCellsPresent bool
	if allowedServingCellsPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var allowedSCS_ListPresent bool
	if allowedSCS_ListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var maxPUSCH_DurationPresent bool
	if maxPUSCH_DurationPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var configuredGrantType1AllowedPresent bool
	if configuredGrantType1AllowedPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var logicalChannelGroupPresent bool
	if logicalChannelGroupPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var schedulingRequestIDPresent bool
	if schedulingRequestIDPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var bitRateQueryProhibitTimerPresent bool
	if bitRateQueryProhibitTimerPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var allowedCG_List_r16Present bool
	if allowedCG_List_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var allowedPHY_PriorityIndex_r16Present bool
	if allowedPHY_PriorityIndex_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var logicalChannelGroupIAB_Ext_r17Present bool
	if logicalChannelGroupIAB_Ext_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var allowedHARQ_mode_r17Present bool
	if allowedHARQ_mode_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var tmp_int_priority int64
	if tmp_int_priority, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 16}, false); err != nil {
		return utils.WrapError("ReadInteger priority", err)
	}
	ie.priority = tmp_int_priority
	if err = ie.prioritisedBitRate.Decode(r); err != nil {
		return utils.WrapError("Decode prioritisedBitRate", err)
	}
	if err = ie.bucketSizeDuration.Decode(r); err != nil {
		return utils.WrapError("Decode bucketSizeDuration", err)
	}
	if allowedServingCellsPresent {
		tmp_allowedServingCells := utils.NewSequence[*ServCellIndex]([]*ServCellIndex{}, uper.Constraint{Lb: 1, Ub: maxNrofServingCells_1}, false)
		fn_allowedServingCells := func() *ServCellIndex {
			return new(ServCellIndex)
		}
		if err = tmp_allowedServingCells.Decode(r, fn_allowedServingCells); err != nil {
			return utils.WrapError("Decode allowedServingCells", err)
		}
		ie.allowedServingCells = []ServCellIndex{}
		for _, i := range tmp_allowedServingCells.Value {
			ie.allowedServingCells = append(ie.allowedServingCells, *i)
		}
	}
	if allowedSCS_ListPresent {
		tmp_allowedSCS_List := utils.NewSequence[*SubcarrierSpacing]([]*SubcarrierSpacing{}, uper.Constraint{Lb: 1, Ub: maxSCSs}, false)
		fn_allowedSCS_List := func() *SubcarrierSpacing {
			return new(SubcarrierSpacing)
		}
		if err = tmp_allowedSCS_List.Decode(r, fn_allowedSCS_List); err != nil {
			return utils.WrapError("Decode allowedSCS_List", err)
		}
		ie.allowedSCS_List = []SubcarrierSpacing{}
		for _, i := range tmp_allowedSCS_List.Value {
			ie.allowedSCS_List = append(ie.allowedSCS_List, *i)
		}
	}
	if maxPUSCH_DurationPresent {
		ie.maxPUSCH_Duration = new(LogicalChannelConfig_ul_SpecificParameters_maxPUSCH_Duration)
		if err = ie.maxPUSCH_Duration.Decode(r); err != nil {
			return utils.WrapError("Decode maxPUSCH_Duration", err)
		}
	}
	if configuredGrantType1AllowedPresent {
		ie.configuredGrantType1Allowed = new(LogicalChannelConfig_ul_SpecificParameters_configuredGrantType1Allowed)
		if err = ie.configuredGrantType1Allowed.Decode(r); err != nil {
			return utils.WrapError("Decode configuredGrantType1Allowed", err)
		}
	}
	if logicalChannelGroupPresent {
		var tmp_int_logicalChannelGroup int64
		if tmp_int_logicalChannelGroup, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: maxLCG_ID}, false); err != nil {
			return utils.WrapError("Decode logicalChannelGroup", err)
		}
		ie.logicalChannelGroup = &tmp_int_logicalChannelGroup
	}
	if schedulingRequestIDPresent {
		ie.schedulingRequestID = new(SchedulingRequestId)
		if err = ie.schedulingRequestID.Decode(r); err != nil {
			return utils.WrapError("Decode schedulingRequestID", err)
		}
	}
	var tmp_bool_logicalChannelSR_Mask bool
	if tmp_bool_logicalChannelSR_Mask, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean logicalChannelSR_Mask", err)
	}
	ie.logicalChannelSR_Mask = tmp_bool_logicalChannelSR_Mask
	var tmp_bool_logicalChannelSR_DelayTimerApplied bool
	if tmp_bool_logicalChannelSR_DelayTimerApplied, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean logicalChannelSR_DelayTimerApplied", err)
	}
	ie.logicalChannelSR_DelayTimerApplied = tmp_bool_logicalChannelSR_DelayTimerApplied
	if bitRateQueryProhibitTimerPresent {
		ie.bitRateQueryProhibitTimer = new(LogicalChannelConfig_ul_SpecificParameters_bitRateQueryProhibitTimer)
		if err = ie.bitRateQueryProhibitTimer.Decode(r); err != nil {
			return utils.WrapError("Decode bitRateQueryProhibitTimer", err)
		}
	}
	if allowedCG_List_r16Present {
		tmp_allowedCG_List_r16 := utils.NewSequence[*ConfiguredGrantConfigIndexMAC_r16]([]*ConfiguredGrantConfigIndexMAC_r16{}, uper.Constraint{Lb: 0, Ub: maxNrofConfiguredGrantConfigMAC_1_r16}, false)
		fn_allowedCG_List_r16 := func() *ConfiguredGrantConfigIndexMAC_r16 {
			return new(ConfiguredGrantConfigIndexMAC_r16)
		}
		if err = tmp_allowedCG_List_r16.Decode(r, fn_allowedCG_List_r16); err != nil {
			return utils.WrapError("Decode allowedCG_List_r16", err)
		}
		ie.allowedCG_List_r16 = []ConfiguredGrantConfigIndexMAC_r16{}
		for _, i := range tmp_allowedCG_List_r16.Value {
			ie.allowedCG_List_r16 = append(ie.allowedCG_List_r16, *i)
		}
	}
	if allowedPHY_PriorityIndex_r16Present {
		ie.allowedPHY_PriorityIndex_r16 = new(LogicalChannelConfig_ul_SpecificParameters_allowedPHY_PriorityIndex_r16)
		if err = ie.allowedPHY_PriorityIndex_r16.Decode(r); err != nil {
			return utils.WrapError("Decode allowedPHY_PriorityIndex_r16", err)
		}
	}
	if logicalChannelGroupIAB_Ext_r17Present {
		var tmp_int_logicalChannelGroupIAB_Ext_r17 int64
		if tmp_int_logicalChannelGroupIAB_Ext_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: maxLCG_ID_IAB_r17}, false); err != nil {
			return utils.WrapError("Decode logicalChannelGroupIAB_Ext_r17", err)
		}
		ie.logicalChannelGroupIAB_Ext_r17 = &tmp_int_logicalChannelGroupIAB_Ext_r17
	}
	if allowedHARQ_mode_r17Present {
		ie.allowedHARQ_mode_r17 = new(LogicalChannelConfig_ul_SpecificParameters_allowedHARQ_mode_r17)
		if err = ie.allowedHARQ_mode_r17.Decode(r); err != nil {
			return utils.WrapError("Decode allowedHARQ_mode_r17", err)
		}
	}
	return nil
}
