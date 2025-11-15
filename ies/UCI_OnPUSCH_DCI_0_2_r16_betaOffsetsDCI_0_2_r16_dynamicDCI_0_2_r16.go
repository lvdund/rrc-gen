package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	UCI_OnPUSCH_DCI_0_2_r16_betaOffsetsDCI_0_2_r16_dynamicDCI_0_2_r16_Choice_nothing uint64 = iota
	UCI_OnPUSCH_DCI_0_2_r16_betaOffsetsDCI_0_2_r16_dynamicDCI_0_2_r16_Choice_oneBit_r16
	UCI_OnPUSCH_DCI_0_2_r16_betaOffsetsDCI_0_2_r16_dynamicDCI_0_2_r16_Choice_twoBits_r16
)

type UCI_OnPUSCH_DCI_0_2_r16_betaOffsetsDCI_0_2_r16_dynamicDCI_0_2_r16 struct {
	Choice      uint64
	oneBit_r16  []BetaOffsets `lb:2,ub:2,madatory`
	twoBits_r16 []BetaOffsets `lb:4,ub:4,madatory`
}

func (ie *UCI_OnPUSCH_DCI_0_2_r16_betaOffsetsDCI_0_2_r16_dynamicDCI_0_2_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case UCI_OnPUSCH_DCI_0_2_r16_betaOffsetsDCI_0_2_r16_dynamicDCI_0_2_r16_Choice_oneBit_r16:
		tmp := utils.NewSequence[*BetaOffsets]([]*BetaOffsets{}, uper.Constraint{Lb: 2, Ub: 2}, false)
		for _, i := range ie.oneBit_r16 {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = utils.WrapError("Encode oneBit_r16", err)
		}
	case UCI_OnPUSCH_DCI_0_2_r16_betaOffsetsDCI_0_2_r16_dynamicDCI_0_2_r16_Choice_twoBits_r16:
		tmp := utils.NewSequence[*BetaOffsets]([]*BetaOffsets{}, uper.Constraint{Lb: 4, Ub: 4}, false)
		for _, i := range ie.twoBits_r16 {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = utils.WrapError("Encode twoBits_r16", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *UCI_OnPUSCH_DCI_0_2_r16_betaOffsetsDCI_0_2_r16_dynamicDCI_0_2_r16) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case UCI_OnPUSCH_DCI_0_2_r16_betaOffsetsDCI_0_2_r16_dynamicDCI_0_2_r16_Choice_oneBit_r16:
		tmp := utils.NewSequence[*BetaOffsets]([]*BetaOffsets{}, uper.Constraint{Lb: 2, Ub: 2}, false)
		fn := func() *BetaOffsets {
			return new(BetaOffsets)
		}
		if err = tmp.Decode(r, fn); err != nil {
			return utils.WrapError("Decode oneBit_r16", err)
		}
		ie.oneBit_r16 = []BetaOffsets{}
		for _, i := range tmp.Value {
			ie.oneBit_r16 = append(ie.oneBit_r16, *i)
		}
	case UCI_OnPUSCH_DCI_0_2_r16_betaOffsetsDCI_0_2_r16_dynamicDCI_0_2_r16_Choice_twoBits_r16:
		tmp := utils.NewSequence[*BetaOffsets]([]*BetaOffsets{}, uper.Constraint{Lb: 4, Ub: 4}, false)
		fn := func() *BetaOffsets {
			return new(BetaOffsets)
		}
		if err = tmp.Decode(r, fn); err != nil {
			return utils.WrapError("Decode twoBits_r16", err)
		}
		ie.twoBits_r16 = []BetaOffsets{}
		for _, i := range tmp.Value {
			ie.twoBits_r16 = append(ie.twoBits_r16, *i)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
