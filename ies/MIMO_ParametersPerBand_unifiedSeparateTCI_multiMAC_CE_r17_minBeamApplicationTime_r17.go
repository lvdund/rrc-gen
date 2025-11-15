package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	MIMO_ParametersPerBand_unifiedSeparateTCI_multiMAC_CE_r17_minBeamApplicationTime_r17_Enum_n1   uper.Enumerated = 0
	MIMO_ParametersPerBand_unifiedSeparateTCI_multiMAC_CE_r17_minBeamApplicationTime_r17_Enum_n2   uper.Enumerated = 1
	MIMO_ParametersPerBand_unifiedSeparateTCI_multiMAC_CE_r17_minBeamApplicationTime_r17_Enum_n4   uper.Enumerated = 2
	MIMO_ParametersPerBand_unifiedSeparateTCI_multiMAC_CE_r17_minBeamApplicationTime_r17_Enum_n7   uper.Enumerated = 3
	MIMO_ParametersPerBand_unifiedSeparateTCI_multiMAC_CE_r17_minBeamApplicationTime_r17_Enum_n14  uper.Enumerated = 4
	MIMO_ParametersPerBand_unifiedSeparateTCI_multiMAC_CE_r17_minBeamApplicationTime_r17_Enum_n28  uper.Enumerated = 5
	MIMO_ParametersPerBand_unifiedSeparateTCI_multiMAC_CE_r17_minBeamApplicationTime_r17_Enum_n42  uper.Enumerated = 6
	MIMO_ParametersPerBand_unifiedSeparateTCI_multiMAC_CE_r17_minBeamApplicationTime_r17_Enum_n56  uper.Enumerated = 7
	MIMO_ParametersPerBand_unifiedSeparateTCI_multiMAC_CE_r17_minBeamApplicationTime_r17_Enum_n70  uper.Enumerated = 8
	MIMO_ParametersPerBand_unifiedSeparateTCI_multiMAC_CE_r17_minBeamApplicationTime_r17_Enum_n84  uper.Enumerated = 9
	MIMO_ParametersPerBand_unifiedSeparateTCI_multiMAC_CE_r17_minBeamApplicationTime_r17_Enum_n98  uper.Enumerated = 10
	MIMO_ParametersPerBand_unifiedSeparateTCI_multiMAC_CE_r17_minBeamApplicationTime_r17_Enum_n112 uper.Enumerated = 11
	MIMO_ParametersPerBand_unifiedSeparateTCI_multiMAC_CE_r17_minBeamApplicationTime_r17_Enum_n224 uper.Enumerated = 12
	MIMO_ParametersPerBand_unifiedSeparateTCI_multiMAC_CE_r17_minBeamApplicationTime_r17_Enum_n336 uper.Enumerated = 13
)

type MIMO_ParametersPerBand_unifiedSeparateTCI_multiMAC_CE_r17_minBeamApplicationTime_r17 struct {
	Value uper.Enumerated `lb:0,ub:13,madatory`
}

func (ie *MIMO_ParametersPerBand_unifiedSeparateTCI_multiMAC_CE_r17_minBeamApplicationTime_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 13}, false); err != nil {
		return utils.WrapError("Encode MIMO_ParametersPerBand_unifiedSeparateTCI_multiMAC_CE_r17_minBeamApplicationTime_r17", err)
	}
	return nil
}

func (ie *MIMO_ParametersPerBand_unifiedSeparateTCI_multiMAC_CE_r17_minBeamApplicationTime_r17) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 13}, false); err != nil {
		return utils.WrapError("Decode MIMO_ParametersPerBand_unifiedSeparateTCI_multiMAC_CE_r17_minBeamApplicationTime_r17", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
