package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	RACH_ConfigCommonTwoStepRA_r16_msgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_two_Enum_n4  uper.Enumerated = 0
	RACH_ConfigCommonTwoStepRA_r16_msgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_two_Enum_n8  uper.Enumerated = 1
	RACH_ConfigCommonTwoStepRA_r16_msgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_two_Enum_n12 uper.Enumerated = 2
	RACH_ConfigCommonTwoStepRA_r16_msgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_two_Enum_n16 uper.Enumerated = 3
	RACH_ConfigCommonTwoStepRA_r16_msgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_two_Enum_n20 uper.Enumerated = 4
	RACH_ConfigCommonTwoStepRA_r16_msgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_two_Enum_n24 uper.Enumerated = 5
	RACH_ConfigCommonTwoStepRA_r16_msgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_two_Enum_n28 uper.Enumerated = 6
	RACH_ConfigCommonTwoStepRA_r16_msgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_two_Enum_n32 uper.Enumerated = 7
)

type RACH_ConfigCommonTwoStepRA_r16_msgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_two struct {
	Value uper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *RACH_ConfigCommonTwoStepRA_r16_msgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_two) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode RACH_ConfigCommonTwoStepRA_r16_msgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_two", err)
	}
	return nil
}

func (ie *RACH_ConfigCommonTwoStepRA_r16_msgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_two) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode RACH_ConfigCommonTwoStepRA_r16_msgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_two", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
