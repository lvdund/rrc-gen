package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CrossCarrierSchedulingConfig_carrierIndicatorSize_r16 struct {
	carrierIndicatorSizeDCI_1_2_r16 int64 `lb:0,ub:3,madatory`
	carrierIndicatorSizeDCI_0_2_r16 int64 `lb:0,ub:3,madatory`
}

func (ie *CrossCarrierSchedulingConfig_carrierIndicatorSize_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(ie.carrierIndicatorSizeDCI_1_2_r16, &uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("WriteInteger carrierIndicatorSizeDCI_1_2_r16", err)
	}
	if err = w.WriteInteger(ie.carrierIndicatorSizeDCI_0_2_r16, &uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("WriteInteger carrierIndicatorSizeDCI_0_2_r16", err)
	}
	return nil
}

func (ie *CrossCarrierSchedulingConfig_carrierIndicatorSize_r16) Decode(r *uper.UperReader) error {
	var err error
	var tmp_int_carrierIndicatorSizeDCI_1_2_r16 int64
	if tmp_int_carrierIndicatorSizeDCI_1_2_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("ReadInteger carrierIndicatorSizeDCI_1_2_r16", err)
	}
	ie.carrierIndicatorSizeDCI_1_2_r16 = tmp_int_carrierIndicatorSizeDCI_1_2_r16
	var tmp_int_carrierIndicatorSizeDCI_0_2_r16 int64
	if tmp_int_carrierIndicatorSizeDCI_0_2_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("ReadInteger carrierIndicatorSizeDCI_0_2_r16", err)
	}
	ie.carrierIndicatorSizeDCI_0_2_r16 = tmp_int_carrierIndicatorSizeDCI_0_2_r16
	return nil
}
