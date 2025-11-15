package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	InitialUE_Identity_Choice_nothing uint64 = iota
	InitialUE_Identity_Choice_ng_5G_S_TMSI_Part1
	InitialUE_Identity_Choice_randomValue
)

type InitialUE_Identity struct {
	Choice             uint64
	ng_5G_S_TMSI_Part1 uper.BitString `lb:39,ub:39,madatory`
	randomValue        uper.BitString `lb:39,ub:39,madatory`
}

func (ie *InitialUE_Identity) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case InitialUE_Identity_Choice_ng_5G_S_TMSI_Part1:
		if err = w.WriteBitString(ie.ng_5G_S_TMSI_Part1.Bytes, uint(ie.ng_5G_S_TMSI_Part1.NumBits), &uper.Constraint{Lb: 39, Ub: 39}, false); err != nil {
			err = utils.WrapError("Encode ng_5G_S_TMSI_Part1", err)
		}
	case InitialUE_Identity_Choice_randomValue:
		if err = w.WriteBitString(ie.randomValue.Bytes, uint(ie.randomValue.NumBits), &uper.Constraint{Lb: 39, Ub: 39}, false); err != nil {
			err = utils.WrapError("Encode randomValue", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *InitialUE_Identity) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case InitialUE_Identity_Choice_ng_5G_S_TMSI_Part1:
		var tmp_bs_ng_5G_S_TMSI_Part1 []byte
		var n_ng_5G_S_TMSI_Part1 uint
		if tmp_bs_ng_5G_S_TMSI_Part1, n_ng_5G_S_TMSI_Part1, err = r.ReadBitString(&uper.Constraint{Lb: 39, Ub: 39}, false); err != nil {
			return utils.WrapError("Decode ng_5G_S_TMSI_Part1", err)
		}
		ie.ng_5G_S_TMSI_Part1 = uper.BitString{
			Bytes:   tmp_bs_ng_5G_S_TMSI_Part1,
			NumBits: uint64(n_ng_5G_S_TMSI_Part1),
		}
	case InitialUE_Identity_Choice_randomValue:
		var tmp_bs_randomValue []byte
		var n_randomValue uint
		if tmp_bs_randomValue, n_randomValue, err = r.ReadBitString(&uper.Constraint{Lb: 39, Ub: 39}, false); err != nil {
			return utils.WrapError("Decode randomValue", err)
		}
		ie.randomValue = uper.BitString{
			Bytes:   tmp_bs_randomValue,
			NumBits: uint64(n_randomValue),
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
