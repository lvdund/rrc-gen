package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	ResourceType_r16_SRS_PosResourceSet_r16_Choice_nothing uint64 = iota
	ResourceType_r16_SRS_PosResourceSet_r16_Choice_aperiodic_r16
	ResourceType_r16_SRS_PosResourceSet_r16_Choice_semi_persistent_r16
	ResourceType_r16_SRS_PosResourceSet_r16_Choice_periodic_r16
)

type ResourceType_r16_SRS_PosResourceSet_r16 struct {
	Choice              uint64
	aperiodic_r16       []int64     `lb:1,ub:maxNrofSRS_TriggerStates_1,e_lb:1,e_ub:maxNrofSRS_TriggerStates_1,optional`
	semi_persistent_r16 interface{} `madatory,ext`
	periodic_r16        interface{} `madatory,ext`
}

func (ie *ResourceType_r16_SRS_PosResourceSet_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 3, false); err != nil {
		return err
	}
	switch ie.Choice {
	case ResourceType_r16_SRS_PosResourceSet_r16_Choice_aperiodic_r16:
		tmp := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofSRS_TriggerStates_1}, false)
		for _, i := range ie.aperiodic_r16 {
			tmp_ie := utils.NewINTEGER(int64(i), uper.Constraint{Lb: 1, Ub: maxNrofSRS_TriggerStates_1}, false)
			tmp.Value = append(tmp.Value, &tmp_ie)
		}
		if err = tmp.Encode(w); err != nil {
			err = utils.WrapError("Encode aperiodic_r16", err)
		}
	case ResourceType_r16_SRS_PosResourceSet_r16_Choice_semi_persistent_r16:
		// interface{} field of choice - nothing to encode
	case ResourceType_r16_SRS_PosResourceSet_r16_Choice_periodic_r16:
		// interface{} field of choice - nothing to encode
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *ResourceType_r16_SRS_PosResourceSet_r16) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(3, false); err != nil {
		return err
	}
	switch ie.Choice {
	case ResourceType_r16_SRS_PosResourceSet_r16_Choice_aperiodic_r16:
		tmp := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofSRS_TriggerStates_1}, false)
		fn := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, uper.Constraint{Lb: 1, Ub: maxNrofSRS_TriggerStates_1}, false)
			return &ie
		}
		if err = tmp.Decode(r, fn); err != nil {
			return utils.WrapError("Decode aperiodic_r16", err)
		}
		ie.aperiodic_r16 = []int64{}
		for _, i := range tmp.Value {
			ie.aperiodic_r16 = append(ie.aperiodic_r16, int64(i.Value))
		}
	case ResourceType_r16_SRS_PosResourceSet_r16_Choice_semi_persistent_r16:
		// interface{} field of choice - nothing to decode
	case ResourceType_r16_SRS_PosResourceSet_r16_Choice_periodic_r16:
		// interface{} field of choice - nothing to decode
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
