package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	TRS_ResourceSet_r17_periodicityAndOffset_r17_Choice_nothing uint64 = iota
	TRS_ResourceSet_r17_periodicityAndOffset_r17_Choice_slots10
	TRS_ResourceSet_r17_periodicityAndOffset_r17_Choice_slots20
	TRS_ResourceSet_r17_periodicityAndOffset_r17_Choice_slots40
	TRS_ResourceSet_r17_periodicityAndOffset_r17_Choice_slots80
)

type TRS_ResourceSet_r17_periodicityAndOffset_r17 struct {
	Choice  uint64
	slots10 int64 `lb:0,ub:9,madatory`
	slots20 int64 `lb:0,ub:19,madatory`
	slots40 int64 `lb:0,ub:39,madatory`
	slots80 int64 `lb:0,ub:79,madatory`
}

func (ie *TRS_ResourceSet_r17_periodicityAndOffset_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 4, false); err != nil {
		return err
	}
	switch ie.Choice {
	case TRS_ResourceSet_r17_periodicityAndOffset_r17_Choice_slots10:
		if err = w.WriteInteger(int64(ie.slots10), &uper.Constraint{Lb: 0, Ub: 9}, false); err != nil {
			err = utils.WrapError("Encode slots10", err)
		}
	case TRS_ResourceSet_r17_periodicityAndOffset_r17_Choice_slots20:
		if err = w.WriteInteger(int64(ie.slots20), &uper.Constraint{Lb: 0, Ub: 19}, false); err != nil {
			err = utils.WrapError("Encode slots20", err)
		}
	case TRS_ResourceSet_r17_periodicityAndOffset_r17_Choice_slots40:
		if err = w.WriteInteger(int64(ie.slots40), &uper.Constraint{Lb: 0, Ub: 39}, false); err != nil {
			err = utils.WrapError("Encode slots40", err)
		}
	case TRS_ResourceSet_r17_periodicityAndOffset_r17_Choice_slots80:
		if err = w.WriteInteger(int64(ie.slots80), &uper.Constraint{Lb: 0, Ub: 79}, false); err != nil {
			err = utils.WrapError("Encode slots80", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *TRS_ResourceSet_r17_periodicityAndOffset_r17) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(4, false); err != nil {
		return err
	}
	switch ie.Choice {
	case TRS_ResourceSet_r17_periodicityAndOffset_r17_Choice_slots10:
		var tmp_int_slots10 int64
		if tmp_int_slots10, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 9}, false); err != nil {
			return utils.WrapError("Decode slots10", err)
		}
		ie.slots10 = tmp_int_slots10
	case TRS_ResourceSet_r17_periodicityAndOffset_r17_Choice_slots20:
		var tmp_int_slots20 int64
		if tmp_int_slots20, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 19}, false); err != nil {
			return utils.WrapError("Decode slots20", err)
		}
		ie.slots20 = tmp_int_slots20
	case TRS_ResourceSet_r17_periodicityAndOffset_r17_Choice_slots40:
		var tmp_int_slots40 int64
		if tmp_int_slots40, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 39}, false); err != nil {
			return utils.WrapError("Decode slots40", err)
		}
		ie.slots40 = tmp_int_slots40
	case TRS_ResourceSet_r17_periodicityAndOffset_r17_Choice_slots80:
		var tmp_int_slots80 int64
		if tmp_int_slots80, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 79}, false); err != nil {
			return utils.WrapError("Decode slots80", err)
		}
		ie.slots80 = tmp_int_slots80
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
