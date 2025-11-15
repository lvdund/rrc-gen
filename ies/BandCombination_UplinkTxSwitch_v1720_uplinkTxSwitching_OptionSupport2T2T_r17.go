package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	BandCombination_UplinkTxSwitch_v1720_uplinkTxSwitching_OptionSupport2T2T_r17_Enum_switchedUL uper.Enumerated = 0
	BandCombination_UplinkTxSwitch_v1720_uplinkTxSwitching_OptionSupport2T2T_r17_Enum_dualUL     uper.Enumerated = 1
	BandCombination_UplinkTxSwitch_v1720_uplinkTxSwitching_OptionSupport2T2T_r17_Enum_both       uper.Enumerated = 2
)

type BandCombination_UplinkTxSwitch_v1720_uplinkTxSwitching_OptionSupport2T2T_r17 struct {
	Value uper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *BandCombination_UplinkTxSwitch_v1720_uplinkTxSwitching_OptionSupport2T2T_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode BandCombination_UplinkTxSwitch_v1720_uplinkTxSwitching_OptionSupport2T2T_r17", err)
	}
	return nil
}

func (ie *BandCombination_UplinkTxSwitch_v1720_uplinkTxSwitching_OptionSupport2T2T_r17) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode BandCombination_UplinkTxSwitch_v1720_uplinkTxSwitching_OptionSupport2T2T_r17", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
