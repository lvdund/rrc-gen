package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	CodebookConfig_r16_codebookType_type2_subType_Choice_nothing uint64 = iota
	CodebookConfig_r16_codebookType_type2_subType_Choice_typeII_r16
	CodebookConfig_r16_codebookType_type2_subType_Choice_typeII_PortSelection_r16
)

type CodebookConfig_r16_codebookType_type2_subType struct {
	Choice                   uint64
	typeII_r16               *CodebookConfig_r16_codebookType_type2_subType_typeII_r16
	typeII_PortSelection_r16 *CodebookConfig_r16_codebookType_type2_subType_typeII_PortSelection_r16
}

func (ie *CodebookConfig_r16_codebookType_type2_subType) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CodebookConfig_r16_codebookType_type2_subType_Choice_typeII_r16:
		if err = ie.typeII_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode typeII_r16", err)
		}
	case CodebookConfig_r16_codebookType_type2_subType_Choice_typeII_PortSelection_r16:
		if err = ie.typeII_PortSelection_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode typeII_PortSelection_r16", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *CodebookConfig_r16_codebookType_type2_subType) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CodebookConfig_r16_codebookType_type2_subType_Choice_typeII_r16:
		ie.typeII_r16 = new(CodebookConfig_r16_codebookType_type2_subType_typeII_r16)
		if err = ie.typeII_r16.Decode(r); err != nil {
			return utils.WrapError("Decode typeII_r16", err)
		}
	case CodebookConfig_r16_codebookType_type2_subType_Choice_typeII_PortSelection_r16:
		ie.typeII_PortSelection_r16 = new(CodebookConfig_r16_codebookType_type2_subType_typeII_PortSelection_r16)
		if err = ie.typeII_PortSelection_r16.Decode(r); err != nil {
			return utils.WrapError("Decode typeII_PortSelection_r16", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
