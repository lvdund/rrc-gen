package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	EventTriggerConfig_coarseLocationRequest_r17_Enum_true uper.Enumerated = 0
)

type EventTriggerConfig_coarseLocationRequest_r17 struct {
	Value uper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *EventTriggerConfig_coarseLocationRequest_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode EventTriggerConfig_coarseLocationRequest_r17", err)
	}
	return nil
}

func (ie *EventTriggerConfig_coarseLocationRequest_r17) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode EventTriggerConfig_coarseLocationRequest_r17", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
