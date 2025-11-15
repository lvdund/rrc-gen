package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	MeasTriggerQuantityOffset_Choice_nothing uint64 = iota
	MeasTriggerQuantityOffset_Choice_rsrp
	MeasTriggerQuantityOffset_Choice_rsrq
	MeasTriggerQuantityOffset_Choice_sinr
)

type MeasTriggerQuantityOffset struct {
	Choice uint64
	rsrp   int64 `lb:-30,ub:30,madatory`
	rsrq   int64 `lb:-30,ub:30,madatory`
	sinr   int64 `lb:-30,ub:30,madatory`
}

func (ie *MeasTriggerQuantityOffset) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 3, false); err != nil {
		return err
	}
	switch ie.Choice {
	case MeasTriggerQuantityOffset_Choice_rsrp:
		if err = w.WriteInteger(int64(ie.rsrp), &uper.Constraint{Lb: -30, Ub: 30}, false); err != nil {
			err = utils.WrapError("Encode rsrp", err)
		}
	case MeasTriggerQuantityOffset_Choice_rsrq:
		if err = w.WriteInteger(int64(ie.rsrq), &uper.Constraint{Lb: -30, Ub: 30}, false); err != nil {
			err = utils.WrapError("Encode rsrq", err)
		}
	case MeasTriggerQuantityOffset_Choice_sinr:
		if err = w.WriteInteger(int64(ie.sinr), &uper.Constraint{Lb: -30, Ub: 30}, false); err != nil {
			err = utils.WrapError("Encode sinr", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *MeasTriggerQuantityOffset) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(3, false); err != nil {
		return err
	}
	switch ie.Choice {
	case MeasTriggerQuantityOffset_Choice_rsrp:
		var tmp_int_rsrp int64
		if tmp_int_rsrp, err = r.ReadInteger(&uper.Constraint{Lb: -30, Ub: 30}, false); err != nil {
			return utils.WrapError("Decode rsrp", err)
		}
		ie.rsrp = tmp_int_rsrp
	case MeasTriggerQuantityOffset_Choice_rsrq:
		var tmp_int_rsrq int64
		if tmp_int_rsrq, err = r.ReadInteger(&uper.Constraint{Lb: -30, Ub: 30}, false); err != nil {
			return utils.WrapError("Decode rsrq", err)
		}
		ie.rsrq = tmp_int_rsrq
	case MeasTriggerQuantityOffset_Choice_sinr:
		var tmp_int_sinr int64
		if tmp_int_sinr, err = r.ReadInteger(&uper.Constraint{Lb: -30, Ub: 30}, false); err != nil {
			return utils.WrapError("Decode sinr", err)
		}
		ie.sinr = tmp_int_sinr
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
