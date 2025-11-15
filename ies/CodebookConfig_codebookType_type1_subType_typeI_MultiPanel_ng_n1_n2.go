package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	CodebookConfig_codebookType_type1_subType_typeI_MultiPanel_ng_n1_n2_Choice_nothing uint64 = iota
	CodebookConfig_codebookType_type1_subType_typeI_MultiPanel_ng_n1_n2_Choice_two_two_one_TypeI_MultiPanel_Restriction
	CodebookConfig_codebookType_type1_subType_typeI_MultiPanel_ng_n1_n2_Choice_two_four_one_TypeI_MultiPanel_Restriction
	CodebookConfig_codebookType_type1_subType_typeI_MultiPanel_ng_n1_n2_Choice_four_two_one_TypeI_MultiPanel_Restriction
	CodebookConfig_codebookType_type1_subType_typeI_MultiPanel_ng_n1_n2_Choice_two_two_two_TypeI_MultiPanel_Restriction
	CodebookConfig_codebookType_type1_subType_typeI_MultiPanel_ng_n1_n2_Choice_two_eight_one_TypeI_MultiPanel_Restriction
	CodebookConfig_codebookType_type1_subType_typeI_MultiPanel_ng_n1_n2_Choice_four_four_one_TypeI_MultiPanel_Restriction
	CodebookConfig_codebookType_type1_subType_typeI_MultiPanel_ng_n1_n2_Choice_two_four_two_TypeI_MultiPanel_Restriction
	CodebookConfig_codebookType_type1_subType_typeI_MultiPanel_ng_n1_n2_Choice_four_two_two_TypeI_MultiPanel_Restriction
)

type CodebookConfig_codebookType_type1_subType_typeI_MultiPanel_ng_n1_n2 struct {
	Choice                                     uint64
	two_two_one_TypeI_MultiPanel_Restriction   uper.BitString `lb:8,ub:8,madatory`
	two_four_one_TypeI_MultiPanel_Restriction  uper.BitString `lb:16,ub:16,madatory`
	four_two_one_TypeI_MultiPanel_Restriction  uper.BitString `lb:8,ub:8,madatory`
	two_two_two_TypeI_MultiPanel_Restriction   uper.BitString `lb:64,ub:64,madatory`
	two_eight_one_TypeI_MultiPanel_Restriction uper.BitString `lb:32,ub:32,madatory`
	four_four_one_TypeI_MultiPanel_Restriction uper.BitString `lb:16,ub:16,madatory`
	two_four_two_TypeI_MultiPanel_Restriction  uper.BitString `lb:128,ub:128,madatory`
	four_two_two_TypeI_MultiPanel_Restriction  uper.BitString `lb:64,ub:64,madatory`
}

func (ie *CodebookConfig_codebookType_type1_subType_typeI_MultiPanel_ng_n1_n2) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 8, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CodebookConfig_codebookType_type1_subType_typeI_MultiPanel_ng_n1_n2_Choice_two_two_one_TypeI_MultiPanel_Restriction:
		if err = w.WriteBitString(ie.two_two_one_TypeI_MultiPanel_Restriction.Bytes, uint(ie.two_two_one_TypeI_MultiPanel_Restriction.NumBits), &uper.Constraint{Lb: 8, Ub: 8}, false); err != nil {
			err = utils.WrapError("Encode two_two_one_TypeI_MultiPanel_Restriction", err)
		}
	case CodebookConfig_codebookType_type1_subType_typeI_MultiPanel_ng_n1_n2_Choice_two_four_one_TypeI_MultiPanel_Restriction:
		if err = w.WriteBitString(ie.two_four_one_TypeI_MultiPanel_Restriction.Bytes, uint(ie.two_four_one_TypeI_MultiPanel_Restriction.NumBits), &uper.Constraint{Lb: 16, Ub: 16}, false); err != nil {
			err = utils.WrapError("Encode two_four_one_TypeI_MultiPanel_Restriction", err)
		}
	case CodebookConfig_codebookType_type1_subType_typeI_MultiPanel_ng_n1_n2_Choice_four_two_one_TypeI_MultiPanel_Restriction:
		if err = w.WriteBitString(ie.four_two_one_TypeI_MultiPanel_Restriction.Bytes, uint(ie.four_two_one_TypeI_MultiPanel_Restriction.NumBits), &uper.Constraint{Lb: 8, Ub: 8}, false); err != nil {
			err = utils.WrapError("Encode four_two_one_TypeI_MultiPanel_Restriction", err)
		}
	case CodebookConfig_codebookType_type1_subType_typeI_MultiPanel_ng_n1_n2_Choice_two_two_two_TypeI_MultiPanel_Restriction:
		if err = w.WriteBitString(ie.two_two_two_TypeI_MultiPanel_Restriction.Bytes, uint(ie.two_two_two_TypeI_MultiPanel_Restriction.NumBits), &uper.Constraint{Lb: 64, Ub: 64}, false); err != nil {
			err = utils.WrapError("Encode two_two_two_TypeI_MultiPanel_Restriction", err)
		}
	case CodebookConfig_codebookType_type1_subType_typeI_MultiPanel_ng_n1_n2_Choice_two_eight_one_TypeI_MultiPanel_Restriction:
		if err = w.WriteBitString(ie.two_eight_one_TypeI_MultiPanel_Restriction.Bytes, uint(ie.two_eight_one_TypeI_MultiPanel_Restriction.NumBits), &uper.Constraint{Lb: 32, Ub: 32}, false); err != nil {
			err = utils.WrapError("Encode two_eight_one_TypeI_MultiPanel_Restriction", err)
		}
	case CodebookConfig_codebookType_type1_subType_typeI_MultiPanel_ng_n1_n2_Choice_four_four_one_TypeI_MultiPanel_Restriction:
		if err = w.WriteBitString(ie.four_four_one_TypeI_MultiPanel_Restriction.Bytes, uint(ie.four_four_one_TypeI_MultiPanel_Restriction.NumBits), &uper.Constraint{Lb: 16, Ub: 16}, false); err != nil {
			err = utils.WrapError("Encode four_four_one_TypeI_MultiPanel_Restriction", err)
		}
	case CodebookConfig_codebookType_type1_subType_typeI_MultiPanel_ng_n1_n2_Choice_two_four_two_TypeI_MultiPanel_Restriction:
		if err = w.WriteBitString(ie.two_four_two_TypeI_MultiPanel_Restriction.Bytes, uint(ie.two_four_two_TypeI_MultiPanel_Restriction.NumBits), &uper.Constraint{Lb: 128, Ub: 128}, false); err != nil {
			err = utils.WrapError("Encode two_four_two_TypeI_MultiPanel_Restriction", err)
		}
	case CodebookConfig_codebookType_type1_subType_typeI_MultiPanel_ng_n1_n2_Choice_four_two_two_TypeI_MultiPanel_Restriction:
		if err = w.WriteBitString(ie.four_two_two_TypeI_MultiPanel_Restriction.Bytes, uint(ie.four_two_two_TypeI_MultiPanel_Restriction.NumBits), &uper.Constraint{Lb: 64, Ub: 64}, false); err != nil {
			err = utils.WrapError("Encode four_two_two_TypeI_MultiPanel_Restriction", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *CodebookConfig_codebookType_type1_subType_typeI_MultiPanel_ng_n1_n2) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(8, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CodebookConfig_codebookType_type1_subType_typeI_MultiPanel_ng_n1_n2_Choice_two_two_one_TypeI_MultiPanel_Restriction:
		var tmp_bs_two_two_one_TypeI_MultiPanel_Restriction []byte
		var n_two_two_one_TypeI_MultiPanel_Restriction uint
		if tmp_bs_two_two_one_TypeI_MultiPanel_Restriction, n_two_two_one_TypeI_MultiPanel_Restriction, err = r.ReadBitString(&uper.Constraint{Lb: 8, Ub: 8}, false); err != nil {
			return utils.WrapError("Decode two_two_one_TypeI_MultiPanel_Restriction", err)
		}
		ie.two_two_one_TypeI_MultiPanel_Restriction = uper.BitString{
			Bytes:   tmp_bs_two_two_one_TypeI_MultiPanel_Restriction,
			NumBits: uint64(n_two_two_one_TypeI_MultiPanel_Restriction),
		}
	case CodebookConfig_codebookType_type1_subType_typeI_MultiPanel_ng_n1_n2_Choice_two_four_one_TypeI_MultiPanel_Restriction:
		var tmp_bs_two_four_one_TypeI_MultiPanel_Restriction []byte
		var n_two_four_one_TypeI_MultiPanel_Restriction uint
		if tmp_bs_two_four_one_TypeI_MultiPanel_Restriction, n_two_four_one_TypeI_MultiPanel_Restriction, err = r.ReadBitString(&uper.Constraint{Lb: 16, Ub: 16}, false); err != nil {
			return utils.WrapError("Decode two_four_one_TypeI_MultiPanel_Restriction", err)
		}
		ie.two_four_one_TypeI_MultiPanel_Restriction = uper.BitString{
			Bytes:   tmp_bs_two_four_one_TypeI_MultiPanel_Restriction,
			NumBits: uint64(n_two_four_one_TypeI_MultiPanel_Restriction),
		}
	case CodebookConfig_codebookType_type1_subType_typeI_MultiPanel_ng_n1_n2_Choice_four_two_one_TypeI_MultiPanel_Restriction:
		var tmp_bs_four_two_one_TypeI_MultiPanel_Restriction []byte
		var n_four_two_one_TypeI_MultiPanel_Restriction uint
		if tmp_bs_four_two_one_TypeI_MultiPanel_Restriction, n_four_two_one_TypeI_MultiPanel_Restriction, err = r.ReadBitString(&uper.Constraint{Lb: 8, Ub: 8}, false); err != nil {
			return utils.WrapError("Decode four_two_one_TypeI_MultiPanel_Restriction", err)
		}
		ie.four_two_one_TypeI_MultiPanel_Restriction = uper.BitString{
			Bytes:   tmp_bs_four_two_one_TypeI_MultiPanel_Restriction,
			NumBits: uint64(n_four_two_one_TypeI_MultiPanel_Restriction),
		}
	case CodebookConfig_codebookType_type1_subType_typeI_MultiPanel_ng_n1_n2_Choice_two_two_two_TypeI_MultiPanel_Restriction:
		var tmp_bs_two_two_two_TypeI_MultiPanel_Restriction []byte
		var n_two_two_two_TypeI_MultiPanel_Restriction uint
		if tmp_bs_two_two_two_TypeI_MultiPanel_Restriction, n_two_two_two_TypeI_MultiPanel_Restriction, err = r.ReadBitString(&uper.Constraint{Lb: 64, Ub: 64}, false); err != nil {
			return utils.WrapError("Decode two_two_two_TypeI_MultiPanel_Restriction", err)
		}
		ie.two_two_two_TypeI_MultiPanel_Restriction = uper.BitString{
			Bytes:   tmp_bs_two_two_two_TypeI_MultiPanel_Restriction,
			NumBits: uint64(n_two_two_two_TypeI_MultiPanel_Restriction),
		}
	case CodebookConfig_codebookType_type1_subType_typeI_MultiPanel_ng_n1_n2_Choice_two_eight_one_TypeI_MultiPanel_Restriction:
		var tmp_bs_two_eight_one_TypeI_MultiPanel_Restriction []byte
		var n_two_eight_one_TypeI_MultiPanel_Restriction uint
		if tmp_bs_two_eight_one_TypeI_MultiPanel_Restriction, n_two_eight_one_TypeI_MultiPanel_Restriction, err = r.ReadBitString(&uper.Constraint{Lb: 32, Ub: 32}, false); err != nil {
			return utils.WrapError("Decode two_eight_one_TypeI_MultiPanel_Restriction", err)
		}
		ie.two_eight_one_TypeI_MultiPanel_Restriction = uper.BitString{
			Bytes:   tmp_bs_two_eight_one_TypeI_MultiPanel_Restriction,
			NumBits: uint64(n_two_eight_one_TypeI_MultiPanel_Restriction),
		}
	case CodebookConfig_codebookType_type1_subType_typeI_MultiPanel_ng_n1_n2_Choice_four_four_one_TypeI_MultiPanel_Restriction:
		var tmp_bs_four_four_one_TypeI_MultiPanel_Restriction []byte
		var n_four_four_one_TypeI_MultiPanel_Restriction uint
		if tmp_bs_four_four_one_TypeI_MultiPanel_Restriction, n_four_four_one_TypeI_MultiPanel_Restriction, err = r.ReadBitString(&uper.Constraint{Lb: 16, Ub: 16}, false); err != nil {
			return utils.WrapError("Decode four_four_one_TypeI_MultiPanel_Restriction", err)
		}
		ie.four_four_one_TypeI_MultiPanel_Restriction = uper.BitString{
			Bytes:   tmp_bs_four_four_one_TypeI_MultiPanel_Restriction,
			NumBits: uint64(n_four_four_one_TypeI_MultiPanel_Restriction),
		}
	case CodebookConfig_codebookType_type1_subType_typeI_MultiPanel_ng_n1_n2_Choice_two_four_two_TypeI_MultiPanel_Restriction:
		var tmp_bs_two_four_two_TypeI_MultiPanel_Restriction []byte
		var n_two_four_two_TypeI_MultiPanel_Restriction uint
		if tmp_bs_two_four_two_TypeI_MultiPanel_Restriction, n_two_four_two_TypeI_MultiPanel_Restriction, err = r.ReadBitString(&uper.Constraint{Lb: 128, Ub: 128}, false); err != nil {
			return utils.WrapError("Decode two_four_two_TypeI_MultiPanel_Restriction", err)
		}
		ie.two_four_two_TypeI_MultiPanel_Restriction = uper.BitString{
			Bytes:   tmp_bs_two_four_two_TypeI_MultiPanel_Restriction,
			NumBits: uint64(n_two_four_two_TypeI_MultiPanel_Restriction),
		}
	case CodebookConfig_codebookType_type1_subType_typeI_MultiPanel_ng_n1_n2_Choice_four_two_two_TypeI_MultiPanel_Restriction:
		var tmp_bs_four_two_two_TypeI_MultiPanel_Restriction []byte
		var n_four_two_two_TypeI_MultiPanel_Restriction uint
		if tmp_bs_four_two_two_TypeI_MultiPanel_Restriction, n_four_two_two_TypeI_MultiPanel_Restriction, err = r.ReadBitString(&uper.Constraint{Lb: 64, Ub: 64}, false); err != nil {
			return utils.WrapError("Decode four_two_two_TypeI_MultiPanel_Restriction", err)
		}
		ie.four_two_two_TypeI_MultiPanel_Restriction = uper.BitString{
			Bytes:   tmp_bs_four_two_two_TypeI_MultiPanel_Restriction,
			NumBits: uint64(n_four_two_two_TypeI_MultiPanel_Restriction),
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
