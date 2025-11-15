package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_TxResourceReq_r16 struct {
	sl_DestinationIdentity_r16           SL_DestinationIdentity_r16           `madatory`
	sl_CastType_r16                      SL_TxResourceReq_r16_sl_CastType_r16 `madatory`
	sl_RLC_ModeIndicationList_r16        []SL_RLC_ModeIndication_r16          `lb:1,ub:maxNrofSLRB_r16,optional`
	sl_QoS_InfoList_r16                  []SL_QoS_Info_r16                    `lb:1,ub:maxNrofSL_QFIsPerDest_r16,optional`
	sl_TypeTxSyncList_r16                []SL_TypeTxSync_r16                  `lb:1,ub:maxNrofFreqSL_r16,optional`
	sl_TxInterestedFreqList_r16          *SL_TxInterestedFreqList_r16         `optional`
	sl_CapabilityInformationSidelink_r16 *[]byte                              `optional`
}

func (ie *SL_TxResourceReq_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{len(ie.sl_RLC_ModeIndicationList_r16) > 0, len(ie.sl_QoS_InfoList_r16) > 0, len(ie.sl_TypeTxSyncList_r16) > 0, ie.sl_TxInterestedFreqList_r16 != nil, ie.sl_CapabilityInformationSidelink_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.sl_DestinationIdentity_r16.Encode(w); err != nil {
		return utils.WrapError("Encode sl_DestinationIdentity_r16", err)
	}
	if err = ie.sl_CastType_r16.Encode(w); err != nil {
		return utils.WrapError("Encode sl_CastType_r16", err)
	}
	if len(ie.sl_RLC_ModeIndicationList_r16) > 0 {
		tmp_sl_RLC_ModeIndicationList_r16 := utils.NewSequence[*SL_RLC_ModeIndication_r16]([]*SL_RLC_ModeIndication_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofSLRB_r16}, false)
		for _, i := range ie.sl_RLC_ModeIndicationList_r16 {
			tmp_sl_RLC_ModeIndicationList_r16.Value = append(tmp_sl_RLC_ModeIndicationList_r16.Value, &i)
		}
		if err = tmp_sl_RLC_ModeIndicationList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_RLC_ModeIndicationList_r16", err)
		}
	}
	if len(ie.sl_QoS_InfoList_r16) > 0 {
		tmp_sl_QoS_InfoList_r16 := utils.NewSequence[*SL_QoS_Info_r16]([]*SL_QoS_Info_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofSL_QFIsPerDest_r16}, false)
		for _, i := range ie.sl_QoS_InfoList_r16 {
			tmp_sl_QoS_InfoList_r16.Value = append(tmp_sl_QoS_InfoList_r16.Value, &i)
		}
		if err = tmp_sl_QoS_InfoList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_QoS_InfoList_r16", err)
		}
	}
	if len(ie.sl_TypeTxSyncList_r16) > 0 {
		tmp_sl_TypeTxSyncList_r16 := utils.NewSequence[*SL_TypeTxSync_r16]([]*SL_TypeTxSync_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofFreqSL_r16}, false)
		for _, i := range ie.sl_TypeTxSyncList_r16 {
			tmp_sl_TypeTxSyncList_r16.Value = append(tmp_sl_TypeTxSyncList_r16.Value, &i)
		}
		if err = tmp_sl_TypeTxSyncList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_TypeTxSyncList_r16", err)
		}
	}
	if ie.sl_TxInterestedFreqList_r16 != nil {
		if err = ie.sl_TxInterestedFreqList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_TxInterestedFreqList_r16", err)
		}
	}
	if ie.sl_CapabilityInformationSidelink_r16 != nil {
		if err = w.WriteOctetString(*ie.sl_CapabilityInformationSidelink_r16, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode sl_CapabilityInformationSidelink_r16", err)
		}
	}
	return nil
}

func (ie *SL_TxResourceReq_r16) Decode(r *uper.UperReader) error {
	var err error
	var sl_RLC_ModeIndicationList_r16Present bool
	if sl_RLC_ModeIndicationList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_QoS_InfoList_r16Present bool
	if sl_QoS_InfoList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_TypeTxSyncList_r16Present bool
	if sl_TypeTxSyncList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_TxInterestedFreqList_r16Present bool
	if sl_TxInterestedFreqList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_CapabilityInformationSidelink_r16Present bool
	if sl_CapabilityInformationSidelink_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.sl_DestinationIdentity_r16.Decode(r); err != nil {
		return utils.WrapError("Decode sl_DestinationIdentity_r16", err)
	}
	if err = ie.sl_CastType_r16.Decode(r); err != nil {
		return utils.WrapError("Decode sl_CastType_r16", err)
	}
	if sl_RLC_ModeIndicationList_r16Present {
		tmp_sl_RLC_ModeIndicationList_r16 := utils.NewSequence[*SL_RLC_ModeIndication_r16]([]*SL_RLC_ModeIndication_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofSLRB_r16}, false)
		fn_sl_RLC_ModeIndicationList_r16 := func() *SL_RLC_ModeIndication_r16 {
			return new(SL_RLC_ModeIndication_r16)
		}
		if err = tmp_sl_RLC_ModeIndicationList_r16.Decode(r, fn_sl_RLC_ModeIndicationList_r16); err != nil {
			return utils.WrapError("Decode sl_RLC_ModeIndicationList_r16", err)
		}
		ie.sl_RLC_ModeIndicationList_r16 = []SL_RLC_ModeIndication_r16{}
		for _, i := range tmp_sl_RLC_ModeIndicationList_r16.Value {
			ie.sl_RLC_ModeIndicationList_r16 = append(ie.sl_RLC_ModeIndicationList_r16, *i)
		}
	}
	if sl_QoS_InfoList_r16Present {
		tmp_sl_QoS_InfoList_r16 := utils.NewSequence[*SL_QoS_Info_r16]([]*SL_QoS_Info_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofSL_QFIsPerDest_r16}, false)
		fn_sl_QoS_InfoList_r16 := func() *SL_QoS_Info_r16 {
			return new(SL_QoS_Info_r16)
		}
		if err = tmp_sl_QoS_InfoList_r16.Decode(r, fn_sl_QoS_InfoList_r16); err != nil {
			return utils.WrapError("Decode sl_QoS_InfoList_r16", err)
		}
		ie.sl_QoS_InfoList_r16 = []SL_QoS_Info_r16{}
		for _, i := range tmp_sl_QoS_InfoList_r16.Value {
			ie.sl_QoS_InfoList_r16 = append(ie.sl_QoS_InfoList_r16, *i)
		}
	}
	if sl_TypeTxSyncList_r16Present {
		tmp_sl_TypeTxSyncList_r16 := utils.NewSequence[*SL_TypeTxSync_r16]([]*SL_TypeTxSync_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofFreqSL_r16}, false)
		fn_sl_TypeTxSyncList_r16 := func() *SL_TypeTxSync_r16 {
			return new(SL_TypeTxSync_r16)
		}
		if err = tmp_sl_TypeTxSyncList_r16.Decode(r, fn_sl_TypeTxSyncList_r16); err != nil {
			return utils.WrapError("Decode sl_TypeTxSyncList_r16", err)
		}
		ie.sl_TypeTxSyncList_r16 = []SL_TypeTxSync_r16{}
		for _, i := range tmp_sl_TypeTxSyncList_r16.Value {
			ie.sl_TypeTxSyncList_r16 = append(ie.sl_TypeTxSyncList_r16, *i)
		}
	}
	if sl_TxInterestedFreqList_r16Present {
		ie.sl_TxInterestedFreqList_r16 = new(SL_TxInterestedFreqList_r16)
		if err = ie.sl_TxInterestedFreqList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_TxInterestedFreqList_r16", err)
		}
	}
	if sl_CapabilityInformationSidelink_r16Present {
		var tmp_os_sl_CapabilityInformationSidelink_r16 []byte
		if tmp_os_sl_CapabilityInformationSidelink_r16, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode sl_CapabilityInformationSidelink_r16", err)
		}
		ie.sl_CapabilityInformationSidelink_r16 = &tmp_os_sl_CapabilityInformationSidelink_r16
	}
	return nil
}
