package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	PUSCH_Config_frequencyHoppingDCI_0_2_r16_Choice_nothing uint64 = iota
	PUSCH_Config_frequencyHoppingDCI_0_2_r16_Choice_pusch_RepTypeA
	PUSCH_Config_frequencyHoppingDCI_0_2_r16_Choice_pusch_RepTypeB
)

type PUSCH_Config_frequencyHoppingDCI_0_2_r16 struct {
	Choice         uint64
	pusch_RepTypeA *PUSCH_Config_frequencyHoppingDCI_0_2_r16_pusch_RepTypeA
	pusch_RepTypeB *PUSCH_Config_frequencyHoppingDCI_0_2_r16_pusch_RepTypeB
}

func (ie *PUSCH_Config_frequencyHoppingDCI_0_2_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case PUSCH_Config_frequencyHoppingDCI_0_2_r16_Choice_pusch_RepTypeA:
		if err = ie.pusch_RepTypeA.Encode(w); err != nil {
			err = utils.WrapError("Encode pusch_RepTypeA", err)
		}
	case PUSCH_Config_frequencyHoppingDCI_0_2_r16_Choice_pusch_RepTypeB:
		if err = ie.pusch_RepTypeB.Encode(w); err != nil {
			err = utils.WrapError("Encode pusch_RepTypeB", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *PUSCH_Config_frequencyHoppingDCI_0_2_r16) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case PUSCH_Config_frequencyHoppingDCI_0_2_r16_Choice_pusch_RepTypeA:
		ie.pusch_RepTypeA = new(PUSCH_Config_frequencyHoppingDCI_0_2_r16_pusch_RepTypeA)
		if err = ie.pusch_RepTypeA.Decode(r); err != nil {
			return utils.WrapError("Decode pusch_RepTypeA", err)
		}
	case PUSCH_Config_frequencyHoppingDCI_0_2_r16_Choice_pusch_RepTypeB:
		ie.pusch_RepTypeB = new(PUSCH_Config_frequencyHoppingDCI_0_2_r16_pusch_RepTypeB)
		if err = ie.pusch_RepTypeB.Decode(r); err != nil {
			return utils.WrapError("Decode pusch_RepTypeB", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
