package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	BandParameters_Choice_nothing uint64 = iota
	BandParameters_Choice_eutra
	BandParameters_Choice_nr
)

type BandParameters struct {
	Choice uint64
	eutra  *BandParameters_eutra
	nr     *BandParameters_nr
}

func (ie *BandParameters) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case BandParameters_Choice_eutra:
		if err = ie.eutra.Encode(w); err != nil {
			err = utils.WrapError("Encode eutra", err)
		}
	case BandParameters_Choice_nr:
		if err = ie.nr.Encode(w); err != nil {
			err = utils.WrapError("Encode nr", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *BandParameters) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case BandParameters_Choice_eutra:
		ie.eutra = new(BandParameters_eutra)
		if err = ie.eutra.Decode(r); err != nil {
			return utils.WrapError("Decode eutra", err)
		}
	case BandParameters_Choice_nr:
		ie.nr = new(BandParameters_nr)
		if err = ie.nr.Decode(r); err != nil {
			return utils.WrapError("Decode nr", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
