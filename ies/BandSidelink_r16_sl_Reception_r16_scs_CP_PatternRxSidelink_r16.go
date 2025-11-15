package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	BandSidelink_r16_sl_Reception_r16_scs_CP_PatternRxSidelink_r16_Choice_nothing uint64 = iota
	BandSidelink_r16_sl_Reception_r16_scs_CP_PatternRxSidelink_r16_Choice_fr1_r16
	BandSidelink_r16_sl_Reception_r16_scs_CP_PatternRxSidelink_r16_Choice_fr2_r16
)

type BandSidelink_r16_sl_Reception_r16_scs_CP_PatternRxSidelink_r16 struct {
	Choice  uint64
	fr1_r16 *Fr1_r16
	fr2_r16 *Fr2_r16
}

func (ie *BandSidelink_r16_sl_Reception_r16_scs_CP_PatternRxSidelink_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case BandSidelink_r16_sl_Reception_r16_scs_CP_PatternRxSidelink_r16_Choice_fr1_r16:
		if err = ie.fr1_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode fr1_r16", err)
		}
	case BandSidelink_r16_sl_Reception_r16_scs_CP_PatternRxSidelink_r16_Choice_fr2_r16:
		if err = ie.fr2_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode fr2_r16", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *BandSidelink_r16_sl_Reception_r16_scs_CP_PatternRxSidelink_r16) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case BandSidelink_r16_sl_Reception_r16_scs_CP_PatternRxSidelink_r16_Choice_fr1_r16:
		ie.fr1_r16 = new(Fr1_r16)
		if err = ie.fr1_r16.Decode(r); err != nil {
			return utils.WrapError("Decode fr1_r16", err)
		}
	case BandSidelink_r16_sl_Reception_r16_scs_CP_PatternRxSidelink_r16_Choice_fr2_r16:
		ie.fr2_r16 = new(Fr2_r16)
		if err = ie.fr2_r16.Decode(r); err != nil {
			return utils.WrapError("Decode fr2_r16", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
