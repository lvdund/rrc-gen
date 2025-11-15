package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SupportedBandwidth_Choice_nothing uint64 = iota
	SupportedBandwidth_Choice_fr1
	SupportedBandwidth_Choice_fr2
)

type SupportedBandwidth struct {
	Choice uint64
	fr1    *SupportedBandwidth_fr1
	fr2    *SupportedBandwidth_fr2
}

func (ie *SupportedBandwidth) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SupportedBandwidth_Choice_fr1:
		if err = ie.fr1.Encode(w); err != nil {
			err = utils.WrapError("Encode fr1", err)
		}
	case SupportedBandwidth_Choice_fr2:
		if err = ie.fr2.Encode(w); err != nil {
			err = utils.WrapError("Encode fr2", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *SupportedBandwidth) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SupportedBandwidth_Choice_fr1:
		ie.fr1 = new(SupportedBandwidth_fr1)
		if err = ie.fr1.Decode(r); err != nil {
			return utils.WrapError("Decode fr1", err)
		}
	case SupportedBandwidth_Choice_fr2:
		ie.fr2 = new(SupportedBandwidth_fr2)
		if err = ie.fr2.Decode(r); err != nil {
			return utils.WrapError("Decode fr2", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
