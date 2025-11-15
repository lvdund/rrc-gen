package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	UCI_OnPUSCH_betaOffsets_Choice_nothing uint64 = iota
	UCI_OnPUSCH_betaOffsets_Choice_dynamic
	UCI_OnPUSCH_betaOffsets_Choice_semiStatic
)

type UCI_OnPUSCH_betaOffsets struct {
	Choice     uint64
	dynamic    []BetaOffsets `lb:4,ub:4,madatory`
	semiStatic *BetaOffsets
}

func (ie *UCI_OnPUSCH_betaOffsets) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case UCI_OnPUSCH_betaOffsets_Choice_dynamic:
		tmp := utils.NewSequence[*BetaOffsets]([]*BetaOffsets{}, uper.Constraint{Lb: 4, Ub: 4}, false)
		for _, i := range ie.dynamic {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = utils.WrapError("Encode dynamic", err)
		}
	case UCI_OnPUSCH_betaOffsets_Choice_semiStatic:
		if err = ie.semiStatic.Encode(w); err != nil {
			err = utils.WrapError("Encode semiStatic", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *UCI_OnPUSCH_betaOffsets) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case UCI_OnPUSCH_betaOffsets_Choice_dynamic:
		tmp := utils.NewSequence[*BetaOffsets]([]*BetaOffsets{}, uper.Constraint{Lb: 4, Ub: 4}, false)
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
	case UCI_OnPUSCH_betaOffsets_Choice_semiStatic:
		ie.semiStatic = new(BetaOffsets)
		if err = ie.semiStatic.Decode(r); err != nil {
			return utils.WrapError("Decode semiStatic", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
