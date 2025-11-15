package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	BandParametersSidelinkDiscovery_r17_sl_TransmissionMode2_PartialSensing_r17_scs_CP_PatternTxSidelinkModeTwo_r17_Choice_nothing uint64 = iota
	BandParametersSidelinkDiscovery_r17_sl_TransmissionMode2_PartialSensing_r17_scs_CP_PatternTxSidelinkModeTwo_r17_Choice_fr1_r17
	BandParametersSidelinkDiscovery_r17_sl_TransmissionMode2_PartialSensing_r17_scs_CP_PatternTxSidelinkModeTwo_r17_Choice_fr2_r17
)

type BandParametersSidelinkDiscovery_r17_sl_TransmissionMode2_PartialSensing_r17_scs_CP_PatternTxSidelinkModeTwo_r17 struct {
	Choice  uint64
	fr1_r17 *Fr1_r17
	fr2_r17 *Fr2_r17
}

func (ie *BandParametersSidelinkDiscovery_r17_sl_TransmissionMode2_PartialSensing_r17_scs_CP_PatternTxSidelinkModeTwo_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case BandParametersSidelinkDiscovery_r17_sl_TransmissionMode2_PartialSensing_r17_scs_CP_PatternTxSidelinkModeTwo_r17_Choice_fr1_r17:
		if err = ie.fr1_r17.Encode(w); err != nil {
			err = utils.WrapError("Encode fr1_r17", err)
		}
	case BandParametersSidelinkDiscovery_r17_sl_TransmissionMode2_PartialSensing_r17_scs_CP_PatternTxSidelinkModeTwo_r17_Choice_fr2_r17:
		if err = ie.fr2_r17.Encode(w); err != nil {
			err = utils.WrapError("Encode fr2_r17", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *BandParametersSidelinkDiscovery_r17_sl_TransmissionMode2_PartialSensing_r17_scs_CP_PatternTxSidelinkModeTwo_r17) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case BandParametersSidelinkDiscovery_r17_sl_TransmissionMode2_PartialSensing_r17_scs_CP_PatternTxSidelinkModeTwo_r17_Choice_fr1_r17:
		ie.fr1_r17 = new(Fr1_r17)
		if err = ie.fr1_r17.Decode(r); err != nil {
			return utils.WrapError("Decode fr1_r17", err)
		}
	case BandParametersSidelinkDiscovery_r17_sl_TransmissionMode2_PartialSensing_r17_scs_CP_PatternTxSidelinkModeTwo_r17_Choice_fr2_r17:
		ie.fr2_r17 = new(Fr2_r17)
		if err = ie.fr2_r17.Decode(r); err != nil {
			return utils.WrapError("Decode fr2_r17", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
