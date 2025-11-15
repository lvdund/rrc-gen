package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	CodebookConfig_codebookType_type2_phaseAlphabetSize_Enum_n4 uper.Enumerated = 0
	CodebookConfig_codebookType_type2_phaseAlphabetSize_Enum_n8 uper.Enumerated = 1
)

type CodebookConfig_codebookType_type2_phaseAlphabetSize struct {
	Value uper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *CodebookConfig_codebookType_type2_phaseAlphabetSize) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode CodebookConfig_codebookType_type2_phaseAlphabetSize", err)
	}
	return nil
}

func (ie *CodebookConfig_codebookType_type2_phaseAlphabetSize) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode CodebookConfig_codebookType_type2_phaseAlphabetSize", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
