package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_InterUE_CoordinationScheme1_r17 struct {
	sl_IUC_Explicit_r17                             *SL_InterUE_CoordinationScheme1_r17_sl_IUC_Explicit_r17          `optional`
	sl_IUC_Condition_r17                            *SL_InterUE_CoordinationScheme1_r17_sl_IUC_Condition_r17         `optional`
	sl_Condition1_A_2_r17                           *SL_InterUE_CoordinationScheme1_r17_sl_Condition1_A_2_r17        `optional`
	sl_ThresholdRSRP_Condition1_B_1_Option1List_r17 []SL_ThresholdRSRP_Condition1_B_1_r17                            `lb:1,ub:8,optional`
	sl_ThresholdRSRP_Condition1_B_1_Option2List_r17 []SL_ThresholdRSRP_Condition1_B_1_r17                            `lb:1,ub:8,optional`
	sl_ContainerCoordInfo_r17                       *SL_InterUE_CoordinationScheme1_r17_sl_ContainerCoordInfo_r17    `optional`
	sl_ContainerRequest_r17                         *SL_InterUE_CoordinationScheme1_r17_sl_ContainerRequest_r17      `optional`
	sl_TriggerConditionCoordInfo_r17                *int64                                                           `lb:0,ub:1,optional`
	sl_TriggerConditionRequest_r17                  *int64                                                           `lb:0,ub:1,optional`
	sl_PriorityCoordInfoExplicit_r17                *int64                                                           `lb:1,ub:8,optional`
	sl_PriorityCoordInfoCondition_r17               *int64                                                           `lb:1,ub:8,optional`
	sl_PriorityRequest_r17                          *int64                                                           `lb:1,ub:8,optional`
	sl_PriorityPreferredResourceSet_r17             *int64                                                           `lb:1,ub:8,optional`
	sl_MaxSlotOffsetTRIV_r17                        *int64                                                           `lb:1,ub:8000,optional`
	sl_NumSubCH_PreferredResourceSet_r17            *int64                                                           `lb:1,ub:27,optional`
	sl_ReservedPeriodPreferredResourceSet_r17       *int64                                                           `lb:1,ub:16,optional`
	sl_DetermineResourceType_r17                    *SL_InterUE_CoordinationScheme1_r17_sl_DetermineResourceType_r17 `optional`
}

func (ie *SL_InterUE_CoordinationScheme1_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.sl_IUC_Explicit_r17 != nil, ie.sl_IUC_Condition_r17 != nil, ie.sl_Condition1_A_2_r17 != nil, len(ie.sl_ThresholdRSRP_Condition1_B_1_Option1List_r17) > 0, len(ie.sl_ThresholdRSRP_Condition1_B_1_Option2List_r17) > 0, ie.sl_ContainerCoordInfo_r17 != nil, ie.sl_ContainerRequest_r17 != nil, ie.sl_TriggerConditionCoordInfo_r17 != nil, ie.sl_TriggerConditionRequest_r17 != nil, ie.sl_PriorityCoordInfoExplicit_r17 != nil, ie.sl_PriorityCoordInfoCondition_r17 != nil, ie.sl_PriorityRequest_r17 != nil, ie.sl_PriorityPreferredResourceSet_r17 != nil, ie.sl_MaxSlotOffsetTRIV_r17 != nil, ie.sl_NumSubCH_PreferredResourceSet_r17 != nil, ie.sl_ReservedPeriodPreferredResourceSet_r17 != nil, ie.sl_DetermineResourceType_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.sl_IUC_Explicit_r17 != nil {
		if err = ie.sl_IUC_Explicit_r17.Encode(w); err != nil {
			return utils.WrapError("Encode sl_IUC_Explicit_r17", err)
		}
	}
	if ie.sl_IUC_Condition_r17 != nil {
		if err = ie.sl_IUC_Condition_r17.Encode(w); err != nil {
			return utils.WrapError("Encode sl_IUC_Condition_r17", err)
		}
	}
	if ie.sl_Condition1_A_2_r17 != nil {
		if err = ie.sl_Condition1_A_2_r17.Encode(w); err != nil {
			return utils.WrapError("Encode sl_Condition1_A_2_r17", err)
		}
	}
	if len(ie.sl_ThresholdRSRP_Condition1_B_1_Option1List_r17) > 0 {
		tmp_sl_ThresholdRSRP_Condition1_B_1_Option1List_r17 := utils.NewSequence[*SL_ThresholdRSRP_Condition1_B_1_r17]([]*SL_ThresholdRSRP_Condition1_B_1_r17{}, uper.Constraint{Lb: 1, Ub: 8}, false)
		for _, i := range ie.sl_ThresholdRSRP_Condition1_B_1_Option1List_r17 {
			tmp_sl_ThresholdRSRP_Condition1_B_1_Option1List_r17.Value = append(tmp_sl_ThresholdRSRP_Condition1_B_1_Option1List_r17.Value, &i)
		}
		if err = tmp_sl_ThresholdRSRP_Condition1_B_1_Option1List_r17.Encode(w); err != nil {
			return utils.WrapError("Encode sl_ThresholdRSRP_Condition1_B_1_Option1List_r17", err)
		}
	}
	if len(ie.sl_ThresholdRSRP_Condition1_B_1_Option2List_r17) > 0 {
		tmp_sl_ThresholdRSRP_Condition1_B_1_Option2List_r17 := utils.NewSequence[*SL_ThresholdRSRP_Condition1_B_1_r17]([]*SL_ThresholdRSRP_Condition1_B_1_r17{}, uper.Constraint{Lb: 1, Ub: 8}, false)
		for _, i := range ie.sl_ThresholdRSRP_Condition1_B_1_Option2List_r17 {
			tmp_sl_ThresholdRSRP_Condition1_B_1_Option2List_r17.Value = append(tmp_sl_ThresholdRSRP_Condition1_B_1_Option2List_r17.Value, &i)
		}
		if err = tmp_sl_ThresholdRSRP_Condition1_B_1_Option2List_r17.Encode(w); err != nil {
			return utils.WrapError("Encode sl_ThresholdRSRP_Condition1_B_1_Option2List_r17", err)
		}
	}
	if ie.sl_ContainerCoordInfo_r17 != nil {
		if err = ie.sl_ContainerCoordInfo_r17.Encode(w); err != nil {
			return utils.WrapError("Encode sl_ContainerCoordInfo_r17", err)
		}
	}
	if ie.sl_ContainerRequest_r17 != nil {
		if err = ie.sl_ContainerRequest_r17.Encode(w); err != nil {
			return utils.WrapError("Encode sl_ContainerRequest_r17", err)
		}
	}
	if ie.sl_TriggerConditionCoordInfo_r17 != nil {
		if err = w.WriteInteger(*ie.sl_TriggerConditionCoordInfo_r17, &uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
			return utils.WrapError("Encode sl_TriggerConditionCoordInfo_r17", err)
		}
	}
	if ie.sl_TriggerConditionRequest_r17 != nil {
		if err = w.WriteInteger(*ie.sl_TriggerConditionRequest_r17, &uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
			return utils.WrapError("Encode sl_TriggerConditionRequest_r17", err)
		}
	}
	if ie.sl_PriorityCoordInfoExplicit_r17 != nil {
		if err = w.WriteInteger(*ie.sl_PriorityCoordInfoExplicit_r17, &uper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
			return utils.WrapError("Encode sl_PriorityCoordInfoExplicit_r17", err)
		}
	}
	if ie.sl_PriorityCoordInfoCondition_r17 != nil {
		if err = w.WriteInteger(*ie.sl_PriorityCoordInfoCondition_r17, &uper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
			return utils.WrapError("Encode sl_PriorityCoordInfoCondition_r17", err)
		}
	}
	if ie.sl_PriorityRequest_r17 != nil {
		if err = w.WriteInteger(*ie.sl_PriorityRequest_r17, &uper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
			return utils.WrapError("Encode sl_PriorityRequest_r17", err)
		}
	}
	if ie.sl_PriorityPreferredResourceSet_r17 != nil {
		if err = w.WriteInteger(*ie.sl_PriorityPreferredResourceSet_r17, &uper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
			return utils.WrapError("Encode sl_PriorityPreferredResourceSet_r17", err)
		}
	}
	if ie.sl_MaxSlotOffsetTRIV_r17 != nil {
		if err = w.WriteInteger(*ie.sl_MaxSlotOffsetTRIV_r17, &uper.Constraint{Lb: 1, Ub: 8000}, false); err != nil {
			return utils.WrapError("Encode sl_MaxSlotOffsetTRIV_r17", err)
		}
	}
	if ie.sl_NumSubCH_PreferredResourceSet_r17 != nil {
		if err = w.WriteInteger(*ie.sl_NumSubCH_PreferredResourceSet_r17, &uper.Constraint{Lb: 1, Ub: 27}, false); err != nil {
			return utils.WrapError("Encode sl_NumSubCH_PreferredResourceSet_r17", err)
		}
	}
	if ie.sl_ReservedPeriodPreferredResourceSet_r17 != nil {
		if err = w.WriteInteger(*ie.sl_ReservedPeriodPreferredResourceSet_r17, &uper.Constraint{Lb: 1, Ub: 16}, false); err != nil {
			return utils.WrapError("Encode sl_ReservedPeriodPreferredResourceSet_r17", err)
		}
	}
	if ie.sl_DetermineResourceType_r17 != nil {
		if err = ie.sl_DetermineResourceType_r17.Encode(w); err != nil {
			return utils.WrapError("Encode sl_DetermineResourceType_r17", err)
		}
	}
	return nil
}

func (ie *SL_InterUE_CoordinationScheme1_r17) Decode(r *uper.UperReader) error {
	var err error
	var sl_IUC_Explicit_r17Present bool
	if sl_IUC_Explicit_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_IUC_Condition_r17Present bool
	if sl_IUC_Condition_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_Condition1_A_2_r17Present bool
	if sl_Condition1_A_2_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_ThresholdRSRP_Condition1_B_1_Option1List_r17Present bool
	if sl_ThresholdRSRP_Condition1_B_1_Option1List_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_ThresholdRSRP_Condition1_B_1_Option2List_r17Present bool
	if sl_ThresholdRSRP_Condition1_B_1_Option2List_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_ContainerCoordInfo_r17Present bool
	if sl_ContainerCoordInfo_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_ContainerRequest_r17Present bool
	if sl_ContainerRequest_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_TriggerConditionCoordInfo_r17Present bool
	if sl_TriggerConditionCoordInfo_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_TriggerConditionRequest_r17Present bool
	if sl_TriggerConditionRequest_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_PriorityCoordInfoExplicit_r17Present bool
	if sl_PriorityCoordInfoExplicit_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_PriorityCoordInfoCondition_r17Present bool
	if sl_PriorityCoordInfoCondition_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_PriorityRequest_r17Present bool
	if sl_PriorityRequest_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_PriorityPreferredResourceSet_r17Present bool
	if sl_PriorityPreferredResourceSet_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_MaxSlotOffsetTRIV_r17Present bool
	if sl_MaxSlotOffsetTRIV_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_NumSubCH_PreferredResourceSet_r17Present bool
	if sl_NumSubCH_PreferredResourceSet_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_ReservedPeriodPreferredResourceSet_r17Present bool
	if sl_ReservedPeriodPreferredResourceSet_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_DetermineResourceType_r17Present bool
	if sl_DetermineResourceType_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if sl_IUC_Explicit_r17Present {
		ie.sl_IUC_Explicit_r17 = new(SL_InterUE_CoordinationScheme1_r17_sl_IUC_Explicit_r17)
		if err = ie.sl_IUC_Explicit_r17.Decode(r); err != nil {
			return utils.WrapError("Decode sl_IUC_Explicit_r17", err)
		}
	}
	if sl_IUC_Condition_r17Present {
		ie.sl_IUC_Condition_r17 = new(SL_InterUE_CoordinationScheme1_r17_sl_IUC_Condition_r17)
		if err = ie.sl_IUC_Condition_r17.Decode(r); err != nil {
			return utils.WrapError("Decode sl_IUC_Condition_r17", err)
		}
	}
	if sl_Condition1_A_2_r17Present {
		ie.sl_Condition1_A_2_r17 = new(SL_InterUE_CoordinationScheme1_r17_sl_Condition1_A_2_r17)
		if err = ie.sl_Condition1_A_2_r17.Decode(r); err != nil {
			return utils.WrapError("Decode sl_Condition1_A_2_r17", err)
		}
	}
	if sl_ThresholdRSRP_Condition1_B_1_Option1List_r17Present {
		tmp_sl_ThresholdRSRP_Condition1_B_1_Option1List_r17 := utils.NewSequence[*SL_ThresholdRSRP_Condition1_B_1_r17]([]*SL_ThresholdRSRP_Condition1_B_1_r17{}, uper.Constraint{Lb: 1, Ub: 8}, false)
		fn_sl_ThresholdRSRP_Condition1_B_1_Option1List_r17 := func() *SL_ThresholdRSRP_Condition1_B_1_r17 {
			return new(SL_ThresholdRSRP_Condition1_B_1_r17)
		}
		if err = tmp_sl_ThresholdRSRP_Condition1_B_1_Option1List_r17.Decode(r, fn_sl_ThresholdRSRP_Condition1_B_1_Option1List_r17); err != nil {
			return utils.WrapError("Decode sl_ThresholdRSRP_Condition1_B_1_Option1List_r17", err)
		}
		ie.sl_ThresholdRSRP_Condition1_B_1_Option1List_r17 = []SL_ThresholdRSRP_Condition1_B_1_r17{}
		for _, i := range tmp_sl_ThresholdRSRP_Condition1_B_1_Option1List_r17.Value {
			ie.sl_ThresholdRSRP_Condition1_B_1_Option1List_r17 = append(ie.sl_ThresholdRSRP_Condition1_B_1_Option1List_r17, *i)
		}
	}
	if sl_ThresholdRSRP_Condition1_B_1_Option2List_r17Present {
		tmp_sl_ThresholdRSRP_Condition1_B_1_Option2List_r17 := utils.NewSequence[*SL_ThresholdRSRP_Condition1_B_1_r17]([]*SL_ThresholdRSRP_Condition1_B_1_r17{}, uper.Constraint{Lb: 1, Ub: 8}, false)
		fn_sl_ThresholdRSRP_Condition1_B_1_Option2List_r17 := func() *SL_ThresholdRSRP_Condition1_B_1_r17 {
			return new(SL_ThresholdRSRP_Condition1_B_1_r17)
		}
		if err = tmp_sl_ThresholdRSRP_Condition1_B_1_Option2List_r17.Decode(r, fn_sl_ThresholdRSRP_Condition1_B_1_Option2List_r17); err != nil {
			return utils.WrapError("Decode sl_ThresholdRSRP_Condition1_B_1_Option2List_r17", err)
		}
		ie.sl_ThresholdRSRP_Condition1_B_1_Option2List_r17 = []SL_ThresholdRSRP_Condition1_B_1_r17{}
		for _, i := range tmp_sl_ThresholdRSRP_Condition1_B_1_Option2List_r17.Value {
			ie.sl_ThresholdRSRP_Condition1_B_1_Option2List_r17 = append(ie.sl_ThresholdRSRP_Condition1_B_1_Option2List_r17, *i)
		}
	}
	if sl_ContainerCoordInfo_r17Present {
		ie.sl_ContainerCoordInfo_r17 = new(SL_InterUE_CoordinationScheme1_r17_sl_ContainerCoordInfo_r17)
		if err = ie.sl_ContainerCoordInfo_r17.Decode(r); err != nil {
			return utils.WrapError("Decode sl_ContainerCoordInfo_r17", err)
		}
	}
	if sl_ContainerRequest_r17Present {
		ie.sl_ContainerRequest_r17 = new(SL_InterUE_CoordinationScheme1_r17_sl_ContainerRequest_r17)
		if err = ie.sl_ContainerRequest_r17.Decode(r); err != nil {
			return utils.WrapError("Decode sl_ContainerRequest_r17", err)
		}
	}
	if sl_TriggerConditionCoordInfo_r17Present {
		var tmp_int_sl_TriggerConditionCoordInfo_r17 int64
		if tmp_int_sl_TriggerConditionCoordInfo_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
			return utils.WrapError("Decode sl_TriggerConditionCoordInfo_r17", err)
		}
		ie.sl_TriggerConditionCoordInfo_r17 = &tmp_int_sl_TriggerConditionCoordInfo_r17
	}
	if sl_TriggerConditionRequest_r17Present {
		var tmp_int_sl_TriggerConditionRequest_r17 int64
		if tmp_int_sl_TriggerConditionRequest_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
			return utils.WrapError("Decode sl_TriggerConditionRequest_r17", err)
		}
		ie.sl_TriggerConditionRequest_r17 = &tmp_int_sl_TriggerConditionRequest_r17
	}
	if sl_PriorityCoordInfoExplicit_r17Present {
		var tmp_int_sl_PriorityCoordInfoExplicit_r17 int64
		if tmp_int_sl_PriorityCoordInfoExplicit_r17, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
			return utils.WrapError("Decode sl_PriorityCoordInfoExplicit_r17", err)
		}
		ie.sl_PriorityCoordInfoExplicit_r17 = &tmp_int_sl_PriorityCoordInfoExplicit_r17
	}
	if sl_PriorityCoordInfoCondition_r17Present {
		var tmp_int_sl_PriorityCoordInfoCondition_r17 int64
		if tmp_int_sl_PriorityCoordInfoCondition_r17, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
			return utils.WrapError("Decode sl_PriorityCoordInfoCondition_r17", err)
		}
		ie.sl_PriorityCoordInfoCondition_r17 = &tmp_int_sl_PriorityCoordInfoCondition_r17
	}
	if sl_PriorityRequest_r17Present {
		var tmp_int_sl_PriorityRequest_r17 int64
		if tmp_int_sl_PriorityRequest_r17, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
			return utils.WrapError("Decode sl_PriorityRequest_r17", err)
		}
		ie.sl_PriorityRequest_r17 = &tmp_int_sl_PriorityRequest_r17
	}
	if sl_PriorityPreferredResourceSet_r17Present {
		var tmp_int_sl_PriorityPreferredResourceSet_r17 int64
		if tmp_int_sl_PriorityPreferredResourceSet_r17, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
			return utils.WrapError("Decode sl_PriorityPreferredResourceSet_r17", err)
		}
		ie.sl_PriorityPreferredResourceSet_r17 = &tmp_int_sl_PriorityPreferredResourceSet_r17
	}
	if sl_MaxSlotOffsetTRIV_r17Present {
		var tmp_int_sl_MaxSlotOffsetTRIV_r17 int64
		if tmp_int_sl_MaxSlotOffsetTRIV_r17, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 8000}, false); err != nil {
			return utils.WrapError("Decode sl_MaxSlotOffsetTRIV_r17", err)
		}
		ie.sl_MaxSlotOffsetTRIV_r17 = &tmp_int_sl_MaxSlotOffsetTRIV_r17
	}
	if sl_NumSubCH_PreferredResourceSet_r17Present {
		var tmp_int_sl_NumSubCH_PreferredResourceSet_r17 int64
		if tmp_int_sl_NumSubCH_PreferredResourceSet_r17, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 27}, false); err != nil {
			return utils.WrapError("Decode sl_NumSubCH_PreferredResourceSet_r17", err)
		}
		ie.sl_NumSubCH_PreferredResourceSet_r17 = &tmp_int_sl_NumSubCH_PreferredResourceSet_r17
	}
	if sl_ReservedPeriodPreferredResourceSet_r17Present {
		var tmp_int_sl_ReservedPeriodPreferredResourceSet_r17 int64
		if tmp_int_sl_ReservedPeriodPreferredResourceSet_r17, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 16}, false); err != nil {
			return utils.WrapError("Decode sl_ReservedPeriodPreferredResourceSet_r17", err)
		}
		ie.sl_ReservedPeriodPreferredResourceSet_r17 = &tmp_int_sl_ReservedPeriodPreferredResourceSet_r17
	}
	if sl_DetermineResourceType_r17Present {
		ie.sl_DetermineResourceType_r17 = new(SL_InterUE_CoordinationScheme1_r17_sl_DetermineResourceType_r17)
		if err = ie.sl_DetermineResourceType_r17.Decode(r); err != nil {
			return utils.WrapError("Decode sl_DetermineResourceType_r17", err)
		}
	}
	return nil
}
