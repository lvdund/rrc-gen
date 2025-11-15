package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_SelectionWindowConfig_r16 struct {
	sl_Priority_r16        int64                                               `lb:1,ub:8,madatory`
	sl_SelectionWindow_r16 SL_SelectionWindowConfig_r16_sl_SelectionWindow_r16 `madatory`
}

func (ie *SL_SelectionWindowConfig_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(ie.sl_Priority_r16, &uper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
		return utils.WrapError("WriteInteger sl_Priority_r16", err)
	}
	if err = ie.sl_SelectionWindow_r16.Encode(w); err != nil {
		return utils.WrapError("Encode sl_SelectionWindow_r16", err)
	}
	return nil
}

func (ie *SL_SelectionWindowConfig_r16) Decode(r *uper.UperReader) error {
	var err error
	var tmp_int_sl_Priority_r16 int64
	if tmp_int_sl_Priority_r16, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
		return utils.WrapError("ReadInteger sl_Priority_r16", err)
	}
	ie.sl_Priority_r16 = tmp_int_sl_Priority_r16
	if err = ie.sl_SelectionWindow_r16.Decode(r); err != nil {
		return utils.WrapError("Decode sl_SelectionWindow_r16", err)
	}
	return nil
}
