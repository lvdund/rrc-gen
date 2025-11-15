package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	CarrierState_r17_Choice_nothing uint64 = iota
	CarrierState_r17_Choice_deActivated_r17
	CarrierState_r17_Choice_activeBWP_r17
)

type CarrierState_r17 struct {
	Choice          uint64
	deActivated_r17 uper.NULL `madatory`
	activeBWP_r17   int64     `lb:0,ub:maxNrofBWPs,madatory`
}

func (ie *CarrierState_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CarrierState_r17_Choice_deActivated_r17:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode deActivated_r17", err)
		}
	case CarrierState_r17_Choice_activeBWP_r17:
		if err = w.WriteInteger(int64(ie.activeBWP_r17), &uper.Constraint{Lb: 0, Ub: maxNrofBWPs}, false); err != nil {
			err = utils.WrapError("Encode activeBWP_r17", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *CarrierState_r17) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CarrierState_r17_Choice_deActivated_r17:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode deActivated_r17", err)
		}
	case CarrierState_r17_Choice_activeBWP_r17:
		var tmp_int_activeBWP_r17 int64
		if tmp_int_activeBWP_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: maxNrofBWPs}, false); err != nil {
			return utils.WrapError("Decode activeBWP_r17", err)
		}
		ie.activeBWP_r17 = tmp_int_activeBWP_r17
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
