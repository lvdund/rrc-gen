package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SRI_PUSCH_PowerControl_sri_PUSCH_ClosedLoopIndex_Enum_i0 uper.Enumerated = 0
	SRI_PUSCH_PowerControl_sri_PUSCH_ClosedLoopIndex_Enum_i1 uper.Enumerated = 1
)

type SRI_PUSCH_PowerControl_sri_PUSCH_ClosedLoopIndex struct {
	Value uper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *SRI_PUSCH_PowerControl_sri_PUSCH_ClosedLoopIndex) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode SRI_PUSCH_PowerControl_sri_PUSCH_ClosedLoopIndex", err)
	}
	return nil
}

func (ie *SRI_PUSCH_PowerControl_sri_PUSCH_ClosedLoopIndex) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode SRI_PUSCH_PowerControl_sri_PUSCH_ClosedLoopIndex", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
