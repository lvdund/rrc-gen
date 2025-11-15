package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UplinkTxSwitching_r16 struct {
	uplinkTxSwitchingPeriodLocation_r16 bool                                               `madatory`
	uplinkTxSwitchingCarrier_r16        UplinkTxSwitching_r16_uplinkTxSwitchingCarrier_r16 `madatory`
}

func (ie *UplinkTxSwitching_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteBoolean(ie.uplinkTxSwitchingPeriodLocation_r16); err != nil {
		return utils.WrapError("WriteBoolean uplinkTxSwitchingPeriodLocation_r16", err)
	}
	if err = ie.uplinkTxSwitchingCarrier_r16.Encode(w); err != nil {
		return utils.WrapError("Encode uplinkTxSwitchingCarrier_r16", err)
	}
	return nil
}

func (ie *UplinkTxSwitching_r16) Decode(r *uper.UperReader) error {
	var err error
	var tmp_bool_uplinkTxSwitchingPeriodLocation_r16 bool
	if tmp_bool_uplinkTxSwitchingPeriodLocation_r16, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean uplinkTxSwitchingPeriodLocation_r16", err)
	}
	ie.uplinkTxSwitchingPeriodLocation_r16 = tmp_bool_uplinkTxSwitchingPeriodLocation_r16
	if err = ie.uplinkTxSwitchingCarrier_r16.Decode(r); err != nil {
		return utils.WrapError("Decode uplinkTxSwitchingCarrier_r16", err)
	}
	return nil
}
