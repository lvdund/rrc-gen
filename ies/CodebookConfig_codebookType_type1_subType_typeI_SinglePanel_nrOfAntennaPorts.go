package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	CodebookConfig_codebookType_type1_subType_typeI_SinglePanel_nrOfAntennaPorts_Choice_nothing uint64 = iota
	CodebookConfig_codebookType_type1_subType_typeI_SinglePanel_nrOfAntennaPorts_Choice_two
	CodebookConfig_codebookType_type1_subType_typeI_SinglePanel_nrOfAntennaPorts_Choice_moreThanTwo
)

type CodebookConfig_codebookType_type1_subType_typeI_SinglePanel_nrOfAntennaPorts struct {
	Choice      uint64
	two         *CodebookConfig_codebookType_type1_subType_typeI_SinglePanel_nrOfAntennaPorts_two
	moreThanTwo *CodebookConfig_codebookType_type1_subType_typeI_SinglePanel_nrOfAntennaPorts_moreThanTwo
}

func (ie *CodebookConfig_codebookType_type1_subType_typeI_SinglePanel_nrOfAntennaPorts) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CodebookConfig_codebookType_type1_subType_typeI_SinglePanel_nrOfAntennaPorts_Choice_two:
		if err = ie.two.Encode(w); err != nil {
			err = utils.WrapError("Encode two", err)
		}
	case CodebookConfig_codebookType_type1_subType_typeI_SinglePanel_nrOfAntennaPorts_Choice_moreThanTwo:
		if err = ie.moreThanTwo.Encode(w); err != nil {
			err = utils.WrapError("Encode moreThanTwo", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *CodebookConfig_codebookType_type1_subType_typeI_SinglePanel_nrOfAntennaPorts) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CodebookConfig_codebookType_type1_subType_typeI_SinglePanel_nrOfAntennaPorts_Choice_two:
		ie.two = new(CodebookConfig_codebookType_type1_subType_typeI_SinglePanel_nrOfAntennaPorts_two)
		if err = ie.two.Decode(r); err != nil {
			return utils.WrapError("Decode two", err)
		}
	case CodebookConfig_codebookType_type1_subType_typeI_SinglePanel_nrOfAntennaPorts_Choice_moreThanTwo:
		ie.moreThanTwo = new(CodebookConfig_codebookType_type1_subType_typeI_SinglePanel_nrOfAntennaPorts_moreThanTwo)
		if err = ie.moreThanTwo.Decode(r); err != nil {
			return utils.WrapError("Decode moreThanTwo", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
