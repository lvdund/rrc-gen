package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	FeatureSetUplink_v1720_phy_PrioritizationHighPriorityDG_LowPriorityCG_r17_additionalCancellationTime_r17_scs_60kHz_r17_Enum_sym0 uper.Enumerated = 0
	FeatureSetUplink_v1720_phy_PrioritizationHighPriorityDG_LowPriorityCG_r17_additionalCancellationTime_r17_scs_60kHz_r17_Enum_sym1 uper.Enumerated = 1
	FeatureSetUplink_v1720_phy_PrioritizationHighPriorityDG_LowPriorityCG_r17_additionalCancellationTime_r17_scs_60kHz_r17_Enum_sym2 uper.Enumerated = 2
	FeatureSetUplink_v1720_phy_PrioritizationHighPriorityDG_LowPriorityCG_r17_additionalCancellationTime_r17_scs_60kHz_r17_Enum_sym3 uper.Enumerated = 3
	FeatureSetUplink_v1720_phy_PrioritizationHighPriorityDG_LowPriorityCG_r17_additionalCancellationTime_r17_scs_60kHz_r17_Enum_sym4 uper.Enumerated = 4
	FeatureSetUplink_v1720_phy_PrioritizationHighPriorityDG_LowPriorityCG_r17_additionalCancellationTime_r17_scs_60kHz_r17_Enum_sym5 uper.Enumerated = 5
	FeatureSetUplink_v1720_phy_PrioritizationHighPriorityDG_LowPriorityCG_r17_additionalCancellationTime_r17_scs_60kHz_r17_Enum_sym6 uper.Enumerated = 6
	FeatureSetUplink_v1720_phy_PrioritizationHighPriorityDG_LowPriorityCG_r17_additionalCancellationTime_r17_scs_60kHz_r17_Enum_sym7 uper.Enumerated = 7
	FeatureSetUplink_v1720_phy_PrioritizationHighPriorityDG_LowPriorityCG_r17_additionalCancellationTime_r17_scs_60kHz_r17_Enum_sym8 uper.Enumerated = 8
)

type FeatureSetUplink_v1720_phy_PrioritizationHighPriorityDG_LowPriorityCG_r17_additionalCancellationTime_r17_scs_60kHz_r17 struct {
	Value uper.Enumerated `lb:0,ub:8,madatory`
}

func (ie *FeatureSetUplink_v1720_phy_PrioritizationHighPriorityDG_LowPriorityCG_r17_additionalCancellationTime_r17_scs_60kHz_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 8}, false); err != nil {
		return utils.WrapError("Encode FeatureSetUplink_v1720_phy_PrioritizationHighPriorityDG_LowPriorityCG_r17_additionalCancellationTime_r17_scs_60kHz_r17", err)
	}
	return nil
}

func (ie *FeatureSetUplink_v1720_phy_PrioritizationHighPriorityDG_LowPriorityCG_r17_additionalCancellationTime_r17_scs_60kHz_r17) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 8}, false); err != nil {
		return utils.WrapError("Decode FeatureSetUplink_v1720_phy_PrioritizationHighPriorityDG_LowPriorityCG_r17_additionalCancellationTime_r17_scs_60kHz_r17", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
