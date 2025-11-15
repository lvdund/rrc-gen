package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BandParametersSidelinkEUTRA_NR_v1710_nr_sl_TransmissionMode2_PartialSensing_r17 struct {
	harq_TxProcessModeTwoSidelink_r17   BandParametersSidelinkEUTRA_NR_v1710_nr_sl_TransmissionMode2_PartialSensing_r17_harq_TxProcessModeTwoSidelink_r17    `madatory`
	scs_CP_PatternTxSidelinkModeTwo_r17 *BandParametersSidelinkEUTRA_NR_v1710_nr_sl_TransmissionMode2_PartialSensing_r17_scs_CP_PatternTxSidelinkModeTwo_r17 `optional`
	extendedCP_Mode2PartialSensing_r17  *BandParametersSidelinkEUTRA_NR_v1710_nr_sl_TransmissionMode2_PartialSensing_r17_extendedCP_Mode2PartialSensing_r17  `optional`
	dl_openLoopPC_Sidelink_r17          *BandParametersSidelinkEUTRA_NR_v1710_nr_sl_TransmissionMode2_PartialSensing_r17_dl_openLoopPC_Sidelink_r17          `optional`
}

func (ie *BandParametersSidelinkEUTRA_NR_v1710_nr_sl_TransmissionMode2_PartialSensing_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.scs_CP_PatternTxSidelinkModeTwo_r17 != nil, ie.extendedCP_Mode2PartialSensing_r17 != nil, ie.dl_openLoopPC_Sidelink_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.harq_TxProcessModeTwoSidelink_r17.Encode(w); err != nil {
		return utils.WrapError("Encode harq_TxProcessModeTwoSidelink_r17", err)
	}
	if ie.scs_CP_PatternTxSidelinkModeTwo_r17 != nil {
		if err = ie.scs_CP_PatternTxSidelinkModeTwo_r17.Encode(w); err != nil {
			return utils.WrapError("Encode scs_CP_PatternTxSidelinkModeTwo_r17", err)
		}
	}
	if ie.extendedCP_Mode2PartialSensing_r17 != nil {
		if err = ie.extendedCP_Mode2PartialSensing_r17.Encode(w); err != nil {
			return utils.WrapError("Encode extendedCP_Mode2PartialSensing_r17", err)
		}
	}
	if ie.dl_openLoopPC_Sidelink_r17 != nil {
		if err = ie.dl_openLoopPC_Sidelink_r17.Encode(w); err != nil {
			return utils.WrapError("Encode dl_openLoopPC_Sidelink_r17", err)
		}
	}
	return nil
}

func (ie *BandParametersSidelinkEUTRA_NR_v1710_nr_sl_TransmissionMode2_PartialSensing_r17) Decode(r *uper.UperReader) error {
	var err error
	var scs_CP_PatternTxSidelinkModeTwo_r17Present bool
	if scs_CP_PatternTxSidelinkModeTwo_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var extendedCP_Mode2PartialSensing_r17Present bool
	if extendedCP_Mode2PartialSensing_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var dl_openLoopPC_Sidelink_r17Present bool
	if dl_openLoopPC_Sidelink_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.harq_TxProcessModeTwoSidelink_r17.Decode(r); err != nil {
		return utils.WrapError("Decode harq_TxProcessModeTwoSidelink_r17", err)
	}
	if scs_CP_PatternTxSidelinkModeTwo_r17Present {
		ie.scs_CP_PatternTxSidelinkModeTwo_r17 = new(BandParametersSidelinkEUTRA_NR_v1710_nr_sl_TransmissionMode2_PartialSensing_r17_scs_CP_PatternTxSidelinkModeTwo_r17)
		if err = ie.scs_CP_PatternTxSidelinkModeTwo_r17.Decode(r); err != nil {
			return utils.WrapError("Decode scs_CP_PatternTxSidelinkModeTwo_r17", err)
		}
	}
	if extendedCP_Mode2PartialSensing_r17Present {
		ie.extendedCP_Mode2PartialSensing_r17 = new(BandParametersSidelinkEUTRA_NR_v1710_nr_sl_TransmissionMode2_PartialSensing_r17_extendedCP_Mode2PartialSensing_r17)
		if err = ie.extendedCP_Mode2PartialSensing_r17.Decode(r); err != nil {
			return utils.WrapError("Decode extendedCP_Mode2PartialSensing_r17", err)
		}
	}
	if dl_openLoopPC_Sidelink_r17Present {
		ie.dl_openLoopPC_Sidelink_r17 = new(BandParametersSidelinkEUTRA_NR_v1710_nr_sl_TransmissionMode2_PartialSensing_r17_dl_openLoopPC_Sidelink_r17)
		if err = ie.dl_openLoopPC_Sidelink_r17.Decode(r); err != nil {
			return utils.WrapError("Decode dl_openLoopPC_Sidelink_r17", err)
		}
	}
	return nil
}
