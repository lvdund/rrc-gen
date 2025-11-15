package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SupportedBandwidth_v1700_Choice_nothing uint64 = iota
	SupportedBandwidth_v1700_Choice_fr1_r17
	SupportedBandwidth_v1700_Choice_fr2_r17
)

type SupportedBandwidth_v1700 struct {
	Choice  uint64
	fr1_r17 *SupportedBandwidth_v1700_fr1_r17
	fr2_r17 *SupportedBandwidth_v1700_fr2_r17
}

func (ie *SupportedBandwidth_v1700) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SupportedBandwidth_v1700_Choice_fr1_r17:
		if err = ie.fr1_r17.Encode(w); err != nil {
			err = utils.WrapError("Encode fr1_r17", err)
		}
	case SupportedBandwidth_v1700_Choice_fr2_r17:
		if err = ie.fr2_r17.Encode(w); err != nil {
			err = utils.WrapError("Encode fr2_r17", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *SupportedBandwidth_v1700) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SupportedBandwidth_v1700_Choice_fr1_r17:
		ie.fr1_r17 = new(SupportedBandwidth_v1700_fr1_r17)
		if err = ie.fr1_r17.Decode(r); err != nil {
			return utils.WrapError("Decode fr1_r17", err)
		}
	case SupportedBandwidth_v1700_Choice_fr2_r17:
		ie.fr2_r17 = new(SupportedBandwidth_v1700_fr2_r17)
		if err = ie.fr2_r17.Decode(r); err != nil {
			return utils.WrapError("Decode fr2_r17", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
