package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	MTCH_SSB_MappingWindowCycleOffset_r17_Choice_nothing uint64 = iota
	MTCH_SSB_MappingWindowCycleOffset_r17_Choice_ms10
	MTCH_SSB_MappingWindowCycleOffset_r17_Choice_ms20
	MTCH_SSB_MappingWindowCycleOffset_r17_Choice_ms32
	MTCH_SSB_MappingWindowCycleOffset_r17_Choice_ms64
	MTCH_SSB_MappingWindowCycleOffset_r17_Choice_ms128
	MTCH_SSB_MappingWindowCycleOffset_r17_Choice_ms256
)

type MTCH_SSB_MappingWindowCycleOffset_r17 struct {
	Choice uint64
	ms10   int64 `lb:0,ub:9,madatory`
	ms20   int64 `lb:0,ub:19,madatory`
	ms32   int64 `lb:0,ub:31,madatory`
	ms64   int64 `lb:0,ub:63,madatory`
	ms128  int64 `lb:0,ub:127,madatory`
	ms256  int64 `lb:0,ub:255,madatory`
}

func (ie *MTCH_SSB_MappingWindowCycleOffset_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 6, false); err != nil {
		return err
	}
	switch ie.Choice {
	case MTCH_SSB_MappingWindowCycleOffset_r17_Choice_ms10:
		if err = w.WriteInteger(int64(ie.ms10), &uper.Constraint{Lb: 0, Ub: 9}, false); err != nil {
			err = utils.WrapError("Encode ms10", err)
		}
	case MTCH_SSB_MappingWindowCycleOffset_r17_Choice_ms20:
		if err = w.WriteInteger(int64(ie.ms20), &uper.Constraint{Lb: 0, Ub: 19}, false); err != nil {
			err = utils.WrapError("Encode ms20", err)
		}
	case MTCH_SSB_MappingWindowCycleOffset_r17_Choice_ms32:
		if err = w.WriteInteger(int64(ie.ms32), &uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
			err = utils.WrapError("Encode ms32", err)
		}
	case MTCH_SSB_MappingWindowCycleOffset_r17_Choice_ms64:
		if err = w.WriteInteger(int64(ie.ms64), &uper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
			err = utils.WrapError("Encode ms64", err)
		}
	case MTCH_SSB_MappingWindowCycleOffset_r17_Choice_ms128:
		if err = w.WriteInteger(int64(ie.ms128), &uper.Constraint{Lb: 0, Ub: 127}, false); err != nil {
			err = utils.WrapError("Encode ms128", err)
		}
	case MTCH_SSB_MappingWindowCycleOffset_r17_Choice_ms256:
		if err = w.WriteInteger(int64(ie.ms256), &uper.Constraint{Lb: 0, Ub: 255}, false); err != nil {
			err = utils.WrapError("Encode ms256", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *MTCH_SSB_MappingWindowCycleOffset_r17) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(6, false); err != nil {
		return err
	}
	switch ie.Choice {
	case MTCH_SSB_MappingWindowCycleOffset_r17_Choice_ms10:
		var tmp_int_ms10 int64
		if tmp_int_ms10, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 9}, false); err != nil {
			return utils.WrapError("Decode ms10", err)
		}
		ie.ms10 = tmp_int_ms10
	case MTCH_SSB_MappingWindowCycleOffset_r17_Choice_ms20:
		var tmp_int_ms20 int64
		if tmp_int_ms20, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 19}, false); err != nil {
			return utils.WrapError("Decode ms20", err)
		}
		ie.ms20 = tmp_int_ms20
	case MTCH_SSB_MappingWindowCycleOffset_r17_Choice_ms32:
		var tmp_int_ms32 int64
		if tmp_int_ms32, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
			return utils.WrapError("Decode ms32", err)
		}
		ie.ms32 = tmp_int_ms32
	case MTCH_SSB_MappingWindowCycleOffset_r17_Choice_ms64:
		var tmp_int_ms64 int64
		if tmp_int_ms64, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
			return utils.WrapError("Decode ms64", err)
		}
		ie.ms64 = tmp_int_ms64
	case MTCH_SSB_MappingWindowCycleOffset_r17_Choice_ms128:
		var tmp_int_ms128 int64
		if tmp_int_ms128, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 127}, false); err != nil {
			return utils.WrapError("Decode ms128", err)
		}
		ie.ms128 = tmp_int_ms128
	case MTCH_SSB_MappingWindowCycleOffset_r17_Choice_ms256:
		var tmp_int_ms256 int64
		if tmp_int_ms256, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 255}, false); err != nil {
			return utils.WrapError("Decode ms256", err)
		}
		ie.ms256 = tmp_int_ms256
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
