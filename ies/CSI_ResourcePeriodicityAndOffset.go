package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	CSI_ResourcePeriodicityAndOffset_Choice_nothing uint64 = iota
	CSI_ResourcePeriodicityAndOffset_Choice_slots4
	CSI_ResourcePeriodicityAndOffset_Choice_slots5
	CSI_ResourcePeriodicityAndOffset_Choice_slots8
	CSI_ResourcePeriodicityAndOffset_Choice_slots10
	CSI_ResourcePeriodicityAndOffset_Choice_slots16
	CSI_ResourcePeriodicityAndOffset_Choice_slots20
	CSI_ResourcePeriodicityAndOffset_Choice_slots32
	CSI_ResourcePeriodicityAndOffset_Choice_slots40
	CSI_ResourcePeriodicityAndOffset_Choice_slots64
	CSI_ResourcePeriodicityAndOffset_Choice_slots80
	CSI_ResourcePeriodicityAndOffset_Choice_slots160
	CSI_ResourcePeriodicityAndOffset_Choice_slots320
	CSI_ResourcePeriodicityAndOffset_Choice_slots640
)

type CSI_ResourcePeriodicityAndOffset struct {
	Choice   uint64
	slots4   int64 `lb:0,ub:3,madatory`
	slots5   int64 `lb:0,ub:4,madatory`
	slots8   int64 `lb:0,ub:7,madatory`
	slots10  int64 `lb:0,ub:9,madatory`
	slots16  int64 `lb:0,ub:15,madatory`
	slots20  int64 `lb:0,ub:19,madatory`
	slots32  int64 `lb:0,ub:31,madatory`
	slots40  int64 `lb:0,ub:39,madatory`
	slots64  int64 `lb:0,ub:63,madatory`
	slots80  int64 `lb:0,ub:79,madatory`
	slots160 int64 `lb:0,ub:159,madatory`
	slots320 int64 `lb:0,ub:319,madatory`
	slots640 int64 `lb:0,ub:639,madatory`
}

func (ie *CSI_ResourcePeriodicityAndOffset) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 13, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CSI_ResourcePeriodicityAndOffset_Choice_slots4:
		if err = w.WriteInteger(int64(ie.slots4), &uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
			err = utils.WrapError("Encode slots4", err)
		}
	case CSI_ResourcePeriodicityAndOffset_Choice_slots5:
		if err = w.WriteInteger(int64(ie.slots5), &uper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
			err = utils.WrapError("Encode slots5", err)
		}
	case CSI_ResourcePeriodicityAndOffset_Choice_slots8:
		if err = w.WriteInteger(int64(ie.slots8), &uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
			err = utils.WrapError("Encode slots8", err)
		}
	case CSI_ResourcePeriodicityAndOffset_Choice_slots10:
		if err = w.WriteInteger(int64(ie.slots10), &uper.Constraint{Lb: 0, Ub: 9}, false); err != nil {
			err = utils.WrapError("Encode slots10", err)
		}
	case CSI_ResourcePeriodicityAndOffset_Choice_slots16:
		if err = w.WriteInteger(int64(ie.slots16), &uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
			err = utils.WrapError("Encode slots16", err)
		}
	case CSI_ResourcePeriodicityAndOffset_Choice_slots20:
		if err = w.WriteInteger(int64(ie.slots20), &uper.Constraint{Lb: 0, Ub: 19}, false); err != nil {
			err = utils.WrapError("Encode slots20", err)
		}
	case CSI_ResourcePeriodicityAndOffset_Choice_slots32:
		if err = w.WriteInteger(int64(ie.slots32), &uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
			err = utils.WrapError("Encode slots32", err)
		}
	case CSI_ResourcePeriodicityAndOffset_Choice_slots40:
		if err = w.WriteInteger(int64(ie.slots40), &uper.Constraint{Lb: 0, Ub: 39}, false); err != nil {
			err = utils.WrapError("Encode slots40", err)
		}
	case CSI_ResourcePeriodicityAndOffset_Choice_slots64:
		if err = w.WriteInteger(int64(ie.slots64), &uper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
			err = utils.WrapError("Encode slots64", err)
		}
	case CSI_ResourcePeriodicityAndOffset_Choice_slots80:
		if err = w.WriteInteger(int64(ie.slots80), &uper.Constraint{Lb: 0, Ub: 79}, false); err != nil {
			err = utils.WrapError("Encode slots80", err)
		}
	case CSI_ResourcePeriodicityAndOffset_Choice_slots160:
		if err = w.WriteInteger(int64(ie.slots160), &uper.Constraint{Lb: 0, Ub: 159}, false); err != nil {
			err = utils.WrapError("Encode slots160", err)
		}
	case CSI_ResourcePeriodicityAndOffset_Choice_slots320:
		if err = w.WriteInteger(int64(ie.slots320), &uper.Constraint{Lb: 0, Ub: 319}, false); err != nil {
			err = utils.WrapError("Encode slots320", err)
		}
	case CSI_ResourcePeriodicityAndOffset_Choice_slots640:
		if err = w.WriteInteger(int64(ie.slots640), &uper.Constraint{Lb: 0, Ub: 639}, false); err != nil {
			err = utils.WrapError("Encode slots640", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *CSI_ResourcePeriodicityAndOffset) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(13, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CSI_ResourcePeriodicityAndOffset_Choice_slots4:
		var tmp_int_slots4 int64
		if tmp_int_slots4, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
			return utils.WrapError("Decode slots4", err)
		}
		ie.slots4 = tmp_int_slots4
	case CSI_ResourcePeriodicityAndOffset_Choice_slots5:
		var tmp_int_slots5 int64
		if tmp_int_slots5, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
			return utils.WrapError("Decode slots5", err)
		}
		ie.slots5 = tmp_int_slots5
	case CSI_ResourcePeriodicityAndOffset_Choice_slots8:
		var tmp_int_slots8 int64
		if tmp_int_slots8, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
			return utils.WrapError("Decode slots8", err)
		}
		ie.slots8 = tmp_int_slots8
	case CSI_ResourcePeriodicityAndOffset_Choice_slots10:
		var tmp_int_slots10 int64
		if tmp_int_slots10, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 9}, false); err != nil {
			return utils.WrapError("Decode slots10", err)
		}
		ie.slots10 = tmp_int_slots10
	case CSI_ResourcePeriodicityAndOffset_Choice_slots16:
		var tmp_int_slots16 int64
		if tmp_int_slots16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
			return utils.WrapError("Decode slots16", err)
		}
		ie.slots16 = tmp_int_slots16
	case CSI_ResourcePeriodicityAndOffset_Choice_slots20:
		var tmp_int_slots20 int64
		if tmp_int_slots20, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 19}, false); err != nil {
			return utils.WrapError("Decode slots20", err)
		}
		ie.slots20 = tmp_int_slots20
	case CSI_ResourcePeriodicityAndOffset_Choice_slots32:
		var tmp_int_slots32 int64
		if tmp_int_slots32, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
			return utils.WrapError("Decode slots32", err)
		}
		ie.slots32 = tmp_int_slots32
	case CSI_ResourcePeriodicityAndOffset_Choice_slots40:
		var tmp_int_slots40 int64
		if tmp_int_slots40, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 39}, false); err != nil {
			return utils.WrapError("Decode slots40", err)
		}
		ie.slots40 = tmp_int_slots40
	case CSI_ResourcePeriodicityAndOffset_Choice_slots64:
		var tmp_int_slots64 int64
		if tmp_int_slots64, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
			return utils.WrapError("Decode slots64", err)
		}
		ie.slots64 = tmp_int_slots64
	case CSI_ResourcePeriodicityAndOffset_Choice_slots80:
		var tmp_int_slots80 int64
		if tmp_int_slots80, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 79}, false); err != nil {
			return utils.WrapError("Decode slots80", err)
		}
		ie.slots80 = tmp_int_slots80
	case CSI_ResourcePeriodicityAndOffset_Choice_slots160:
		var tmp_int_slots160 int64
		if tmp_int_slots160, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 159}, false); err != nil {
			return utils.WrapError("Decode slots160", err)
		}
		ie.slots160 = tmp_int_slots160
	case CSI_ResourcePeriodicityAndOffset_Choice_slots320:
		var tmp_int_slots320 int64
		if tmp_int_slots320, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 319}, false); err != nil {
			return utils.WrapError("Decode slots320", err)
		}
		ie.slots320 = tmp_int_slots320
	case CSI_ResourcePeriodicityAndOffset_Choice_slots640:
		var tmp_int_slots640 int64
		if tmp_int_slots640, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 639}, false); err != nil {
			return utils.WrapError("Decode slots640", err)
		}
		ie.slots640 = tmp_int_slots640
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
