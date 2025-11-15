package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CodebookConfig_r16_codebookType_type2 struct {
	subType                                CodebookConfig_r16_codebookType_type2_subType `lb:16,ub:16,madatory`
	numberOfPMI_SubbandsPerCQI_Subband_r16 int64                                         `lb:1,ub:2,madatory`
	paramCombination_r16                   int64                                         `lb:1,ub:8,madatory`
}

func (ie *CodebookConfig_r16_codebookType_type2) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.subType.Encode(w); err != nil {
		return utils.WrapError("Encode subType", err)
	}
	if err = w.WriteInteger(ie.numberOfPMI_SubbandsPerCQI_Subband_r16, &uper.Constraint{Lb: 1, Ub: 2}, false); err != nil {
		return utils.WrapError("WriteInteger numberOfPMI_SubbandsPerCQI_Subband_r16", err)
	}
	if err = w.WriteInteger(ie.paramCombination_r16, &uper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
		return utils.WrapError("WriteInteger paramCombination_r16", err)
	}
	return nil
}

func (ie *CodebookConfig_r16_codebookType_type2) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.subType.Decode(r); err != nil {
		return utils.WrapError("Decode subType", err)
	}
	var tmp_int_numberOfPMI_SubbandsPerCQI_Subband_r16 int64
	if tmp_int_numberOfPMI_SubbandsPerCQI_Subband_r16, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 2}, false); err != nil {
		return utils.WrapError("ReadInteger numberOfPMI_SubbandsPerCQI_Subband_r16", err)
	}
	ie.numberOfPMI_SubbandsPerCQI_Subband_r16 = tmp_int_numberOfPMI_SubbandsPerCQI_Subband_r16
	var tmp_int_paramCombination_r16 int64
	if tmp_int_paramCombination_r16, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
		return utils.WrapError("ReadInteger paramCombination_r16", err)
	}
	ie.paramCombination_r16 = tmp_int_paramCombination_r16
	return nil
}
