package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SSB_ToMeasure_Choice_nothing uint64 = iota
	SSB_ToMeasure_Choice_shortBitmap
	SSB_ToMeasure_Choice_mediumBitmap
	SSB_ToMeasure_Choice_longBitmap
)

type SSB_ToMeasure struct {
	Choice       uint64
	shortBitmap  uper.BitString `lb:4,ub:4,madatory`
	mediumBitmap uper.BitString `lb:8,ub:8,madatory`
	longBitmap   uper.BitString `lb:64,ub:64,madatory`
}

func (ie *SSB_ToMeasure) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 3, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SSB_ToMeasure_Choice_shortBitmap:
		if err = w.WriteBitString(ie.shortBitmap.Bytes, uint(ie.shortBitmap.NumBits), &uper.Constraint{Lb: 4, Ub: 4}, false); err != nil {
			err = utils.WrapError("Encode shortBitmap", err)
		}
	case SSB_ToMeasure_Choice_mediumBitmap:
		if err = w.WriteBitString(ie.mediumBitmap.Bytes, uint(ie.mediumBitmap.NumBits), &uper.Constraint{Lb: 8, Ub: 8}, false); err != nil {
			err = utils.WrapError("Encode mediumBitmap", err)
		}
	case SSB_ToMeasure_Choice_longBitmap:
		if err = w.WriteBitString(ie.longBitmap.Bytes, uint(ie.longBitmap.NumBits), &uper.Constraint{Lb: 64, Ub: 64}, false); err != nil {
			err = utils.WrapError("Encode longBitmap", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *SSB_ToMeasure) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(3, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SSB_ToMeasure_Choice_shortBitmap:
		var tmp_bs_shortBitmap []byte
		var n_shortBitmap uint
		if tmp_bs_shortBitmap, n_shortBitmap, err = r.ReadBitString(&uper.Constraint{Lb: 4, Ub: 4}, false); err != nil {
			return utils.WrapError("Decode shortBitmap", err)
		}
		ie.shortBitmap = uper.BitString{
			Bytes:   tmp_bs_shortBitmap,
			NumBits: uint64(n_shortBitmap),
		}
	case SSB_ToMeasure_Choice_mediumBitmap:
		var tmp_bs_mediumBitmap []byte
		var n_mediumBitmap uint
		if tmp_bs_mediumBitmap, n_mediumBitmap, err = r.ReadBitString(&uper.Constraint{Lb: 8, Ub: 8}, false); err != nil {
			return utils.WrapError("Decode mediumBitmap", err)
		}
		ie.mediumBitmap = uper.BitString{
			Bytes:   tmp_bs_mediumBitmap,
			NumBits: uint64(n_mediumBitmap),
		}
	case SSB_ToMeasure_Choice_longBitmap:
		var tmp_bs_longBitmap []byte
		var n_longBitmap uint
		if tmp_bs_longBitmap, n_longBitmap, err = r.ReadBitString(&uper.Constraint{Lb: 64, Ub: 64}, false); err != nil {
			return utils.WrapError("Decode longBitmap", err)
		}
		ie.longBitmap = uper.BitString{
			Bytes:   tmp_bs_longBitmap,
			NumBits: uint64(n_longBitmap),
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
