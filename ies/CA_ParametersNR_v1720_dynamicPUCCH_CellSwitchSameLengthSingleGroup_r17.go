package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CA_ParametersNR_v1720_dynamicPUCCH_CellSwitchSameLengthSingleGroup_r17 struct {
	pucch_Group_r17        CA_ParametersNR_v1720_dynamicPUCCH_CellSwitchSameLengthSingleGroup_r17_pucch_Group_r17 `madatory`
	pucch_Group_Config_r17 PUCCH_Group_Config_r17                                                                 `madatory`
}

func (ie *CA_ParametersNR_v1720_dynamicPUCCH_CellSwitchSameLengthSingleGroup_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.pucch_Group_r17.Encode(w); err != nil {
		return utils.WrapError("Encode pucch_Group_r17", err)
	}
	if err = ie.pucch_Group_Config_r17.Encode(w); err != nil {
		return utils.WrapError("Encode pucch_Group_Config_r17", err)
	}
	return nil
}

func (ie *CA_ParametersNR_v1720_dynamicPUCCH_CellSwitchSameLengthSingleGroup_r17) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.pucch_Group_r17.Decode(r); err != nil {
		return utils.WrapError("Decode pucch_Group_r17", err)
	}
	if err = ie.pucch_Group_Config_r17.Decode(r); err != nil {
		return utils.WrapError("Decode pucch_Group_Config_r17", err)
	}
	return nil
}
