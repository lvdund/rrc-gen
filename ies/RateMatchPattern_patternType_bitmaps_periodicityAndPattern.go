package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	RateMatchPattern_patternType_bitmaps_periodicityAndPattern_Choice_nothing uint64 = iota
	RateMatchPattern_patternType_bitmaps_periodicityAndPattern_Choice_n2
	RateMatchPattern_patternType_bitmaps_periodicityAndPattern_Choice_n4
	RateMatchPattern_patternType_bitmaps_periodicityAndPattern_Choice_n5
	RateMatchPattern_patternType_bitmaps_periodicityAndPattern_Choice_n8
	RateMatchPattern_patternType_bitmaps_periodicityAndPattern_Choice_n10
	RateMatchPattern_patternType_bitmaps_periodicityAndPattern_Choice_n20
	RateMatchPattern_patternType_bitmaps_periodicityAndPattern_Choice_n40
)

type RateMatchPattern_patternType_bitmaps_periodicityAndPattern struct {
	Choice uint64
	n2     uper.BitString `lb:2,ub:2,madatory`
	n4     uper.BitString `lb:4,ub:4,madatory`
	n5     uper.BitString `lb:5,ub:5,madatory`
	n8     uper.BitString `lb:8,ub:8,madatory`
	n10    uper.BitString `lb:10,ub:10,madatory`
	n20    uper.BitString `lb:20,ub:20,madatory`
	n40    uper.BitString `lb:40,ub:40,madatory`
}

func (ie *RateMatchPattern_patternType_bitmaps_periodicityAndPattern) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 7, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RateMatchPattern_patternType_bitmaps_periodicityAndPattern_Choice_n2:
		if err = w.WriteBitString(ie.n2.Bytes, uint(ie.n2.NumBits), &uper.Constraint{Lb: 2, Ub: 2}, false); err != nil {
			err = utils.WrapError("Encode n2", err)
		}
	case RateMatchPattern_patternType_bitmaps_periodicityAndPattern_Choice_n4:
		if err = w.WriteBitString(ie.n4.Bytes, uint(ie.n4.NumBits), &uper.Constraint{Lb: 4, Ub: 4}, false); err != nil {
			err = utils.WrapError("Encode n4", err)
		}
	case RateMatchPattern_patternType_bitmaps_periodicityAndPattern_Choice_n5:
		if err = w.WriteBitString(ie.n5.Bytes, uint(ie.n5.NumBits), &uper.Constraint{Lb: 5, Ub: 5}, false); err != nil {
			err = utils.WrapError("Encode n5", err)
		}
	case RateMatchPattern_patternType_bitmaps_periodicityAndPattern_Choice_n8:
		if err = w.WriteBitString(ie.n8.Bytes, uint(ie.n8.NumBits), &uper.Constraint{Lb: 8, Ub: 8}, false); err != nil {
			err = utils.WrapError("Encode n8", err)
		}
	case RateMatchPattern_patternType_bitmaps_periodicityAndPattern_Choice_n10:
		if err = w.WriteBitString(ie.n10.Bytes, uint(ie.n10.NumBits), &uper.Constraint{Lb: 10, Ub: 10}, false); err != nil {
			err = utils.WrapError("Encode n10", err)
		}
	case RateMatchPattern_patternType_bitmaps_periodicityAndPattern_Choice_n20:
		if err = w.WriteBitString(ie.n20.Bytes, uint(ie.n20.NumBits), &uper.Constraint{Lb: 20, Ub: 20}, false); err != nil {
			err = utils.WrapError("Encode n20", err)
		}
	case RateMatchPattern_patternType_bitmaps_periodicityAndPattern_Choice_n40:
		if err = w.WriteBitString(ie.n40.Bytes, uint(ie.n40.NumBits), &uper.Constraint{Lb: 40, Ub: 40}, false); err != nil {
			err = utils.WrapError("Encode n40", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *RateMatchPattern_patternType_bitmaps_periodicityAndPattern) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(7, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RateMatchPattern_patternType_bitmaps_periodicityAndPattern_Choice_n2:
		var tmp_bs_n2 []byte
		var n_n2 uint
		if tmp_bs_n2, n_n2, err = r.ReadBitString(&uper.Constraint{Lb: 2, Ub: 2}, false); err != nil {
			return utils.WrapError("Decode n2", err)
		}
		ie.n2 = uper.BitString{
			Bytes:   tmp_bs_n2,
			NumBits: uint64(n_n2),
		}
	case RateMatchPattern_patternType_bitmaps_periodicityAndPattern_Choice_n4:
		var tmp_bs_n4 []byte
		var n_n4 uint
		if tmp_bs_n4, n_n4, err = r.ReadBitString(&uper.Constraint{Lb: 4, Ub: 4}, false); err != nil {
			return utils.WrapError("Decode n4", err)
		}
		ie.n4 = uper.BitString{
			Bytes:   tmp_bs_n4,
			NumBits: uint64(n_n4),
		}
	case RateMatchPattern_patternType_bitmaps_periodicityAndPattern_Choice_n5:
		var tmp_bs_n5 []byte
		var n_n5 uint
		if tmp_bs_n5, n_n5, err = r.ReadBitString(&uper.Constraint{Lb: 5, Ub: 5}, false); err != nil {
			return utils.WrapError("Decode n5", err)
		}
		ie.n5 = uper.BitString{
			Bytes:   tmp_bs_n5,
			NumBits: uint64(n_n5),
		}
	case RateMatchPattern_patternType_bitmaps_periodicityAndPattern_Choice_n8:
		var tmp_bs_n8 []byte
		var n_n8 uint
		if tmp_bs_n8, n_n8, err = r.ReadBitString(&uper.Constraint{Lb: 8, Ub: 8}, false); err != nil {
			return utils.WrapError("Decode n8", err)
		}
		ie.n8 = uper.BitString{
			Bytes:   tmp_bs_n8,
			NumBits: uint64(n_n8),
		}
	case RateMatchPattern_patternType_bitmaps_periodicityAndPattern_Choice_n10:
		var tmp_bs_n10 []byte
		var n_n10 uint
		if tmp_bs_n10, n_n10, err = r.ReadBitString(&uper.Constraint{Lb: 10, Ub: 10}, false); err != nil {
			return utils.WrapError("Decode n10", err)
		}
		ie.n10 = uper.BitString{
			Bytes:   tmp_bs_n10,
			NumBits: uint64(n_n10),
		}
	case RateMatchPattern_patternType_bitmaps_periodicityAndPattern_Choice_n20:
		var tmp_bs_n20 []byte
		var n_n20 uint
		if tmp_bs_n20, n_n20, err = r.ReadBitString(&uper.Constraint{Lb: 20, Ub: 20}, false); err != nil {
			return utils.WrapError("Decode n20", err)
		}
		ie.n20 = uper.BitString{
			Bytes:   tmp_bs_n20,
			NumBits: uint64(n_n20),
		}
	case RateMatchPattern_patternType_bitmaps_periodicityAndPattern_Choice_n40:
		var tmp_bs_n40 []byte
		var n_n40 uint
		if tmp_bs_n40, n_n40, err = r.ReadBitString(&uper.Constraint{Lb: 40, Ub: 40}, false); err != nil {
			return utils.WrapError("Decode n40", err)
		}
		ie.n40 = uper.BitString{
			Bytes:   tmp_bs_n40,
			NumBits: uint64(n_n40),
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
