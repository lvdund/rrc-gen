package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CodebookConfig_codebookType_type1_subType_typeI_SinglePanel struct {
	nrOfAntennaPorts                 *CodebookConfig_codebookType_type1_subType_typeI_SinglePanel_nrOfAntennaPorts `lb:6,ub:6,optional`
	typeI_SinglePanel_ri_Restriction uper.BitString                                                                `lb:8,ub:8,madatory`
}

func (ie *CodebookConfig_codebookType_type1_subType_typeI_SinglePanel) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.nrOfAntennaPorts != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.nrOfAntennaPorts != nil {
		if err = ie.nrOfAntennaPorts.Encode(w); err != nil {
			return utils.WrapError("Encode nrOfAntennaPorts", err)
		}
	}
	if err = w.WriteBitString(ie.typeI_SinglePanel_ri_Restriction.Bytes, uint(ie.typeI_SinglePanel_ri_Restriction.NumBits), &uper.Constraint{Lb: 8, Ub: 8}, false); err != nil {
		return utils.WrapError("WriteBitString typeI_SinglePanel_ri_Restriction", err)
	}
	return nil
}

func (ie *CodebookConfig_codebookType_type1_subType_typeI_SinglePanel) Decode(r *uper.UperReader) error {
	var err error
	var nrOfAntennaPortsPresent bool
	if nrOfAntennaPortsPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if nrOfAntennaPortsPresent {
		ie.nrOfAntennaPorts = new(CodebookConfig_codebookType_type1_subType_typeI_SinglePanel_nrOfAntennaPorts)
		if err = ie.nrOfAntennaPorts.Decode(r); err != nil {
			return utils.WrapError("Decode nrOfAntennaPorts", err)
		}
	}
	var tmp_bs_typeI_SinglePanel_ri_Restriction []byte
	var n_typeI_SinglePanel_ri_Restriction uint
	if tmp_bs_typeI_SinglePanel_ri_Restriction, n_typeI_SinglePanel_ri_Restriction, err = r.ReadBitString(&uper.Constraint{Lb: 8, Ub: 8}, false); err != nil {
		return utils.WrapError("ReadBitString typeI_SinglePanel_ri_Restriction", err)
	}
	ie.typeI_SinglePanel_ri_Restriction = uper.BitString{
		Bytes:   tmp_bs_typeI_SinglePanel_ri_Restriction,
		NumBits: uint64(n_typeI_SinglePanel_ri_Restriction),
	}
	return nil
}
