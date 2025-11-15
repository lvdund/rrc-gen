package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	CodebookConfig_codebookType_type2_subType_typeII_PortSelection_portSelectionSamplingSize_Enum_n1 uper.Enumerated = 0
	CodebookConfig_codebookType_type2_subType_typeII_PortSelection_portSelectionSamplingSize_Enum_n2 uper.Enumerated = 1
	CodebookConfig_codebookType_type2_subType_typeII_PortSelection_portSelectionSamplingSize_Enum_n3 uper.Enumerated = 2
	CodebookConfig_codebookType_type2_subType_typeII_PortSelection_portSelectionSamplingSize_Enum_n4 uper.Enumerated = 3
)

type CodebookConfig_codebookType_type2_subType_typeII_PortSelection_portSelectionSamplingSize struct {
	Value uper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *CodebookConfig_codebookType_type2_subType_typeII_PortSelection_portSelectionSamplingSize) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode CodebookConfig_codebookType_type2_subType_typeII_PortSelection_portSelectionSamplingSize", err)
	}
	return nil
}

func (ie *CodebookConfig_codebookType_type2_subType_typeII_PortSelection_portSelectionSamplingSize) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode CodebookConfig_codebookType_type2_subType_typeII_PortSelection_portSelectionSamplingSize", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
