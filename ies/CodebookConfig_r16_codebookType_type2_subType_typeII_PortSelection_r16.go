package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CodebookConfig_r16_codebookType_type2_subType_typeII_PortSelection_r16 struct {
	portSelectionSamplingSize_r16          CodebookConfig_r16_codebookType_type2_subType_typeII_PortSelection_r16_portSelectionSamplingSize_r16 `madatory`
	typeII_PortSelectionRI_Restriction_r16 uper.BitString                                                                                       `lb:4,ub:4,madatory`
}

func (ie *CodebookConfig_r16_codebookType_type2_subType_typeII_PortSelection_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.portSelectionSamplingSize_r16.Encode(w); err != nil {
		return utils.WrapError("Encode portSelectionSamplingSize_r16", err)
	}
	if err = w.WriteBitString(ie.typeII_PortSelectionRI_Restriction_r16.Bytes, uint(ie.typeII_PortSelectionRI_Restriction_r16.NumBits), &uper.Constraint{Lb: 4, Ub: 4}, false); err != nil {
		return utils.WrapError("WriteBitString typeII_PortSelectionRI_Restriction_r16", err)
	}
	return nil
}

func (ie *CodebookConfig_r16_codebookType_type2_subType_typeII_PortSelection_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.portSelectionSamplingSize_r16.Decode(r); err != nil {
		return utils.WrapError("Decode portSelectionSamplingSize_r16", err)
	}
	var tmp_bs_typeII_PortSelectionRI_Restriction_r16 []byte
	var n_typeII_PortSelectionRI_Restriction_r16 uint
	if tmp_bs_typeII_PortSelectionRI_Restriction_r16, n_typeII_PortSelectionRI_Restriction_r16, err = r.ReadBitString(&uper.Constraint{Lb: 4, Ub: 4}, false); err != nil {
		return utils.WrapError("ReadBitString typeII_PortSelectionRI_Restriction_r16", err)
	}
	ie.typeII_PortSelectionRI_Restriction_r16 = uper.BitString{
		Bytes:   tmp_bs_typeII_PortSelectionRI_Restriction_r16,
		NumBits: uint64(n_typeII_PortSelectionRI_Restriction_r16),
	}
	return nil
}
