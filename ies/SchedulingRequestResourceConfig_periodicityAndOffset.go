package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SchedulingRequestResourceConfig_periodicityAndOffset_Choice_nothing uint64 = iota
	SchedulingRequestResourceConfig_periodicityAndOffset_Choice_sym2
	SchedulingRequestResourceConfig_periodicityAndOffset_Choice_sym6or7
	SchedulingRequestResourceConfig_periodicityAndOffset_Choice_sl1
	SchedulingRequestResourceConfig_periodicityAndOffset_Choice_sl2
	SchedulingRequestResourceConfig_periodicityAndOffset_Choice_sl4
	SchedulingRequestResourceConfig_periodicityAndOffset_Choice_sl5
	SchedulingRequestResourceConfig_periodicityAndOffset_Choice_sl8
	SchedulingRequestResourceConfig_periodicityAndOffset_Choice_sl10
	SchedulingRequestResourceConfig_periodicityAndOffset_Choice_sl16
	SchedulingRequestResourceConfig_periodicityAndOffset_Choice_sl20
	SchedulingRequestResourceConfig_periodicityAndOffset_Choice_sl40
	SchedulingRequestResourceConfig_periodicityAndOffset_Choice_sl80
	SchedulingRequestResourceConfig_periodicityAndOffset_Choice_sl160
	SchedulingRequestResourceConfig_periodicityAndOffset_Choice_sl320
	SchedulingRequestResourceConfig_periodicityAndOffset_Choice_sl640
)

type SchedulingRequestResourceConfig_periodicityAndOffset struct {
	Choice  uint64
	sym2    uper.NULL `madatory`
	sym6or7 uper.NULL `madatory`
	sl1     uper.NULL `madatory`
	sl2     int64     `lb:0,ub:1,madatory`
	sl4     int64     `lb:0,ub:3,madatory`
	sl5     int64     `lb:0,ub:4,madatory`
	sl8     int64     `lb:0,ub:7,madatory`
	sl10    int64     `lb:0,ub:9,madatory`
	sl16    int64     `lb:0,ub:15,madatory`
	sl20    int64     `lb:0,ub:19,madatory`
	sl40    int64     `lb:0,ub:39,madatory`
	sl80    int64     `lb:0,ub:79,madatory`
	sl160   int64     `lb:0,ub:159,madatory`
	sl320   int64     `lb:0,ub:319,madatory`
	sl640   int64     `lb:0,ub:639,madatory`
}

func (ie *SchedulingRequestResourceConfig_periodicityAndOffset) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 15, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SchedulingRequestResourceConfig_periodicityAndOffset_Choice_sym2:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode sym2", err)
		}
	case SchedulingRequestResourceConfig_periodicityAndOffset_Choice_sym6or7:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode sym6or7", err)
		}
	case SchedulingRequestResourceConfig_periodicityAndOffset_Choice_sl1:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode sl1", err)
		}
	case SchedulingRequestResourceConfig_periodicityAndOffset_Choice_sl2:
		if err = w.WriteInteger(int64(ie.sl2), &uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
			err = utils.WrapError("Encode sl2", err)
		}
	case SchedulingRequestResourceConfig_periodicityAndOffset_Choice_sl4:
		if err = w.WriteInteger(int64(ie.sl4), &uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
			err = utils.WrapError("Encode sl4", err)
		}
	case SchedulingRequestResourceConfig_periodicityAndOffset_Choice_sl5:
		if err = w.WriteInteger(int64(ie.sl5), &uper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
			err = utils.WrapError("Encode sl5", err)
		}
	case SchedulingRequestResourceConfig_periodicityAndOffset_Choice_sl8:
		if err = w.WriteInteger(int64(ie.sl8), &uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
			err = utils.WrapError("Encode sl8", err)
		}
	case SchedulingRequestResourceConfig_periodicityAndOffset_Choice_sl10:
		if err = w.WriteInteger(int64(ie.sl10), &uper.Constraint{Lb: 0, Ub: 9}, false); err != nil {
			err = utils.WrapError("Encode sl10", err)
		}
	case SchedulingRequestResourceConfig_periodicityAndOffset_Choice_sl16:
		if err = w.WriteInteger(int64(ie.sl16), &uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
			err = utils.WrapError("Encode sl16", err)
		}
	case SchedulingRequestResourceConfig_periodicityAndOffset_Choice_sl20:
		if err = w.WriteInteger(int64(ie.sl20), &uper.Constraint{Lb: 0, Ub: 19}, false); err != nil {
			err = utils.WrapError("Encode sl20", err)
		}
	case SchedulingRequestResourceConfig_periodicityAndOffset_Choice_sl40:
		if err = w.WriteInteger(int64(ie.sl40), &uper.Constraint{Lb: 0, Ub: 39}, false); err != nil {
			err = utils.WrapError("Encode sl40", err)
		}
	case SchedulingRequestResourceConfig_periodicityAndOffset_Choice_sl80:
		if err = w.WriteInteger(int64(ie.sl80), &uper.Constraint{Lb: 0, Ub: 79}, false); err != nil {
			err = utils.WrapError("Encode sl80", err)
		}
	case SchedulingRequestResourceConfig_periodicityAndOffset_Choice_sl160:
		if err = w.WriteInteger(int64(ie.sl160), &uper.Constraint{Lb: 0, Ub: 159}, false); err != nil {
			err = utils.WrapError("Encode sl160", err)
		}
	case SchedulingRequestResourceConfig_periodicityAndOffset_Choice_sl320:
		if err = w.WriteInteger(int64(ie.sl320), &uper.Constraint{Lb: 0, Ub: 319}, false); err != nil {
			err = utils.WrapError("Encode sl320", err)
		}
	case SchedulingRequestResourceConfig_periodicityAndOffset_Choice_sl640:
		if err = w.WriteInteger(int64(ie.sl640), &uper.Constraint{Lb: 0, Ub: 639}, false); err != nil {
			err = utils.WrapError("Encode sl640", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *SchedulingRequestResourceConfig_periodicityAndOffset) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(15, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SchedulingRequestResourceConfig_periodicityAndOffset_Choice_sym2:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode sym2", err)
		}
	case SchedulingRequestResourceConfig_periodicityAndOffset_Choice_sym6or7:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode sym6or7", err)
		}
	case SchedulingRequestResourceConfig_periodicityAndOffset_Choice_sl1:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode sl1", err)
		}
	case SchedulingRequestResourceConfig_periodicityAndOffset_Choice_sl2:
		var tmp_int_sl2 int64
		if tmp_int_sl2, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
			return utils.WrapError("Decode sl2", err)
		}
		ie.sl2 = tmp_int_sl2
	case SchedulingRequestResourceConfig_periodicityAndOffset_Choice_sl4:
		var tmp_int_sl4 int64
		if tmp_int_sl4, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
			return utils.WrapError("Decode sl4", err)
		}
		ie.sl4 = tmp_int_sl4
	case SchedulingRequestResourceConfig_periodicityAndOffset_Choice_sl5:
		var tmp_int_sl5 int64
		if tmp_int_sl5, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
			return utils.WrapError("Decode sl5", err)
		}
		ie.sl5 = tmp_int_sl5
	case SchedulingRequestResourceConfig_periodicityAndOffset_Choice_sl8:
		var tmp_int_sl8 int64
		if tmp_int_sl8, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
			return utils.WrapError("Decode sl8", err)
		}
		ie.sl8 = tmp_int_sl8
	case SchedulingRequestResourceConfig_periodicityAndOffset_Choice_sl10:
		var tmp_int_sl10 int64
		if tmp_int_sl10, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 9}, false); err != nil {
			return utils.WrapError("Decode sl10", err)
		}
		ie.sl10 = tmp_int_sl10
	case SchedulingRequestResourceConfig_periodicityAndOffset_Choice_sl16:
		var tmp_int_sl16 int64
		if tmp_int_sl16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
			return utils.WrapError("Decode sl16", err)
		}
		ie.sl16 = tmp_int_sl16
	case SchedulingRequestResourceConfig_periodicityAndOffset_Choice_sl20:
		var tmp_int_sl20 int64
		if tmp_int_sl20, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 19}, false); err != nil {
			return utils.WrapError("Decode sl20", err)
		}
		ie.sl20 = tmp_int_sl20
	case SchedulingRequestResourceConfig_periodicityAndOffset_Choice_sl40:
		var tmp_int_sl40 int64
		if tmp_int_sl40, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 39}, false); err != nil {
			return utils.WrapError("Decode sl40", err)
		}
		ie.sl40 = tmp_int_sl40
	case SchedulingRequestResourceConfig_periodicityAndOffset_Choice_sl80:
		var tmp_int_sl80 int64
		if tmp_int_sl80, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 79}, false); err != nil {
			return utils.WrapError("Decode sl80", err)
		}
		ie.sl80 = tmp_int_sl80
	case SchedulingRequestResourceConfig_periodicityAndOffset_Choice_sl160:
		var tmp_int_sl160 int64
		if tmp_int_sl160, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 159}, false); err != nil {
			return utils.WrapError("Decode sl160", err)
		}
		ie.sl160 = tmp_int_sl160
	case SchedulingRequestResourceConfig_periodicityAndOffset_Choice_sl320:
		var tmp_int_sl320 int64
		if tmp_int_sl320, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 319}, false); err != nil {
			return utils.WrapError("Decode sl320", err)
		}
		ie.sl320 = tmp_int_sl320
	case SchedulingRequestResourceConfig_periodicityAndOffset_Choice_sl640:
		var tmp_int_sl640 int64
		if tmp_int_sl640, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 639}, false); err != nil {
			return utils.WrapError("Decode sl640", err)
		}
		ie.sl640 = tmp_int_sl640
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
