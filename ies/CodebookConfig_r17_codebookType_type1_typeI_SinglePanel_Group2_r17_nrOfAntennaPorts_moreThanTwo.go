package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group2_r17_nrOfAntennaPorts_moreThanTwo struct {
	n1_n2 CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group2_r17_nrOfAntennaPorts_moreThanTwo_n1_n2 `lb:8,ub:8,madatory`
}

func (ie *CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group2_r17_nrOfAntennaPorts_moreThanTwo) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.n1_n2.Encode(w); err != nil {
		return utils.WrapError("Encode n1_n2", err)
	}
	return nil
}

func (ie *CodebookConfig_r17_codebookType_type1_typeI_SinglePanel_Group2_r17_nrOfAntennaPorts_moreThanTwo) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.n1_n2.Decode(r); err != nil {
		return utils.WrapError("Decode n1_n2", err)
	}
	return nil
}
