package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	PUCCH_Config_subslotLengthForPUCCH_r16_Choice_nothing uint64 = iota
	PUCCH_Config_subslotLengthForPUCCH_r16_Choice_normalCP_r16
	PUCCH_Config_subslotLengthForPUCCH_r16_Choice_extendedCP_r16
)

type PUCCH_Config_subslotLengthForPUCCH_r16 struct {
	Choice         uint64
	normalCP_r16   *PUCCH_Config_subslotLengthForPUCCH_r16_normalCP_r16
	extendedCP_r16 *PUCCH_Config_subslotLengthForPUCCH_r16_extendedCP_r16
}

func (ie *PUCCH_Config_subslotLengthForPUCCH_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case PUCCH_Config_subslotLengthForPUCCH_r16_Choice_normalCP_r16:
		if err = ie.normalCP_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode normalCP_r16", err)
		}
	case PUCCH_Config_subslotLengthForPUCCH_r16_Choice_extendedCP_r16:
		if err = ie.extendedCP_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode extendedCP_r16", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *PUCCH_Config_subslotLengthForPUCCH_r16) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case PUCCH_Config_subslotLengthForPUCCH_r16_Choice_normalCP_r16:
		ie.normalCP_r16 = new(PUCCH_Config_subslotLengthForPUCCH_r16_normalCP_r16)
		if err = ie.normalCP_r16.Decode(r); err != nil {
			return utils.WrapError("Decode normalCP_r16", err)
		}
	case PUCCH_Config_subslotLengthForPUCCH_r16_Choice_extendedCP_r16:
		ie.extendedCP_r16 = new(PUCCH_Config_subslotLengthForPUCCH_r16_extendedCP_r16)
		if err = ie.extendedCP_r16.Decode(r); err != nil {
			return utils.WrapError("Decode extendedCP_r16", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
