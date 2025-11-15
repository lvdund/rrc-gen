package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	BandNR_channelBWs_UL_v1590_Choice_nothing uint64 = iota
	BandNR_channelBWs_UL_v1590_Choice_fr1
	BandNR_channelBWs_UL_v1590_Choice_fr2
)

type BandNR_channelBWs_UL_v1590 struct {
	Choice uint64
	fr1    *BandNR_channelBWs_UL_v1590_fr1
	fr2    *BandNR_channelBWs_UL_v1590_fr2
}

func (ie *BandNR_channelBWs_UL_v1590) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case BandNR_channelBWs_UL_v1590_Choice_fr1:
		if err = ie.fr1.Encode(w); err != nil {
			err = utils.WrapError("Encode fr1", err)
		}
	case BandNR_channelBWs_UL_v1590_Choice_fr2:
		if err = ie.fr2.Encode(w); err != nil {
			err = utils.WrapError("Encode fr2", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *BandNR_channelBWs_UL_v1590) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case BandNR_channelBWs_UL_v1590_Choice_fr1:
		ie.fr1 = new(BandNR_channelBWs_UL_v1590_fr1)
		if err = ie.fr1.Decode(r); err != nil {
			return utils.WrapError("Decode fr1", err)
		}
	case BandNR_channelBWs_UL_v1590_Choice_fr2:
		ie.fr2 = new(BandNR_channelBWs_UL_v1590_fr2)
		if err = ie.fr2.Decode(r); err != nil {
			return utils.WrapError("Decode fr2", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
