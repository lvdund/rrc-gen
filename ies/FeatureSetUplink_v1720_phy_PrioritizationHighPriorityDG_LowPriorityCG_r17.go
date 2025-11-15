package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type FeatureSetUplink_v1720_phy_PrioritizationHighPriorityDG_LowPriorityCG_r17 struct {
	pusch_PreparationLowPriority_r17 FeatureSetUplink_v1720_phy_PrioritizationHighPriorityDG_LowPriorityCG_r17_pusch_PreparationLowPriority_r17 `madatory`
	additionalCancellationTime_r17   *FeatureSetUplink_v1720_phy_PrioritizationHighPriorityDG_LowPriorityCG_r17_additionalCancellationTime_r17  `optional`
	maxNumberCarriers_r17            int64                                                                                                      `lb:1,ub:16,madatory`
}

func (ie *FeatureSetUplink_v1720_phy_PrioritizationHighPriorityDG_LowPriorityCG_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.additionalCancellationTime_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.pusch_PreparationLowPriority_r17.Encode(w); err != nil {
		return utils.WrapError("Encode pusch_PreparationLowPriority_r17", err)
	}
	if ie.additionalCancellationTime_r17 != nil {
		if err = ie.additionalCancellationTime_r17.Encode(w); err != nil {
			return utils.WrapError("Encode additionalCancellationTime_r17", err)
		}
	}
	if err = w.WriteInteger(ie.maxNumberCarriers_r17, &uper.Constraint{Lb: 1, Ub: 16}, false); err != nil {
		return utils.WrapError("WriteInteger maxNumberCarriers_r17", err)
	}
	return nil
}

func (ie *FeatureSetUplink_v1720_phy_PrioritizationHighPriorityDG_LowPriorityCG_r17) Decode(r *uper.UperReader) error {
	var err error
	var additionalCancellationTime_r17Present bool
	if additionalCancellationTime_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.pusch_PreparationLowPriority_r17.Decode(r); err != nil {
		return utils.WrapError("Decode pusch_PreparationLowPriority_r17", err)
	}
	if additionalCancellationTime_r17Present {
		ie.additionalCancellationTime_r17 = new(FeatureSetUplink_v1720_phy_PrioritizationHighPriorityDG_LowPriorityCG_r17_additionalCancellationTime_r17)
		if err = ie.additionalCancellationTime_r17.Decode(r); err != nil {
			return utils.WrapError("Decode additionalCancellationTime_r17", err)
		}
	}
	var tmp_int_maxNumberCarriers_r17 int64
	if tmp_int_maxNumberCarriers_r17, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 16}, false); err != nil {
		return utils.WrapError("ReadInteger maxNumberCarriers_r17", err)
	}
	ie.maxNumberCarriers_r17 = tmp_int_maxNumberCarriers_r17
	return nil
}
