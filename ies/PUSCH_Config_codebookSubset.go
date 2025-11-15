package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	PUSCH_Config_codebookSubset_Enum_fullyAndPartialAndNonCoherent uper.Enumerated = 0
	PUSCH_Config_codebookSubset_Enum_partialAndNonCoherent         uper.Enumerated = 1
	PUSCH_Config_codebookSubset_Enum_nonCoherent                   uper.Enumerated = 2
)

type PUSCH_Config_codebookSubset struct {
	Value uper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *PUSCH_Config_codebookSubset) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode PUSCH_Config_codebookSubset", err)
	}
	return nil
}

func (ie *PUSCH_Config_codebookSubset) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode PUSCH_Config_codebookSubset", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
