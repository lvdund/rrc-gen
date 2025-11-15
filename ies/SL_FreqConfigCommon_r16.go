package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_FreqConfigCommon_r16 struct {
	sl_SCS_SpecificCarrierList_r16 []SCS_SpecificCarrier                               `lb:1,ub:maxSCSs,madatory`
	sl_AbsoluteFrequencyPointA_r16 ARFCN_ValueNR                                       `madatory`
	sl_AbsoluteFrequencySSB_r16    *ARFCN_ValueNR                                      `optional`
	frequencyShift7p5khzSL_r16     *SL_FreqConfigCommon_r16_frequencyShift7p5khzSL_r16 `optional`
	valueN_r16                     int64                                               `lb:-1,ub:1,madatory`
	sl_BWP_List_r16                []SL_BWP_ConfigCommon_r16                           `lb:1,ub:maxNrofSL_BWPs_r16,optional`
	sl_SyncPriority_r16            *SL_FreqConfigCommon_r16_sl_SyncPriority_r16        `optional`
	sl_NbAsSync_r16                *bool                                               `optional`
	sl_SyncConfigList_r16          *SL_SyncConfigList_r16                              `optional`
}

func (ie *SL_FreqConfigCommon_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.sl_AbsoluteFrequencySSB_r16 != nil, ie.frequencyShift7p5khzSL_r16 != nil, len(ie.sl_BWP_List_r16) > 0, ie.sl_SyncPriority_r16 != nil, ie.sl_NbAsSync_r16 != nil, ie.sl_SyncConfigList_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	tmp_sl_SCS_SpecificCarrierList_r16 := utils.NewSequence[*SCS_SpecificCarrier]([]*SCS_SpecificCarrier{}, uper.Constraint{Lb: 1, Ub: maxSCSs}, false)
	for _, i := range ie.sl_SCS_SpecificCarrierList_r16 {
		tmp_sl_SCS_SpecificCarrierList_r16.Value = append(tmp_sl_SCS_SpecificCarrierList_r16.Value, &i)
	}
	if err = tmp_sl_SCS_SpecificCarrierList_r16.Encode(w); err != nil {
		return utils.WrapError("Encode sl_SCS_SpecificCarrierList_r16", err)
	}
	if err = ie.sl_AbsoluteFrequencyPointA_r16.Encode(w); err != nil {
		return utils.WrapError("Encode sl_AbsoluteFrequencyPointA_r16", err)
	}
	if ie.sl_AbsoluteFrequencySSB_r16 != nil {
		if err = ie.sl_AbsoluteFrequencySSB_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_AbsoluteFrequencySSB_r16", err)
		}
	}
	if ie.frequencyShift7p5khzSL_r16 != nil {
		if err = ie.frequencyShift7p5khzSL_r16.Encode(w); err != nil {
			return utils.WrapError("Encode frequencyShift7p5khzSL_r16", err)
		}
	}
	if err = w.WriteInteger(ie.valueN_r16, &uper.Constraint{Lb: -1, Ub: 1}, false); err != nil {
		return utils.WrapError("WriteInteger valueN_r16", err)
	}
	if len(ie.sl_BWP_List_r16) > 0 {
		tmp_sl_BWP_List_r16 := utils.NewSequence[*SL_BWP_ConfigCommon_r16]([]*SL_BWP_ConfigCommon_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofSL_BWPs_r16}, false)
		for _, i := range ie.sl_BWP_List_r16 {
			tmp_sl_BWP_List_r16.Value = append(tmp_sl_BWP_List_r16.Value, &i)
		}
		if err = tmp_sl_BWP_List_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_BWP_List_r16", err)
		}
	}
	if ie.sl_SyncPriority_r16 != nil {
		if err = ie.sl_SyncPriority_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_SyncPriority_r16", err)
		}
	}
	if ie.sl_NbAsSync_r16 != nil {
		if err = w.WriteBoolean(*ie.sl_NbAsSync_r16); err != nil {
			return utils.WrapError("Encode sl_NbAsSync_r16", err)
		}
	}
	if ie.sl_SyncConfigList_r16 != nil {
		if err = ie.sl_SyncConfigList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_SyncConfigList_r16", err)
		}
	}
	return nil
}

func (ie *SL_FreqConfigCommon_r16) Decode(r *uper.UperReader) error {
	var err error
	var sl_AbsoluteFrequencySSB_r16Present bool
	if sl_AbsoluteFrequencySSB_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var frequencyShift7p5khzSL_r16Present bool
	if frequencyShift7p5khzSL_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_BWP_List_r16Present bool
	if sl_BWP_List_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_SyncPriority_r16Present bool
	if sl_SyncPriority_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_NbAsSync_r16Present bool
	if sl_NbAsSync_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_SyncConfigList_r16Present bool
	if sl_SyncConfigList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	tmp_sl_SCS_SpecificCarrierList_r16 := utils.NewSequence[*SCS_SpecificCarrier]([]*SCS_SpecificCarrier{}, uper.Constraint{Lb: 1, Ub: maxSCSs}, false)
	fn_sl_SCS_SpecificCarrierList_r16 := func() *SCS_SpecificCarrier {
		return new(SCS_SpecificCarrier)
	}
	if err = tmp_sl_SCS_SpecificCarrierList_r16.Decode(r, fn_sl_SCS_SpecificCarrierList_r16); err != nil {
		return utils.WrapError("Decode sl_SCS_SpecificCarrierList_r16", err)
	}
	ie.sl_SCS_SpecificCarrierList_r16 = []SCS_SpecificCarrier{}
	for _, i := range tmp_sl_SCS_SpecificCarrierList_r16.Value {
		ie.sl_SCS_SpecificCarrierList_r16 = append(ie.sl_SCS_SpecificCarrierList_r16, *i)
	}
	if err = ie.sl_AbsoluteFrequencyPointA_r16.Decode(r); err != nil {
		return utils.WrapError("Decode sl_AbsoluteFrequencyPointA_r16", err)
	}
	if sl_AbsoluteFrequencySSB_r16Present {
		ie.sl_AbsoluteFrequencySSB_r16 = new(ARFCN_ValueNR)
		if err = ie.sl_AbsoluteFrequencySSB_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_AbsoluteFrequencySSB_r16", err)
		}
	}
	if frequencyShift7p5khzSL_r16Present {
		ie.frequencyShift7p5khzSL_r16 = new(SL_FreqConfigCommon_r16_frequencyShift7p5khzSL_r16)
		if err = ie.frequencyShift7p5khzSL_r16.Decode(r); err != nil {
			return utils.WrapError("Decode frequencyShift7p5khzSL_r16", err)
		}
	}
	var tmp_int_valueN_r16 int64
	if tmp_int_valueN_r16, err = r.ReadInteger(&uper.Constraint{Lb: -1, Ub: 1}, false); err != nil {
		return utils.WrapError("ReadInteger valueN_r16", err)
	}
	ie.valueN_r16 = tmp_int_valueN_r16
	if sl_BWP_List_r16Present {
		tmp_sl_BWP_List_r16 := utils.NewSequence[*SL_BWP_ConfigCommon_r16]([]*SL_BWP_ConfigCommon_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofSL_BWPs_r16}, false)
		fn_sl_BWP_List_r16 := func() *SL_BWP_ConfigCommon_r16 {
			return new(SL_BWP_ConfigCommon_r16)
		}
		if err = tmp_sl_BWP_List_r16.Decode(r, fn_sl_BWP_List_r16); err != nil {
			return utils.WrapError("Decode sl_BWP_List_r16", err)
		}
		ie.sl_BWP_List_r16 = []SL_BWP_ConfigCommon_r16{}
		for _, i := range tmp_sl_BWP_List_r16.Value {
			ie.sl_BWP_List_r16 = append(ie.sl_BWP_List_r16, *i)
		}
	}
	if sl_SyncPriority_r16Present {
		ie.sl_SyncPriority_r16 = new(SL_FreqConfigCommon_r16_sl_SyncPriority_r16)
		if err = ie.sl_SyncPriority_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_SyncPriority_r16", err)
		}
	}
	if sl_NbAsSync_r16Present {
		var tmp_bool_sl_NbAsSync_r16 bool
		if tmp_bool_sl_NbAsSync_r16, err = r.ReadBoolean(); err != nil {
			return utils.WrapError("Decode sl_NbAsSync_r16", err)
		}
		ie.sl_NbAsSync_r16 = &tmp_bool_sl_NbAsSync_r16
	}
	if sl_SyncConfigList_r16Present {
		ie.sl_SyncConfigList_r16 = new(SL_SyncConfigList_r16)
		if err = ie.sl_SyncConfigList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_SyncConfigList_r16", err)
		}
	}
	return nil
}
