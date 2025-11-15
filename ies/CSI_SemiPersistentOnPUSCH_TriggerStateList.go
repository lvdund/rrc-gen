package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CSI_SemiPersistentOnPUSCH_TriggerStateList struct {
	Value []CSI_SemiPersistentOnPUSCH_TriggerState `lb:1,ub:maxNrOfSemiPersistentPUSCH_Triggers,madatory`
}

func (ie *CSI_SemiPersistentOnPUSCH_TriggerStateList) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*CSI_SemiPersistentOnPUSCH_TriggerState]([]*CSI_SemiPersistentOnPUSCH_TriggerState{}, uper.Constraint{Lb: 1, Ub: maxNrOfSemiPersistentPUSCH_Triggers}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode CSI_SemiPersistentOnPUSCH_TriggerStateList", err)
	}
	return nil
}

func (ie *CSI_SemiPersistentOnPUSCH_TriggerStateList) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*CSI_SemiPersistentOnPUSCH_TriggerState]([]*CSI_SemiPersistentOnPUSCH_TriggerState{}, uper.Constraint{Lb: 1, Ub: maxNrOfSemiPersistentPUSCH_Triggers}, false)
	fn := func() *CSI_SemiPersistentOnPUSCH_TriggerState {
		return new(CSI_SemiPersistentOnPUSCH_TriggerState)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode CSI_SemiPersistentOnPUSCH_TriggerStateList", err)
	}
	ie.Value = []CSI_SemiPersistentOnPUSCH_TriggerState{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
