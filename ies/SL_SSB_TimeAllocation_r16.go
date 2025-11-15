package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_SSB_TimeAllocation_r16 struct {
	sl_NumSSB_WithinPeriod_r16 *SL_SSB_TimeAllocation_r16_sl_NumSSB_WithinPeriod_r16 `optional`
	sl_TimeOffsetSSB_r16       *int64                                                `lb:0,ub:1279,optional`
	sl_TimeInterval_r16        *int64                                                `lb:0,ub:639,optional`
}

func (ie *SL_SSB_TimeAllocation_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.sl_NumSSB_WithinPeriod_r16 != nil, ie.sl_TimeOffsetSSB_r16 != nil, ie.sl_TimeInterval_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.sl_NumSSB_WithinPeriod_r16 != nil {
		if err = ie.sl_NumSSB_WithinPeriod_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_NumSSB_WithinPeriod_r16", err)
		}
	}
	if ie.sl_TimeOffsetSSB_r16 != nil {
		if err = w.WriteInteger(*ie.sl_TimeOffsetSSB_r16, &uper.Constraint{Lb: 0, Ub: 1279}, false); err != nil {
			return utils.WrapError("Encode sl_TimeOffsetSSB_r16", err)
		}
	}
	if ie.sl_TimeInterval_r16 != nil {
		if err = w.WriteInteger(*ie.sl_TimeInterval_r16, &uper.Constraint{Lb: 0, Ub: 639}, false); err != nil {
			return utils.WrapError("Encode sl_TimeInterval_r16", err)
		}
	}
	return nil
}

func (ie *SL_SSB_TimeAllocation_r16) Decode(r *uper.UperReader) error {
	var err error
	var sl_NumSSB_WithinPeriod_r16Present bool
	if sl_NumSSB_WithinPeriod_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_TimeOffsetSSB_r16Present bool
	if sl_TimeOffsetSSB_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_TimeInterval_r16Present bool
	if sl_TimeInterval_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if sl_NumSSB_WithinPeriod_r16Present {
		ie.sl_NumSSB_WithinPeriod_r16 = new(SL_SSB_TimeAllocation_r16_sl_NumSSB_WithinPeriod_r16)
		if err = ie.sl_NumSSB_WithinPeriod_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_NumSSB_WithinPeriod_r16", err)
		}
	}
	if sl_TimeOffsetSSB_r16Present {
		var tmp_int_sl_TimeOffsetSSB_r16 int64
		if tmp_int_sl_TimeOffsetSSB_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 1279}, false); err != nil {
			return utils.WrapError("Decode sl_TimeOffsetSSB_r16", err)
		}
		ie.sl_TimeOffsetSSB_r16 = &tmp_int_sl_TimeOffsetSSB_r16
	}
	if sl_TimeInterval_r16Present {
		var tmp_int_sl_TimeInterval_r16 int64
		if tmp_int_sl_TimeInterval_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 639}, false); err != nil {
			return utils.WrapError("Decode sl_TimeInterval_r16", err)
		}
		ie.sl_TimeInterval_r16 = &tmp_int_sl_TimeInterval_r16
	}
	return nil
}
