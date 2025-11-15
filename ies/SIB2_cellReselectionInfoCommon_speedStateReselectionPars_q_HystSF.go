package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SIB2_cellReselectionInfoCommon_speedStateReselectionPars_q_HystSF struct {
	sf_Medium SIB2_cellReselectionInfoCommon_speedStateReselectionPars_q_HystSF_sf_Medium `madatory`
	sf_High   SIB2_cellReselectionInfoCommon_speedStateReselectionPars_q_HystSF_sf_High   `madatory`
}

func (ie *SIB2_cellReselectionInfoCommon_speedStateReselectionPars_q_HystSF) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.sf_Medium.Encode(w); err != nil {
		return utils.WrapError("Encode sf_Medium", err)
	}
	if err = ie.sf_High.Encode(w); err != nil {
		return utils.WrapError("Encode sf_High", err)
	}
	return nil
}

func (ie *SIB2_cellReselectionInfoCommon_speedStateReselectionPars_q_HystSF) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.sf_Medium.Decode(r); err != nil {
		return utils.WrapError("Decode sf_Medium", err)
	}
	if err = ie.sf_High.Decode(r); err != nil {
		return utils.WrapError("Decode sf_High", err)
	}
	return nil
}
