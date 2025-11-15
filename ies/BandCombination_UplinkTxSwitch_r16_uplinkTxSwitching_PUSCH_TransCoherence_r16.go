package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	BandCombination_UplinkTxSwitch_r16_uplinkTxSwitching_PUSCH_TransCoherence_r16_Enum_nonCoherent  uper.Enumerated = 0
	BandCombination_UplinkTxSwitch_r16_uplinkTxSwitching_PUSCH_TransCoherence_r16_Enum_fullCoherent uper.Enumerated = 1
)

type BandCombination_UplinkTxSwitch_r16_uplinkTxSwitching_PUSCH_TransCoherence_r16 struct {
	Value uper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *BandCombination_UplinkTxSwitch_r16_uplinkTxSwitching_PUSCH_TransCoherence_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode BandCombination_UplinkTxSwitch_r16_uplinkTxSwitching_PUSCH_TransCoherence_r16", err)
	}
	return nil
}

func (ie *BandCombination_UplinkTxSwitch_r16_uplinkTxSwitching_PUSCH_TransCoherence_r16) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode BandCombination_UplinkTxSwitch_r16_uplinkTxSwitching_PUSCH_TransCoherence_r16", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
