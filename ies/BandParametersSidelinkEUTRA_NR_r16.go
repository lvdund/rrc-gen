package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	BandParametersSidelinkEUTRA_NR_r16_Choice_nothing uint64 = iota
	BandParametersSidelinkEUTRA_NR_r16_Choice_eutra
	BandParametersSidelinkEUTRA_NR_r16_Choice_nr
)

type BandParametersSidelinkEUTRA_NR_r16 struct {
	Choice uint64
	eutra  *BandParametersSidelinkEUTRA_NR_r16_eutra
	nr     *BandParametersSidelinkEUTRA_NR_r16_nr
}

func (ie *BandParametersSidelinkEUTRA_NR_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case BandParametersSidelinkEUTRA_NR_r16_Choice_eutra:
		if err = ie.eutra.Encode(w); err != nil {
			err = utils.WrapError("Encode eutra", err)
		}
	case BandParametersSidelinkEUTRA_NR_r16_Choice_nr:
		if err = ie.nr.Encode(w); err != nil {
			err = utils.WrapError("Encode nr", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *BandParametersSidelinkEUTRA_NR_r16) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case BandParametersSidelinkEUTRA_NR_r16_Choice_eutra:
		ie.eutra = new(BandParametersSidelinkEUTRA_NR_r16_eutra)
		if err = ie.eutra.Decode(r); err != nil {
			return utils.WrapError("Decode eutra", err)
		}
	case BandParametersSidelinkEUTRA_NR_r16_Choice_nr:
		ie.nr = new(BandParametersSidelinkEUTRA_NR_r16_nr)
		if err = ie.nr.Decode(r); err != nil {
			return utils.WrapError("Decode nr", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
