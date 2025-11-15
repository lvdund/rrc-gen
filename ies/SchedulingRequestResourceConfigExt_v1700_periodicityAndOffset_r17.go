package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SchedulingRequestResourceConfigExt_v1700_periodicityAndOffset_r17_Choice_nothing uint64 = iota
	SchedulingRequestResourceConfigExt_v1700_periodicityAndOffset_r17_Choice_sl1280
	SchedulingRequestResourceConfigExt_v1700_periodicityAndOffset_r17_Choice_sl2560
	SchedulingRequestResourceConfigExt_v1700_periodicityAndOffset_r17_Choice_sl5120
)

type SchedulingRequestResourceConfigExt_v1700_periodicityAndOffset_r17 struct {
	Choice uint64
	sl1280 int64 `lb:0,ub:1279,madatory`
	sl2560 int64 `lb:0,ub:2559,madatory`
	sl5120 int64 `lb:0,ub:5119,madatory`
}

func (ie *SchedulingRequestResourceConfigExt_v1700_periodicityAndOffset_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 3, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SchedulingRequestResourceConfigExt_v1700_periodicityAndOffset_r17_Choice_sl1280:
		if err = w.WriteInteger(int64(ie.sl1280), &uper.Constraint{Lb: 0, Ub: 1279}, false); err != nil {
			err = utils.WrapError("Encode sl1280", err)
		}
	case SchedulingRequestResourceConfigExt_v1700_periodicityAndOffset_r17_Choice_sl2560:
		if err = w.WriteInteger(int64(ie.sl2560), &uper.Constraint{Lb: 0, Ub: 2559}, false); err != nil {
			err = utils.WrapError("Encode sl2560", err)
		}
	case SchedulingRequestResourceConfigExt_v1700_periodicityAndOffset_r17_Choice_sl5120:
		if err = w.WriteInteger(int64(ie.sl5120), &uper.Constraint{Lb: 0, Ub: 5119}, false); err != nil {
			err = utils.WrapError("Encode sl5120", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *SchedulingRequestResourceConfigExt_v1700_periodicityAndOffset_r17) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(3, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SchedulingRequestResourceConfigExt_v1700_periodicityAndOffset_r17_Choice_sl1280:
		var tmp_int_sl1280 int64
		if tmp_int_sl1280, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 1279}, false); err != nil {
			return utils.WrapError("Decode sl1280", err)
		}
		ie.sl1280 = tmp_int_sl1280
	case SchedulingRequestResourceConfigExt_v1700_periodicityAndOffset_r17_Choice_sl2560:
		var tmp_int_sl2560 int64
		if tmp_int_sl2560, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 2559}, false); err != nil {
			return utils.WrapError("Decode sl2560", err)
		}
		ie.sl2560 = tmp_int_sl2560
	case SchedulingRequestResourceConfigExt_v1700_periodicityAndOffset_r17_Choice_sl5120:
		var tmp_int_sl5120 int64
		if tmp_int_sl5120, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 5119}, false); err != nil {
			return utils.WrapError("Decode sl5120", err)
		}
		ie.sl5120 = tmp_int_sl5120
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
