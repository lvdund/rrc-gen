package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	CodebookConfig_codebookType_type2_numberOfBeams_Enum_two   uper.Enumerated = 0
	CodebookConfig_codebookType_type2_numberOfBeams_Enum_three uper.Enumerated = 1
	CodebookConfig_codebookType_type2_numberOfBeams_Enum_four  uper.Enumerated = 2
)

type CodebookConfig_codebookType_type2_numberOfBeams struct {
	Value uper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *CodebookConfig_codebookType_type2_numberOfBeams) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode CodebookConfig_codebookType_type2_numberOfBeams", err)
	}
	return nil
}

func (ie *CodebookConfig_codebookType_type2_numberOfBeams) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode CodebookConfig_codebookType_type2_numberOfBeams", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
