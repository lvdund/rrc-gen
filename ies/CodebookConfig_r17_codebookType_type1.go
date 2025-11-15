package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CodebookConfig_r17_codebookType_type1 struct {
	typeI_SinglePanel_Group1_r17             *CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group1_r17 `lb:6,ub:6,optional`
	typeI_SinglePanel_Group2_r17             *CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group2_r17 `lb:6,ub:6,optional`
	typeI_SinglePanel_ri_RestrictionSTRP_r17 *uper.BitString                                                     `lb:8,ub:8,optional`
	typeI_SinglePanel_ri_RestrictionSDM_r17  *uper.BitString                                                     `lb:4,ub:4,optional`
}

func (ie *CodebookConfig_r17_codebookType_type1) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.typeI_SinglePanel_Group1_r17 != nil, ie.typeI_SinglePanel_Group2_r17 != nil, ie.typeI_SinglePanel_ri_RestrictionSTRP_r17 != nil, ie.typeI_SinglePanel_ri_RestrictionSDM_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.typeI_SinglePanel_Group1_r17 != nil {
		if err = ie.typeI_SinglePanel_Group1_r17.Encode(w); err != nil {
			return utils.WrapError("Encode typeI_SinglePanel_Group1_r17", err)
		}
	}
	if ie.typeI_SinglePanel_Group2_r17 != nil {
		if err = ie.typeI_SinglePanel_Group2_r17.Encode(w); err != nil {
			return utils.WrapError("Encode typeI_SinglePanel_Group2_r17", err)
		}
	}
	if ie.typeI_SinglePanel_ri_RestrictionSTRP_r17 != nil {
		if err = w.WriteBitString(ie.typeI_SinglePanel_ri_RestrictionSTRP_r17.Bytes, uint(ie.typeI_SinglePanel_ri_RestrictionSTRP_r17.NumBits), &uper.Constraint{Lb: 8, Ub: 8}, false); err != nil {
			return utils.WrapError("Encode typeI_SinglePanel_ri_RestrictionSTRP_r17", err)
		}
	}
	if ie.typeI_SinglePanel_ri_RestrictionSDM_r17 != nil {
		if err = w.WriteBitString(ie.typeI_SinglePanel_ri_RestrictionSDM_r17.Bytes, uint(ie.typeI_SinglePanel_ri_RestrictionSDM_r17.NumBits), &uper.Constraint{Lb: 4, Ub: 4}, false); err != nil {
			return utils.WrapError("Encode typeI_SinglePanel_ri_RestrictionSDM_r17", err)
		}
	}
	return nil
}

func (ie *CodebookConfig_r17_codebookType_type1) Decode(r *uper.UperReader) error {
	var err error
	var typeI_SinglePanel_Group1_r17Present bool
	if typeI_SinglePanel_Group1_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var typeI_SinglePanel_Group2_r17Present bool
	if typeI_SinglePanel_Group2_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var typeI_SinglePanel_ri_RestrictionSTRP_r17Present bool
	if typeI_SinglePanel_ri_RestrictionSTRP_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var typeI_SinglePanel_ri_RestrictionSDM_r17Present bool
	if typeI_SinglePanel_ri_RestrictionSDM_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if typeI_SinglePanel_Group1_r17Present {
		ie.typeI_SinglePanel_Group1_r17 = new(CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group1_r17)
		if err = ie.typeI_SinglePanel_Group1_r17.Decode(r); err != nil {
			return utils.WrapError("Decode typeI_SinglePanel_Group1_r17", err)
		}
	}
	if typeI_SinglePanel_Group2_r17Present {
		ie.typeI_SinglePanel_Group2_r17 = new(CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group2_r17)
		if err = ie.typeI_SinglePanel_Group2_r17.Decode(r); err != nil {
			return utils.WrapError("Decode typeI_SinglePanel_Group2_r17", err)
		}
	}
	if typeI_SinglePanel_ri_RestrictionSTRP_r17Present {
		var tmp_bs_typeI_SinglePanel_ri_RestrictionSTRP_r17 []byte
		var n_typeI_SinglePanel_ri_RestrictionSTRP_r17 uint
		if tmp_bs_typeI_SinglePanel_ri_RestrictionSTRP_r17, n_typeI_SinglePanel_ri_RestrictionSTRP_r17, err = r.ReadBitString(&uper.Constraint{Lb: 8, Ub: 8}, false); err != nil {
			return utils.WrapError("Decode typeI_SinglePanel_ri_RestrictionSTRP_r17", err)
		}
		tmp_bitstring := uper.BitString{
			Bytes:   tmp_bs_typeI_SinglePanel_ri_RestrictionSTRP_r17,
			NumBits: uint64(n_typeI_SinglePanel_ri_RestrictionSTRP_r17),
		}
		ie.typeI_SinglePanel_ri_RestrictionSTRP_r17 = &tmp_bitstring
	}
	if typeI_SinglePanel_ri_RestrictionSDM_r17Present {
		var tmp_bs_typeI_SinglePanel_ri_RestrictionSDM_r17 []byte
		var n_typeI_SinglePanel_ri_RestrictionSDM_r17 uint
		if tmp_bs_typeI_SinglePanel_ri_RestrictionSDM_r17, n_typeI_SinglePanel_ri_RestrictionSDM_r17, err = r.ReadBitString(&uper.Constraint{Lb: 4, Ub: 4}, false); err != nil {
			return utils.WrapError("Decode typeI_SinglePanel_ri_RestrictionSDM_r17", err)
		}
		tmp_bitstring := uper.BitString{
			Bytes:   tmp_bs_typeI_SinglePanel_ri_RestrictionSDM_r17,
			NumBits: uint64(n_typeI_SinglePanel_ri_RestrictionSDM_r17),
		}
		ie.typeI_SinglePanel_ri_RestrictionSDM_r17 = &tmp_bitstring
	}
	return nil
}
