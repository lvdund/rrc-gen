package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type WLAN_RTT_r16 struct {
	rttValue_r16    int64                     `lb:0,ub:16777215,madatory`
	rttUnits_r16    WLAN_RTT_r16_rttUnits_r16 `madatory`
	rttAccuracy_r16 *int64                    `lb:0,ub:255,optional`
}

func (ie *WLAN_RTT_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.rttAccuracy_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = w.WriteInteger(ie.rttValue_r16, &uper.Constraint{Lb: 0, Ub: 16777215}, false); err != nil {
		return utils.WrapError("WriteInteger rttValue_r16", err)
	}
	if err = ie.rttUnits_r16.Encode(w); err != nil {
		return utils.WrapError("Encode rttUnits_r16", err)
	}
	if ie.rttAccuracy_r16 != nil {
		if err = w.WriteInteger(*ie.rttAccuracy_r16, &uper.Constraint{Lb: 0, Ub: 255}, false); err != nil {
			return utils.WrapError("Encode rttAccuracy_r16", err)
		}
	}
	return nil
}

func (ie *WLAN_RTT_r16) Decode(r *uper.UperReader) error {
	var err error
	var rttAccuracy_r16Present bool
	if rttAccuracy_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var tmp_int_rttValue_r16 int64
	if tmp_int_rttValue_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 16777215}, false); err != nil {
		return utils.WrapError("ReadInteger rttValue_r16", err)
	}
	ie.rttValue_r16 = tmp_int_rttValue_r16
	if err = ie.rttUnits_r16.Decode(r); err != nil {
		return utils.WrapError("Decode rttUnits_r16", err)
	}
	if rttAccuracy_r16Present {
		var tmp_int_rttAccuracy_r16 int64
		if tmp_int_rttAccuracy_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 255}, false); err != nil {
			return utils.WrapError("Decode rttAccuracy_r16", err)
		}
		ie.rttAccuracy_r16 = &tmp_int_rttAccuracy_r16
	}
	return nil
}
