package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CodebookConfig_codebookType_type1_subType_typeI_SinglePanel_nrOfAntennaPorts_two struct {
	twoTX_CodebookSubsetRestriction uper.BitString `lb:6,ub:6,madatory`
}

func (ie *CodebookConfig_codebookType_type1_subType_typeI_SinglePanel_nrOfAntennaPorts_two) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteBitString(ie.twoTX_CodebookSubsetRestriction.Bytes, uint(ie.twoTX_CodebookSubsetRestriction.NumBits), &uper.Constraint{Lb: 6, Ub: 6}, false); err != nil {
		return utils.WrapError("WriteBitString twoTX_CodebookSubsetRestriction", err)
	}
	return nil
}

func (ie *CodebookConfig_codebookType_type1_subType_typeI_SinglePanel_nrOfAntennaPorts_two) Decode(r *uper.UperReader) error {
	var err error
	var tmp_bs_twoTX_CodebookSubsetRestriction []byte
	var n_twoTX_CodebookSubsetRestriction uint
	if tmp_bs_twoTX_CodebookSubsetRestriction, n_twoTX_CodebookSubsetRestriction, err = r.ReadBitString(&uper.Constraint{Lb: 6, Ub: 6}, false); err != nil {
		return utils.WrapError("ReadBitString twoTX_CodebookSubsetRestriction", err)
	}
	ie.twoTX_CodebookSubsetRestriction = uper.BitString{
		Bytes:   tmp_bs_twoTX_CodebookSubsetRestriction,
		NumBits: uint64(n_twoTX_CodebookSubsetRestriction),
	}
	return nil
}
