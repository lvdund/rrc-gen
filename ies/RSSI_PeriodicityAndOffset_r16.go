package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	RSSI_PeriodicityAndOffset_r16_Choice_nothing uint64 = iota
	RSSI_PeriodicityAndOffset_r16_Choice_sl10
	RSSI_PeriodicityAndOffset_r16_Choice_sl20
	RSSI_PeriodicityAndOffset_r16_Choice_sl40
	RSSI_PeriodicityAndOffset_r16_Choice_sl80
	RSSI_PeriodicityAndOffset_r16_Choice_sl160
	RSSI_PeriodicityAndOffset_r16_Choice_sl320
	RSSI_PeriodicityAndOffset_r16_Choice_s1640
)

type RSSI_PeriodicityAndOffset_r16 struct {
	Choice uint64
	sl10   int64 `lb:0,ub:9,madatory`
	sl20   int64 `lb:0,ub:19,madatory`
	sl40   int64 `lb:0,ub:39,madatory`
	sl80   int64 `lb:0,ub:79,madatory`
	sl160  int64 `lb:0,ub:159,madatory`
	sl320  int64 `lb:0,ub:319,madatory`
	s1640  int64 `lb:0,ub:639,madatory`
}

func (ie *RSSI_PeriodicityAndOffset_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 7, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RSSI_PeriodicityAndOffset_r16_Choice_sl10:
		if err = w.WriteInteger(int64(ie.sl10), &uper.Constraint{Lb: 0, Ub: 9}, false); err != nil {
			err = utils.WrapError("Encode sl10", err)
		}
	case RSSI_PeriodicityAndOffset_r16_Choice_sl20:
		if err = w.WriteInteger(int64(ie.sl20), &uper.Constraint{Lb: 0, Ub: 19}, false); err != nil {
			err = utils.WrapError("Encode sl20", err)
		}
	case RSSI_PeriodicityAndOffset_r16_Choice_sl40:
		if err = w.WriteInteger(int64(ie.sl40), &uper.Constraint{Lb: 0, Ub: 39}, false); err != nil {
			err = utils.WrapError("Encode sl40", err)
		}
	case RSSI_PeriodicityAndOffset_r16_Choice_sl80:
		if err = w.WriteInteger(int64(ie.sl80), &uper.Constraint{Lb: 0, Ub: 79}, false); err != nil {
			err = utils.WrapError("Encode sl80", err)
		}
	case RSSI_PeriodicityAndOffset_r16_Choice_sl160:
		if err = w.WriteInteger(int64(ie.sl160), &uper.Constraint{Lb: 0, Ub: 159}, false); err != nil {
			err = utils.WrapError("Encode sl160", err)
		}
	case RSSI_PeriodicityAndOffset_r16_Choice_sl320:
		if err = w.WriteInteger(int64(ie.sl320), &uper.Constraint{Lb: 0, Ub: 319}, false); err != nil {
			err = utils.WrapError("Encode sl320", err)
		}
	case RSSI_PeriodicityAndOffset_r16_Choice_s1640:
		if err = w.WriteInteger(int64(ie.s1640), &uper.Constraint{Lb: 0, Ub: 639}, false); err != nil {
			err = utils.WrapError("Encode s1640", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *RSSI_PeriodicityAndOffset_r16) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(7, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RSSI_PeriodicityAndOffset_r16_Choice_sl10:
		var tmp_int_sl10 int64
		if tmp_int_sl10, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 9}, false); err != nil {
			return utils.WrapError("Decode sl10", err)
		}
		ie.sl10 = tmp_int_sl10
	case RSSI_PeriodicityAndOffset_r16_Choice_sl20:
		var tmp_int_sl20 int64
		if tmp_int_sl20, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 19}, false); err != nil {
			return utils.WrapError("Decode sl20", err)
		}
		ie.sl20 = tmp_int_sl20
	case RSSI_PeriodicityAndOffset_r16_Choice_sl40:
		var tmp_int_sl40 int64
		if tmp_int_sl40, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 39}, false); err != nil {
			return utils.WrapError("Decode sl40", err)
		}
		ie.sl40 = tmp_int_sl40
	case RSSI_PeriodicityAndOffset_r16_Choice_sl80:
		var tmp_int_sl80 int64
		if tmp_int_sl80, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 79}, false); err != nil {
			return utils.WrapError("Decode sl80", err)
		}
		ie.sl80 = tmp_int_sl80
	case RSSI_PeriodicityAndOffset_r16_Choice_sl160:
		var tmp_int_sl160 int64
		if tmp_int_sl160, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 159}, false); err != nil {
			return utils.WrapError("Decode sl160", err)
		}
		ie.sl160 = tmp_int_sl160
	case RSSI_PeriodicityAndOffset_r16_Choice_sl320:
		var tmp_int_sl320 int64
		if tmp_int_sl320, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 319}, false); err != nil {
			return utils.WrapError("Decode sl320", err)
		}
		ie.sl320 = tmp_int_sl320
	case RSSI_PeriodicityAndOffset_r16_Choice_s1640:
		var tmp_int_s1640 int64
		if tmp_int_s1640, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 639}, false); err != nil {
			return utils.WrapError("Decode s1640", err)
		}
		ie.s1640 = tmp_int_s1640
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
