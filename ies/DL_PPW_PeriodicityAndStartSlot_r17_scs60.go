package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	DL_PPW_PeriodicityAndStartSlot_r17_scs60_Choice_nothing uint64 = iota
	DL_PPW_PeriodicityAndStartSlot_r17_scs60_Choice_n16
	DL_PPW_PeriodicityAndStartSlot_r17_scs60_Choice_n20
	DL_PPW_PeriodicityAndStartSlot_r17_scs60_Choice_n32
	DL_PPW_PeriodicityAndStartSlot_r17_scs60_Choice_n40
	DL_PPW_PeriodicityAndStartSlot_r17_scs60_Choice_n64
	DL_PPW_PeriodicityAndStartSlot_r17_scs60_Choice_n80
	DL_PPW_PeriodicityAndStartSlot_r17_scs60_Choice_n128
	DL_PPW_PeriodicityAndStartSlot_r17_scs60_Choice_n160
	DL_PPW_PeriodicityAndStartSlot_r17_scs60_Choice_n256
	DL_PPW_PeriodicityAndStartSlot_r17_scs60_Choice_n320
	DL_PPW_PeriodicityAndStartSlot_r17_scs60_Choice_n640
	DL_PPW_PeriodicityAndStartSlot_r17_scs60_Choice_n1280
	DL_PPW_PeriodicityAndStartSlot_r17_scs60_Choice_n2560
	DL_PPW_PeriodicityAndStartSlot_r17_scs60_Choice_n5120
	DL_PPW_PeriodicityAndStartSlot_r17_scs60_Choice_n10240
	DL_PPW_PeriodicityAndStartSlot_r17_scs60_Choice_n20480
	DL_PPW_PeriodicityAndStartSlot_r17_scs60_Choice_n40960
)

type DL_PPW_PeriodicityAndStartSlot_r17_scs60 struct {
	Choice uint64
	n16    int64 `lb:0,ub:15,madatory`
	n20    int64 `lb:0,ub:19,madatory`
	n32    int64 `lb:0,ub:31,madatory`
	n40    int64 `lb:0,ub:39,madatory`
	n64    int64 `lb:0,ub:63,madatory`
	n80    int64 `lb:0,ub:79,madatory`
	n128   int64 `lb:0,ub:127,madatory`
	n160   int64 `lb:0,ub:159,madatory`
	n256   int64 `lb:0,ub:255,madatory`
	n320   int64 `lb:0,ub:319,madatory`
	n640   int64 `lb:0,ub:639,madatory`
	n1280  int64 `lb:0,ub:1279,madatory`
	n2560  int64 `lb:0,ub:2559,madatory`
	n5120  int64 `lb:0,ub:5119,madatory`
	n10240 int64 `lb:0,ub:10239,madatory`
	n20480 int64 `lb:0,ub:20479,madatory`
	n40960 int64 `lb:0,ub:40959,madatory`
}

func (ie *DL_PPW_PeriodicityAndStartSlot_r17_scs60) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 17, false); err != nil {
		return err
	}
	switch ie.Choice {
	case DL_PPW_PeriodicityAndStartSlot_r17_scs60_Choice_n16:
		if err = w.WriteInteger(int64(ie.n16), &uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
			err = utils.WrapError("Encode n16", err)
		}
	case DL_PPW_PeriodicityAndStartSlot_r17_scs60_Choice_n20:
		if err = w.WriteInteger(int64(ie.n20), &uper.Constraint{Lb: 0, Ub: 19}, false); err != nil {
			err = utils.WrapError("Encode n20", err)
		}
	case DL_PPW_PeriodicityAndStartSlot_r17_scs60_Choice_n32:
		if err = w.WriteInteger(int64(ie.n32), &uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
			err = utils.WrapError("Encode n32", err)
		}
	case DL_PPW_PeriodicityAndStartSlot_r17_scs60_Choice_n40:
		if err = w.WriteInteger(int64(ie.n40), &uper.Constraint{Lb: 0, Ub: 39}, false); err != nil {
			err = utils.WrapError("Encode n40", err)
		}
	case DL_PPW_PeriodicityAndStartSlot_r17_scs60_Choice_n64:
		if err = w.WriteInteger(int64(ie.n64), &uper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
			err = utils.WrapError("Encode n64", err)
		}
	case DL_PPW_PeriodicityAndStartSlot_r17_scs60_Choice_n80:
		if err = w.WriteInteger(int64(ie.n80), &uper.Constraint{Lb: 0, Ub: 79}, false); err != nil {
			err = utils.WrapError("Encode n80", err)
		}
	case DL_PPW_PeriodicityAndStartSlot_r17_scs60_Choice_n128:
		if err = w.WriteInteger(int64(ie.n128), &uper.Constraint{Lb: 0, Ub: 127}, false); err != nil {
			err = utils.WrapError("Encode n128", err)
		}
	case DL_PPW_PeriodicityAndStartSlot_r17_scs60_Choice_n160:
		if err = w.WriteInteger(int64(ie.n160), &uper.Constraint{Lb: 0, Ub: 159}, false); err != nil {
			err = utils.WrapError("Encode n160", err)
		}
	case DL_PPW_PeriodicityAndStartSlot_r17_scs60_Choice_n256:
		if err = w.WriteInteger(int64(ie.n256), &uper.Constraint{Lb: 0, Ub: 255}, false); err != nil {
			err = utils.WrapError("Encode n256", err)
		}
	case DL_PPW_PeriodicityAndStartSlot_r17_scs60_Choice_n320:
		if err = w.WriteInteger(int64(ie.n320), &uper.Constraint{Lb: 0, Ub: 319}, false); err != nil {
			err = utils.WrapError("Encode n320", err)
		}
	case DL_PPW_PeriodicityAndStartSlot_r17_scs60_Choice_n640:
		if err = w.WriteInteger(int64(ie.n640), &uper.Constraint{Lb: 0, Ub: 639}, false); err != nil {
			err = utils.WrapError("Encode n640", err)
		}
	case DL_PPW_PeriodicityAndStartSlot_r17_scs60_Choice_n1280:
		if err = w.WriteInteger(int64(ie.n1280), &uper.Constraint{Lb: 0, Ub: 1279}, false); err != nil {
			err = utils.WrapError("Encode n1280", err)
		}
	case DL_PPW_PeriodicityAndStartSlot_r17_scs60_Choice_n2560:
		if err = w.WriteInteger(int64(ie.n2560), &uper.Constraint{Lb: 0, Ub: 2559}, false); err != nil {
			err = utils.WrapError("Encode n2560", err)
		}
	case DL_PPW_PeriodicityAndStartSlot_r17_scs60_Choice_n5120:
		if err = w.WriteInteger(int64(ie.n5120), &uper.Constraint{Lb: 0, Ub: 5119}, false); err != nil {
			err = utils.WrapError("Encode n5120", err)
		}
	case DL_PPW_PeriodicityAndStartSlot_r17_scs60_Choice_n10240:
		if err = w.WriteInteger(int64(ie.n10240), &uper.Constraint{Lb: 0, Ub: 10239}, false); err != nil {
			err = utils.WrapError("Encode n10240", err)
		}
	case DL_PPW_PeriodicityAndStartSlot_r17_scs60_Choice_n20480:
		if err = w.WriteInteger(int64(ie.n20480), &uper.Constraint{Lb: 0, Ub: 20479}, false); err != nil {
			err = utils.WrapError("Encode n20480", err)
		}
	case DL_PPW_PeriodicityAndStartSlot_r17_scs60_Choice_n40960:
		if err = w.WriteInteger(int64(ie.n40960), &uper.Constraint{Lb: 0, Ub: 40959}, false); err != nil {
			err = utils.WrapError("Encode n40960", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *DL_PPW_PeriodicityAndStartSlot_r17_scs60) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(17, false); err != nil {
		return err
	}
	switch ie.Choice {
	case DL_PPW_PeriodicityAndStartSlot_r17_scs60_Choice_n16:
		var tmp_int_n16 int64
		if tmp_int_n16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
			return utils.WrapError("Decode n16", err)
		}
		ie.n16 = tmp_int_n16
	case DL_PPW_PeriodicityAndStartSlot_r17_scs60_Choice_n20:
		var tmp_int_n20 int64
		if tmp_int_n20, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 19}, false); err != nil {
			return utils.WrapError("Decode n20", err)
		}
		ie.n20 = tmp_int_n20
	case DL_PPW_PeriodicityAndStartSlot_r17_scs60_Choice_n32:
		var tmp_int_n32 int64
		if tmp_int_n32, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
			return utils.WrapError("Decode n32", err)
		}
		ie.n32 = tmp_int_n32
	case DL_PPW_PeriodicityAndStartSlot_r17_scs60_Choice_n40:
		var tmp_int_n40 int64
		if tmp_int_n40, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 39}, false); err != nil {
			return utils.WrapError("Decode n40", err)
		}
		ie.n40 = tmp_int_n40
	case DL_PPW_PeriodicityAndStartSlot_r17_scs60_Choice_n64:
		var tmp_int_n64 int64
		if tmp_int_n64, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
			return utils.WrapError("Decode n64", err)
		}
		ie.n64 = tmp_int_n64
	case DL_PPW_PeriodicityAndStartSlot_r17_scs60_Choice_n80:
		var tmp_int_n80 int64
		if tmp_int_n80, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 79}, false); err != nil {
			return utils.WrapError("Decode n80", err)
		}
		ie.n80 = tmp_int_n80
	case DL_PPW_PeriodicityAndStartSlot_r17_scs60_Choice_n128:
		var tmp_int_n128 int64
		if tmp_int_n128, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 127}, false); err != nil {
			return utils.WrapError("Decode n128", err)
		}
		ie.n128 = tmp_int_n128
	case DL_PPW_PeriodicityAndStartSlot_r17_scs60_Choice_n160:
		var tmp_int_n160 int64
		if tmp_int_n160, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 159}, false); err != nil {
			return utils.WrapError("Decode n160", err)
		}
		ie.n160 = tmp_int_n160
	case DL_PPW_PeriodicityAndStartSlot_r17_scs60_Choice_n256:
		var tmp_int_n256 int64
		if tmp_int_n256, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 255}, false); err != nil {
			return utils.WrapError("Decode n256", err)
		}
		ie.n256 = tmp_int_n256
	case DL_PPW_PeriodicityAndStartSlot_r17_scs60_Choice_n320:
		var tmp_int_n320 int64
		if tmp_int_n320, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 319}, false); err != nil {
			return utils.WrapError("Decode n320", err)
		}
		ie.n320 = tmp_int_n320
	case DL_PPW_PeriodicityAndStartSlot_r17_scs60_Choice_n640:
		var tmp_int_n640 int64
		if tmp_int_n640, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 639}, false); err != nil {
			return utils.WrapError("Decode n640", err)
		}
		ie.n640 = tmp_int_n640
	case DL_PPW_PeriodicityAndStartSlot_r17_scs60_Choice_n1280:
		var tmp_int_n1280 int64
		if tmp_int_n1280, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 1279}, false); err != nil {
			return utils.WrapError("Decode n1280", err)
		}
		ie.n1280 = tmp_int_n1280
	case DL_PPW_PeriodicityAndStartSlot_r17_scs60_Choice_n2560:
		var tmp_int_n2560 int64
		if tmp_int_n2560, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 2559}, false); err != nil {
			return utils.WrapError("Decode n2560", err)
		}
		ie.n2560 = tmp_int_n2560
	case DL_PPW_PeriodicityAndStartSlot_r17_scs60_Choice_n5120:
		var tmp_int_n5120 int64
		if tmp_int_n5120, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 5119}, false); err != nil {
			return utils.WrapError("Decode n5120", err)
		}
		ie.n5120 = tmp_int_n5120
	case DL_PPW_PeriodicityAndStartSlot_r17_scs60_Choice_n10240:
		var tmp_int_n10240 int64
		if tmp_int_n10240, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 10239}, false); err != nil {
			return utils.WrapError("Decode n10240", err)
		}
		ie.n10240 = tmp_int_n10240
	case DL_PPW_PeriodicityAndStartSlot_r17_scs60_Choice_n20480:
		var tmp_int_n20480 int64
		if tmp_int_n20480, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 20479}, false); err != nil {
			return utils.WrapError("Decode n20480", err)
		}
		ie.n20480 = tmp_int_n20480
	case DL_PPW_PeriodicityAndStartSlot_r17_scs60_Choice_n40960:
		var tmp_int_n40960 int64
		if tmp_int_n40960, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 40959}, false); err != nil {
			return utils.WrapError("Decode n40960", err)
		}
		ie.n40960 = tmp_int_n40960
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
