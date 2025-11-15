package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	CodebookConfig_codebookType_type2_subType_Choice_nothing uint64 = iota
	CodebookConfig_codebookType_type2_subType_Choice_typeII
	CodebookConfig_codebookType_type2_subType_Choice_typeII_PortSelection
)

type CodebookConfig_codebookType_type2_subType struct {
	Choice               uint64
	typeII               *CodebookConfig_codebookType_type2_subType_typeII
	typeII_PortSelection *CodebookConfig_codebookType_type2_subType_typeII_PortSelection
}

func (ie *CodebookConfig_codebookType_type2_subType) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CodebookConfig_codebookType_type2_subType_Choice_typeII:
		if err = ie.typeII.Encode(w); err != nil {
			err = utils.WrapError("Encode typeII", err)
		}
	case CodebookConfig_codebookType_type2_subType_Choice_typeII_PortSelection:
		if err = ie.typeII_PortSelection.Encode(w); err != nil {
			err = utils.WrapError("Encode typeII_PortSelection", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *CodebookConfig_codebookType_type2_subType) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CodebookConfig_codebookType_type2_subType_Choice_typeII:
		ie.typeII = new(CodebookConfig_codebookType_type2_subType_typeII)
		if err = ie.typeII.Decode(r); err != nil {
			return utils.WrapError("Decode typeII", err)
		}
	case CodebookConfig_codebookType_type2_subType_Choice_typeII_PortSelection:
		ie.typeII_PortSelection = new(CodebookConfig_codebookType_type2_subType_typeII_PortSelection)
		if err = ie.typeII_PortSelection.Decode(r); err != nil {
			return utils.WrapError("Decode typeII_PortSelection", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
