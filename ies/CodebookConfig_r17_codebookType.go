package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	CodebookConfig_r17_codebookType_Choice_nothing uint64 = iota
	CodebookConfig_r17_codebookType_Choice_type1
	CodebookConfig_r17_codebookType_Choice_type2
)

type CodebookConfig_r17_codebookType struct {
	Choice uint64
	type1  *CodebookConfig_r17_codebookType_type1
	type2  *CodebookConfig_r17_codebookType_type2
}

func (ie *CodebookConfig_r17_codebookType) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CodebookConfig_r17_codebookType_Choice_type1:
		if err = ie.type1.Encode(w); err != nil {
			err = utils.WrapError("Encode type1", err)
		}
	case CodebookConfig_r17_codebookType_Choice_type2:
		if err = ie.type2.Encode(w); err != nil {
			err = utils.WrapError("Encode type2", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *CodebookConfig_r17_codebookType) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CodebookConfig_r17_codebookType_Choice_type1:
		ie.type1 = new(CodebookConfig_r17_codebookType_type1)
		if err = ie.type1.Decode(r); err != nil {
			return utils.WrapError("Decode type1", err)
		}
	case CodebookConfig_r17_codebookType_Choice_type2:
		ie.type2 = new(CodebookConfig_r17_codebookType_type2)
		if err = ie.type2.Decode(r); err != nil {
			return utils.WrapError("Decode type2", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
