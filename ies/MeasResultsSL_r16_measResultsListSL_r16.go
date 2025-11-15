package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	MeasResultsSL_r16_measResultsListSL_r16_Choice_nothing uint64 = iota
	MeasResultsSL_r16_measResultsListSL_r16_Choice_measResultNR_SL_r16
)

type MeasResultsSL_r16_measResultsListSL_r16 struct {
	Choice              uint64
	measResultNR_SL_r16 *MeasResultNR_SL_r16
}

func (ie *MeasResultsSL_r16_measResultsListSL_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 1, false); err != nil {
		return err
	}
	switch ie.Choice {
	case MeasResultsSL_r16_measResultsListSL_r16_Choice_measResultNR_SL_r16:
		if err = ie.measResultNR_SL_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode measResultNR_SL_r16", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *MeasResultsSL_r16_measResultsListSL_r16) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(1, false); err != nil {
		return err
	}
	switch ie.Choice {
	case MeasResultsSL_r16_measResultsListSL_r16_Choice_measResultNR_SL_r16:
		ie.measResultNR_SL_r16 = new(MeasResultNR_SL_r16)
		if err = ie.measResultNR_SL_r16.Decode(r); err != nil {
			return utils.WrapError("Decode measResultNR_SL_r16", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
