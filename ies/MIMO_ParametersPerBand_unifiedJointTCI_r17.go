package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MIMO_ParametersPerBand_unifiedJointTCI_r17 struct {
	maxConfiguredJointTCI_r17   MIMO_ParametersPerBand_unifiedJointTCI_r17_maxConfiguredJointTCI_r17   `madatory`
	maxActivatedTCIAcrossCC_r17 MIMO_ParametersPerBand_unifiedJointTCI_r17_maxActivatedTCIAcrossCC_r17 `madatory`
}

func (ie *MIMO_ParametersPerBand_unifiedJointTCI_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.maxConfiguredJointTCI_r17.Encode(w); err != nil {
		return utils.WrapError("Encode maxConfiguredJointTCI_r17", err)
	}
	if err = ie.maxActivatedTCIAcrossCC_r17.Encode(w); err != nil {
		return utils.WrapError("Encode maxActivatedTCIAcrossCC_r17", err)
	}
	return nil
}

func (ie *MIMO_ParametersPerBand_unifiedJointTCI_r17) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.maxConfiguredJointTCI_r17.Decode(r); err != nil {
		return utils.WrapError("Decode maxConfiguredJointTCI_r17", err)
	}
	if err = ie.maxActivatedTCIAcrossCC_r17.Decode(r); err != nil {
		return utils.WrapError("Decode maxActivatedTCIAcrossCC_r17", err)
	}
	return nil
}
