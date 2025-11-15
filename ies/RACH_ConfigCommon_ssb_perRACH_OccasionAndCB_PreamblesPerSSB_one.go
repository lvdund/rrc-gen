package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	RACH_ConfigCommon_ssb_perRACH_OccasionAndCB_PreamblesPerSSB_one_Enum_n4  uper.Enumerated = 0
	RACH_ConfigCommon_ssb_perRACH_OccasionAndCB_PreamblesPerSSB_one_Enum_n8  uper.Enumerated = 1
	RACH_ConfigCommon_ssb_perRACH_OccasionAndCB_PreamblesPerSSB_one_Enum_n12 uper.Enumerated = 2
	RACH_ConfigCommon_ssb_perRACH_OccasionAndCB_PreamblesPerSSB_one_Enum_n16 uper.Enumerated = 3
	RACH_ConfigCommon_ssb_perRACH_OccasionAndCB_PreamblesPerSSB_one_Enum_n20 uper.Enumerated = 4
	RACH_ConfigCommon_ssb_perRACH_OccasionAndCB_PreamblesPerSSB_one_Enum_n24 uper.Enumerated = 5
	RACH_ConfigCommon_ssb_perRACH_OccasionAndCB_PreamblesPerSSB_one_Enum_n28 uper.Enumerated = 6
	RACH_ConfigCommon_ssb_perRACH_OccasionAndCB_PreamblesPerSSB_one_Enum_n32 uper.Enumerated = 7
	RACH_ConfigCommon_ssb_perRACH_OccasionAndCB_PreamblesPerSSB_one_Enum_n36 uper.Enumerated = 8
	RACH_ConfigCommon_ssb_perRACH_OccasionAndCB_PreamblesPerSSB_one_Enum_n40 uper.Enumerated = 9
	RACH_ConfigCommon_ssb_perRACH_OccasionAndCB_PreamblesPerSSB_one_Enum_n44 uper.Enumerated = 10
	RACH_ConfigCommon_ssb_perRACH_OccasionAndCB_PreamblesPerSSB_one_Enum_n48 uper.Enumerated = 11
	RACH_ConfigCommon_ssb_perRACH_OccasionAndCB_PreamblesPerSSB_one_Enum_n52 uper.Enumerated = 12
	RACH_ConfigCommon_ssb_perRACH_OccasionAndCB_PreamblesPerSSB_one_Enum_n56 uper.Enumerated = 13
	RACH_ConfigCommon_ssb_perRACH_OccasionAndCB_PreamblesPerSSB_one_Enum_n60 uper.Enumerated = 14
	RACH_ConfigCommon_ssb_perRACH_OccasionAndCB_PreamblesPerSSB_one_Enum_n64 uper.Enumerated = 15
)

type RACH_ConfigCommon_ssb_perRACH_OccasionAndCB_PreamblesPerSSB_one struct {
	Value uper.Enumerated `lb:0,ub:15,madatory`
}

func (ie *RACH_ConfigCommon_ssb_perRACH_OccasionAndCB_PreamblesPerSSB_one) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Encode RACH_ConfigCommon_ssb_perRACH_OccasionAndCB_PreamblesPerSSB_one", err)
	}
	return nil
}

func (ie *RACH_ConfigCommon_ssb_perRACH_OccasionAndCB_PreamblesPerSSB_one) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Decode RACH_ConfigCommon_ssb_perRACH_OccasionAndCB_PreamblesPerSSB_one", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
