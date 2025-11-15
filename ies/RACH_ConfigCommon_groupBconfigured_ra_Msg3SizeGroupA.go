package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	RACH_ConfigCommon_groupBconfigured_ra_Msg3SizeGroupA_Enum_b56    uper.Enumerated = 0
	RACH_ConfigCommon_groupBconfigured_ra_Msg3SizeGroupA_Enum_b144   uper.Enumerated = 1
	RACH_ConfigCommon_groupBconfigured_ra_Msg3SizeGroupA_Enum_b208   uper.Enumerated = 2
	RACH_ConfigCommon_groupBconfigured_ra_Msg3SizeGroupA_Enum_b256   uper.Enumerated = 3
	RACH_ConfigCommon_groupBconfigured_ra_Msg3SizeGroupA_Enum_b282   uper.Enumerated = 4
	RACH_ConfigCommon_groupBconfigured_ra_Msg3SizeGroupA_Enum_b480   uper.Enumerated = 5
	RACH_ConfigCommon_groupBconfigured_ra_Msg3SizeGroupA_Enum_b640   uper.Enumerated = 6
	RACH_ConfigCommon_groupBconfigured_ra_Msg3SizeGroupA_Enum_b800   uper.Enumerated = 7
	RACH_ConfigCommon_groupBconfigured_ra_Msg3SizeGroupA_Enum_b1000  uper.Enumerated = 8
	RACH_ConfigCommon_groupBconfigured_ra_Msg3SizeGroupA_Enum_b72    uper.Enumerated = 9
	RACH_ConfigCommon_groupBconfigured_ra_Msg3SizeGroupA_Enum_spare6 uper.Enumerated = 10
	RACH_ConfigCommon_groupBconfigured_ra_Msg3SizeGroupA_Enum_spare5 uper.Enumerated = 11
	RACH_ConfigCommon_groupBconfigured_ra_Msg3SizeGroupA_Enum_spare4 uper.Enumerated = 12
	RACH_ConfigCommon_groupBconfigured_ra_Msg3SizeGroupA_Enum_spare3 uper.Enumerated = 13
	RACH_ConfigCommon_groupBconfigured_ra_Msg3SizeGroupA_Enum_spare2 uper.Enumerated = 14
	RACH_ConfigCommon_groupBconfigured_ra_Msg3SizeGroupA_Enum_spare1 uper.Enumerated = 15
)

type RACH_ConfigCommon_groupBconfigured_ra_Msg3SizeGroupA struct {
	Value uper.Enumerated `lb:0,ub:15,madatory`
}

func (ie *RACH_ConfigCommon_groupBconfigured_ra_Msg3SizeGroupA) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Encode RACH_ConfigCommon_groupBconfigured_ra_Msg3SizeGroupA", err)
	}
	return nil
}

func (ie *RACH_ConfigCommon_groupBconfigured_ra_Msg3SizeGroupA) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Decode RACH_ConfigCommon_groupBconfigured_ra_Msg3SizeGroupA", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
