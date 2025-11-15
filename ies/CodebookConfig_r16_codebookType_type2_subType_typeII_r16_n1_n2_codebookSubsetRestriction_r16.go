package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	CodebookConfig_r16_codebookType_type2_subType_typeII_r16_n1_n2_codebookSubsetRestriction_r16_Choice_nothing uint64 = iota
	CodebookConfig_r16_codebookType_type2_subType_typeII_r16_n1_n2_codebookSubsetRestriction_r16_Choice_two_one
	CodebookConfig_r16_codebookType_type2_subType_typeII_r16_n1_n2_codebookSubsetRestriction_r16_Choice_two_two
	CodebookConfig_r16_codebookType_type2_subType_typeII_r16_n1_n2_codebookSubsetRestriction_r16_Choice_four_one
	CodebookConfig_r16_codebookType_type2_subType_typeII_r16_n1_n2_codebookSubsetRestriction_r16_Choice_three_two
	CodebookConfig_r16_codebookType_type2_subType_typeII_r16_n1_n2_codebookSubsetRestriction_r16_Choice_six_one
	CodebookConfig_r16_codebookType_type2_subType_typeII_r16_n1_n2_codebookSubsetRestriction_r16_Choice_four_two
	CodebookConfig_r16_codebookType_type2_subType_typeII_r16_n1_n2_codebookSubsetRestriction_r16_Choice_eight_one
	CodebookConfig_r16_codebookType_type2_subType_typeII_r16_n1_n2_codebookSubsetRestriction_r16_Choice_four_three
	CodebookConfig_r16_codebookType_type2_subType_typeII_r16_n1_n2_codebookSubsetRestriction_r16_Choice_six_two
	CodebookConfig_r16_codebookType_type2_subType_typeII_r16_n1_n2_codebookSubsetRestriction_r16_Choice_twelve_one
	CodebookConfig_r16_codebookType_type2_subType_typeII_r16_n1_n2_codebookSubsetRestriction_r16_Choice_four_four
	CodebookConfig_r16_codebookType_type2_subType_typeII_r16_n1_n2_codebookSubsetRestriction_r16_Choice_eight_two
	CodebookConfig_r16_codebookType_type2_subType_typeII_r16_n1_n2_codebookSubsetRestriction_r16_Choice_sixteen_one
)

type CodebookConfig_r16_codebookType_type2_subType_typeII_r16_n1_n2_codebookSubsetRestriction_r16 struct {
	Choice      uint64
	two_one     uper.BitString `lb:16,ub:16,madatory`
	two_two     uper.BitString `lb:43,ub:43,madatory`
	four_one    uper.BitString `lb:32,ub:32,madatory`
	three_two   uper.BitString `lb:59,ub:59,madatory`
	six_one     uper.BitString `lb:48,ub:48,madatory`
	four_two    uper.BitString `lb:75,ub:75,madatory`
	eight_one   uper.BitString `lb:64,ub:64,madatory`
	four_three  uper.BitString `lb:107,ub:107,madatory`
	six_two     uper.BitString `lb:107,ub:107,madatory`
	twelve_one  uper.BitString `lb:96,ub:96,madatory`
	four_four   uper.BitString `lb:139,ub:139,madatory`
	eight_two   uper.BitString `lb:139,ub:139,madatory`
	sixteen_one uper.BitString `lb:128,ub:128,madatory`
}

func (ie *CodebookConfig_r16_codebookType_type2_subType_typeII_r16_n1_n2_codebookSubsetRestriction_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 13, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CodebookConfig_r16_codebookType_type2_subType_typeII_r16_n1_n2_codebookSubsetRestriction_r16_Choice_two_one:
		if err = w.WriteBitString(ie.two_one.Bytes, uint(ie.two_one.NumBits), &uper.Constraint{Lb: 16, Ub: 16}, false); err != nil {
			err = utils.WrapError("Encode two_one", err)
		}
	case CodebookConfig_r16_codebookType_type2_subType_typeII_r16_n1_n2_codebookSubsetRestriction_r16_Choice_two_two:
		if err = w.WriteBitString(ie.two_two.Bytes, uint(ie.two_two.NumBits), &uper.Constraint{Lb: 43, Ub: 43}, false); err != nil {
			err = utils.WrapError("Encode two_two", err)
		}
	case CodebookConfig_r16_codebookType_type2_subType_typeII_r16_n1_n2_codebookSubsetRestriction_r16_Choice_four_one:
		if err = w.WriteBitString(ie.four_one.Bytes, uint(ie.four_one.NumBits), &uper.Constraint{Lb: 32, Ub: 32}, false); err != nil {
			err = utils.WrapError("Encode four_one", err)
		}
	case CodebookConfig_r16_codebookType_type2_subType_typeII_r16_n1_n2_codebookSubsetRestriction_r16_Choice_three_two:
		if err = w.WriteBitString(ie.three_two.Bytes, uint(ie.three_two.NumBits), &uper.Constraint{Lb: 59, Ub: 59}, false); err != nil {
			err = utils.WrapError("Encode three_two", err)
		}
	case CodebookConfig_r16_codebookType_type2_subType_typeII_r16_n1_n2_codebookSubsetRestriction_r16_Choice_six_one:
		if err = w.WriteBitString(ie.six_one.Bytes, uint(ie.six_one.NumBits), &uper.Constraint{Lb: 48, Ub: 48}, false); err != nil {
			err = utils.WrapError("Encode six_one", err)
		}
	case CodebookConfig_r16_codebookType_type2_subType_typeII_r16_n1_n2_codebookSubsetRestriction_r16_Choice_four_two:
		if err = w.WriteBitString(ie.four_two.Bytes, uint(ie.four_two.NumBits), &uper.Constraint{Lb: 75, Ub: 75}, false); err != nil {
			err = utils.WrapError("Encode four_two", err)
		}
	case CodebookConfig_r16_codebookType_type2_subType_typeII_r16_n1_n2_codebookSubsetRestriction_r16_Choice_eight_one:
		if err = w.WriteBitString(ie.eight_one.Bytes, uint(ie.eight_one.NumBits), &uper.Constraint{Lb: 64, Ub: 64}, false); err != nil {
			err = utils.WrapError("Encode eight_one", err)
		}
	case CodebookConfig_r16_codebookType_type2_subType_typeII_r16_n1_n2_codebookSubsetRestriction_r16_Choice_four_three:
		if err = w.WriteBitString(ie.four_three.Bytes, uint(ie.four_three.NumBits), &uper.Constraint{Lb: 107, Ub: 107}, false); err != nil {
			err = utils.WrapError("Encode four_three", err)
		}
	case CodebookConfig_r16_codebookType_type2_subType_typeII_r16_n1_n2_codebookSubsetRestriction_r16_Choice_six_two:
		if err = w.WriteBitString(ie.six_two.Bytes, uint(ie.six_two.NumBits), &uper.Constraint{Lb: 107, Ub: 107}, false); err != nil {
			err = utils.WrapError("Encode six_two", err)
		}
	case CodebookConfig_r16_codebookType_type2_subType_typeII_r16_n1_n2_codebookSubsetRestriction_r16_Choice_twelve_one:
		if err = w.WriteBitString(ie.twelve_one.Bytes, uint(ie.twelve_one.NumBits), &uper.Constraint{Lb: 96, Ub: 96}, false); err != nil {
			err = utils.WrapError("Encode twelve_one", err)
		}
	case CodebookConfig_r16_codebookType_type2_subType_typeII_r16_n1_n2_codebookSubsetRestriction_r16_Choice_four_four:
		if err = w.WriteBitString(ie.four_four.Bytes, uint(ie.four_four.NumBits), &uper.Constraint{Lb: 139, Ub: 139}, false); err != nil {
			err = utils.WrapError("Encode four_four", err)
		}
	case CodebookConfig_r16_codebookType_type2_subType_typeII_r16_n1_n2_codebookSubsetRestriction_r16_Choice_eight_two:
		if err = w.WriteBitString(ie.eight_two.Bytes, uint(ie.eight_two.NumBits), &uper.Constraint{Lb: 139, Ub: 139}, false); err != nil {
			err = utils.WrapError("Encode eight_two", err)
		}
	case CodebookConfig_r16_codebookType_type2_subType_typeII_r16_n1_n2_codebookSubsetRestriction_r16_Choice_sixteen_one:
		if err = w.WriteBitString(ie.sixteen_one.Bytes, uint(ie.sixteen_one.NumBits), &uper.Constraint{Lb: 128, Ub: 128}, false); err != nil {
			err = utils.WrapError("Encode sixteen_one", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *CodebookConfig_r16_codebookType_type2_subType_typeII_r16_n1_n2_codebookSubsetRestriction_r16) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(13, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CodebookConfig_r16_codebookType_type2_subType_typeII_r16_n1_n2_codebookSubsetRestriction_r16_Choice_two_one:
		var tmp_bs_two_one []byte
		var n_two_one uint
		if tmp_bs_two_one, n_two_one, err = r.ReadBitString(&uper.Constraint{Lb: 16, Ub: 16}, false); err != nil {
			return utils.WrapError("Decode two_one", err)
		}
		ie.two_one = uper.BitString{
			Bytes:   tmp_bs_two_one,
			NumBits: uint64(n_two_one),
		}
	case CodebookConfig_r16_codebookType_type2_subType_typeII_r16_n1_n2_codebookSubsetRestriction_r16_Choice_two_two:
		var tmp_bs_two_two []byte
		var n_two_two uint
		if tmp_bs_two_two, n_two_two, err = r.ReadBitString(&uper.Constraint{Lb: 43, Ub: 43}, false); err != nil {
			return utils.WrapError("Decode two_two", err)
		}
		ie.two_two = uper.BitString{
			Bytes:   tmp_bs_two_two,
			NumBits: uint64(n_two_two),
		}
	case CodebookConfig_r16_codebookType_type2_subType_typeII_r16_n1_n2_codebookSubsetRestriction_r16_Choice_four_one:
		var tmp_bs_four_one []byte
		var n_four_one uint
		if tmp_bs_four_one, n_four_one, err = r.ReadBitString(&uper.Constraint{Lb: 32, Ub: 32}, false); err != nil {
			return utils.WrapError("Decode four_one", err)
		}
		ie.four_one = uper.BitString{
			Bytes:   tmp_bs_four_one,
			NumBits: uint64(n_four_one),
		}
	case CodebookConfig_r16_codebookType_type2_subType_typeII_r16_n1_n2_codebookSubsetRestriction_r16_Choice_three_two:
		var tmp_bs_three_two []byte
		var n_three_two uint
		if tmp_bs_three_two, n_three_two, err = r.ReadBitString(&uper.Constraint{Lb: 59, Ub: 59}, false); err != nil {
			return utils.WrapError("Decode three_two", err)
		}
		ie.three_two = uper.BitString{
			Bytes:   tmp_bs_three_two,
			NumBits: uint64(n_three_two),
		}
	case CodebookConfig_r16_codebookType_type2_subType_typeII_r16_n1_n2_codebookSubsetRestriction_r16_Choice_six_one:
		var tmp_bs_six_one []byte
		var n_six_one uint
		if tmp_bs_six_one, n_six_one, err = r.ReadBitString(&uper.Constraint{Lb: 48, Ub: 48}, false); err != nil {
			return utils.WrapError("Decode six_one", err)
		}
		ie.six_one = uper.BitString{
			Bytes:   tmp_bs_six_one,
			NumBits: uint64(n_six_one),
		}
	case CodebookConfig_r16_codebookType_type2_subType_typeII_r16_n1_n2_codebookSubsetRestriction_r16_Choice_four_two:
		var tmp_bs_four_two []byte
		var n_four_two uint
		if tmp_bs_four_two, n_four_two, err = r.ReadBitString(&uper.Constraint{Lb: 75, Ub: 75}, false); err != nil {
			return utils.WrapError("Decode four_two", err)
		}
		ie.four_two = uper.BitString{
			Bytes:   tmp_bs_four_two,
			NumBits: uint64(n_four_two),
		}
	case CodebookConfig_r16_codebookType_type2_subType_typeII_r16_n1_n2_codebookSubsetRestriction_r16_Choice_eight_one:
		var tmp_bs_eight_one []byte
		var n_eight_one uint
		if tmp_bs_eight_one, n_eight_one, err = r.ReadBitString(&uper.Constraint{Lb: 64, Ub: 64}, false); err != nil {
			return utils.WrapError("Decode eight_one", err)
		}
		ie.eight_one = uper.BitString{
			Bytes:   tmp_bs_eight_one,
			NumBits: uint64(n_eight_one),
		}
	case CodebookConfig_r16_codebookType_type2_subType_typeII_r16_n1_n2_codebookSubsetRestriction_r16_Choice_four_three:
		var tmp_bs_four_three []byte
		var n_four_three uint
		if tmp_bs_four_three, n_four_three, err = r.ReadBitString(&uper.Constraint{Lb: 107, Ub: 107}, false); err != nil {
			return utils.WrapError("Decode four_three", err)
		}
		ie.four_three = uper.BitString{
			Bytes:   tmp_bs_four_three,
			NumBits: uint64(n_four_three),
		}
	case CodebookConfig_r16_codebookType_type2_subType_typeII_r16_n1_n2_codebookSubsetRestriction_r16_Choice_six_two:
		var tmp_bs_six_two []byte
		var n_six_two uint
		if tmp_bs_six_two, n_six_two, err = r.ReadBitString(&uper.Constraint{Lb: 107, Ub: 107}, false); err != nil {
			return utils.WrapError("Decode six_two", err)
		}
		ie.six_two = uper.BitString{
			Bytes:   tmp_bs_six_two,
			NumBits: uint64(n_six_two),
		}
	case CodebookConfig_r16_codebookType_type2_subType_typeII_r16_n1_n2_codebookSubsetRestriction_r16_Choice_twelve_one:
		var tmp_bs_twelve_one []byte
		var n_twelve_one uint
		if tmp_bs_twelve_one, n_twelve_one, err = r.ReadBitString(&uper.Constraint{Lb: 96, Ub: 96}, false); err != nil {
			return utils.WrapError("Decode twelve_one", err)
		}
		ie.twelve_one = uper.BitString{
			Bytes:   tmp_bs_twelve_one,
			NumBits: uint64(n_twelve_one),
		}
	case CodebookConfig_r16_codebookType_type2_subType_typeII_r16_n1_n2_codebookSubsetRestriction_r16_Choice_four_four:
		var tmp_bs_four_four []byte
		var n_four_four uint
		if tmp_bs_four_four, n_four_four, err = r.ReadBitString(&uper.Constraint{Lb: 139, Ub: 139}, false); err != nil {
			return utils.WrapError("Decode four_four", err)
		}
		ie.four_four = uper.BitString{
			Bytes:   tmp_bs_four_four,
			NumBits: uint64(n_four_four),
		}
	case CodebookConfig_r16_codebookType_type2_subType_typeII_r16_n1_n2_codebookSubsetRestriction_r16_Choice_eight_two:
		var tmp_bs_eight_two []byte
		var n_eight_two uint
		if tmp_bs_eight_two, n_eight_two, err = r.ReadBitString(&uper.Constraint{Lb: 139, Ub: 139}, false); err != nil {
			return utils.WrapError("Decode eight_two", err)
		}
		ie.eight_two = uper.BitString{
			Bytes:   tmp_bs_eight_two,
			NumBits: uint64(n_eight_two),
		}
	case CodebookConfig_r16_codebookType_type2_subType_typeII_r16_n1_n2_codebookSubsetRestriction_r16_Choice_sixteen_one:
		var tmp_bs_sixteen_one []byte
		var n_sixteen_one uint
		if tmp_bs_sixteen_one, n_sixteen_one, err = r.ReadBitString(&uper.Constraint{Lb: 128, Ub: 128}, false); err != nil {
			return utils.WrapError("Decode sixteen_one", err)
		}
		ie.sixteen_one = uper.BitString{
			Bytes:   tmp_bs_sixteen_one,
			NumBits: uint64(n_sixteen_one),
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
