package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	BandNR_channelBW_DL_IAB_r16_Choice_nothing uint64 = iota
	BandNR_channelBW_DL_IAB_r16_Choice_fr1_100mhz
	BandNR_channelBW_DL_IAB_r16_Choice_fr2_200mhz
)

type BandNR_channelBW_DL_IAB_r16 struct {
	Choice     uint64
	fr1_100mhz *Fr1_100mhz
	fr2_200mhz *Fr2_200mhz
}

func (ie *BandNR_channelBW_DL_IAB_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case BandNR_channelBW_DL_IAB_r16_Choice_fr1_100mhz:
		if err = ie.fr1_100mhz.Encode(w); err != nil {
			err = utils.WrapError("Encode fr1_100mhz", err)
		}
	case BandNR_channelBW_DL_IAB_r16_Choice_fr2_200mhz:
		if err = ie.fr2_200mhz.Encode(w); err != nil {
			err = utils.WrapError("Encode fr2_200mhz", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *BandNR_channelBW_DL_IAB_r16) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case BandNR_channelBW_DL_IAB_r16_Choice_fr1_100mhz:
		ie.fr1_100mhz = new(Fr1_100mhz)
		if err = ie.fr1_100mhz.Decode(r); err != nil {
			return utils.WrapError("Decode fr1_100mhz", err)
		}
	case BandNR_channelBW_DL_IAB_r16_Choice_fr2_200mhz:
		ie.fr2_200mhz = new(Fr2_200mhz)
		if err = ie.fr2_200mhz.Decode(r); err != nil {
			return utils.WrapError("Decode fr2_200mhz", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
