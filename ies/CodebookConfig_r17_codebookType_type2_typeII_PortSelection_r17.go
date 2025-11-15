package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CodebookConfig_r17_codebookType_type2_typeII_PortSelection_r17 struct {
	paramCombination_r17                   int64                                                                        `lb:1,ub:8,madatory`
	valueOfN_r17                           *CodebookConfig_r17_codebookType_type2_typeII_PortSelection_r17_valueOfN_r17 `optional`
	numberOfPMI_SubbandsPerCQI_Subband_r17 *int64                                                                       `lb:1,ub:2,optional`
	typeII_PortSelectionRI_Restriction_r17 uper.BitString                                                               `lb:4,ub:4,madatory`
}

func (ie *CodebookConfig_r17_codebookType_type2_typeII_PortSelection_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.valueOfN_r17 != nil, ie.numberOfPMI_SubbandsPerCQI_Subband_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = w.WriteInteger(ie.paramCombination_r17, &uper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
		return utils.WrapError("WriteInteger paramCombination_r17", err)
	}
	if ie.valueOfN_r17 != nil {
		if err = ie.valueOfN_r17.Encode(w); err != nil {
			return utils.WrapError("Encode valueOfN_r17", err)
		}
	}
	if ie.numberOfPMI_SubbandsPerCQI_Subband_r17 != nil {
		if err = w.WriteInteger(*ie.numberOfPMI_SubbandsPerCQI_Subband_r17, &uper.Constraint{Lb: 1, Ub: 2}, false); err != nil {
			return utils.WrapError("Encode numberOfPMI_SubbandsPerCQI_Subband_r17", err)
		}
	}
	if err = w.WriteBitString(ie.typeII_PortSelectionRI_Restriction_r17.Bytes, uint(ie.typeII_PortSelectionRI_Restriction_r17.NumBits), &uper.Constraint{Lb: 4, Ub: 4}, false); err != nil {
		return utils.WrapError("WriteBitString typeII_PortSelectionRI_Restriction_r17", err)
	}
	return nil
}

func (ie *CodebookConfig_r17_codebookType_type2_typeII_PortSelection_r17) Decode(r *uper.UperReader) error {
	var err error
	var valueOfN_r17Present bool
	if valueOfN_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var numberOfPMI_SubbandsPerCQI_Subband_r17Present bool
	if numberOfPMI_SubbandsPerCQI_Subband_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var tmp_int_paramCombination_r17 int64
	if tmp_int_paramCombination_r17, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
		return utils.WrapError("ReadInteger paramCombination_r17", err)
	}
	ie.paramCombination_r17 = tmp_int_paramCombination_r17
	if valueOfN_r17Present {
		ie.valueOfN_r17 = new(CodebookConfig_r17_codebookType_type2_typeII_PortSelection_r17_valueOfN_r17)
		if err = ie.valueOfN_r17.Decode(r); err != nil {
			return utils.WrapError("Decode valueOfN_r17", err)
		}
	}
	if numberOfPMI_SubbandsPerCQI_Subband_r17Present {
		var tmp_int_numberOfPMI_SubbandsPerCQI_Subband_r17 int64
		if tmp_int_numberOfPMI_SubbandsPerCQI_Subband_r17, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 2}, false); err != nil {
			return utils.WrapError("Decode numberOfPMI_SubbandsPerCQI_Subband_r17", err)
		}
		ie.numberOfPMI_SubbandsPerCQI_Subband_r17 = &tmp_int_numberOfPMI_SubbandsPerCQI_Subband_r17
	}
	var tmp_bs_typeII_PortSelectionRI_Restriction_r17 []byte
	var n_typeII_PortSelectionRI_Restriction_r17 uint
	if tmp_bs_typeII_PortSelectionRI_Restriction_r17, n_typeII_PortSelectionRI_Restriction_r17, err = r.ReadBitString(&uper.Constraint{Lb: 4, Ub: 4}, false); err != nil {
		return utils.WrapError("ReadBitString typeII_PortSelectionRI_Restriction_r17", err)
	}
	ie.typeII_PortSelectionRI_Restriction_r17 = uper.BitString{
		Bytes:   tmp_bs_typeII_PortSelectionRI_Restriction_r17,
		NumBits: uint64(n_typeII_PortSelectionRI_Restriction_r17),
	}
	return nil
}
