package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	TDD_UL_DL_SlotConfig_symbols_Choice_nothing uint64 = iota
	TDD_UL_DL_SlotConfig_symbols_Choice_allDownlink
	TDD_UL_DL_SlotConfig_symbols_Choice_allUplink
	TDD_UL_DL_SlotConfig_symbols_Choice_explicit
)

type TDD_UL_DL_SlotConfig_symbols struct {
	Choice      uint64
	allDownlink uper.NULL `madatory`
	allUplink   uper.NULL `madatory`
	explicit    *TDD_UL_DL_SlotConfig_symbols_explicit
}

func (ie *TDD_UL_DL_SlotConfig_symbols) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 3, false); err != nil {
		return err
	}
	switch ie.Choice {
	case TDD_UL_DL_SlotConfig_symbols_Choice_allDownlink:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode allDownlink", err)
		}
	case TDD_UL_DL_SlotConfig_symbols_Choice_allUplink:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode allUplink", err)
		}
	case TDD_UL_DL_SlotConfig_symbols_Choice_explicit:
		if err = ie.explicit.Encode(w); err != nil {
			err = utils.WrapError("Encode explicit", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *TDD_UL_DL_SlotConfig_symbols) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(3, false); err != nil {
		return err
	}
	switch ie.Choice {
	case TDD_UL_DL_SlotConfig_symbols_Choice_allDownlink:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode allDownlink", err)
		}
	case TDD_UL_DL_SlotConfig_symbols_Choice_allUplink:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode allUplink", err)
		}
	case TDD_UL_DL_SlotConfig_symbols_Choice_explicit:
		ie.explicit = new(TDD_UL_DL_SlotConfig_symbols_explicit)
		if err = ie.explicit.Decode(r); err != nil {
			return utils.WrapError("Decode explicit", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
