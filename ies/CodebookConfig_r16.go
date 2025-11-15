package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CodebookConfig_r16 struct {
	codebookType CodebookConfig_r16_codebookType `lb:16,ub:16,madatory`
}

func (ie *CodebookConfig_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.codebookType.Encode(w); err != nil {
		return utils.WrapError("Encode codebookType", err)
	}
	return nil
}

func (ie *CodebookConfig_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.codebookType.Decode(r); err != nil {
		return utils.WrapError("Decode codebookType", err)
	}
	return nil
}
