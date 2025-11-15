package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CodebookConfig_r17 struct {
	codebookType *CodebookConfig_r17_codebookType `lb:6,ub:6,optional`
}

func (ie *CodebookConfig_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.codebookType != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.codebookType != nil {
		if err = ie.codebookType.Encode(w); err != nil {
			return utils.WrapError("Encode codebookType", err)
		}
	}
	return nil
}

func (ie *CodebookConfig_r17) Decode(r *uper.UperReader) error {
	var err error
	var codebookTypePresent bool
	if codebookTypePresent, err = r.ReadBool(); err != nil {
		return err
	}
	if codebookTypePresent {
		ie.codebookType = new(CodebookConfig_r17_codebookType)
		if err = ie.codebookType.Decode(r); err != nil {
			return utils.WrapError("Decode codebookType", err)
		}
	}
	return nil
}
