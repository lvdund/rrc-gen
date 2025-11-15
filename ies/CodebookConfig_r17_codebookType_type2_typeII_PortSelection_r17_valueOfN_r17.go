package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	CodebookConfig_r17_codebookType_type2_typeII_PortSelection_r17_valueOfN_r17_Enum_n2 uper.Enumerated = 0
	CodebookConfig_r17_codebookType_type2_typeII_PortSelection_r17_valueOfN_r17_Enum_n4 uper.Enumerated = 1
)

type CodebookConfig_r17_codebookType_type2_typeII_PortSelection_r17_valueOfN_r17 struct {
	Value uper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *CodebookConfig_r17_codebookType_type2_typeII_PortSelection_r17_valueOfN_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode CodebookConfig_r17_codebookType_type2_typeII_PortSelection_r17_valueOfN_r17", err)
	}
	return nil
}

func (ie *CodebookConfig_r17_codebookType_type2_typeII_PortSelection_r17_valueOfN_r17) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode CodebookConfig_r17_codebookType_type2_typeII_PortSelection_r17_valueOfN_r17", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
