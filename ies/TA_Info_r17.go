package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type TA_Info_r17 struct {
	ta_Common_r17             int64  `lb:0,ub:66485757,madatory`
	ta_CommonDrift_r17        *int64 `lb:-257303,ub:257303,optional`
	ta_CommonDriftVariant_r17 *int64 `lb:0,ub:28949,optional`
}

func (ie *TA_Info_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.ta_CommonDrift_r17 != nil, ie.ta_CommonDriftVariant_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = w.WriteInteger(ie.ta_Common_r17, &uper.Constraint{Lb: 0, Ub: 66485757}, false); err != nil {
		return utils.WrapError("WriteInteger ta_Common_r17", err)
	}
	if ie.ta_CommonDrift_r17 != nil {
		if err = w.WriteInteger(*ie.ta_CommonDrift_r17, &uper.Constraint{Lb: -257303, Ub: 257303}, false); err != nil {
			return utils.WrapError("Encode ta_CommonDrift_r17", err)
		}
	}
	if ie.ta_CommonDriftVariant_r17 != nil {
		if err = w.WriteInteger(*ie.ta_CommonDriftVariant_r17, &uper.Constraint{Lb: 0, Ub: 28949}, false); err != nil {
			return utils.WrapError("Encode ta_CommonDriftVariant_r17", err)
		}
	}
	return nil
}

func (ie *TA_Info_r17) Decode(r *uper.UperReader) error {
	var err error
	var ta_CommonDrift_r17Present bool
	if ta_CommonDrift_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ta_CommonDriftVariant_r17Present bool
	if ta_CommonDriftVariant_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var tmp_int_ta_Common_r17 int64
	if tmp_int_ta_Common_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 66485757}, false); err != nil {
		return utils.WrapError("ReadInteger ta_Common_r17", err)
	}
	ie.ta_Common_r17 = tmp_int_ta_Common_r17
	if ta_CommonDrift_r17Present {
		var tmp_int_ta_CommonDrift_r17 int64
		if tmp_int_ta_CommonDrift_r17, err = r.ReadInteger(&uper.Constraint{Lb: -257303, Ub: 257303}, false); err != nil {
			return utils.WrapError("Decode ta_CommonDrift_r17", err)
		}
		ie.ta_CommonDrift_r17 = &tmp_int_ta_CommonDrift_r17
	}
	if ta_CommonDriftVariant_r17Present {
		var tmp_int_ta_CommonDriftVariant_r17 int64
		if tmp_int_ta_CommonDriftVariant_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 28949}, false); err != nil {
			return utils.WrapError("Decode ta_CommonDriftVariant_r17", err)
		}
		ie.ta_CommonDriftVariant_r17 = &tmp_int_ta_CommonDriftVariant_r17
	}
	return nil
}
