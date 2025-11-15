package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MIMO_ParametersPerBand_unifiedSeparateTCI_InterCell_r17 struct {
	k_DL_PerCC_r17    MIMO_ParametersPerBand_unifiedSeparateTCI_InterCell_r17_k_DL_PerCC_r17    `madatory`
	k_UL_PerCC_r17    MIMO_ParametersPerBand_unifiedSeparateTCI_InterCell_r17_k_UL_PerCC_r17    `madatory`
	k_DL_AcrossCC_r17 MIMO_ParametersPerBand_unifiedSeparateTCI_InterCell_r17_k_DL_AcrossCC_r17 `madatory`
	k_UL_AcrossCC_r17 MIMO_ParametersPerBand_unifiedSeparateTCI_InterCell_r17_k_UL_AcrossCC_r17 `madatory`
}

func (ie *MIMO_ParametersPerBand_unifiedSeparateTCI_InterCell_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.k_DL_PerCC_r17.Encode(w); err != nil {
		return utils.WrapError("Encode k_DL_PerCC_r17", err)
	}
	if err = ie.k_UL_PerCC_r17.Encode(w); err != nil {
		return utils.WrapError("Encode k_UL_PerCC_r17", err)
	}
	if err = ie.k_DL_AcrossCC_r17.Encode(w); err != nil {
		return utils.WrapError("Encode k_DL_AcrossCC_r17", err)
	}
	if err = ie.k_UL_AcrossCC_r17.Encode(w); err != nil {
		return utils.WrapError("Encode k_UL_AcrossCC_r17", err)
	}
	return nil
}

func (ie *MIMO_ParametersPerBand_unifiedSeparateTCI_InterCell_r17) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.k_DL_PerCC_r17.Decode(r); err != nil {
		return utils.WrapError("Decode k_DL_PerCC_r17", err)
	}
	if err = ie.k_UL_PerCC_r17.Decode(r); err != nil {
		return utils.WrapError("Decode k_UL_PerCC_r17", err)
	}
	if err = ie.k_DL_AcrossCC_r17.Decode(r); err != nil {
		return utils.WrapError("Decode k_DL_AcrossCC_r17", err)
	}
	if err = ie.k_UL_AcrossCC_r17.Decode(r); err != nil {
		return utils.WrapError("Decode k_UL_AcrossCC_r17", err)
	}
	return nil
}
