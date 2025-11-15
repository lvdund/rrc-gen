package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	RACH_ConfigCommon_groupBconfigured_messagePowerOffsetGroupB_Enum_minusinfinity uper.Enumerated = 0
	RACH_ConfigCommon_groupBconfigured_messagePowerOffsetGroupB_Enum_dB0           uper.Enumerated = 1
	RACH_ConfigCommon_groupBconfigured_messagePowerOffsetGroupB_Enum_dB5           uper.Enumerated = 2
	RACH_ConfigCommon_groupBconfigured_messagePowerOffsetGroupB_Enum_dB8           uper.Enumerated = 3
	RACH_ConfigCommon_groupBconfigured_messagePowerOffsetGroupB_Enum_dB10          uper.Enumerated = 4
	RACH_ConfigCommon_groupBconfigured_messagePowerOffsetGroupB_Enum_dB12          uper.Enumerated = 5
	RACH_ConfigCommon_groupBconfigured_messagePowerOffsetGroupB_Enum_dB15          uper.Enumerated = 6
	RACH_ConfigCommon_groupBconfigured_messagePowerOffsetGroupB_Enum_dB18          uper.Enumerated = 7
)

type RACH_ConfigCommon_groupBconfigured_messagePowerOffsetGroupB struct {
	Value uper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *RACH_ConfigCommon_groupBconfigured_messagePowerOffsetGroupB) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode RACH_ConfigCommon_groupBconfigured_messagePowerOffsetGroupB", err)
	}
	return nil
}

func (ie *RACH_ConfigCommon_groupBconfigured_messagePowerOffsetGroupB) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode RACH_ConfigCommon_groupBconfigured_messagePowerOffsetGroupB", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
