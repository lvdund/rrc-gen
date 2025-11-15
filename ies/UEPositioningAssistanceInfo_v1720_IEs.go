package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UEPositioningAssistanceInfo_v1720_IEs struct {
	ue_TxTEG_TimingErrorMarginValue_r17 *UEPositioningAssistanceInfo_v1720_IEs_ue_TxTEG_TimingErrorMarginValue_r17 `optional`
	nonCriticalExtension                interface{}                                                                `optional`
}

func (ie *UEPositioningAssistanceInfo_v1720_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.ue_TxTEG_TimingErrorMarginValue_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.ue_TxTEG_TimingErrorMarginValue_r17 != nil {
		if err = ie.ue_TxTEG_TimingErrorMarginValue_r17.Encode(w); err != nil {
			return utils.WrapError("Encode ue_TxTEG_TimingErrorMarginValue_r17", err)
		}
	}
	return nil
}

func (ie *UEPositioningAssistanceInfo_v1720_IEs) Decode(r *uper.UperReader) error {
	var err error
	var ue_TxTEG_TimingErrorMarginValue_r17Present bool
	if ue_TxTEG_TimingErrorMarginValue_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if ue_TxTEG_TimingErrorMarginValue_r17Present {
		ie.ue_TxTEG_TimingErrorMarginValue_r17 = new(UEPositioningAssistanceInfo_v1720_IEs_ue_TxTEG_TimingErrorMarginValue_r17)
		if err = ie.ue_TxTEG_TimingErrorMarginValue_r17.Decode(r); err != nil {
			return utils.WrapError("Decode ue_TxTEG_TimingErrorMarginValue_r17", err)
		}
	}
	return nil
}
