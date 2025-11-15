package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CSI_AperiodicTriggerStateList struct {
	Value []CSI_AperiodicTriggerState `lb:1,ub:maxNrOfCSI_AperiodicTriggers,madatory`
}

func (ie *CSI_AperiodicTriggerStateList) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*CSI_AperiodicTriggerState]([]*CSI_AperiodicTriggerState{}, uper.Constraint{Lb: 1, Ub: maxNrOfCSI_AperiodicTriggers}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode CSI_AperiodicTriggerStateList", err)
	}
	return nil
}

func (ie *CSI_AperiodicTriggerStateList) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*CSI_AperiodicTriggerState]([]*CSI_AperiodicTriggerState{}, uper.Constraint{Lb: 1, Ub: maxNrOfCSI_AperiodicTriggers}, false)
	fn := func() *CSI_AperiodicTriggerState {
		return new(CSI_AperiodicTriggerState)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode CSI_AperiodicTriggerStateList", err)
	}
	ie.Value = []CSI_AperiodicTriggerState{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
