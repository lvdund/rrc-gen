package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	UEPositioningAssistanceInfo_v1720_IEs_ue_TxTEG_TimingErrorMarginValue_r17_Enum_tc0  uper.Enumerated = 0
	UEPositioningAssistanceInfo_v1720_IEs_ue_TxTEG_TimingErrorMarginValue_r17_Enum_tc2  uper.Enumerated = 1
	UEPositioningAssistanceInfo_v1720_IEs_ue_TxTEG_TimingErrorMarginValue_r17_Enum_tc4  uper.Enumerated = 2
	UEPositioningAssistanceInfo_v1720_IEs_ue_TxTEG_TimingErrorMarginValue_r17_Enum_tc6  uper.Enumerated = 3
	UEPositioningAssistanceInfo_v1720_IEs_ue_TxTEG_TimingErrorMarginValue_r17_Enum_tc8  uper.Enumerated = 4
	UEPositioningAssistanceInfo_v1720_IEs_ue_TxTEG_TimingErrorMarginValue_r17_Enum_tc12 uper.Enumerated = 5
	UEPositioningAssistanceInfo_v1720_IEs_ue_TxTEG_TimingErrorMarginValue_r17_Enum_tc16 uper.Enumerated = 6
	UEPositioningAssistanceInfo_v1720_IEs_ue_TxTEG_TimingErrorMarginValue_r17_Enum_tc20 uper.Enumerated = 7
	UEPositioningAssistanceInfo_v1720_IEs_ue_TxTEG_TimingErrorMarginValue_r17_Enum_tc24 uper.Enumerated = 8
	UEPositioningAssistanceInfo_v1720_IEs_ue_TxTEG_TimingErrorMarginValue_r17_Enum_tc32 uper.Enumerated = 9
	UEPositioningAssistanceInfo_v1720_IEs_ue_TxTEG_TimingErrorMarginValue_r17_Enum_tc40 uper.Enumerated = 10
	UEPositioningAssistanceInfo_v1720_IEs_ue_TxTEG_TimingErrorMarginValue_r17_Enum_tc48 uper.Enumerated = 11
	UEPositioningAssistanceInfo_v1720_IEs_ue_TxTEG_TimingErrorMarginValue_r17_Enum_tc56 uper.Enumerated = 12
	UEPositioningAssistanceInfo_v1720_IEs_ue_TxTEG_TimingErrorMarginValue_r17_Enum_tc64 uper.Enumerated = 13
	UEPositioningAssistanceInfo_v1720_IEs_ue_TxTEG_TimingErrorMarginValue_r17_Enum_tc72 uper.Enumerated = 14
	UEPositioningAssistanceInfo_v1720_IEs_ue_TxTEG_TimingErrorMarginValue_r17_Enum_tc80 uper.Enumerated = 15
)

type UEPositioningAssistanceInfo_v1720_IEs_ue_TxTEG_TimingErrorMarginValue_r17 struct {
	Value uper.Enumerated `lb:0,ub:15,madatory`
}

func (ie *UEPositioningAssistanceInfo_v1720_IEs_ue_TxTEG_TimingErrorMarginValue_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Encode UEPositioningAssistanceInfo_v1720_IEs_ue_TxTEG_TimingErrorMarginValue_r17", err)
	}
	return nil
}

func (ie *UEPositioningAssistanceInfo_v1720_IEs_ue_TxTEG_TimingErrorMarginValue_r17) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Decode UEPositioningAssistanceInfo_v1720_IEs_ue_TxTEG_TimingErrorMarginValue_r17", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
