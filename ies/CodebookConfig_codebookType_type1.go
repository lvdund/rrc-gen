package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CodebookConfig_codebookType_type1 struct {
	subType      *CodebookConfig_codebookType_type1_subType `lb:6,ub:6,optional`
	codebookMode int64                                      `lb:1,ub:2,madatory`
}

func (ie *CodebookConfig_codebookType_type1) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.subType != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.subType != nil {
		if err = ie.subType.Encode(w); err != nil {
			return utils.WrapError("Encode subType", err)
		}
	}
	if err = w.WriteInteger(ie.codebookMode, &uper.Constraint{Lb: 1, Ub: 2}, false); err != nil {
		return utils.WrapError("WriteInteger codebookMode", err)
	}
	return nil
}

func (ie *CodebookConfig_codebookType_type1) Decode(r *uper.UperReader) error {
	var err error
	var subTypePresent bool
	if subTypePresent, err = r.ReadBool(); err != nil {
		return err
	}
	if subTypePresent {
		ie.subType = new(CodebookConfig_codebookType_type1_subType)
		if err = ie.subType.Decode(r); err != nil {
			return utils.WrapError("Decode subType", err)
		}
	}
	var tmp_int_codebookMode int64
	if tmp_int_codebookMode, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 2}, false); err != nil {
		return utils.WrapError("ReadInteger codebookMode", err)
	}
	ie.codebookMode = tmp_int_codebookMode
	return nil
}
