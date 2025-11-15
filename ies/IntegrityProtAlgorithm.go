package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	IntegrityProtAlgorithm_Enum_nia0   uper.Enumerated = 0
	IntegrityProtAlgorithm_Enum_nia1   uper.Enumerated = 1
	IntegrityProtAlgorithm_Enum_nia2   uper.Enumerated = 2
	IntegrityProtAlgorithm_Enum_nia3   uper.Enumerated = 3
	IntegrityProtAlgorithm_Enum_spare4 uper.Enumerated = 4
	IntegrityProtAlgorithm_Enum_spare3 uper.Enumerated = 5
	IntegrityProtAlgorithm_Enum_spare2 uper.Enumerated = 6
	IntegrityProtAlgorithm_Enum_spare1 uper.Enumerated = 7
)

type IntegrityProtAlgorithm struct {
	Value uper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *IntegrityProtAlgorithm) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode IntegrityProtAlgorithm", err)
	}
	return nil
}

func (ie *IntegrityProtAlgorithm) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode IntegrityProtAlgorithm", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
