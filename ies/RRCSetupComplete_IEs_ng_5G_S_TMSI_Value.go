package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	RRCSetupComplete_IEs_ng_5G_S_TMSI_Value_Choice_nothing uint64 = iota
	RRCSetupComplete_IEs_ng_5G_S_TMSI_Value_Choice_ng_5G_S_TMSI
	RRCSetupComplete_IEs_ng_5G_S_TMSI_Value_Choice_ng_5G_S_TMSI_Part2
)

type RRCSetupComplete_IEs_ng_5G_S_TMSI_Value struct {
	Choice             uint64
	ng_5G_S_TMSI       *NG_5G_S_TMSI
	ng_5G_S_TMSI_Part2 uper.BitString `lb:9,ub:9,madatory`
}

func (ie *RRCSetupComplete_IEs_ng_5G_S_TMSI_Value) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RRCSetupComplete_IEs_ng_5G_S_TMSI_Value_Choice_ng_5G_S_TMSI:
		if err = ie.ng_5G_S_TMSI.Encode(w); err != nil {
			err = utils.WrapError("Encode ng_5G_S_TMSI", err)
		}
	case RRCSetupComplete_IEs_ng_5G_S_TMSI_Value_Choice_ng_5G_S_TMSI_Part2:
		if err = w.WriteBitString(ie.ng_5G_S_TMSI_Part2.Bytes, uint(ie.ng_5G_S_TMSI_Part2.NumBits), &uper.Constraint{Lb: 9, Ub: 9}, false); err != nil {
			err = utils.WrapError("Encode ng_5G_S_TMSI_Part2", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *RRCSetupComplete_IEs_ng_5G_S_TMSI_Value) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RRCSetupComplete_IEs_ng_5G_S_TMSI_Value_Choice_ng_5G_S_TMSI:
		ie.ng_5G_S_TMSI = new(NG_5G_S_TMSI)
		if err = ie.ng_5G_S_TMSI.Decode(r); err != nil {
			return utils.WrapError("Decode ng_5G_S_TMSI", err)
		}
	case RRCSetupComplete_IEs_ng_5G_S_TMSI_Value_Choice_ng_5G_S_TMSI_Part2:
		var tmp_bs_ng_5G_S_TMSI_Part2 []byte
		var n_ng_5G_S_TMSI_Part2 uint
		if tmp_bs_ng_5G_S_TMSI_Part2, n_ng_5G_S_TMSI_Part2, err = r.ReadBitString(&uper.Constraint{Lb: 9, Ub: 9}, false); err != nil {
			return utils.WrapError("Decode ng_5G_S_TMSI_Part2", err)
		}
		ie.ng_5G_S_TMSI_Part2 = uper.BitString{
			Bytes:   tmp_bs_ng_5G_S_TMSI_Part2,
			NumBits: uint64(n_ng_5G_S_TMSI_Part2),
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
