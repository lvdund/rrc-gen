package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RxTxTimeDiff_r17 struct {
	result_k5_r17 *int64 `lb:0,ub:61565,optional`
}

func (ie *RxTxTimeDiff_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.result_k5_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.result_k5_r17 != nil {
		if err = w.WriteInteger(*ie.result_k5_r17, &uper.Constraint{Lb: 0, Ub: 61565}, false); err != nil {
			return utils.WrapError("Encode result_k5_r17", err)
		}
	}
	return nil
}

func (ie *RxTxTimeDiff_r17) Decode(r *uper.UperReader) error {
	var err error
	var result_k5_r17Present bool
	if result_k5_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if result_k5_r17Present {
		var tmp_int_result_k5_r17 int64
		if tmp_int_result_k5_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 61565}, false); err != nil {
			return utils.WrapError("Decode result_k5_r17", err)
		}
		ie.result_k5_r17 = &tmp_int_result_k5_r17
	}
	return nil
}
