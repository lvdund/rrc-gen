package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	CodebookConfig_codebookType_type1_subType_Choice_nothing uint64 = iota
	CodebookConfig_codebookType_type1_subType_Choice_typeI_SinglePanel
	CodebookConfig_codebookType_type1_subType_Choice_typeI_MultiPanel
)

type CodebookConfig_codebookType_type1_subType struct {
	Choice            uint64
	typeI_SinglePanel *CodebookConfig_codebookType_type1_subType_typeI_SinglePanel
	typeI_MultiPanel  *CodebookConfig_codebookType_type1_subType_typeI_MultiPanel
}

func (ie *CodebookConfig_codebookType_type1_subType) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CodebookConfig_codebookType_type1_subType_Choice_typeI_SinglePanel:
		if err = ie.typeI_SinglePanel.Encode(w); err != nil {
			err = utils.WrapError("Encode typeI_SinglePanel", err)
		}
	case CodebookConfig_codebookType_type1_subType_Choice_typeI_MultiPanel:
		if err = ie.typeI_MultiPanel.Encode(w); err != nil {
			err = utils.WrapError("Encode typeI_MultiPanel", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *CodebookConfig_codebookType_type1_subType) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CodebookConfig_codebookType_type1_subType_Choice_typeI_SinglePanel:
		ie.typeI_SinglePanel = new(CodebookConfig_codebookType_type1_subType_typeI_SinglePanel)
		if err = ie.typeI_SinglePanel.Decode(r); err != nil {
			return utils.WrapError("Decode typeI_SinglePanel", err)
		}
	case CodebookConfig_codebookType_type1_subType_Choice_typeI_MultiPanel:
		ie.typeI_MultiPanel = new(CodebookConfig_codebookType_type1_subType_typeI_MultiPanel)
		if err = ie.typeI_MultiPanel.Decode(r); err != nil {
			return utils.WrapError("Decode typeI_MultiPanel", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
