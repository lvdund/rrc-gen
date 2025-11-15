package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	RLC_Config_Choice_nothing uint64 = iota
	RLC_Config_Choice_am
	RLC_Config_Choice_um_Bi_Directional
	RLC_Config_Choice_um_Uni_Directional_UL
	RLC_Config_Choice_um_Uni_Directional_DL
)

type RLC_Config struct {
	Choice                uint64
	am                    *RLC_Config_am
	um_Bi_Directional     *RLC_Config_um_Bi_Directional
	um_Uni_Directional_UL *RLC_Config_um_Uni_Directional_UL
	um_Uni_Directional_DL *RLC_Config_um_Uni_Directional_DL
}

func (ie *RLC_Config) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 4, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RLC_Config_Choice_am:
		if err = ie.am.Encode(w); err != nil {
			err = utils.WrapError("Encode am", err)
		}
	case RLC_Config_Choice_um_Bi_Directional:
		if err = ie.um_Bi_Directional.Encode(w); err != nil {
			err = utils.WrapError("Encode um_Bi_Directional", err)
		}
	case RLC_Config_Choice_um_Uni_Directional_UL:
		if err = ie.um_Uni_Directional_UL.Encode(w); err != nil {
			err = utils.WrapError("Encode um_Uni_Directional_UL", err)
		}
	case RLC_Config_Choice_um_Uni_Directional_DL:
		if err = ie.um_Uni_Directional_DL.Encode(w); err != nil {
			err = utils.WrapError("Encode um_Uni_Directional_DL", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *RLC_Config) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(4, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RLC_Config_Choice_am:
		ie.am = new(RLC_Config_am)
		if err = ie.am.Decode(r); err != nil {
			return utils.WrapError("Decode am", err)
		}
	case RLC_Config_Choice_um_Bi_Directional:
		ie.um_Bi_Directional = new(RLC_Config_um_Bi_Directional)
		if err = ie.um_Bi_Directional.Decode(r); err != nil {
			return utils.WrapError("Decode um_Bi_Directional", err)
		}
	case RLC_Config_Choice_um_Uni_Directional_UL:
		ie.um_Uni_Directional_UL = new(RLC_Config_um_Uni_Directional_UL)
		if err = ie.um_Uni_Directional_UL.Decode(r); err != nil {
			return utils.WrapError("Decode um_Uni_Directional_UL", err)
		}
	case RLC_Config_Choice_um_Uni_Directional_DL:
		ie.um_Uni_Directional_DL = new(RLC_Config_um_Uni_Directional_DL)
		if err = ie.um_Uni_Directional_DL.Decode(r); err != nil {
			return utils.WrapError("Decode um_Uni_Directional_DL", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
