package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SearchSpaceExt_v1700_monitoringSlotPeriodicityAndOffset_v1710_Choice_nothing uint64 = iota
	SearchSpaceExt_v1700_monitoringSlotPeriodicityAndOffset_v1710_Choice_sl32
	SearchSpaceExt_v1700_monitoringSlotPeriodicityAndOffset_v1710_Choice_sl64
	SearchSpaceExt_v1700_monitoringSlotPeriodicityAndOffset_v1710_Choice_sl128
	SearchSpaceExt_v1700_monitoringSlotPeriodicityAndOffset_v1710_Choice_sl5120
	SearchSpaceExt_v1700_monitoringSlotPeriodicityAndOffset_v1710_Choice_sl10240
	SearchSpaceExt_v1700_monitoringSlotPeriodicityAndOffset_v1710_Choice_sl20480
)

type SearchSpaceExt_v1700_monitoringSlotPeriodicityAndOffset_v1710 struct {
	Choice  uint64
	sl32    int64 `lb:0,ub:31,madatory`
	sl64    int64 `lb:0,ub:63,madatory`
	sl128   int64 `lb:0,ub:127,madatory`
	sl5120  int64 `lb:0,ub:5119,madatory`
	sl10240 int64 `lb:0,ub:10239,madatory`
	sl20480 int64 `lb:0,ub:20479,madatory`
}

func (ie *SearchSpaceExt_v1700_monitoringSlotPeriodicityAndOffset_v1710) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 6, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SearchSpaceExt_v1700_monitoringSlotPeriodicityAndOffset_v1710_Choice_sl32:
		if err = w.WriteInteger(int64(ie.sl32), &uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
			err = utils.WrapError("Encode sl32", err)
		}
	case SearchSpaceExt_v1700_monitoringSlotPeriodicityAndOffset_v1710_Choice_sl64:
		if err = w.WriteInteger(int64(ie.sl64), &uper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
			err = utils.WrapError("Encode sl64", err)
		}
	case SearchSpaceExt_v1700_monitoringSlotPeriodicityAndOffset_v1710_Choice_sl128:
		if err = w.WriteInteger(int64(ie.sl128), &uper.Constraint{Lb: 0, Ub: 127}, false); err != nil {
			err = utils.WrapError("Encode sl128", err)
		}
	case SearchSpaceExt_v1700_monitoringSlotPeriodicityAndOffset_v1710_Choice_sl5120:
		if err = w.WriteInteger(int64(ie.sl5120), &uper.Constraint{Lb: 0, Ub: 5119}, false); err != nil {
			err = utils.WrapError("Encode sl5120", err)
		}
	case SearchSpaceExt_v1700_monitoringSlotPeriodicityAndOffset_v1710_Choice_sl10240:
		if err = w.WriteInteger(int64(ie.sl10240), &uper.Constraint{Lb: 0, Ub: 10239}, false); err != nil {
			err = utils.WrapError("Encode sl10240", err)
		}
	case SearchSpaceExt_v1700_monitoringSlotPeriodicityAndOffset_v1710_Choice_sl20480:
		if err = w.WriteInteger(int64(ie.sl20480), &uper.Constraint{Lb: 0, Ub: 20479}, false); err != nil {
			err = utils.WrapError("Encode sl20480", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *SearchSpaceExt_v1700_monitoringSlotPeriodicityAndOffset_v1710) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(6, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SearchSpaceExt_v1700_monitoringSlotPeriodicityAndOffset_v1710_Choice_sl32:
		var tmp_int_sl32 int64
		if tmp_int_sl32, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
			return utils.WrapError("Decode sl32", err)
		}
		ie.sl32 = tmp_int_sl32
	case SearchSpaceExt_v1700_monitoringSlotPeriodicityAndOffset_v1710_Choice_sl64:
		var tmp_int_sl64 int64
		if tmp_int_sl64, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
			return utils.WrapError("Decode sl64", err)
		}
		ie.sl64 = tmp_int_sl64
	case SearchSpaceExt_v1700_monitoringSlotPeriodicityAndOffset_v1710_Choice_sl128:
		var tmp_int_sl128 int64
		if tmp_int_sl128, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 127}, false); err != nil {
			return utils.WrapError("Decode sl128", err)
		}
		ie.sl128 = tmp_int_sl128
	case SearchSpaceExt_v1700_monitoringSlotPeriodicityAndOffset_v1710_Choice_sl5120:
		var tmp_int_sl5120 int64
		if tmp_int_sl5120, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 5119}, false); err != nil {
			return utils.WrapError("Decode sl5120", err)
		}
		ie.sl5120 = tmp_int_sl5120
	case SearchSpaceExt_v1700_monitoringSlotPeriodicityAndOffset_v1710_Choice_sl10240:
		var tmp_int_sl10240 int64
		if tmp_int_sl10240, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 10239}, false); err != nil {
			return utils.WrapError("Decode sl10240", err)
		}
		ie.sl10240 = tmp_int_sl10240
	case SearchSpaceExt_v1700_monitoringSlotPeriodicityAndOffset_v1710_Choice_sl20480:
		var tmp_int_sl20480 int64
		if tmp_int_sl20480, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 20479}, false); err != nil {
			return utils.WrapError("Decode sl20480", err)
		}
		ie.sl20480 = tmp_int_sl20480
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
