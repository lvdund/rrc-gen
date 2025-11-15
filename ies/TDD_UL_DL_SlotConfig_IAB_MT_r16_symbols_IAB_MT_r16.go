package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	TDD_UL_DL_SlotConfig_IAB_MT_r16_symbols_IAB_MT_r16_Choice_nothing uint64 = iota
	TDD_UL_DL_SlotConfig_IAB_MT_r16_symbols_IAB_MT_r16_Choice_allDownlink_r16
	TDD_UL_DL_SlotConfig_IAB_MT_r16_symbols_IAB_MT_r16_Choice_allUplink_r16
	TDD_UL_DL_SlotConfig_IAB_MT_r16_symbols_IAB_MT_r16_Choice_explicit_r16
	TDD_UL_DL_SlotConfig_IAB_MT_r16_symbols_IAB_MT_r16_Choice_explicit_IAB_MT_r16
)

type TDD_UL_DL_SlotConfig_IAB_MT_r16_symbols_IAB_MT_r16 struct {
	Choice              uint64
	allDownlink_r16     uper.NULL `madatory`
	allUplink_r16       uper.NULL `madatory`
	explicit_r16        *TDD_UL_DL_SlotConfig_IAB_MT_r16_symbols_IAB_MT_r16_explicit_r16
	explicit_IAB_MT_r16 *TDD_UL_DL_SlotConfig_IAB_MT_r16_symbols_IAB_MT_r16_explicit_IAB_MT_r16
}

func (ie *TDD_UL_DL_SlotConfig_IAB_MT_r16_symbols_IAB_MT_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 4, false); err != nil {
		return err
	}
	switch ie.Choice {
	case TDD_UL_DL_SlotConfig_IAB_MT_r16_symbols_IAB_MT_r16_Choice_allDownlink_r16:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode allDownlink_r16", err)
		}
	case TDD_UL_DL_SlotConfig_IAB_MT_r16_symbols_IAB_MT_r16_Choice_allUplink_r16:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode allUplink_r16", err)
		}
	case TDD_UL_DL_SlotConfig_IAB_MT_r16_symbols_IAB_MT_r16_Choice_explicit_r16:
		if err = ie.explicit_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode explicit_r16", err)
		}
	case TDD_UL_DL_SlotConfig_IAB_MT_r16_symbols_IAB_MT_r16_Choice_explicit_IAB_MT_r16:
		if err = ie.explicit_IAB_MT_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode explicit_IAB_MT_r16", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *TDD_UL_DL_SlotConfig_IAB_MT_r16_symbols_IAB_MT_r16) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(4, false); err != nil {
		return err
	}
	switch ie.Choice {
	case TDD_UL_DL_SlotConfig_IAB_MT_r16_symbols_IAB_MT_r16_Choice_allDownlink_r16:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode allDownlink_r16", err)
		}
	case TDD_UL_DL_SlotConfig_IAB_MT_r16_symbols_IAB_MT_r16_Choice_allUplink_r16:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode allUplink_r16", err)
		}
	case TDD_UL_DL_SlotConfig_IAB_MT_r16_symbols_IAB_MT_r16_Choice_explicit_r16:
		ie.explicit_r16 = new(TDD_UL_DL_SlotConfig_IAB_MT_r16_symbols_IAB_MT_r16_explicit_r16)
		if err = ie.explicit_r16.Decode(r); err != nil {
			return utils.WrapError("Decode explicit_r16", err)
		}
	case TDD_UL_DL_SlotConfig_IAB_MT_r16_symbols_IAB_MT_r16_Choice_explicit_IAB_MT_r16:
		ie.explicit_IAB_MT_r16 = new(TDD_UL_DL_SlotConfig_IAB_MT_r16_symbols_IAB_MT_r16_explicit_IAB_MT_r16)
		if err = ie.explicit_IAB_MT_r16.Decode(r); err != nil {
			return utils.WrapError("Decode explicit_IAB_MT_r16", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
