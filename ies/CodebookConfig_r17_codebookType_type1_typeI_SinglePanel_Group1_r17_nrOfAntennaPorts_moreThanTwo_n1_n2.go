package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group1_r17_nrOfAntennaPorts_moreThanTwo_n1_n2_Choice_nothing uint64 = iota
	CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group1_r17_nrOfAntennaPorts_moreThanTwo_n1_n2_Choice_two_one_TypeI_SinglePanel_Restriction1_r17
	CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group1_r17_nrOfAntennaPorts_moreThanTwo_n1_n2_Choice_two_two_TypeI_SinglePanel_Restriction1_r17
	CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group1_r17_nrOfAntennaPorts_moreThanTwo_n1_n2_Choice_four_one_TypeI_SinglePanel_Restriction1_r17
	CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group1_r17_nrOfAntennaPorts_moreThanTwo_n1_n2_Choice_three_two_TypeI_SinglePanel_Restriction1_r17
	CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group1_r17_nrOfAntennaPorts_moreThanTwo_n1_n2_Choice_six_one_TypeI_SinglePanel_Restriction1_r17
	CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group1_r17_nrOfAntennaPorts_moreThanTwo_n1_n2_Choice_four_two_TypeI_SinglePanel_Restriction1_r17
	CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group1_r17_nrOfAntennaPorts_moreThanTwo_n1_n2_Choice_eight_one_TypeI_SinglePanel_Restriction1_r17
	CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group1_r17_nrOfAntennaPorts_moreThanTwo_n1_n2_Choice_four_three_TypeI_SinglePanel_Restriction1_r17
	CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group1_r17_nrOfAntennaPorts_moreThanTwo_n1_n2_Choice_six_two_TypeI_SinglePanel_Restriction1_r17
	CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group1_r17_nrOfAntennaPorts_moreThanTwo_n1_n2_Choice_twelve_one_TypeI_SinglePanel_Restriction1_r17
	CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group1_r17_nrOfAntennaPorts_moreThanTwo_n1_n2_Choice_four_four_TypeI_SinglePanel_Restriction1_r17
	CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group1_r17_nrOfAntennaPorts_moreThanTwo_n1_n2_Choice_eight_two_TypeI_SinglePanel_Restriction1_r17
	CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group1_r17_nrOfAntennaPorts_moreThanTwo_n1_n2_Choice_sixteen_one_TypeI_SinglePanel_Restriction1_r17
)

type CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group1_r17_nrOfAntennaPorts_moreThanTwo_n1_n2 struct {
	Choice                                         uint64
	two_one_TypeI_SinglePanel_Restriction1_r17     uper.BitString `lb:8,ub:8,madatory`
	two_two_TypeI_SinglePanel_Restriction1_r17     uper.BitString `lb:64,ub:64,madatory`
	four_one_TypeI_SinglePanel_Restriction1_r17    uper.BitString `lb:16,ub:16,madatory`
	three_two_TypeI_SinglePanel_Restriction1_r17   uper.BitString `lb:96,ub:96,madatory`
	six_one_TypeI_SinglePanel_Restriction1_r17     uper.BitString `lb:24,ub:24,madatory`
	four_two_TypeI_SinglePanel_Restriction1_r17    uper.BitString `lb:128,ub:128,madatory`
	eight_one_TypeI_SinglePanel_Restriction1_r17   uper.BitString `lb:32,ub:32,madatory`
	four_three_TypeI_SinglePanel_Restriction1_r17  uper.BitString `lb:192,ub:192,madatory`
	six_two_TypeI_SinglePanel_Restriction1_r17     uper.BitString `lb:192,ub:192,madatory`
	twelve_one_TypeI_SinglePanel_Restriction1_r17  uper.BitString `lb:48,ub:48,madatory`
	four_four_TypeI_SinglePanel_Restriction1_r17   uper.BitString `lb:256,ub:256,madatory`
	eight_two_TypeI_SinglePanel_Restriction1_r17   uper.BitString `lb:256,ub:256,madatory`
	sixteen_one_TypeI_SinglePanel_Restriction1_r17 uper.BitString `lb:64,ub:64,madatory`
}

func (ie *CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group1_r17_nrOfAntennaPorts_moreThanTwo_n1_n2) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 13, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group1_r17_nrOfAntennaPorts_moreThanTwo_n1_n2_Choice_two_one_TypeI_SinglePanel_Restriction1_r17:
		if err = w.WriteBitString(ie.two_one_TypeI_SinglePanel_Restriction1_r17.Bytes, uint(ie.two_one_TypeI_SinglePanel_Restriction1_r17.NumBits), &uper.Constraint{Lb: 8, Ub: 8}, false); err != nil {
			err = utils.WrapError("Encode two_one_TypeI_SinglePanel_Restriction1_r17", err)
		}
	case CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group1_r17_nrOfAntennaPorts_moreThanTwo_n1_n2_Choice_two_two_TypeI_SinglePanel_Restriction1_r17:
		if err = w.WriteBitString(ie.two_two_TypeI_SinglePanel_Restriction1_r17.Bytes, uint(ie.two_two_TypeI_SinglePanel_Restriction1_r17.NumBits), &uper.Constraint{Lb: 64, Ub: 64}, false); err != nil {
			err = utils.WrapError("Encode two_two_TypeI_SinglePanel_Restriction1_r17", err)
		}
	case CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group1_r17_nrOfAntennaPorts_moreThanTwo_n1_n2_Choice_four_one_TypeI_SinglePanel_Restriction1_r17:
		if err = w.WriteBitString(ie.four_one_TypeI_SinglePanel_Restriction1_r17.Bytes, uint(ie.four_one_TypeI_SinglePanel_Restriction1_r17.NumBits), &uper.Constraint{Lb: 16, Ub: 16}, false); err != nil {
			err = utils.WrapError("Encode four_one_TypeI_SinglePanel_Restriction1_r17", err)
		}
	case CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group1_r17_nrOfAntennaPorts_moreThanTwo_n1_n2_Choice_three_two_TypeI_SinglePanel_Restriction1_r17:
		if err = w.WriteBitString(ie.three_two_TypeI_SinglePanel_Restriction1_r17.Bytes, uint(ie.three_two_TypeI_SinglePanel_Restriction1_r17.NumBits), &uper.Constraint{Lb: 96, Ub: 96}, false); err != nil {
			err = utils.WrapError("Encode three_two_TypeI_SinglePanel_Restriction1_r17", err)
		}
	case CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group1_r17_nrOfAntennaPorts_moreThanTwo_n1_n2_Choice_six_one_TypeI_SinglePanel_Restriction1_r17:
		if err = w.WriteBitString(ie.six_one_TypeI_SinglePanel_Restriction1_r17.Bytes, uint(ie.six_one_TypeI_SinglePanel_Restriction1_r17.NumBits), &uper.Constraint{Lb: 24, Ub: 24}, false); err != nil {
			err = utils.WrapError("Encode six_one_TypeI_SinglePanel_Restriction1_r17", err)
		}
	case CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group1_r17_nrOfAntennaPorts_moreThanTwo_n1_n2_Choice_four_two_TypeI_SinglePanel_Restriction1_r17:
		if err = w.WriteBitString(ie.four_two_TypeI_SinglePanel_Restriction1_r17.Bytes, uint(ie.four_two_TypeI_SinglePanel_Restriction1_r17.NumBits), &uper.Constraint{Lb: 128, Ub: 128}, false); err != nil {
			err = utils.WrapError("Encode four_two_TypeI_SinglePanel_Restriction1_r17", err)
		}
	case CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group1_r17_nrOfAntennaPorts_moreThanTwo_n1_n2_Choice_eight_one_TypeI_SinglePanel_Restriction1_r17:
		if err = w.WriteBitString(ie.eight_one_TypeI_SinglePanel_Restriction1_r17.Bytes, uint(ie.eight_one_TypeI_SinglePanel_Restriction1_r17.NumBits), &uper.Constraint{Lb: 32, Ub: 32}, false); err != nil {
			err = utils.WrapError("Encode eight_one_TypeI_SinglePanel_Restriction1_r17", err)
		}
	case CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group1_r17_nrOfAntennaPorts_moreThanTwo_n1_n2_Choice_four_three_TypeI_SinglePanel_Restriction1_r17:
		if err = w.WriteBitString(ie.four_three_TypeI_SinglePanel_Restriction1_r17.Bytes, uint(ie.four_three_TypeI_SinglePanel_Restriction1_r17.NumBits), &uper.Constraint{Lb: 192, Ub: 192}, false); err != nil {
			err = utils.WrapError("Encode four_three_TypeI_SinglePanel_Restriction1_r17", err)
		}
	case CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group1_r17_nrOfAntennaPorts_moreThanTwo_n1_n2_Choice_six_two_TypeI_SinglePanel_Restriction1_r17:
		if err = w.WriteBitString(ie.six_two_TypeI_SinglePanel_Restriction1_r17.Bytes, uint(ie.six_two_TypeI_SinglePanel_Restriction1_r17.NumBits), &uper.Constraint{Lb: 192, Ub: 192}, false); err != nil {
			err = utils.WrapError("Encode six_two_TypeI_SinglePanel_Restriction1_r17", err)
		}
	case CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group1_r17_nrOfAntennaPorts_moreThanTwo_n1_n2_Choice_twelve_one_TypeI_SinglePanel_Restriction1_r17:
		if err = w.WriteBitString(ie.twelve_one_TypeI_SinglePanel_Restriction1_r17.Bytes, uint(ie.twelve_one_TypeI_SinglePanel_Restriction1_r17.NumBits), &uper.Constraint{Lb: 48, Ub: 48}, false); err != nil {
			err = utils.WrapError("Encode twelve_one_TypeI_SinglePanel_Restriction1_r17", err)
		}
	case CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group1_r17_nrOfAntennaPorts_moreThanTwo_n1_n2_Choice_four_four_TypeI_SinglePanel_Restriction1_r17:
		if err = w.WriteBitString(ie.four_four_TypeI_SinglePanel_Restriction1_r17.Bytes, uint(ie.four_four_TypeI_SinglePanel_Restriction1_r17.NumBits), &uper.Constraint{Lb: 256, Ub: 256}, false); err != nil {
			err = utils.WrapError("Encode four_four_TypeI_SinglePanel_Restriction1_r17", err)
		}
	case CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group1_r17_nrOfAntennaPorts_moreThanTwo_n1_n2_Choice_eight_two_TypeI_SinglePanel_Restriction1_r17:
		if err = w.WriteBitString(ie.eight_two_TypeI_SinglePanel_Restriction1_r17.Bytes, uint(ie.eight_two_TypeI_SinglePanel_Restriction1_r17.NumBits), &uper.Constraint{Lb: 256, Ub: 256}, false); err != nil {
			err = utils.WrapError("Encode eight_two_TypeI_SinglePanel_Restriction1_r17", err)
		}
	case CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group1_r17_nrOfAntennaPorts_moreThanTwo_n1_n2_Choice_sixteen_one_TypeI_SinglePanel_Restriction1_r17:
		if err = w.WriteBitString(ie.sixteen_one_TypeI_SinglePanel_Restriction1_r17.Bytes, uint(ie.sixteen_one_TypeI_SinglePanel_Restriction1_r17.NumBits), &uper.Constraint{Lb: 64, Ub: 64}, false); err != nil {
			err = utils.WrapError("Encode sixteen_one_TypeI_SinglePanel_Restriction1_r17", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group1_r17_nrOfAntennaPorts_moreThanTwo_n1_n2) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(13, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group1_r17_nrOfAntennaPorts_moreThanTwo_n1_n2_Choice_two_one_TypeI_SinglePanel_Restriction1_r17:
		var tmp_bs_two_one_TypeI_SinglePanel_Restriction1_r17 []byte
		var n_two_one_TypeI_SinglePanel_Restriction1_r17 uint
		if tmp_bs_two_one_TypeI_SinglePanel_Restriction1_r17, n_two_one_TypeI_SinglePanel_Restriction1_r17, err = r.ReadBitString(&uper.Constraint{Lb: 8, Ub: 8}, false); err != nil {
			return utils.WrapError("Decode two_one_TypeI_SinglePanel_Restriction1_r17", err)
		}
		ie.two_one_TypeI_SinglePanel_Restriction1_r17 = uper.BitString{
			Bytes:   tmp_bs_two_one_TypeI_SinglePanel_Restriction1_r17,
			NumBits: uint64(n_two_one_TypeI_SinglePanel_Restriction1_r17),
		}
	case CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group1_r17_nrOfAntennaPorts_moreThanTwo_n1_n2_Choice_two_two_TypeI_SinglePanel_Restriction1_r17:
		var tmp_bs_two_two_TypeI_SinglePanel_Restriction1_r17 []byte
		var n_two_two_TypeI_SinglePanel_Restriction1_r17 uint
		if tmp_bs_two_two_TypeI_SinglePanel_Restriction1_r17, n_two_two_TypeI_SinglePanel_Restriction1_r17, err = r.ReadBitString(&uper.Constraint{Lb: 64, Ub: 64}, false); err != nil {
			return utils.WrapError("Decode two_two_TypeI_SinglePanel_Restriction1_r17", err)
		}
		ie.two_two_TypeI_SinglePanel_Restriction1_r17 = uper.BitString{
			Bytes:   tmp_bs_two_two_TypeI_SinglePanel_Restriction1_r17,
			NumBits: uint64(n_two_two_TypeI_SinglePanel_Restriction1_r17),
		}
	case CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group1_r17_nrOfAntennaPorts_moreThanTwo_n1_n2_Choice_four_one_TypeI_SinglePanel_Restriction1_r17:
		var tmp_bs_four_one_TypeI_SinglePanel_Restriction1_r17 []byte
		var n_four_one_TypeI_SinglePanel_Restriction1_r17 uint
		if tmp_bs_four_one_TypeI_SinglePanel_Restriction1_r17, n_four_one_TypeI_SinglePanel_Restriction1_r17, err = r.ReadBitString(&uper.Constraint{Lb: 16, Ub: 16}, false); err != nil {
			return utils.WrapError("Decode four_one_TypeI_SinglePanel_Restriction1_r17", err)
		}
		ie.four_one_TypeI_SinglePanel_Restriction1_r17 = uper.BitString{
			Bytes:   tmp_bs_four_one_TypeI_SinglePanel_Restriction1_r17,
			NumBits: uint64(n_four_one_TypeI_SinglePanel_Restriction1_r17),
		}
	case CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group1_r17_nrOfAntennaPorts_moreThanTwo_n1_n2_Choice_three_two_TypeI_SinglePanel_Restriction1_r17:
		var tmp_bs_three_two_TypeI_SinglePanel_Restriction1_r17 []byte
		var n_three_two_TypeI_SinglePanel_Restriction1_r17 uint
		if tmp_bs_three_two_TypeI_SinglePanel_Restriction1_r17, n_three_two_TypeI_SinglePanel_Restriction1_r17, err = r.ReadBitString(&uper.Constraint{Lb: 96, Ub: 96}, false); err != nil {
			return utils.WrapError("Decode three_two_TypeI_SinglePanel_Restriction1_r17", err)
		}
		ie.three_two_TypeI_SinglePanel_Restriction1_r17 = uper.BitString{
			Bytes:   tmp_bs_three_two_TypeI_SinglePanel_Restriction1_r17,
			NumBits: uint64(n_three_two_TypeI_SinglePanel_Restriction1_r17),
		}
	case CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group1_r17_nrOfAntennaPorts_moreThanTwo_n1_n2_Choice_six_one_TypeI_SinglePanel_Restriction1_r17:
		var tmp_bs_six_one_TypeI_SinglePanel_Restriction1_r17 []byte
		var n_six_one_TypeI_SinglePanel_Restriction1_r17 uint
		if tmp_bs_six_one_TypeI_SinglePanel_Restriction1_r17, n_six_one_TypeI_SinglePanel_Restriction1_r17, err = r.ReadBitString(&uper.Constraint{Lb: 24, Ub: 24}, false); err != nil {
			return utils.WrapError("Decode six_one_TypeI_SinglePanel_Restriction1_r17", err)
		}
		ie.six_one_TypeI_SinglePanel_Restriction1_r17 = uper.BitString{
			Bytes:   tmp_bs_six_one_TypeI_SinglePanel_Restriction1_r17,
			NumBits: uint64(n_six_one_TypeI_SinglePanel_Restriction1_r17),
		}
	case CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group1_r17_nrOfAntennaPorts_moreThanTwo_n1_n2_Choice_four_two_TypeI_SinglePanel_Restriction1_r17:
		var tmp_bs_four_two_TypeI_SinglePanel_Restriction1_r17 []byte
		var n_four_two_TypeI_SinglePanel_Restriction1_r17 uint
		if tmp_bs_four_two_TypeI_SinglePanel_Restriction1_r17, n_four_two_TypeI_SinglePanel_Restriction1_r17, err = r.ReadBitString(&uper.Constraint{Lb: 128, Ub: 128}, false); err != nil {
			return utils.WrapError("Decode four_two_TypeI_SinglePanel_Restriction1_r17", err)
		}
		ie.four_two_TypeI_SinglePanel_Restriction1_r17 = uper.BitString{
			Bytes:   tmp_bs_four_two_TypeI_SinglePanel_Restriction1_r17,
			NumBits: uint64(n_four_two_TypeI_SinglePanel_Restriction1_r17),
		}
	case CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group1_r17_nrOfAntennaPorts_moreThanTwo_n1_n2_Choice_eight_one_TypeI_SinglePanel_Restriction1_r17:
		var tmp_bs_eight_one_TypeI_SinglePanel_Restriction1_r17 []byte
		var n_eight_one_TypeI_SinglePanel_Restriction1_r17 uint
		if tmp_bs_eight_one_TypeI_SinglePanel_Restriction1_r17, n_eight_one_TypeI_SinglePanel_Restriction1_r17, err = r.ReadBitString(&uper.Constraint{Lb: 32, Ub: 32}, false); err != nil {
			return utils.WrapError("Decode eight_one_TypeI_SinglePanel_Restriction1_r17", err)
		}
		ie.eight_one_TypeI_SinglePanel_Restriction1_r17 = uper.BitString{
			Bytes:   tmp_bs_eight_one_TypeI_SinglePanel_Restriction1_r17,
			NumBits: uint64(n_eight_one_TypeI_SinglePanel_Restriction1_r17),
		}
	case CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group1_r17_nrOfAntennaPorts_moreThanTwo_n1_n2_Choice_four_three_TypeI_SinglePanel_Restriction1_r17:
		var tmp_bs_four_three_TypeI_SinglePanel_Restriction1_r17 []byte
		var n_four_three_TypeI_SinglePanel_Restriction1_r17 uint
		if tmp_bs_four_three_TypeI_SinglePanel_Restriction1_r17, n_four_three_TypeI_SinglePanel_Restriction1_r17, err = r.ReadBitString(&uper.Constraint{Lb: 192, Ub: 192}, false); err != nil {
			return utils.WrapError("Decode four_three_TypeI_SinglePanel_Restriction1_r17", err)
		}
		ie.four_three_TypeI_SinglePanel_Restriction1_r17 = uper.BitString{
			Bytes:   tmp_bs_four_three_TypeI_SinglePanel_Restriction1_r17,
			NumBits: uint64(n_four_three_TypeI_SinglePanel_Restriction1_r17),
		}
	case CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group1_r17_nrOfAntennaPorts_moreThanTwo_n1_n2_Choice_six_two_TypeI_SinglePanel_Restriction1_r17:
		var tmp_bs_six_two_TypeI_SinglePanel_Restriction1_r17 []byte
		var n_six_two_TypeI_SinglePanel_Restriction1_r17 uint
		if tmp_bs_six_two_TypeI_SinglePanel_Restriction1_r17, n_six_two_TypeI_SinglePanel_Restriction1_r17, err = r.ReadBitString(&uper.Constraint{Lb: 192, Ub: 192}, false); err != nil {
			return utils.WrapError("Decode six_two_TypeI_SinglePanel_Restriction1_r17", err)
		}
		ie.six_two_TypeI_SinglePanel_Restriction1_r17 = uper.BitString{
			Bytes:   tmp_bs_six_two_TypeI_SinglePanel_Restriction1_r17,
			NumBits: uint64(n_six_two_TypeI_SinglePanel_Restriction1_r17),
		}
	case CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group1_r17_nrOfAntennaPorts_moreThanTwo_n1_n2_Choice_twelve_one_TypeI_SinglePanel_Restriction1_r17:
		var tmp_bs_twelve_one_TypeI_SinglePanel_Restriction1_r17 []byte
		var n_twelve_one_TypeI_SinglePanel_Restriction1_r17 uint
		if tmp_bs_twelve_one_TypeI_SinglePanel_Restriction1_r17, n_twelve_one_TypeI_SinglePanel_Restriction1_r17, err = r.ReadBitString(&uper.Constraint{Lb: 48, Ub: 48}, false); err != nil {
			return utils.WrapError("Decode twelve_one_TypeI_SinglePanel_Restriction1_r17", err)
		}
		ie.twelve_one_TypeI_SinglePanel_Restriction1_r17 = uper.BitString{
			Bytes:   tmp_bs_twelve_one_TypeI_SinglePanel_Restriction1_r17,
			NumBits: uint64(n_twelve_one_TypeI_SinglePanel_Restriction1_r17),
		}
	case CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group1_r17_nrOfAntennaPorts_moreThanTwo_n1_n2_Choice_four_four_TypeI_SinglePanel_Restriction1_r17:
		var tmp_bs_four_four_TypeI_SinglePanel_Restriction1_r17 []byte
		var n_four_four_TypeI_SinglePanel_Restriction1_r17 uint
		if tmp_bs_four_four_TypeI_SinglePanel_Restriction1_r17, n_four_four_TypeI_SinglePanel_Restriction1_r17, err = r.ReadBitString(&uper.Constraint{Lb: 256, Ub: 256}, false); err != nil {
			return utils.WrapError("Decode four_four_TypeI_SinglePanel_Restriction1_r17", err)
		}
		ie.four_four_TypeI_SinglePanel_Restriction1_r17 = uper.BitString{
			Bytes:   tmp_bs_four_four_TypeI_SinglePanel_Restriction1_r17,
			NumBits: uint64(n_four_four_TypeI_SinglePanel_Restriction1_r17),
		}
	case CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group1_r17_nrOfAntennaPorts_moreThanTwo_n1_n2_Choice_eight_two_TypeI_SinglePanel_Restriction1_r17:
		var tmp_bs_eight_two_TypeI_SinglePanel_Restriction1_r17 []byte
		var n_eight_two_TypeI_SinglePanel_Restriction1_r17 uint
		if tmp_bs_eight_two_TypeI_SinglePanel_Restriction1_r17, n_eight_two_TypeI_SinglePanel_Restriction1_r17, err = r.ReadBitString(&uper.Constraint{Lb: 256, Ub: 256}, false); err != nil {
			return utils.WrapError("Decode eight_two_TypeI_SinglePanel_Restriction1_r17", err)
		}
		ie.eight_two_TypeI_SinglePanel_Restriction1_r17 = uper.BitString{
			Bytes:   tmp_bs_eight_two_TypeI_SinglePanel_Restriction1_r17,
			NumBits: uint64(n_eight_two_TypeI_SinglePanel_Restriction1_r17),
		}
	case CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group1_r17_nrOfAntennaPorts_moreThanTwo_n1_n2_Choice_sixteen_one_TypeI_SinglePanel_Restriction1_r17:
		var tmp_bs_sixteen_one_TypeI_SinglePanel_Restriction1_r17 []byte
		var n_sixteen_one_TypeI_SinglePanel_Restriction1_r17 uint
		if tmp_bs_sixteen_one_TypeI_SinglePanel_Restriction1_r17, n_sixteen_one_TypeI_SinglePanel_Restriction1_r17, err = r.ReadBitString(&uper.Constraint{Lb: 64, Ub: 64}, false); err != nil {
			return utils.WrapError("Decode sixteen_one_TypeI_SinglePanel_Restriction1_r17", err)
		}
		ie.sixteen_one_TypeI_SinglePanel_Restriction1_r17 = uper.BitString{
			Bytes:   tmp_bs_sixteen_one_TypeI_SinglePanel_Restriction1_r17,
			NumBits: uint64(n_sixteen_one_TypeI_SinglePanel_Restriction1_r17),
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
