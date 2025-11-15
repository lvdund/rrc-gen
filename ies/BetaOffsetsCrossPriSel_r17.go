package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	BetaOffsetsCrossPriSel_r17_Choice_nothing uint64 = iota
	BetaOffsetsCrossPriSel_r17_Choice_dynamic_r17
	BetaOffsetsCrossPriSel_r17_Choice_semiStatic_r17
)

type BetaOffsetsCrossPriSel_r17 struct {
	Choice         uint64
	dynamic_r17    []BetaOffsetsCrossPri_r17 `lb:4,ub:4,madatory`
	semiStatic_r17 *BetaOffsetsCrossPri_r17
}

func (ie *BetaOffsetsCrossPriSel_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case BetaOffsetsCrossPriSel_r17_Choice_dynamic_r17:
		tmp := utils.NewSequence[*BetaOffsetsCrossPri_r17]([]*BetaOffsetsCrossPri_r17{}, uper.Constraint{Lb: 4, Ub: 4}, false)
		for _, i := range ie.dynamic_r17 {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = utils.WrapError("Encode dynamic_r17", err)
		}
	case BetaOffsetsCrossPriSel_r17_Choice_semiStatic_r17:
		if err = ie.semiStatic_r17.Encode(w); err != nil {
			err = utils.WrapError("Encode semiStatic_r17", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *BetaOffsetsCrossPriSel_r17) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case BetaOffsetsCrossPriSel_r17_Choice_dynamic_r17:
		tmp := utils.NewSequence[*BetaOffsetsCrossPri_r17]([]*BetaOffsetsCrossPri_r17{}, uper.Constraint{Lb: 4, Ub: 4}, false)
		fn := func() *BetaOffsetsCrossPri_r17 {
			return new(BetaOffsetsCrossPri_r17)
		}
		if err = tmp.Decode(r, fn); err != nil {
			return utils.WrapError("Decode dynamic_r17", err)
		}
		ie.dynamic_r17 = []BetaOffsetsCrossPri_r17{}
		for _, i := range tmp.Value {
			ie.dynamic_r17 = append(ie.dynamic_r17, *i)
		}
	case BetaOffsetsCrossPriSel_r17_Choice_semiStatic_r17:
		ie.semiStatic_r17 = new(BetaOffsetsCrossPri_r17)
		if err = ie.semiStatic_r17.Decode(r); err != nil {
			return utils.WrapError("Decode semiStatic_r17", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
