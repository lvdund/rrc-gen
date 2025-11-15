package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	CG_UCI_OnPUSCH_Choice_nothing uint64 = iota
	CG_UCI_OnPUSCH_Choice_dynamic
	CG_UCI_OnPUSCH_Choice_semiStatic
)

type CG_UCI_OnPUSCH struct {
	Choice     uint64
	dynamic    []BetaOffsets `lb:1,ub:4,madatory`
	semiStatic *BetaOffsets
}

func (ie *CG_UCI_OnPUSCH) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CG_UCI_OnPUSCH_Choice_dynamic:
		tmp := utils.NewSequence[*BetaOffsets]([]*BetaOffsets{}, uper.Constraint{Lb: 1, Ub: 4}, false)
		for _, i := range ie.dynamic {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = utils.WrapError("Encode dynamic", err)
		}
	case CG_UCI_OnPUSCH_Choice_semiStatic:
		if err = ie.semiStatic.Encode(w); err != nil {
			err = utils.WrapError("Encode semiStatic", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *CG_UCI_OnPUSCH) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CG_UCI_OnPUSCH_Choice_dynamic:
		tmp := utils.NewSequence[*BetaOffsets]([]*BetaOffsets{}, uper.Constraint{Lb: 1, Ub: 4}, false)
		fn := func() *BetaOffsets {
			return new(BetaOffsets)
		}
		if err = tmp.Decode(r, fn); err != nil {
			return utils.WrapError("Decode dynamic", err)
		}
		ie.dynamic = []BetaOffsets{}
		for _, i := range tmp.Value {
			ie.dynamic = append(ie.dynamic, *i)
		}
	case CG_UCI_OnPUSCH_Choice_semiStatic:
		ie.semiStatic = new(BetaOffsets)
		if err = ie.semiStatic.Decode(r); err != nil {
			return utils.WrapError("Decode semiStatic", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
