package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	ULTxSwitchingBandPair_v1700_uplinkTxSwitchingPeriod2T2T_r17_Enum_n35us  uper.Enumerated = 0
	ULTxSwitchingBandPair_v1700_uplinkTxSwitchingPeriod2T2T_r17_Enum_n140us uper.Enumerated = 1
	ULTxSwitchingBandPair_v1700_uplinkTxSwitchingPeriod2T2T_r17_Enum_n210us uper.Enumerated = 2
)

type ULTxSwitchingBandPair_v1700_uplinkTxSwitchingPeriod2T2T_r17 struct {
	Value uper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *ULTxSwitchingBandPair_v1700_uplinkTxSwitchingPeriod2T2T_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode ULTxSwitchingBandPair_v1700_uplinkTxSwitchingPeriod2T2T_r17", err)
	}
	return nil
}

func (ie *ULTxSwitchingBandPair_v1700_uplinkTxSwitchingPeriod2T2T_r17) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode ULTxSwitchingBandPair_v1700_uplinkTxSwitchingPeriod2T2T_r17", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
