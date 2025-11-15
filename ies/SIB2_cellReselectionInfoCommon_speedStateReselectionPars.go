package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SIB2_cellReselectionInfoCommon_speedStateReselectionPars struct {
	mobilityStateParameters MobilityStateParameters                                           `madatory`
	q_HystSF                SIB2_cellReselectionInfoCommon_speedStateReselectionPars_q_HystSF `madatory`
}

func (ie *SIB2_cellReselectionInfoCommon_speedStateReselectionPars) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.mobilityStateParameters.Encode(w); err != nil {
		return utils.WrapError("Encode mobilityStateParameters", err)
	}
	if err = ie.q_HystSF.Encode(w); err != nil {
		return utils.WrapError("Encode q_HystSF", err)
	}
	return nil
}

func (ie *SIB2_cellReselectionInfoCommon_speedStateReselectionPars) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.mobilityStateParameters.Decode(r); err != nil {
		return utils.WrapError("Decode mobilityStateParameters", err)
	}
	if err = ie.q_HystSF.Decode(r); err != nil {
		return utils.WrapError("Decode q_HystSF", err)
	}
	return nil
}
