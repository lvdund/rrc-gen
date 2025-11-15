package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	CA_ParametersNR_v1700_maxCC_32_DL_HARQ_ProcessFR2_2_r17_Enum_n1  uper.Enumerated = 0
	CA_ParametersNR_v1700_maxCC_32_DL_HARQ_ProcessFR2_2_r17_Enum_n2  uper.Enumerated = 1
	CA_ParametersNR_v1700_maxCC_32_DL_HARQ_ProcessFR2_2_r17_Enum_n3  uper.Enumerated = 2
	CA_ParametersNR_v1700_maxCC_32_DL_HARQ_ProcessFR2_2_r17_Enum_n4  uper.Enumerated = 3
	CA_ParametersNR_v1700_maxCC_32_DL_HARQ_ProcessFR2_2_r17_Enum_n6  uper.Enumerated = 4
	CA_ParametersNR_v1700_maxCC_32_DL_HARQ_ProcessFR2_2_r17_Enum_n8  uper.Enumerated = 5
	CA_ParametersNR_v1700_maxCC_32_DL_HARQ_ProcessFR2_2_r17_Enum_n16 uper.Enumerated = 6
	CA_ParametersNR_v1700_maxCC_32_DL_HARQ_ProcessFR2_2_r17_Enum_n32 uper.Enumerated = 7
)

type CA_ParametersNR_v1700_maxCC_32_DL_HARQ_ProcessFR2_2_r17 struct {
	Value uper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *CA_ParametersNR_v1700_maxCC_32_DL_HARQ_ProcessFR2_2_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode CA_ParametersNR_v1700_maxCC_32_DL_HARQ_ProcessFR2_2_r17", err)
	}
	return nil
}

func (ie *CA_ParametersNR_v1700_maxCC_32_DL_HARQ_ProcessFR2_2_r17) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode CA_ParametersNR_v1700_maxCC_32_DL_HARQ_ProcessFR2_2_r17", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
