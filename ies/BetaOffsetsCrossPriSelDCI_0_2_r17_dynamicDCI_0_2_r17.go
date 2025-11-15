package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	BetaOffsetsCrossPriSelDCI_0_2_r17_dynamicDCI_0_2_r17_Choice_nothing uint64 = iota
	BetaOffsetsCrossPriSelDCI_0_2_r17_dynamicDCI_0_2_r17_Choice_oneBit_r17
	BetaOffsetsCrossPriSelDCI_0_2_r17_dynamicDCI_0_2_r17_Choice_twoBits_r17
)

type BetaOffsetsCrossPriSelDCI_0_2_r17_dynamicDCI_0_2_r17 struct {
	Choice      uint64
	oneBit_r17  []BetaOffsetsCrossPri_r17 `lb:2,ub:2,madatory`
	twoBits_r17 []BetaOffsetsCrossPri_r17 `lb:4,ub:4,madatory`
}

func (ie *BetaOffsetsCrossPriSelDCI_0_2_r17_dynamicDCI_0_2_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case BetaOffsetsCrossPriSelDCI_0_2_r17_dynamicDCI_0_2_r17_Choice_oneBit_r17:
		tmp := utils.NewSequence[*BetaOffsetsCrossPri_r17]([]*BetaOffsetsCrossPri_r17{}, uper.Constraint{Lb: 2, Ub: 2}, false)
		for _, i := range ie.oneBit_r17 {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = utils.WrapError("Encode oneBit_r17", err)
		}
	case BetaOffsetsCrossPriSelDCI_0_2_r17_dynamicDCI_0_2_r17_Choice_twoBits_r17:
		tmp := utils.NewSequence[*BetaOffsetsCrossPri_r17]([]*BetaOffsetsCrossPri_r17{}, uper.Constraint{Lb: 4, Ub: 4}, false)
		for _, i := range ie.twoBits_r17 {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = utils.WrapError("Encode twoBits_r17", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *BetaOffsetsCrossPriSelDCI_0_2_r17_dynamicDCI_0_2_r17) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case BetaOffsetsCrossPriSelDCI_0_2_r17_dynamicDCI_0_2_r17_Choice_oneBit_r17:
		tmp := utils.NewSequence[*BetaOffsetsCrossPri_r17]([]*BetaOffsetsCrossPri_r17{}, uper.Constraint{Lb: 2, Ub: 2}, false)
		fn := func() *BetaOffsetsCrossPri_r17 {
			return new(BetaOffsetsCrossPri_r17)
		}
		if err = tmp.Decode(r, fn); err != nil {
			return utils.WrapError("Decode oneBit_r17", err)
		}
		ie.oneBit_r17 = []BetaOffsetsCrossPri_r17{}
		for _, i := range tmp.Value {
			ie.oneBit_r17 = append(ie.oneBit_r17, *i)
		}
	case BetaOffsetsCrossPriSelDCI_0_2_r17_dynamicDCI_0_2_r17_Choice_twoBits_r17:
		tmp := utils.NewSequence[*BetaOffsetsCrossPri_r17]([]*BetaOffsetsCrossPri_r17{}, uper.Constraint{Lb: 4, Ub: 4}, false)
		fn := func() *BetaOffsetsCrossPri_r17 {
			return new(BetaOffsetsCrossPri_r17)
		}
		if err = tmp.Decode(r, fn); err != nil {
			return utils.WrapError("Decode twoBits_r17", err)
		}
		ie.twoBits_r17 = []BetaOffsetsCrossPri_r17{}
		for _, i := range tmp.Value {
			ie.twoBits_r17 = append(ie.twoBits_r17, *i)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
