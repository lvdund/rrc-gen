package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BandParametersSidelinkEUTRA_NR_v1710_nr struct {
	sl_TransmissionMode2_PartialSensing_r17 *BandParametersSidelinkEUTRA_NR_v1710_nr_sl_TransmissionMode2_PartialSensing_r17 `optional`
	rx_sidelinkPSFCH_r17                    *BandParametersSidelinkEUTRA_NR_v1710_nr_rx_sidelinkPSFCH_r17                    `optional`
	tx_IUC_Scheme1_Mode2Sidelink_r17        *BandParametersSidelinkEUTRA_NR_v1710_nr_tx_IUC_Scheme1_Mode2Sidelink_r17        `optional`
	tx_IUC_Scheme2_Mode2Sidelink_r17        *BandParametersSidelinkEUTRA_NR_v1710_nr_tx_IUC_Scheme2_Mode2Sidelink_r17        `optional`
}

func (ie *BandParametersSidelinkEUTRA_NR_v1710_nr) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.sl_TransmissionMode2_PartialSensing_r17 != nil, ie.rx_sidelinkPSFCH_r17 != nil, ie.tx_IUC_Scheme1_Mode2Sidelink_r17 != nil, ie.tx_IUC_Scheme2_Mode2Sidelink_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.sl_TransmissionMode2_PartialSensing_r17 != nil {
		if err = ie.sl_TransmissionMode2_PartialSensing_r17.Encode(w); err != nil {
			return utils.WrapError("Encode sl_TransmissionMode2_PartialSensing_r17", err)
		}
	}
	if ie.rx_sidelinkPSFCH_r17 != nil {
		if err = ie.rx_sidelinkPSFCH_r17.Encode(w); err != nil {
			return utils.WrapError("Encode rx_sidelinkPSFCH_r17", err)
		}
	}
	if ie.tx_IUC_Scheme1_Mode2Sidelink_r17 != nil {
		if err = ie.tx_IUC_Scheme1_Mode2Sidelink_r17.Encode(w); err != nil {
			return utils.WrapError("Encode tx_IUC_Scheme1_Mode2Sidelink_r17", err)
		}
	}
	if ie.tx_IUC_Scheme2_Mode2Sidelink_r17 != nil {
		if err = ie.tx_IUC_Scheme2_Mode2Sidelink_r17.Encode(w); err != nil {
			return utils.WrapError("Encode tx_IUC_Scheme2_Mode2Sidelink_r17", err)
		}
	}
	return nil
}

func (ie *BandParametersSidelinkEUTRA_NR_v1710_nr) Decode(r *uper.UperReader) error {
	var err error
	var sl_TransmissionMode2_PartialSensing_r17Present bool
	if sl_TransmissionMode2_PartialSensing_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var rx_sidelinkPSFCH_r17Present bool
	if rx_sidelinkPSFCH_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var tx_IUC_Scheme1_Mode2Sidelink_r17Present bool
	if tx_IUC_Scheme1_Mode2Sidelink_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var tx_IUC_Scheme2_Mode2Sidelink_r17Present bool
	if tx_IUC_Scheme2_Mode2Sidelink_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if sl_TransmissionMode2_PartialSensing_r17Present {
		ie.sl_TransmissionMode2_PartialSensing_r17 = new(BandParametersSidelinkEUTRA_NR_v1710_nr_sl_TransmissionMode2_PartialSensing_r17)
		if err = ie.sl_TransmissionMode2_PartialSensing_r17.Decode(r); err != nil {
			return utils.WrapError("Decode sl_TransmissionMode2_PartialSensing_r17", err)
		}
	}
	if rx_sidelinkPSFCH_r17Present {
		ie.rx_sidelinkPSFCH_r17 = new(BandParametersSidelinkEUTRA_NR_v1710_nr_rx_sidelinkPSFCH_r17)
		if err = ie.rx_sidelinkPSFCH_r17.Decode(r); err != nil {
			return utils.WrapError("Decode rx_sidelinkPSFCH_r17", err)
		}
	}
	if tx_IUC_Scheme1_Mode2Sidelink_r17Present {
		ie.tx_IUC_Scheme1_Mode2Sidelink_r17 = new(BandParametersSidelinkEUTRA_NR_v1710_nr_tx_IUC_Scheme1_Mode2Sidelink_r17)
		if err = ie.tx_IUC_Scheme1_Mode2Sidelink_r17.Decode(r); err != nil {
			return utils.WrapError("Decode tx_IUC_Scheme1_Mode2Sidelink_r17", err)
		}
	}
	if tx_IUC_Scheme2_Mode2Sidelink_r17Present {
		ie.tx_IUC_Scheme2_Mode2Sidelink_r17 = new(BandParametersSidelinkEUTRA_NR_v1710_nr_tx_IUC_Scheme2_Mode2Sidelink_r17)
		if err = ie.tx_IUC_Scheme2_Mode2Sidelink_r17.Decode(r); err != nil {
			return utils.WrapError("Decode tx_IUC_Scheme2_Mode2Sidelink_r17", err)
		}
	}
	return nil
}
