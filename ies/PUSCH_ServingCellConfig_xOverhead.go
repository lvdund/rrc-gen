package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	PUSCH_ServingCellConfig_xOverhead_Enum_xoh6  uper.Enumerated = 0
	PUSCH_ServingCellConfig_xOverhead_Enum_xoh12 uper.Enumerated = 1
	PUSCH_ServingCellConfig_xOverhead_Enum_xoh18 uper.Enumerated = 2
)

type PUSCH_ServingCellConfig_xOverhead struct {
	Value uper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *PUSCH_ServingCellConfig_xOverhead) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode PUSCH_ServingCellConfig_xOverhead", err)
	}
	return nil
}

func (ie *PUSCH_ServingCellConfig_xOverhead) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode PUSCH_ServingCellConfig_xOverhead", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
