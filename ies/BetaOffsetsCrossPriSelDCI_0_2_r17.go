package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	BetaOffsetsCrossPriSelDCI_0_2_r17_Choice_nothing uint64 = iota
	BetaOffsetsCrossPriSelDCI_0_2_r17_Choice_dynamicDCI_0_2_r17
	BetaOffsetsCrossPriSelDCI_0_2_r17_Choice_semiStaticDCI_0_2_r17
)

type BetaOffsetsCrossPriSelDCI_0_2_r17 struct {
	Choice                uint64
	dynamicDCI_0_2_r17    *BetaOffsetsCrossPriSelDCI_0_2_r17_dynamicDCI_0_2_r17
	semiStaticDCI_0_2_r17 *BetaOffsetsCrossPri_r17
}

func (ie *BetaOffsetsCrossPriSelDCI_0_2_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case BetaOffsetsCrossPriSelDCI_0_2_r17_Choice_dynamicDCI_0_2_r17:
		if err = ie.dynamicDCI_0_2_r17.Encode(w); err != nil {
			err = utils.WrapError("Encode dynamicDCI_0_2_r17", err)
		}
	case BetaOffsetsCrossPriSelDCI_0_2_r17_Choice_semiStaticDCI_0_2_r17:
		if err = ie.semiStaticDCI_0_2_r17.Encode(w); err != nil {
			err = utils.WrapError("Encode semiStaticDCI_0_2_r17", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *BetaOffsetsCrossPriSelDCI_0_2_r17) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case BetaOffsetsCrossPriSelDCI_0_2_r17_Choice_dynamicDCI_0_2_r17:
		ie.dynamicDCI_0_2_r17 = new(BetaOffsetsCrossPriSelDCI_0_2_r17_dynamicDCI_0_2_r17)
		if err = ie.dynamicDCI_0_2_r17.Decode(r); err != nil {
			return utils.WrapError("Decode dynamicDCI_0_2_r17", err)
		}
	case BetaOffsetsCrossPriSelDCI_0_2_r17_Choice_semiStaticDCI_0_2_r17:
		ie.semiStaticDCI_0_2_r17 = new(BetaOffsetsCrossPri_r17)
		if err = ie.semiStaticDCI_0_2_r17.Decode(r); err != nil {
			return utils.WrapError("Decode semiStaticDCI_0_2_r17", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
