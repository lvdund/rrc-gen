package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	PhysicalCellGroupConfig_xScale_Enum_dB0    uper.Enumerated = 0
	PhysicalCellGroupConfig_xScale_Enum_dB6    uper.Enumerated = 1
	PhysicalCellGroupConfig_xScale_Enum_spare2 uper.Enumerated = 2
	PhysicalCellGroupConfig_xScale_Enum_spare1 uper.Enumerated = 3
)

type PhysicalCellGroupConfig_xScale struct {
	Value uper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *PhysicalCellGroupConfig_xScale) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode PhysicalCellGroupConfig_xScale", err)
	}
	return nil
}

func (ie *PhysicalCellGroupConfig_xScale) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode PhysicalCellGroupConfig_xScale", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
