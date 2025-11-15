package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BandSidelink_r16_sl_TransmissionMode2_r16 struct {
	harq_TxProcessModeTwoSidelink_r16   BandSidelink_r16_sl_TransmissionMode2_r16_harq_TxProcessModeTwoSidelink_r16    `madatory`
	scs_CP_PatternTxSidelinkModeTwo_r16 *BandSidelink_r16_sl_TransmissionMode2_r16_scs_CP_PatternTxSidelinkModeTwo_r16 `optional`
	dl_openLoopPC_Sidelink_r16          *BandSidelink_r16_sl_TransmissionMode2_r16_dl_openLoopPC_Sidelink_r16          `optional`
}

func (ie *BandSidelink_r16_sl_TransmissionMode2_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.scs_CP_PatternTxSidelinkModeTwo_r16 != nil, ie.dl_openLoopPC_Sidelink_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.harq_TxProcessModeTwoSidelink_r16.Encode(w); err != nil {
		return utils.WrapError("Encode harq_TxProcessModeTwoSidelink_r16", err)
	}
	if ie.scs_CP_PatternTxSidelinkModeTwo_r16 != nil {
		if err = ie.scs_CP_PatternTxSidelinkModeTwo_r16.Encode(w); err != nil {
			return utils.WrapError("Encode scs_CP_PatternTxSidelinkModeTwo_r16", err)
		}
	}
	if ie.dl_openLoopPC_Sidelink_r16 != nil {
		if err = ie.dl_openLoopPC_Sidelink_r16.Encode(w); err != nil {
			return utils.WrapError("Encode dl_openLoopPC_Sidelink_r16", err)
		}
	}
	return nil
}

func (ie *BandSidelink_r16_sl_TransmissionMode2_r16) Decode(r *uper.UperReader) error {
	var err error
	var scs_CP_PatternTxSidelinkModeTwo_r16Present bool
	if scs_CP_PatternTxSidelinkModeTwo_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var dl_openLoopPC_Sidelink_r16Present bool
	if dl_openLoopPC_Sidelink_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.harq_TxProcessModeTwoSidelink_r16.Decode(r); err != nil {
		return utils.WrapError("Decode harq_TxProcessModeTwoSidelink_r16", err)
	}
	if scs_CP_PatternTxSidelinkModeTwo_r16Present {
		ie.scs_CP_PatternTxSidelinkModeTwo_r16 = new(BandSidelink_r16_sl_TransmissionMode2_r16_scs_CP_PatternTxSidelinkModeTwo_r16)
		if err = ie.scs_CP_PatternTxSidelinkModeTwo_r16.Decode(r); err != nil {
			return utils.WrapError("Decode scs_CP_PatternTxSidelinkModeTwo_r16", err)
		}
	}
	if dl_openLoopPC_Sidelink_r16Present {
		ie.dl_openLoopPC_Sidelink_r16 = new(BandSidelink_r16_sl_TransmissionMode2_r16_dl_openLoopPC_Sidelink_r16)
		if err = ie.dl_openLoopPC_Sidelink_r16.Decode(r); err != nil {
			return utils.WrapError("Decode dl_openLoopPC_Sidelink_r16", err)
		}
	}
	return nil
}
