package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CodebookConfig_codebookType_type2_subType_typeII struct {
	n1_n2_codebookSubsetRestriction CodebookConfig_codebookType_type2_subType_typeII_n1_n2_codebookSubsetRestriction `lb:16,ub:16,madatory`
	typeII_RI_Restriction           uper.BitString                                                                   `lb:2,ub:2,madatory`
}

func (ie *CodebookConfig_codebookType_type2_subType_typeII) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.n1_n2_codebookSubsetRestriction.Encode(w); err != nil {
		return utils.WrapError("Encode n1_n2_codebookSubsetRestriction", err)
	}
	if err = w.WriteBitString(ie.typeII_RI_Restriction.Bytes, uint(ie.typeII_RI_Restriction.NumBits), &uper.Constraint{Lb: 2, Ub: 2}, false); err != nil {
		return utils.WrapError("WriteBitString typeII_RI_Restriction", err)
	}
	return nil
}

func (ie *CodebookConfig_codebookType_type2_subType_typeII) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.n1_n2_codebookSubsetRestriction.Decode(r); err != nil {
		return utils.WrapError("Decode n1_n2_codebookSubsetRestriction", err)
	}
	var tmp_bs_typeII_RI_Restriction []byte
	var n_typeII_RI_Restriction uint
	if tmp_bs_typeII_RI_Restriction, n_typeII_RI_Restriction, err = r.ReadBitString(&uper.Constraint{Lb: 2, Ub: 2}, false); err != nil {
		return utils.WrapError("ReadBitString typeII_RI_Restriction", err)
	}
	ie.typeII_RI_Restriction = uper.BitString{
		Bytes:   tmp_bs_typeII_RI_Restriction,
		NumBits: uint64(n_typeII_RI_Restriction),
	}
	return nil
}
