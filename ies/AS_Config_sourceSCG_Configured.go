package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	AS_Config_sourceSCG_Configured_Enum_true uper.Enumerated = 0
)

type AS_Config_sourceSCG_Configured struct {
	Value uper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *AS_Config_sourceSCG_Configured) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode AS_Config_sourceSCG_Configured", err)
	}
	return nil
}

func (ie *AS_Config_sourceSCG_Configured) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode AS_Config_sourceSCG_Configured", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
