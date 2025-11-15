package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MIMO_ParametersPerBand_unifiedJointTCI_InterCell_r17 struct {
	additionalMAC_CE_PerCC_r17    MIMO_ParametersPerBand_unifiedJointTCI_InterCell_r17_additionalMAC_CE_PerCC_r17    `madatory`
	additionalMAC_CE_AcrossCC_r17 MIMO_ParametersPerBand_unifiedJointTCI_InterCell_r17_additionalMAC_CE_AcrossCC_r17 `madatory`
}

func (ie *MIMO_ParametersPerBand_unifiedJointTCI_InterCell_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.additionalMAC_CE_PerCC_r17.Encode(w); err != nil {
		return utils.WrapError("Encode additionalMAC_CE_PerCC_r17", err)
	}
	if err = ie.additionalMAC_CE_AcrossCC_r17.Encode(w); err != nil {
		return utils.WrapError("Encode additionalMAC_CE_AcrossCC_r17", err)
	}
	return nil
}

func (ie *MIMO_ParametersPerBand_unifiedJointTCI_InterCell_r17) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.additionalMAC_CE_PerCC_r17.Decode(r); err != nil {
		return utils.WrapError("Decode additionalMAC_CE_PerCC_r17", err)
	}
	if err = ie.additionalMAC_CE_AcrossCC_r17.Decode(r); err != nil {
		return utils.WrapError("Decode additionalMAC_CE_AcrossCC_r17", err)
	}
	return nil
}
