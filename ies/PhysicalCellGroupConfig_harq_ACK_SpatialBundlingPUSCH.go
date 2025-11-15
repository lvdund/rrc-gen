package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	PhysicalCellGroupConfig_harq_ACK_SpatialBundlingPUSCH_Enum_true uper.Enumerated = 0
)

type PhysicalCellGroupConfig_harq_ACK_SpatialBundlingPUSCH struct {
	Value uper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *PhysicalCellGroupConfig_harq_ACK_SpatialBundlingPUSCH) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode PhysicalCellGroupConfig_harq_ACK_SpatialBundlingPUSCH", err)
	}
	return nil
}

func (ie *PhysicalCellGroupConfig_harq_ACK_SpatialBundlingPUSCH) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode PhysicalCellGroupConfig_harq_ACK_SpatialBundlingPUSCH", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
