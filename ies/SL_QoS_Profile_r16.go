package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_QoS_Profile_r16 struct {
	sl_PQI_r16   *SL_PQI_r16 `optional`
	sl_GFBR_r16  *int64      `lb:0,ub:4000000000,optional`
	sl_MFBR_r16  *int64      `lb:0,ub:4000000000,optional`
	sl_Range_r16 *int64      `lb:1,ub:1000,optional`
}

func (ie *SL_QoS_Profile_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.sl_PQI_r16 != nil, ie.sl_GFBR_r16 != nil, ie.sl_MFBR_r16 != nil, ie.sl_Range_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.sl_PQI_r16 != nil {
		if err = ie.sl_PQI_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_PQI_r16", err)
		}
	}
	if ie.sl_GFBR_r16 != nil {
		if err = w.WriteInteger(*ie.sl_GFBR_r16, &uper.Constraint{Lb: 0, Ub: 4000000000}, false); err != nil {
			return utils.WrapError("Encode sl_GFBR_r16", err)
		}
	}
	if ie.sl_MFBR_r16 != nil {
		if err = w.WriteInteger(*ie.sl_MFBR_r16, &uper.Constraint{Lb: 0, Ub: 4000000000}, false); err != nil {
			return utils.WrapError("Encode sl_MFBR_r16", err)
		}
	}
	if ie.sl_Range_r16 != nil {
		if err = w.WriteInteger(*ie.sl_Range_r16, &uper.Constraint{Lb: 1, Ub: 1000}, false); err != nil {
			return utils.WrapError("Encode sl_Range_r16", err)
		}
	}
	return nil
}

func (ie *SL_QoS_Profile_r16) Decode(r *uper.UperReader) error {
	var err error
	var sl_PQI_r16Present bool
	if sl_PQI_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_GFBR_r16Present bool
	if sl_GFBR_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_MFBR_r16Present bool
	if sl_MFBR_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_Range_r16Present bool
	if sl_Range_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if sl_PQI_r16Present {
		ie.sl_PQI_r16 = new(SL_PQI_r16)
		if err = ie.sl_PQI_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_PQI_r16", err)
		}
	}
	if sl_GFBR_r16Present {
		var tmp_int_sl_GFBR_r16 int64
		if tmp_int_sl_GFBR_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 4000000000}, false); err != nil {
			return utils.WrapError("Decode sl_GFBR_r16", err)
		}
		ie.sl_GFBR_r16 = &tmp_int_sl_GFBR_r16
	}
	if sl_MFBR_r16Present {
		var tmp_int_sl_MFBR_r16 int64
		if tmp_int_sl_MFBR_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 4000000000}, false); err != nil {
			return utils.WrapError("Decode sl_MFBR_r16", err)
		}
		ie.sl_MFBR_r16 = &tmp_int_sl_MFBR_r16
	}
	if sl_Range_r16Present {
		var tmp_int_sl_Range_r16 int64
		if tmp_int_sl_Range_r16, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 1000}, false); err != nil {
			return utils.WrapError("Decode sl_Range_r16", err)
		}
		ie.sl_Range_r16 = &tmp_int_sl_Range_r16
	}
	return nil
}
