package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	FeatureCombinationPreambles_r17_groupBconfigured_r17_messagePowerOffsetGroupB_r17_Enum_minusinfinity uper.Enumerated = 0
	FeatureCombinationPreambles_r17_groupBconfigured_r17_messagePowerOffsetGroupB_r17_Enum_dB0           uper.Enumerated = 1
	FeatureCombinationPreambles_r17_groupBconfigured_r17_messagePowerOffsetGroupB_r17_Enum_dB5           uper.Enumerated = 2
	FeatureCombinationPreambles_r17_groupBconfigured_r17_messagePowerOffsetGroupB_r17_Enum_dB8           uper.Enumerated = 3
	FeatureCombinationPreambles_r17_groupBconfigured_r17_messagePowerOffsetGroupB_r17_Enum_dB10          uper.Enumerated = 4
	FeatureCombinationPreambles_r17_groupBconfigured_r17_messagePowerOffsetGroupB_r17_Enum_dB12          uper.Enumerated = 5
	FeatureCombinationPreambles_r17_groupBconfigured_r17_messagePowerOffsetGroupB_r17_Enum_dB15          uper.Enumerated = 6
	FeatureCombinationPreambles_r17_groupBconfigured_r17_messagePowerOffsetGroupB_r17_Enum_dB18          uper.Enumerated = 7
)

type FeatureCombinationPreambles_r17_groupBconfigured_r17_messagePowerOffsetGroupB_r17 struct {
	Value uper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *FeatureCombinationPreambles_r17_groupBconfigured_r17_messagePowerOffsetGroupB_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode FeatureCombinationPreambles_r17_groupBconfigured_r17_messagePowerOffsetGroupB_r17", err)
	}
	return nil
}

func (ie *FeatureCombinationPreambles_r17_groupBconfigured_r17_messagePowerOffsetGroupB_r17) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode FeatureCombinationPreambles_r17_groupBconfigured_r17_messagePowerOffsetGroupB_r17", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
