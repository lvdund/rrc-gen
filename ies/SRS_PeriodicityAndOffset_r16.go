package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SRS_PeriodicityAndOffset_r16_Choice_nothing uint64 = iota
	SRS_PeriodicityAndOffset_r16_Choice_sl1
	SRS_PeriodicityAndOffset_r16_Choice_sl2
	SRS_PeriodicityAndOffset_r16_Choice_sl4
	SRS_PeriodicityAndOffset_r16_Choice_sl5
	SRS_PeriodicityAndOffset_r16_Choice_sl8
	SRS_PeriodicityAndOffset_r16_Choice_sl10
	SRS_PeriodicityAndOffset_r16_Choice_sl16
	SRS_PeriodicityAndOffset_r16_Choice_sl20
	SRS_PeriodicityAndOffset_r16_Choice_sl32
	SRS_PeriodicityAndOffset_r16_Choice_sl40
	SRS_PeriodicityAndOffset_r16_Choice_sl64
	SRS_PeriodicityAndOffset_r16_Choice_sl80
	SRS_PeriodicityAndOffset_r16_Choice_sl160
	SRS_PeriodicityAndOffset_r16_Choice_sl320
	SRS_PeriodicityAndOffset_r16_Choice_sl640
	SRS_PeriodicityAndOffset_r16_Choice_sl1280
	SRS_PeriodicityAndOffset_r16_Choice_sl2560
	SRS_PeriodicityAndOffset_r16_Choice_sl5120
	SRS_PeriodicityAndOffset_r16_Choice_sl10240
	SRS_PeriodicityAndOffset_r16_Choice_sl40960
	SRS_PeriodicityAndOffset_r16_Choice_sl81920
)

type SRS_PeriodicityAndOffset_r16 struct {
	Choice  uint64
	sl1     uper.NULL `madatory`
	sl2     int64     `lb:0,ub:1,madatory`
	sl4     int64     `lb:0,ub:3,madatory`
	sl5     int64     `lb:0,ub:4,madatory`
	sl8     int64     `lb:0,ub:7,madatory`
	sl10    int64     `lb:0,ub:9,madatory`
	sl16    int64     `lb:0,ub:15,madatory`
	sl20    int64     `lb:0,ub:19,madatory`
	sl32    int64     `lb:0,ub:31,madatory`
	sl40    int64     `lb:0,ub:39,madatory`
	sl64    int64     `lb:0,ub:63,madatory`
	sl80    int64     `lb:0,ub:79,madatory`
	sl160   int64     `lb:0,ub:159,madatory`
	sl320   int64     `lb:0,ub:319,madatory`
	sl640   int64     `lb:0,ub:639,madatory`
	sl1280  int64     `lb:0,ub:1279,madatory`
	sl2560  int64     `lb:0,ub:2559,madatory`
	sl5120  int64     `lb:0,ub:5119,madatory`
	sl10240 int64     `lb:0,ub:10239,madatory`
	sl40960 int64     `lb:0,ub:40959,madatory`
	sl81920 int64     `lb:0,ub:81919,madatory`
}

func (ie *SRS_PeriodicityAndOffset_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 21, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SRS_PeriodicityAndOffset_r16_Choice_sl1:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode sl1", err)
		}
	case SRS_PeriodicityAndOffset_r16_Choice_sl2:
		if err = w.WriteInteger(int64(ie.sl2), &uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
			err = utils.WrapError("Encode sl2", err)
		}
	case SRS_PeriodicityAndOffset_r16_Choice_sl4:
		if err = w.WriteInteger(int64(ie.sl4), &uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
			err = utils.WrapError("Encode sl4", err)
		}
	case SRS_PeriodicityAndOffset_r16_Choice_sl5:
		if err = w.WriteInteger(int64(ie.sl5), &uper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
			err = utils.WrapError("Encode sl5", err)
		}
	case SRS_PeriodicityAndOffset_r16_Choice_sl8:
		if err = w.WriteInteger(int64(ie.sl8), &uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
			err = utils.WrapError("Encode sl8", err)
		}
	case SRS_PeriodicityAndOffset_r16_Choice_sl10:
		if err = w.WriteInteger(int64(ie.sl10), &uper.Constraint{Lb: 0, Ub: 9}, false); err != nil {
			err = utils.WrapError("Encode sl10", err)
		}
	case SRS_PeriodicityAndOffset_r16_Choice_sl16:
		if err = w.WriteInteger(int64(ie.sl16), &uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
			err = utils.WrapError("Encode sl16", err)
		}
	case SRS_PeriodicityAndOffset_r16_Choice_sl20:
		if err = w.WriteInteger(int64(ie.sl20), &uper.Constraint{Lb: 0, Ub: 19}, false); err != nil {
			err = utils.WrapError("Encode sl20", err)
		}
	case SRS_PeriodicityAndOffset_r16_Choice_sl32:
		if err = w.WriteInteger(int64(ie.sl32), &uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
			err = utils.WrapError("Encode sl32", err)
		}
	case SRS_PeriodicityAndOffset_r16_Choice_sl40:
		if err = w.WriteInteger(int64(ie.sl40), &uper.Constraint{Lb: 0, Ub: 39}, false); err != nil {
			err = utils.WrapError("Encode sl40", err)
		}
	case SRS_PeriodicityAndOffset_r16_Choice_sl64:
		if err = w.WriteInteger(int64(ie.sl64), &uper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
			err = utils.WrapError("Encode sl64", err)
		}
	case SRS_PeriodicityAndOffset_r16_Choice_sl80:
		if err = w.WriteInteger(int64(ie.sl80), &uper.Constraint{Lb: 0, Ub: 79}, false); err != nil {
			err = utils.WrapError("Encode sl80", err)
		}
	case SRS_PeriodicityAndOffset_r16_Choice_sl160:
		if err = w.WriteInteger(int64(ie.sl160), &uper.Constraint{Lb: 0, Ub: 159}, false); err != nil {
			err = utils.WrapError("Encode sl160", err)
		}
	case SRS_PeriodicityAndOffset_r16_Choice_sl320:
		if err = w.WriteInteger(int64(ie.sl320), &uper.Constraint{Lb: 0, Ub: 319}, false); err != nil {
			err = utils.WrapError("Encode sl320", err)
		}
	case SRS_PeriodicityAndOffset_r16_Choice_sl640:
		if err = w.WriteInteger(int64(ie.sl640), &uper.Constraint{Lb: 0, Ub: 639}, false); err != nil {
			err = utils.WrapError("Encode sl640", err)
		}
	case SRS_PeriodicityAndOffset_r16_Choice_sl1280:
		if err = w.WriteInteger(int64(ie.sl1280), &uper.Constraint{Lb: 0, Ub: 1279}, false); err != nil {
			err = utils.WrapError("Encode sl1280", err)
		}
	case SRS_PeriodicityAndOffset_r16_Choice_sl2560:
		if err = w.WriteInteger(int64(ie.sl2560), &uper.Constraint{Lb: 0, Ub: 2559}, false); err != nil {
			err = utils.WrapError("Encode sl2560", err)
		}
	case SRS_PeriodicityAndOffset_r16_Choice_sl5120:
		if err = w.WriteInteger(int64(ie.sl5120), &uper.Constraint{Lb: 0, Ub: 5119}, false); err != nil {
			err = utils.WrapError("Encode sl5120", err)
		}
	case SRS_PeriodicityAndOffset_r16_Choice_sl10240:
		if err = w.WriteInteger(int64(ie.sl10240), &uper.Constraint{Lb: 0, Ub: 10239}, false); err != nil {
			err = utils.WrapError("Encode sl10240", err)
		}
	case SRS_PeriodicityAndOffset_r16_Choice_sl40960:
		if err = w.WriteInteger(int64(ie.sl40960), &uper.Constraint{Lb: 0, Ub: 40959}, false); err != nil {
			err = utils.WrapError("Encode sl40960", err)
		}
	case SRS_PeriodicityAndOffset_r16_Choice_sl81920:
		if err = w.WriteInteger(int64(ie.sl81920), &uper.Constraint{Lb: 0, Ub: 81919}, false); err != nil {
			err = utils.WrapError("Encode sl81920", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *SRS_PeriodicityAndOffset_r16) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(21, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SRS_PeriodicityAndOffset_r16_Choice_sl1:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode sl1", err)
		}
	case SRS_PeriodicityAndOffset_r16_Choice_sl2:
		var tmp_int_sl2 int64
		if tmp_int_sl2, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
			return utils.WrapError("Decode sl2", err)
		}
		ie.sl2 = tmp_int_sl2
	case SRS_PeriodicityAndOffset_r16_Choice_sl4:
		var tmp_int_sl4 int64
		if tmp_int_sl4, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
			return utils.WrapError("Decode sl4", err)
		}
		ie.sl4 = tmp_int_sl4
	case SRS_PeriodicityAndOffset_r16_Choice_sl5:
		var tmp_int_sl5 int64
		if tmp_int_sl5, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
			return utils.WrapError("Decode sl5", err)
		}
		ie.sl5 = tmp_int_sl5
	case SRS_PeriodicityAndOffset_r16_Choice_sl8:
		var tmp_int_sl8 int64
		if tmp_int_sl8, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
			return utils.WrapError("Decode sl8", err)
		}
		ie.sl8 = tmp_int_sl8
	case SRS_PeriodicityAndOffset_r16_Choice_sl10:
		var tmp_int_sl10 int64
		if tmp_int_sl10, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 9}, false); err != nil {
			return utils.WrapError("Decode sl10", err)
		}
		ie.sl10 = tmp_int_sl10
	case SRS_PeriodicityAndOffset_r16_Choice_sl16:
		var tmp_int_sl16 int64
		if tmp_int_sl16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
			return utils.WrapError("Decode sl16", err)
		}
		ie.sl16 = tmp_int_sl16
	case SRS_PeriodicityAndOffset_r16_Choice_sl20:
		var tmp_int_sl20 int64
		if tmp_int_sl20, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 19}, false); err != nil {
			return utils.WrapError("Decode sl20", err)
		}
		ie.sl20 = tmp_int_sl20
	case SRS_PeriodicityAndOffset_r16_Choice_sl32:
		var tmp_int_sl32 int64
		if tmp_int_sl32, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
			return utils.WrapError("Decode sl32", err)
		}
		ie.sl32 = tmp_int_sl32
	case SRS_PeriodicityAndOffset_r16_Choice_sl40:
		var tmp_int_sl40 int64
		if tmp_int_sl40, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 39}, false); err != nil {
			return utils.WrapError("Decode sl40", err)
		}
		ie.sl40 = tmp_int_sl40
	case SRS_PeriodicityAndOffset_r16_Choice_sl64:
		var tmp_int_sl64 int64
		if tmp_int_sl64, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
			return utils.WrapError("Decode sl64", err)
		}
		ie.sl64 = tmp_int_sl64
	case SRS_PeriodicityAndOffset_r16_Choice_sl80:
		var tmp_int_sl80 int64
		if tmp_int_sl80, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 79}, false); err != nil {
			return utils.WrapError("Decode sl80", err)
		}
		ie.sl80 = tmp_int_sl80
	case SRS_PeriodicityAndOffset_r16_Choice_sl160:
		var tmp_int_sl160 int64
		if tmp_int_sl160, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 159}, false); err != nil {
			return utils.WrapError("Decode sl160", err)
		}
		ie.sl160 = tmp_int_sl160
	case SRS_PeriodicityAndOffset_r16_Choice_sl320:
		var tmp_int_sl320 int64
		if tmp_int_sl320, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 319}, false); err != nil {
			return utils.WrapError("Decode sl320", err)
		}
		ie.sl320 = tmp_int_sl320
	case SRS_PeriodicityAndOffset_r16_Choice_sl640:
		var tmp_int_sl640 int64
		if tmp_int_sl640, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 639}, false); err != nil {
			return utils.WrapError("Decode sl640", err)
		}
		ie.sl640 = tmp_int_sl640
	case SRS_PeriodicityAndOffset_r16_Choice_sl1280:
		var tmp_int_sl1280 int64
		if tmp_int_sl1280, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 1279}, false); err != nil {
			return utils.WrapError("Decode sl1280", err)
		}
		ie.sl1280 = tmp_int_sl1280
	case SRS_PeriodicityAndOffset_r16_Choice_sl2560:
		var tmp_int_sl2560 int64
		if tmp_int_sl2560, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 2559}, false); err != nil {
			return utils.WrapError("Decode sl2560", err)
		}
		ie.sl2560 = tmp_int_sl2560
	case SRS_PeriodicityAndOffset_r16_Choice_sl5120:
		var tmp_int_sl5120 int64
		if tmp_int_sl5120, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 5119}, false); err != nil {
			return utils.WrapError("Decode sl5120", err)
		}
		ie.sl5120 = tmp_int_sl5120
	case SRS_PeriodicityAndOffset_r16_Choice_sl10240:
		var tmp_int_sl10240 int64
		if tmp_int_sl10240, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 10239}, false); err != nil {
			return utils.WrapError("Decode sl10240", err)
		}
		ie.sl10240 = tmp_int_sl10240
	case SRS_PeriodicityAndOffset_r16_Choice_sl40960:
		var tmp_int_sl40960 int64
		if tmp_int_sl40960, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 40959}, false); err != nil {
			return utils.WrapError("Decode sl40960", err)
		}
		ie.sl40960 = tmp_int_sl40960
	case SRS_PeriodicityAndOffset_r16_Choice_sl81920:
		var tmp_int_sl81920 int64
		if tmp_int_sl81920, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 81919}, false); err != nil {
			return utils.WrapError("Decode sl81920", err)
		}
		ie.sl81920 = tmp_int_sl81920
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
