package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	CodebookConfig_v1730_codebookType_Choice_nothing uint64 = iota
	CodebookConfig_v1730_codebookType_Choice_type1
)

type CodebookConfig_v1730_codebookType struct {
	Choice uint64
	type1  *CodebookConfig_v1730_codebookType_type1
}

func (ie *CodebookConfig_v1730_codebookType) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 1, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CodebookConfig_v1730_codebookType_Choice_type1:
		if err = ie.type1.Encode(w); err != nil {
			err = utils.WrapError("Encode type1", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *CodebookConfig_v1730_codebookType) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(1, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CodebookConfig_v1730_codebookType_Choice_type1:
		ie.type1 = new(CodebookConfig_v1730_codebookType_type1)
		if err = ie.type1.Decode(r); err != nil {
			return utils.WrapError("Decode type1", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
