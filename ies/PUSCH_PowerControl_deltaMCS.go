package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	PUSCH_PowerControl_deltaMCS_Enum_enabled uper.Enumerated = 0
)

type PUSCH_PowerControl_deltaMCS struct {
	Value uper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *PUSCH_PowerControl_deltaMCS) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode PUSCH_PowerControl_deltaMCS", err)
	}
	return nil
}

func (ie *PUSCH_PowerControl_deltaMCS) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode PUSCH_PowerControl_deltaMCS", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
