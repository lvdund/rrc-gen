package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CodebookConfig_codebookType_type2_subType_typeII_PortSelection struct {
	portSelectionSamplingSize          *CodebookConfig_codebookType_type2_subType_typeII_PortSelection_portSelectionSamplingSize `optional`
	typeII_PortSelectionRI_Restriction uper.BitString                                                                            `lb:2,ub:2,madatory`
}

func (ie *CodebookConfig_codebookType_type2_subType_typeII_PortSelection) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.portSelectionSamplingSize != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.portSelectionSamplingSize != nil {
		if err = ie.portSelectionSamplingSize.Encode(w); err != nil {
			return utils.WrapError("Encode portSelectionSamplingSize", err)
		}
	}
	if err = w.WriteBitString(ie.typeII_PortSelectionRI_Restriction.Bytes, uint(ie.typeII_PortSelectionRI_Restriction.NumBits), &uper.Constraint{Lb: 2, Ub: 2}, false); err != nil {
		return utils.WrapError("WriteBitString typeII_PortSelectionRI_Restriction", err)
	}
	return nil
}

func (ie *CodebookConfig_codebookType_type2_subType_typeII_PortSelection) Decode(r *uper.UperReader) error {
	var err error
	var portSelectionSamplingSizePresent bool
	if portSelectionSamplingSizePresent, err = r.ReadBool(); err != nil {
		return err
	}
	if portSelectionSamplingSizePresent {
		ie.portSelectionSamplingSize = new(CodebookConfig_codebookType_type2_subType_typeII_PortSelection_portSelectionSamplingSize)
		if err = ie.portSelectionSamplingSize.Decode(r); err != nil {
			return utils.WrapError("Decode portSelectionSamplingSize", err)
		}
	}
	var tmp_bs_typeII_PortSelectionRI_Restriction []byte
	var n_typeII_PortSelectionRI_Restriction uint
	if tmp_bs_typeII_PortSelectionRI_Restriction, n_typeII_PortSelectionRI_Restriction, err = r.ReadBitString(&uper.Constraint{Lb: 2, Ub: 2}, false); err != nil {
		return utils.WrapError("ReadBitString typeII_PortSelectionRI_Restriction", err)
	}
	ie.typeII_PortSelectionRI_Restriction = uper.BitString{
		Bytes:   tmp_bs_typeII_PortSelectionRI_Restriction,
		NumBits: uint64(n_typeII_PortSelectionRI_Restriction),
	}
	return nil
}
