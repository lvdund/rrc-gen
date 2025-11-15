package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	RepetitionSchemeConfig_r16_Choice_nothing uint64 = iota
	RepetitionSchemeConfig_r16_Choice_fdm_TDM_r16
	RepetitionSchemeConfig_r16_Choice_slotBased_r16
)

type RepetitionSchemeConfig_r16 struct {
	Choice        uint64
	fdm_TDM_r16   *FDM_TDM_r16
	slotBased_r16 *SlotBased_r16
}

func (ie *RepetitionSchemeConfig_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RepetitionSchemeConfig_r16_Choice_fdm_TDM_r16:
		if err = ie.fdm_TDM_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode fdm_TDM_r16", err)
		}
	case RepetitionSchemeConfig_r16_Choice_slotBased_r16:
		if err = ie.slotBased_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode slotBased_r16", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *RepetitionSchemeConfig_r16) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RepetitionSchemeConfig_r16_Choice_fdm_TDM_r16:
		ie.fdm_TDM_r16 = new(FDM_TDM_r16)
		if err = ie.fdm_TDM_r16.Decode(r); err != nil {
			return utils.WrapError("Decode fdm_TDM_r16", err)
		}
	case RepetitionSchemeConfig_r16_Choice_slotBased_r16:
		ie.slotBased_r16 = new(SlotBased_r16)
		if err = ie.slotBased_r16.Decode(r); err != nil {
			return utils.WrapError("Decode slotBased_r16", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
