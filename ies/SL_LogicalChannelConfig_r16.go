package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_LogicalChannelConfig_r16 struct {
	sl_Priority_r16                           int64                                                           `lb:1,ub:8,madatory`
	sl_PrioritisedBitRate_r16                 SL_LogicalChannelConfig_r16_sl_PrioritisedBitRate_r16           `madatory`
	sl_BucketSizeDuration_r16                 SL_LogicalChannelConfig_r16_sl_BucketSizeDuration_r16           `madatory`
	sl_ConfiguredGrantType1Allowed_r16        *SL_LogicalChannelConfig_r16_sl_ConfiguredGrantType1Allowed_r16 `optional`
	sl_HARQ_FeedbackEnabled_r16               *SL_LogicalChannelConfig_r16_sl_HARQ_FeedbackEnabled_r16        `optional`
	sl_AllowedCG_List_r16                     []SL_ConfigIndexCG_r16                                          `lb:0,ub:maxNrofCG_SL_1_r16,optional`
	sl_AllowedSCS_List_r16                    []SubcarrierSpacing                                             `lb:1,ub:maxSCSs,optional`
	sl_MaxPUSCH_Duration_r16                  *SL_LogicalChannelConfig_r16_sl_MaxPUSCH_Duration_r16           `optional`
	sl_LogicalChannelGroup_r16                *int64                                                          `lb:0,ub:maxLCG_ID,optional`
	sl_SchedulingRequestId_r16                *SchedulingRequestId                                            `optional`
	sl_LogicalChannelSR_DelayTimerApplied_r16 *bool                                                           `optional`
}

func (ie *SL_LogicalChannelConfig_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.sl_ConfiguredGrantType1Allowed_r16 != nil, ie.sl_HARQ_FeedbackEnabled_r16 != nil, len(ie.sl_AllowedCG_List_r16) > 0, len(ie.sl_AllowedSCS_List_r16) > 0, ie.sl_MaxPUSCH_Duration_r16 != nil, ie.sl_LogicalChannelGroup_r16 != nil, ie.sl_SchedulingRequestId_r16 != nil, ie.sl_LogicalChannelSR_DelayTimerApplied_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = w.WriteInteger(ie.sl_Priority_r16, &uper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
		return utils.WrapError("WriteInteger sl_Priority_r16", err)
	}
	if err = ie.sl_PrioritisedBitRate_r16.Encode(w); err != nil {
		return utils.WrapError("Encode sl_PrioritisedBitRate_r16", err)
	}
	if err = ie.sl_BucketSizeDuration_r16.Encode(w); err != nil {
		return utils.WrapError("Encode sl_BucketSizeDuration_r16", err)
	}
	if ie.sl_ConfiguredGrantType1Allowed_r16 != nil {
		if err = ie.sl_ConfiguredGrantType1Allowed_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_ConfiguredGrantType1Allowed_r16", err)
		}
	}
	if ie.sl_HARQ_FeedbackEnabled_r16 != nil {
		if err = ie.sl_HARQ_FeedbackEnabled_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_HARQ_FeedbackEnabled_r16", err)
		}
	}
	if len(ie.sl_AllowedCG_List_r16) > 0 {
		tmp_sl_AllowedCG_List_r16 := utils.NewSequence[*SL_ConfigIndexCG_r16]([]*SL_ConfigIndexCG_r16{}, uper.Constraint{Lb: 0, Ub: maxNrofCG_SL_1_r16}, false)
		for _, i := range ie.sl_AllowedCG_List_r16 {
			tmp_sl_AllowedCG_List_r16.Value = append(tmp_sl_AllowedCG_List_r16.Value, &i)
		}
		if err = tmp_sl_AllowedCG_List_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_AllowedCG_List_r16", err)
		}
	}
	if len(ie.sl_AllowedSCS_List_r16) > 0 {
		tmp_sl_AllowedSCS_List_r16 := utils.NewSequence[*SubcarrierSpacing]([]*SubcarrierSpacing{}, uper.Constraint{Lb: 1, Ub: maxSCSs}, false)
		for _, i := range ie.sl_AllowedSCS_List_r16 {
			tmp_sl_AllowedSCS_List_r16.Value = append(tmp_sl_AllowedSCS_List_r16.Value, &i)
		}
		if err = tmp_sl_AllowedSCS_List_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_AllowedSCS_List_r16", err)
		}
	}
	if ie.sl_MaxPUSCH_Duration_r16 != nil {
		if err = ie.sl_MaxPUSCH_Duration_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_MaxPUSCH_Duration_r16", err)
		}
	}
	if ie.sl_LogicalChannelGroup_r16 != nil {
		if err = w.WriteInteger(*ie.sl_LogicalChannelGroup_r16, &uper.Constraint{Lb: 0, Ub: maxLCG_ID}, false); err != nil {
			return utils.WrapError("Encode sl_LogicalChannelGroup_r16", err)
		}
	}
	if ie.sl_SchedulingRequestId_r16 != nil {
		if err = ie.sl_SchedulingRequestId_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_SchedulingRequestId_r16", err)
		}
	}
	if ie.sl_LogicalChannelSR_DelayTimerApplied_r16 != nil {
		if err = w.WriteBoolean(*ie.sl_LogicalChannelSR_DelayTimerApplied_r16); err != nil {
			return utils.WrapError("Encode sl_LogicalChannelSR_DelayTimerApplied_r16", err)
		}
	}
	return nil
}

func (ie *SL_LogicalChannelConfig_r16) Decode(r *uper.UperReader) error {
	var err error
	var sl_ConfiguredGrantType1Allowed_r16Present bool
	if sl_ConfiguredGrantType1Allowed_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_HARQ_FeedbackEnabled_r16Present bool
	if sl_HARQ_FeedbackEnabled_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_AllowedCG_List_r16Present bool
	if sl_AllowedCG_List_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_AllowedSCS_List_r16Present bool
	if sl_AllowedSCS_List_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_MaxPUSCH_Duration_r16Present bool
	if sl_MaxPUSCH_Duration_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_LogicalChannelGroup_r16Present bool
	if sl_LogicalChannelGroup_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_SchedulingRequestId_r16Present bool
	if sl_SchedulingRequestId_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_LogicalChannelSR_DelayTimerApplied_r16Present bool
	if sl_LogicalChannelSR_DelayTimerApplied_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var tmp_int_sl_Priority_r16 int64
	if tmp_int_sl_Priority_r16, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
		return utils.WrapError("ReadInteger sl_Priority_r16", err)
	}
	ie.sl_Priority_r16 = tmp_int_sl_Priority_r16
	if err = ie.sl_PrioritisedBitRate_r16.Decode(r); err != nil {
		return utils.WrapError("Decode sl_PrioritisedBitRate_r16", err)
	}
	if err = ie.sl_BucketSizeDuration_r16.Decode(r); err != nil {
		return utils.WrapError("Decode sl_BucketSizeDuration_r16", err)
	}
	if sl_ConfiguredGrantType1Allowed_r16Present {
		ie.sl_ConfiguredGrantType1Allowed_r16 = new(SL_LogicalChannelConfig_r16_sl_ConfiguredGrantType1Allowed_r16)
		if err = ie.sl_ConfiguredGrantType1Allowed_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_ConfiguredGrantType1Allowed_r16", err)
		}
	}
	if sl_HARQ_FeedbackEnabled_r16Present {
		ie.sl_HARQ_FeedbackEnabled_r16 = new(SL_LogicalChannelConfig_r16_sl_HARQ_FeedbackEnabled_r16)
		if err = ie.sl_HARQ_FeedbackEnabled_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_HARQ_FeedbackEnabled_r16", err)
		}
	}
	if sl_AllowedCG_List_r16Present {
		tmp_sl_AllowedCG_List_r16 := utils.NewSequence[*SL_ConfigIndexCG_r16]([]*SL_ConfigIndexCG_r16{}, uper.Constraint{Lb: 0, Ub: maxNrofCG_SL_1_r16}, false)
		fn_sl_AllowedCG_List_r16 := func() *SL_ConfigIndexCG_r16 {
			return new(SL_ConfigIndexCG_r16)
		}
		if err = tmp_sl_AllowedCG_List_r16.Decode(r, fn_sl_AllowedCG_List_r16); err != nil {
			return utils.WrapError("Decode sl_AllowedCG_List_r16", err)
		}
		ie.sl_AllowedCG_List_r16 = []SL_ConfigIndexCG_r16{}
		for _, i := range tmp_sl_AllowedCG_List_r16.Value {
			ie.sl_AllowedCG_List_r16 = append(ie.sl_AllowedCG_List_r16, *i)
		}
	}
	if sl_AllowedSCS_List_r16Present {
		tmp_sl_AllowedSCS_List_r16 := utils.NewSequence[*SubcarrierSpacing]([]*SubcarrierSpacing{}, uper.Constraint{Lb: 1, Ub: maxSCSs}, false)
		fn_sl_AllowedSCS_List_r16 := func() *SubcarrierSpacing {
			return new(SubcarrierSpacing)
		}
		if err = tmp_sl_AllowedSCS_List_r16.Decode(r, fn_sl_AllowedSCS_List_r16); err != nil {
			return utils.WrapError("Decode sl_AllowedSCS_List_r16", err)
		}
		ie.sl_AllowedSCS_List_r16 = []SubcarrierSpacing{}
		for _, i := range tmp_sl_AllowedSCS_List_r16.Value {
			ie.sl_AllowedSCS_List_r16 = append(ie.sl_AllowedSCS_List_r16, *i)
		}
	}
	if sl_MaxPUSCH_Duration_r16Present {
		ie.sl_MaxPUSCH_Duration_r16 = new(SL_LogicalChannelConfig_r16_sl_MaxPUSCH_Duration_r16)
		if err = ie.sl_MaxPUSCH_Duration_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_MaxPUSCH_Duration_r16", err)
		}
	}
	if sl_LogicalChannelGroup_r16Present {
		var tmp_int_sl_LogicalChannelGroup_r16 int64
		if tmp_int_sl_LogicalChannelGroup_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: maxLCG_ID}, false); err != nil {
			return utils.WrapError("Decode sl_LogicalChannelGroup_r16", err)
		}
		ie.sl_LogicalChannelGroup_r16 = &tmp_int_sl_LogicalChannelGroup_r16
	}
	if sl_SchedulingRequestId_r16Present {
		ie.sl_SchedulingRequestId_r16 = new(SchedulingRequestId)
		if err = ie.sl_SchedulingRequestId_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_SchedulingRequestId_r16", err)
		}
	}
	if sl_LogicalChannelSR_DelayTimerApplied_r16Present {
		var tmp_bool_sl_LogicalChannelSR_DelayTimerApplied_r16 bool
		if tmp_bool_sl_LogicalChannelSR_DelayTimerApplied_r16, err = r.ReadBoolean(); err != nil {
			return utils.WrapError("Decode sl_LogicalChannelSR_DelayTimerApplied_r16", err)
		}
		ie.sl_LogicalChannelSR_DelayTimerApplied_r16 = &tmp_bool_sl_LogicalChannelSR_DelayTimerApplied_r16
	}
	return nil
}
