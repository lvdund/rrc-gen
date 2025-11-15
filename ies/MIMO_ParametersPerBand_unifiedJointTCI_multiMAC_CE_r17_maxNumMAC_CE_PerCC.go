package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	MIMO_ParametersPerBand_unifiedJointTCI_multiMAC_CE_r17_maxNumMAC_CE_PerCC_Enum_n2 uper.Enumerated = 0
	MIMO_ParametersPerBand_unifiedJointTCI_multiMAC_CE_r17_maxNumMAC_CE_PerCC_Enum_n3 uper.Enumerated = 1
	MIMO_ParametersPerBand_unifiedJointTCI_multiMAC_CE_r17_maxNumMAC_CE_PerCC_Enum_n4 uper.Enumerated = 2
	MIMO_ParametersPerBand_unifiedJointTCI_multiMAC_CE_r17_maxNumMAC_CE_PerCC_Enum_n5 uper.Enumerated = 3
	MIMO_ParametersPerBand_unifiedJointTCI_multiMAC_CE_r17_maxNumMAC_CE_PerCC_Enum_n6 uper.Enumerated = 4
	MIMO_ParametersPerBand_unifiedJointTCI_multiMAC_CE_r17_maxNumMAC_CE_PerCC_Enum_n7 uper.Enumerated = 5
	MIMO_ParametersPerBand_unifiedJointTCI_multiMAC_CE_r17_maxNumMAC_CE_PerCC_Enum_n8 uper.Enumerated = 6
)

type MIMO_ParametersPerBand_unifiedJointTCI_multiMAC_CE_r17_maxNumMAC_CE_PerCC struct {
	Value uper.Enumerated `lb:0,ub:6,madatory`
}

func (ie *MIMO_ParametersPerBand_unifiedJointTCI_multiMAC_CE_r17_maxNumMAC_CE_PerCC) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 6}, false); err != nil {
		return utils.WrapError("Encode MIMO_ParametersPerBand_unifiedJointTCI_multiMAC_CE_r17_maxNumMAC_CE_PerCC", err)
	}
	return nil
}

func (ie *MIMO_ParametersPerBand_unifiedJointTCI_multiMAC_CE_r17_maxNumMAC_CE_PerCC) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 6}, false); err != nil {
		return utils.WrapError("Decode MIMO_ParametersPerBand_unifiedJointTCI_multiMAC_CE_r17_maxNumMAC_CE_PerCC", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
