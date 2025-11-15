package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	MeasResultNR_triggeredEvent_r17_firstTriggeredEvent_Enum_condFirstEvent  uper.Enumerated = 0
	MeasResultNR_triggeredEvent_r17_firstTriggeredEvent_Enum_condSecondEvent uper.Enumerated = 1
)

type MeasResultNR_triggeredEvent_r17_firstTriggeredEvent struct {
	Value uper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *MeasResultNR_triggeredEvent_r17_firstTriggeredEvent) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode MeasResultNR_triggeredEvent_r17_firstTriggeredEvent", err)
	}
	return nil
}

func (ie *MeasResultNR_triggeredEvent_r17_firstTriggeredEvent) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode MeasResultNR_triggeredEvent_r17_firstTriggeredEvent", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
