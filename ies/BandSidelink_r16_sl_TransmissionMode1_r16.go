package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BandSidelink_r16_sl_TransmissionMode1_r16 struct {
	harq_TxProcessModeOneSidelink_r16   BandSidelink_r16_sl_TransmissionMode1_r16_harq_TxProcessModeOneSidelink_r16   `madatory`
	scs_CP_PatternTxSidelinkModeOne_r16 BandSidelink_r16_sl_TransmissionMode1_r16_scs_CP_PatternTxSidelinkModeOne_r16 `madatory`
	extendedCP_TxSidelink_r16           *BandSidelink_r16_sl_TransmissionMode1_r16_extendedCP_TxSidelink_r16          `optional`
	harq_ReportOnPUCCH_r16              *BandSidelink_r16_sl_TransmissionMode1_r16_harq_ReportOnPUCCH_r16             `optional`
}

func (ie *BandSidelink_r16_sl_TransmissionMode1_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.extendedCP_TxSidelink_r16 != nil, ie.harq_ReportOnPUCCH_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.harq_TxProcessModeOneSidelink_r16.Encode(w); err != nil {
		return utils.WrapError("Encode harq_TxProcessModeOneSidelink_r16", err)
	}
	if err = ie.scs_CP_PatternTxSidelinkModeOne_r16.Encode(w); err != nil {
		return utils.WrapError("Encode scs_CP_PatternTxSidelinkModeOne_r16", err)
	}
	if ie.extendedCP_TxSidelink_r16 != nil {
		if err = ie.extendedCP_TxSidelink_r16.Encode(w); err != nil {
			return utils.WrapError("Encode extendedCP_TxSidelink_r16", err)
		}
	}
	if ie.harq_ReportOnPUCCH_r16 != nil {
		if err = ie.harq_ReportOnPUCCH_r16.Encode(w); err != nil {
			return utils.WrapError("Encode harq_ReportOnPUCCH_r16", err)
		}
	}
	return nil
}

func (ie *BandSidelink_r16_sl_TransmissionMode1_r16) Decode(r *uper.UperReader) error {
	var err error
	var extendedCP_TxSidelink_r16Present bool
	if extendedCP_TxSidelink_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var harq_ReportOnPUCCH_r16Present bool
	if harq_ReportOnPUCCH_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.harq_TxProcessModeOneSidelink_r16.Decode(r); err != nil {
		return utils.WrapError("Decode harq_TxProcessModeOneSidelink_r16", err)
	}
	if err = ie.scs_CP_PatternTxSidelinkModeOne_r16.Decode(r); err != nil {
		return utils.WrapError("Decode scs_CP_PatternTxSidelinkModeOne_r16", err)
	}
	if extendedCP_TxSidelink_r16Present {
		ie.extendedCP_TxSidelink_r16 = new(BandSidelink_r16_sl_TransmissionMode1_r16_extendedCP_TxSidelink_r16)
		if err = ie.extendedCP_TxSidelink_r16.Decode(r); err != nil {
			return utils.WrapError("Decode extendedCP_TxSidelink_r16", err)
		}
	}
	if harq_ReportOnPUCCH_r16Present {
		ie.harq_ReportOnPUCCH_r16 = new(BandSidelink_r16_sl_TransmissionMode1_r16_harq_ReportOnPUCCH_r16)
		if err = ie.harq_ReportOnPUCCH_r16.Decode(r); err != nil {
			return utils.WrapError("Decode harq_ReportOnPUCCH_r16", err)
		}
	}
	return nil
}
