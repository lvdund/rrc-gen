package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	UCI_OnPUSCH_DCI_0_2_r16_betaOffsetsDCI_0_2_r16_Choice_nothing uint64 = iota
	UCI_OnPUSCH_DCI_0_2_r16_betaOffsetsDCI_0_2_r16_Choice_dynamicDCI_0_2_r16
	UCI_OnPUSCH_DCI_0_2_r16_betaOffsetsDCI_0_2_r16_Choice_semiStaticDCI_0_2_r16
)

type UCI_OnPUSCH_DCI_0_2_r16_betaOffsetsDCI_0_2_r16 struct {
	Choice                uint64
	dynamicDCI_0_2_r16    *UCI_OnPUSCH_DCI_0_2_r16_betaOffsetsDCI_0_2_r16_dynamicDCI_0_2_r16
	semiStaticDCI_0_2_r16 *BetaOffsets
}

func (ie *UCI_OnPUSCH_DCI_0_2_r16_betaOffsetsDCI_0_2_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case UCI_OnPUSCH_DCI_0_2_r16_betaOffsetsDCI_0_2_r16_Choice_dynamicDCI_0_2_r16:
		if err = ie.dynamicDCI_0_2_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode dynamicDCI_0_2_r16", err)
		}
	case UCI_OnPUSCH_DCI_0_2_r16_betaOffsetsDCI_0_2_r16_Choice_semiStaticDCI_0_2_r16:
		if err = ie.semiStaticDCI_0_2_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode semiStaticDCI_0_2_r16", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *UCI_OnPUSCH_DCI_0_2_r16_betaOffsetsDCI_0_2_r16) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case UCI_OnPUSCH_DCI_0_2_r16_betaOffsetsDCI_0_2_r16_Choice_dynamicDCI_0_2_r16:
		ie.dynamicDCI_0_2_r16 = new(UCI_OnPUSCH_DCI_0_2_r16_betaOffsetsDCI_0_2_r16_dynamicDCI_0_2_r16)
		if err = ie.dynamicDCI_0_2_r16.Decode(r); err != nil {
			return utils.WrapError("Decode dynamicDCI_0_2_r16", err)
		}
	case UCI_OnPUSCH_DCI_0_2_r16_betaOffsetsDCI_0_2_r16_Choice_semiStaticDCI_0_2_r16:
		ie.semiStaticDCI_0_2_r16 = new(BetaOffsets)
		if err = ie.semiStaticDCI_0_2_r16.Decode(r); err != nil {
			return utils.WrapError("Decode semiStaticDCI_0_2_r16", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
