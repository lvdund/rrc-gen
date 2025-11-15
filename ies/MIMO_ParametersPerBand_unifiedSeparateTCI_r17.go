package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MIMO_ParametersPerBand_unifiedSeparateTCI_r17 struct {
	maxConfiguredDL_TCI_r17        MIMO_ParametersPerBand_unifiedSeparateTCI_r17_maxConfiguredDL_TCI_r17        `madatory`
	maxConfiguredUL_TCI_r17        MIMO_ParametersPerBand_unifiedSeparateTCI_r17_maxConfiguredUL_TCI_r17        `madatory`
	maxActivatedDL_TCIAcrossCC_r17 MIMO_ParametersPerBand_unifiedSeparateTCI_r17_maxActivatedDL_TCIAcrossCC_r17 `madatory`
	maxActivatedUL_TCIAcrossCC_r17 MIMO_ParametersPerBand_unifiedSeparateTCI_r17_maxActivatedUL_TCIAcrossCC_r17 `madatory`
}

func (ie *MIMO_ParametersPerBand_unifiedSeparateTCI_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.maxConfiguredDL_TCI_r17.Encode(w); err != nil {
		return utils.WrapError("Encode maxConfiguredDL_TCI_r17", err)
	}
	if err = ie.maxConfiguredUL_TCI_r17.Encode(w); err != nil {
		return utils.WrapError("Encode maxConfiguredUL_TCI_r17", err)
	}
	if err = ie.maxActivatedDL_TCIAcrossCC_r17.Encode(w); err != nil {
		return utils.WrapError("Encode maxActivatedDL_TCIAcrossCC_r17", err)
	}
	if err = ie.maxActivatedUL_TCIAcrossCC_r17.Encode(w); err != nil {
		return utils.WrapError("Encode maxActivatedUL_TCIAcrossCC_r17", err)
	}
	return nil
}

func (ie *MIMO_ParametersPerBand_unifiedSeparateTCI_r17) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.maxConfiguredDL_TCI_r17.Decode(r); err != nil {
		return utils.WrapError("Decode maxConfiguredDL_TCI_r17", err)
	}
	if err = ie.maxConfiguredUL_TCI_r17.Decode(r); err != nil {
		return utils.WrapError("Decode maxConfiguredUL_TCI_r17", err)
	}
	if err = ie.maxActivatedDL_TCIAcrossCC_r17.Decode(r); err != nil {
		return utils.WrapError("Decode maxActivatedDL_TCIAcrossCC_r17", err)
	}
	if err = ie.maxActivatedUL_TCIAcrossCC_r17.Decode(r); err != nil {
		return utils.WrapError("Decode maxActivatedUL_TCIAcrossCC_r17", err)
	}
	return nil
}
