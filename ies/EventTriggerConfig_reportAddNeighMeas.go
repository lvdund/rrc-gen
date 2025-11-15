package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	EventTriggerConfig_reportAddNeighMeas_Enum_setup uper.Enumerated = 0
)

type EventTriggerConfig_reportAddNeighMeas struct {
	Value uper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *EventTriggerConfig_reportAddNeighMeas) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode EventTriggerConfig_reportAddNeighMeas", err)
	}
	return nil
}

func (ie *EventTriggerConfig_reportAddNeighMeas) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode EventTriggerConfig_reportAddNeighMeas", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
