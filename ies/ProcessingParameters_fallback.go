package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	ProcessingParameters_fallback_Enum_sc        uper.Enumerated = 0
	ProcessingParameters_fallback_Enum_cap1_only uper.Enumerated = 1
)

type ProcessingParameters_fallback struct {
	Value uper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *ProcessingParameters_fallback) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode ProcessingParameters_fallback", err)
	}
	return nil
}

func (ie *ProcessingParameters_fallback) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode ProcessingParameters_fallback", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
