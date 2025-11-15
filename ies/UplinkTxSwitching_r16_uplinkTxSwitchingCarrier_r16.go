package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	UplinkTxSwitching_r16_uplinkTxSwitchingCarrier_r16_Enum_carrier1 uper.Enumerated = 0
	UplinkTxSwitching_r16_uplinkTxSwitchingCarrier_r16_Enum_carrier2 uper.Enumerated = 1
)

type UplinkTxSwitching_r16_uplinkTxSwitchingCarrier_r16 struct {
	Value uper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *UplinkTxSwitching_r16_uplinkTxSwitchingCarrier_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode UplinkTxSwitching_r16_uplinkTxSwitchingCarrier_r16", err)
	}
	return nil
}

func (ie *UplinkTxSwitching_r16_uplinkTxSwitchingCarrier_r16) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode UplinkTxSwitching_r16_uplinkTxSwitchingCarrier_r16", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
