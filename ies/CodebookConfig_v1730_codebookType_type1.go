package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CodebookConfig_v1730_codebookType_type1 struct {
	codebookMode *int64 `lb:1,ub:2,optional`
}

func (ie *CodebookConfig_v1730_codebookType_type1) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.codebookMode != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.codebookMode != nil {
		if err = w.WriteInteger(*ie.codebookMode, &uper.Constraint{Lb: 1, Ub: 2}, false); err != nil {
			return utils.WrapError("Encode codebookMode", err)
		}
	}
	return nil
}

func (ie *CodebookConfig_v1730_codebookType_type1) Decode(r *uper.UperReader) error {
	var err error
	var codebookModePresent bool
	if codebookModePresent, err = r.ReadBool(); err != nil {
		return err
	}
	if codebookModePresent {
		var tmp_int_codebookMode int64
		if tmp_int_codebookMode, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 2}, false); err != nil {
			return utils.WrapError("Decode codebookMode", err)
		}
		ie.codebookMode = &tmp_int_codebookMode
	}
	return nil
}
