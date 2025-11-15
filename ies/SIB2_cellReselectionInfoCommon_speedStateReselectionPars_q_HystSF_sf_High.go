package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SIB2_cellReselectionInfoCommon_speedStateReselectionPars_q_HystSF_sf_High_Enum_dB_6 uper.Enumerated = 0
	SIB2_cellReselectionInfoCommon_speedStateReselectionPars_q_HystSF_sf_High_Enum_dB_4 uper.Enumerated = 1
	SIB2_cellReselectionInfoCommon_speedStateReselectionPars_q_HystSF_sf_High_Enum_dB_2 uper.Enumerated = 2
	SIB2_cellReselectionInfoCommon_speedStateReselectionPars_q_HystSF_sf_High_Enum_dB0  uper.Enumerated = 3
)

type SIB2_cellReselectionInfoCommon_speedStateReselectionPars_q_HystSF_sf_High struct {
	Value uper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *SIB2_cellReselectionInfoCommon_speedStateReselectionPars_q_HystSF_sf_High) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode SIB2_cellReselectionInfoCommon_speedStateReselectionPars_q_HystSF_sf_High", err)
	}
	return nil
}

func (ie *SIB2_cellReselectionInfoCommon_speedStateReselectionPars_q_HystSF_sf_High) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode SIB2_cellReselectionInfoCommon_speedStateReselectionPars_q_HystSF_sf_High", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
