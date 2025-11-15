package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	DRX_Info_drx_LongCycleStartOffset_Choice_nothing uint64 = iota
	DRX_Info_drx_LongCycleStartOffset_Choice_ms10
	DRX_Info_drx_LongCycleStartOffset_Choice_ms20
	DRX_Info_drx_LongCycleStartOffset_Choice_ms32
	DRX_Info_drx_LongCycleStartOffset_Choice_ms40
	DRX_Info_drx_LongCycleStartOffset_Choice_ms60
	DRX_Info_drx_LongCycleStartOffset_Choice_ms64
	DRX_Info_drx_LongCycleStartOffset_Choice_ms70
	DRX_Info_drx_LongCycleStartOffset_Choice_ms80
	DRX_Info_drx_LongCycleStartOffset_Choice_ms128
	DRX_Info_drx_LongCycleStartOffset_Choice_ms160
	DRX_Info_drx_LongCycleStartOffset_Choice_ms256
	DRX_Info_drx_LongCycleStartOffset_Choice_ms320
	DRX_Info_drx_LongCycleStartOffset_Choice_ms512
	DRX_Info_drx_LongCycleStartOffset_Choice_ms640
	DRX_Info_drx_LongCycleStartOffset_Choice_ms1024
	DRX_Info_drx_LongCycleStartOffset_Choice_ms1280
	DRX_Info_drx_LongCycleStartOffset_Choice_ms2048
	DRX_Info_drx_LongCycleStartOffset_Choice_ms2560
	DRX_Info_drx_LongCycleStartOffset_Choice_ms5120
	DRX_Info_drx_LongCycleStartOffset_Choice_ms10240
)

type DRX_Info_drx_LongCycleStartOffset struct {
	Choice  uint64
	ms10    int64 `lb:0,ub:9,madatory`
	ms20    int64 `lb:0,ub:19,madatory`
	ms32    int64 `lb:0,ub:31,madatory`
	ms40    int64 `lb:0,ub:39,madatory`
	ms60    int64 `lb:0,ub:59,madatory`
	ms64    int64 `lb:0,ub:63,madatory`
	ms70    int64 `lb:0,ub:69,madatory`
	ms80    int64 `lb:0,ub:79,madatory`
	ms128   int64 `lb:0,ub:127,madatory`
	ms160   int64 `lb:0,ub:159,madatory`
	ms256   int64 `lb:0,ub:255,madatory`
	ms320   int64 `lb:0,ub:319,madatory`
	ms512   int64 `lb:0,ub:511,madatory`
	ms640   int64 `lb:0,ub:639,madatory`
	ms1024  int64 `lb:0,ub:1023,madatory`
	ms1280  int64 `lb:0,ub:1279,madatory`
	ms2048  int64 `lb:0,ub:2047,madatory`
	ms2560  int64 `lb:0,ub:2559,madatory`
	ms5120  int64 `lb:0,ub:5119,madatory`
	ms10240 int64 `lb:0,ub:10239,madatory`
}

func (ie *DRX_Info_drx_LongCycleStartOffset) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 20, false); err != nil {
		return err
	}
	switch ie.Choice {
	case DRX_Info_drx_LongCycleStartOffset_Choice_ms10:
		if err = w.WriteInteger(int64(ie.ms10), &uper.Constraint{Lb: 0, Ub: 9}, false); err != nil {
			err = utils.WrapError("Encode ms10", err)
		}
	case DRX_Info_drx_LongCycleStartOffset_Choice_ms20:
		if err = w.WriteInteger(int64(ie.ms20), &uper.Constraint{Lb: 0, Ub: 19}, false); err != nil {
			err = utils.WrapError("Encode ms20", err)
		}
	case DRX_Info_drx_LongCycleStartOffset_Choice_ms32:
		if err = w.WriteInteger(int64(ie.ms32), &uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
			err = utils.WrapError("Encode ms32", err)
		}
	case DRX_Info_drx_LongCycleStartOffset_Choice_ms40:
		if err = w.WriteInteger(int64(ie.ms40), &uper.Constraint{Lb: 0, Ub: 39}, false); err != nil {
			err = utils.WrapError("Encode ms40", err)
		}
	case DRX_Info_drx_LongCycleStartOffset_Choice_ms60:
		if err = w.WriteInteger(int64(ie.ms60), &uper.Constraint{Lb: 0, Ub: 59}, false); err != nil {
			err = utils.WrapError("Encode ms60", err)
		}
	case DRX_Info_drx_LongCycleStartOffset_Choice_ms64:
		if err = w.WriteInteger(int64(ie.ms64), &uper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
			err = utils.WrapError("Encode ms64", err)
		}
	case DRX_Info_drx_LongCycleStartOffset_Choice_ms70:
		if err = w.WriteInteger(int64(ie.ms70), &uper.Constraint{Lb: 0, Ub: 69}, false); err != nil {
			err = utils.WrapError("Encode ms70", err)
		}
	case DRX_Info_drx_LongCycleStartOffset_Choice_ms80:
		if err = w.WriteInteger(int64(ie.ms80), &uper.Constraint{Lb: 0, Ub: 79}, false); err != nil {
			err = utils.WrapError("Encode ms80", err)
		}
	case DRX_Info_drx_LongCycleStartOffset_Choice_ms128:
		if err = w.WriteInteger(int64(ie.ms128), &uper.Constraint{Lb: 0, Ub: 127}, false); err != nil {
			err = utils.WrapError("Encode ms128", err)
		}
	case DRX_Info_drx_LongCycleStartOffset_Choice_ms160:
		if err = w.WriteInteger(int64(ie.ms160), &uper.Constraint{Lb: 0, Ub: 159}, false); err != nil {
			err = utils.WrapError("Encode ms160", err)
		}
	case DRX_Info_drx_LongCycleStartOffset_Choice_ms256:
		if err = w.WriteInteger(int64(ie.ms256), &uper.Constraint{Lb: 0, Ub: 255}, false); err != nil {
			err = utils.WrapError("Encode ms256", err)
		}
	case DRX_Info_drx_LongCycleStartOffset_Choice_ms320:
		if err = w.WriteInteger(int64(ie.ms320), &uper.Constraint{Lb: 0, Ub: 319}, false); err != nil {
			err = utils.WrapError("Encode ms320", err)
		}
	case DRX_Info_drx_LongCycleStartOffset_Choice_ms512:
		if err = w.WriteInteger(int64(ie.ms512), &uper.Constraint{Lb: 0, Ub: 511}, false); err != nil {
			err = utils.WrapError("Encode ms512", err)
		}
	case DRX_Info_drx_LongCycleStartOffset_Choice_ms640:
		if err = w.WriteInteger(int64(ie.ms640), &uper.Constraint{Lb: 0, Ub: 639}, false); err != nil {
			err = utils.WrapError("Encode ms640", err)
		}
	case DRX_Info_drx_LongCycleStartOffset_Choice_ms1024:
		if err = w.WriteInteger(int64(ie.ms1024), &uper.Constraint{Lb: 0, Ub: 1023}, false); err != nil {
			err = utils.WrapError("Encode ms1024", err)
		}
	case DRX_Info_drx_LongCycleStartOffset_Choice_ms1280:
		if err = w.WriteInteger(int64(ie.ms1280), &uper.Constraint{Lb: 0, Ub: 1279}, false); err != nil {
			err = utils.WrapError("Encode ms1280", err)
		}
	case DRX_Info_drx_LongCycleStartOffset_Choice_ms2048:
		if err = w.WriteInteger(int64(ie.ms2048), &uper.Constraint{Lb: 0, Ub: 2047}, false); err != nil {
			err = utils.WrapError("Encode ms2048", err)
		}
	case DRX_Info_drx_LongCycleStartOffset_Choice_ms2560:
		if err = w.WriteInteger(int64(ie.ms2560), &uper.Constraint{Lb: 0, Ub: 2559}, false); err != nil {
			err = utils.WrapError("Encode ms2560", err)
		}
	case DRX_Info_drx_LongCycleStartOffset_Choice_ms5120:
		if err = w.WriteInteger(int64(ie.ms5120), &uper.Constraint{Lb: 0, Ub: 5119}, false); err != nil {
			err = utils.WrapError("Encode ms5120", err)
		}
	case DRX_Info_drx_LongCycleStartOffset_Choice_ms10240:
		if err = w.WriteInteger(int64(ie.ms10240), &uper.Constraint{Lb: 0, Ub: 10239}, false); err != nil {
			err = utils.WrapError("Encode ms10240", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *DRX_Info_drx_LongCycleStartOffset) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(20, false); err != nil {
		return err
	}
	switch ie.Choice {
	case DRX_Info_drx_LongCycleStartOffset_Choice_ms10:
		var tmp_int_ms10 int64
		if tmp_int_ms10, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 9}, false); err != nil {
			return utils.WrapError("Decode ms10", err)
		}
		ie.ms10 = tmp_int_ms10
	case DRX_Info_drx_LongCycleStartOffset_Choice_ms20:
		var tmp_int_ms20 int64
		if tmp_int_ms20, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 19}, false); err != nil {
			return utils.WrapError("Decode ms20", err)
		}
		ie.ms20 = tmp_int_ms20
	case DRX_Info_drx_LongCycleStartOffset_Choice_ms32:
		var tmp_int_ms32 int64
		if tmp_int_ms32, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
			return utils.WrapError("Decode ms32", err)
		}
		ie.ms32 = tmp_int_ms32
	case DRX_Info_drx_LongCycleStartOffset_Choice_ms40:
		var tmp_int_ms40 int64
		if tmp_int_ms40, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 39}, false); err != nil {
			return utils.WrapError("Decode ms40", err)
		}
		ie.ms40 = tmp_int_ms40
	case DRX_Info_drx_LongCycleStartOffset_Choice_ms60:
		var tmp_int_ms60 int64
		if tmp_int_ms60, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 59}, false); err != nil {
			return utils.WrapError("Decode ms60", err)
		}
		ie.ms60 = tmp_int_ms60
	case DRX_Info_drx_LongCycleStartOffset_Choice_ms64:
		var tmp_int_ms64 int64
		if tmp_int_ms64, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
			return utils.WrapError("Decode ms64", err)
		}
		ie.ms64 = tmp_int_ms64
	case DRX_Info_drx_LongCycleStartOffset_Choice_ms70:
		var tmp_int_ms70 int64
		if tmp_int_ms70, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 69}, false); err != nil {
			return utils.WrapError("Decode ms70", err)
		}
		ie.ms70 = tmp_int_ms70
	case DRX_Info_drx_LongCycleStartOffset_Choice_ms80:
		var tmp_int_ms80 int64
		if tmp_int_ms80, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 79}, false); err != nil {
			return utils.WrapError("Decode ms80", err)
		}
		ie.ms80 = tmp_int_ms80
	case DRX_Info_drx_LongCycleStartOffset_Choice_ms128:
		var tmp_int_ms128 int64
		if tmp_int_ms128, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 127}, false); err != nil {
			return utils.WrapError("Decode ms128", err)
		}
		ie.ms128 = tmp_int_ms128
	case DRX_Info_drx_LongCycleStartOffset_Choice_ms160:
		var tmp_int_ms160 int64
		if tmp_int_ms160, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 159}, false); err != nil {
			return utils.WrapError("Decode ms160", err)
		}
		ie.ms160 = tmp_int_ms160
	case DRX_Info_drx_LongCycleStartOffset_Choice_ms256:
		var tmp_int_ms256 int64
		if tmp_int_ms256, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 255}, false); err != nil {
			return utils.WrapError("Decode ms256", err)
		}
		ie.ms256 = tmp_int_ms256
	case DRX_Info_drx_LongCycleStartOffset_Choice_ms320:
		var tmp_int_ms320 int64
		if tmp_int_ms320, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 319}, false); err != nil {
			return utils.WrapError("Decode ms320", err)
		}
		ie.ms320 = tmp_int_ms320
	case DRX_Info_drx_LongCycleStartOffset_Choice_ms512:
		var tmp_int_ms512 int64
		if tmp_int_ms512, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 511}, false); err != nil {
			return utils.WrapError("Decode ms512", err)
		}
		ie.ms512 = tmp_int_ms512
	case DRX_Info_drx_LongCycleStartOffset_Choice_ms640:
		var tmp_int_ms640 int64
		if tmp_int_ms640, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 639}, false); err != nil {
			return utils.WrapError("Decode ms640", err)
		}
		ie.ms640 = tmp_int_ms640
	case DRX_Info_drx_LongCycleStartOffset_Choice_ms1024:
		var tmp_int_ms1024 int64
		if tmp_int_ms1024, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 1023}, false); err != nil {
			return utils.WrapError("Decode ms1024", err)
		}
		ie.ms1024 = tmp_int_ms1024
	case DRX_Info_drx_LongCycleStartOffset_Choice_ms1280:
		var tmp_int_ms1280 int64
		if tmp_int_ms1280, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 1279}, false); err != nil {
			return utils.WrapError("Decode ms1280", err)
		}
		ie.ms1280 = tmp_int_ms1280
	case DRX_Info_drx_LongCycleStartOffset_Choice_ms2048:
		var tmp_int_ms2048 int64
		if tmp_int_ms2048, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 2047}, false); err != nil {
			return utils.WrapError("Decode ms2048", err)
		}
		ie.ms2048 = tmp_int_ms2048
	case DRX_Info_drx_LongCycleStartOffset_Choice_ms2560:
		var tmp_int_ms2560 int64
		if tmp_int_ms2560, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 2559}, false); err != nil {
			return utils.WrapError("Decode ms2560", err)
		}
		ie.ms2560 = tmp_int_ms2560
	case DRX_Info_drx_LongCycleStartOffset_Choice_ms5120:
		var tmp_int_ms5120 int64
		if tmp_int_ms5120, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 5119}, false); err != nil {
			return utils.WrapError("Decode ms5120", err)
		}
		ie.ms5120 = tmp_int_ms5120
	case DRX_Info_drx_LongCycleStartOffset_Choice_ms10240:
		var tmp_int_ms10240 int64
		if tmp_int_ms10240, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 10239}, false); err != nil {
			return utils.WrapError("Decode ms10240", err)
		}
		ie.ms10240 = tmp_int_ms10240
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
