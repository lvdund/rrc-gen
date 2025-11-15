package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	ServingCellConfig_ca_SlotOffset_r16_Choice_nothing uint64 = iota
	ServingCellConfig_ca_SlotOffset_r16_Choice_refSCS15kHz
	ServingCellConfig_ca_SlotOffset_r16_Choice_refSCS30KHz
	ServingCellConfig_ca_SlotOffset_r16_Choice_refSCS60KHz
	ServingCellConfig_ca_SlotOffset_r16_Choice_refSCS120KHz
)

type ServingCellConfig_ca_SlotOffset_r16 struct {
	Choice       uint64
	refSCS15kHz  int64 `lb:-2,ub:2,madatory`
	refSCS30KHz  int64 `lb:-5,ub:5,madatory`
	refSCS60KHz  int64 `lb:-10,ub:10,madatory`
	refSCS120KHz int64 `lb:-20,ub:20,madatory`
}

func (ie *ServingCellConfig_ca_SlotOffset_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 4, false); err != nil {
		return err
	}
	switch ie.Choice {
	case ServingCellConfig_ca_SlotOffset_r16_Choice_refSCS15kHz:
		if err = w.WriteInteger(int64(ie.refSCS15kHz), &uper.Constraint{Lb: -2, Ub: 2}, false); err != nil {
			err = utils.WrapError("Encode refSCS15kHz", err)
		}
	case ServingCellConfig_ca_SlotOffset_r16_Choice_refSCS30KHz:
		if err = w.WriteInteger(int64(ie.refSCS30KHz), &uper.Constraint{Lb: -5, Ub: 5}, false); err != nil {
			err = utils.WrapError("Encode refSCS30KHz", err)
		}
	case ServingCellConfig_ca_SlotOffset_r16_Choice_refSCS60KHz:
		if err = w.WriteInteger(int64(ie.refSCS60KHz), &uper.Constraint{Lb: -10, Ub: 10}, false); err != nil {
			err = utils.WrapError("Encode refSCS60KHz", err)
		}
	case ServingCellConfig_ca_SlotOffset_r16_Choice_refSCS120KHz:
		if err = w.WriteInteger(int64(ie.refSCS120KHz), &uper.Constraint{Lb: -20, Ub: 20}, false); err != nil {
			err = utils.WrapError("Encode refSCS120KHz", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *ServingCellConfig_ca_SlotOffset_r16) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(4, false); err != nil {
		return err
	}
	switch ie.Choice {
	case ServingCellConfig_ca_SlotOffset_r16_Choice_refSCS15kHz:
		var tmp_int_refSCS15kHz int64
		if tmp_int_refSCS15kHz, err = r.ReadInteger(&uper.Constraint{Lb: -2, Ub: 2}, false); err != nil {
			return utils.WrapError("Decode refSCS15kHz", err)
		}
		ie.refSCS15kHz = tmp_int_refSCS15kHz
	case ServingCellConfig_ca_SlotOffset_r16_Choice_refSCS30KHz:
		var tmp_int_refSCS30KHz int64
		if tmp_int_refSCS30KHz, err = r.ReadInteger(&uper.Constraint{Lb: -5, Ub: 5}, false); err != nil {
			return utils.WrapError("Decode refSCS30KHz", err)
		}
		ie.refSCS30KHz = tmp_int_refSCS30KHz
	case ServingCellConfig_ca_SlotOffset_r16_Choice_refSCS60KHz:
		var tmp_int_refSCS60KHz int64
		if tmp_int_refSCS60KHz, err = r.ReadInteger(&uper.Constraint{Lb: -10, Ub: 10}, false); err != nil {
			return utils.WrapError("Decode refSCS60KHz", err)
		}
		ie.refSCS60KHz = tmp_int_refSCS60KHz
	case ServingCellConfig_ca_SlotOffset_r16_Choice_refSCS120KHz:
		var tmp_int_refSCS120KHz int64
		if tmp_int_refSCS120KHz, err = r.ReadInteger(&uper.Constraint{Lb: -20, Ub: 20}, false); err != nil {
			return utils.WrapError("Decode refSCS120KHz", err)
		}
		ie.refSCS120KHz = tmp_int_refSCS120KHz
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
