package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UplinkTxDirectCurrentTwoCarrierInfo_r16 struct {
	referenceCarrierIndex_r16   ServCellIndex `madatory`
	shift7dot5kHz_r16           bool          `madatory`
	txDirectCurrentLocation_r16 int64         `lb:0,ub:3301,madatory`
}

func (ie *UplinkTxDirectCurrentTwoCarrierInfo_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.referenceCarrierIndex_r16.Encode(w); err != nil {
		return utils.WrapError("Encode referenceCarrierIndex_r16", err)
	}
	if err = w.WriteBoolean(ie.shift7dot5kHz_r16); err != nil {
		return utils.WrapError("WriteBoolean shift7dot5kHz_r16", err)
	}
	if err = w.WriteInteger(ie.txDirectCurrentLocation_r16, &uper.Constraint{Lb: 0, Ub: 3301}, false); err != nil {
		return utils.WrapError("WriteInteger txDirectCurrentLocation_r16", err)
	}
	return nil
}

func (ie *UplinkTxDirectCurrentTwoCarrierInfo_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.referenceCarrierIndex_r16.Decode(r); err != nil {
		return utils.WrapError("Decode referenceCarrierIndex_r16", err)
	}
	var tmp_bool_shift7dot5kHz_r16 bool
	if tmp_bool_shift7dot5kHz_r16, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean shift7dot5kHz_r16", err)
	}
	ie.shift7dot5kHz_r16 = tmp_bool_shift7dot5kHz_r16
	var tmp_int_txDirectCurrentLocation_r16 int64
	if tmp_int_txDirectCurrentLocation_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 3301}, false); err != nil {
		return utils.WrapError("ReadInteger txDirectCurrentLocation_r16", err)
	}
	ie.txDirectCurrentLocation_r16 = tmp_int_txDirectCurrentLocation_r16
	return nil
}
