package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CodebookConfig_codebookType_type1_subType_typeI_MultiPanel struct {
	ng_n1_n2       CodebookConfig_codebookType_type1_subType_typeI_MultiPanel_ng_n1_n2 `lb:8,ub:8,madatory`
	ri_Restriction uper.BitString                                                      `lb:4,ub:4,madatory`
}

func (ie *CodebookConfig_codebookType_type1_subType_typeI_MultiPanel) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.ng_n1_n2.Encode(w); err != nil {
		return utils.WrapError("Encode ng_n1_n2", err)
	}
	if err = w.WriteBitString(ie.ri_Restriction.Bytes, uint(ie.ri_Restriction.NumBits), &uper.Constraint{Lb: 4, Ub: 4}, false); err != nil {
		return utils.WrapError("WriteBitString ri_Restriction", err)
	}
	return nil
}

func (ie *CodebookConfig_codebookType_type1_subType_typeI_MultiPanel) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.ng_n1_n2.Decode(r); err != nil {
		return utils.WrapError("Decode ng_n1_n2", err)
	}
	var tmp_bs_ri_Restriction []byte
	var n_ri_Restriction uint
	if tmp_bs_ri_Restriction, n_ri_Restriction, err = r.ReadBitString(&uper.Constraint{Lb: 4, Ub: 4}, false); err != nil {
		return utils.WrapError("ReadBitString ri_Restriction", err)
	}
	ie.ri_Restriction = uper.BitString{
		Bytes:   tmp_bs_ri_Restriction,
		NumBits: uint64(n_ri_Restriction),
	}
	return nil
}
