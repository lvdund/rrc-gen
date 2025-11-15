package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type FeatureSetUplink_v1720 struct {
	pucch_Repetition_F0_1_2_3_4_RRC_Config_r17         *FeatureSetUplink_v1720_pucch_Repetition_F0_1_2_3_4_RRC_Config_r17         `optional`
	pucch_Repetition_F0_1_2_3_4_DynamicIndication_r17  *FeatureSetUplink_v1720_pucch_Repetition_F0_1_2_3_4_DynamicIndication_r17  `optional`
	interSubslotFreqHopping_PUCCH_r17                  *FeatureSetUplink_v1720_interSubslotFreqHopping_PUCCH_r17                  `optional`
	semiStaticHARQ_ACK_CodebookSub_SlotPUCCH_r17       *FeatureSetUplink_v1720_semiStaticHARQ_ACK_CodebookSub_SlotPUCCH_r17       `optional`
	phy_PrioritizationLowPriorityDG_HighPriorityCG_r17 *int64                                                                     `lb:1,ub:16,optional`
	phy_PrioritizationHighPriorityDG_LowPriorityCG_r17 *FeatureSetUplink_v1720_phy_PrioritizationHighPriorityDG_LowPriorityCG_r17 `lb:1,ub:16,optional`
	extendedDC_LocationReport_r17                      *FeatureSetUplink_v1720_extendedDC_LocationReport_r17                      `optional`
}

func (ie *FeatureSetUplink_v1720) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.pucch_Repetition_F0_1_2_3_4_RRC_Config_r17 != nil, ie.pucch_Repetition_F0_1_2_3_4_DynamicIndication_r17 != nil, ie.interSubslotFreqHopping_PUCCH_r17 != nil, ie.semiStaticHARQ_ACK_CodebookSub_SlotPUCCH_r17 != nil, ie.phy_PrioritizationLowPriorityDG_HighPriorityCG_r17 != nil, ie.phy_PrioritizationHighPriorityDG_LowPriorityCG_r17 != nil, ie.extendedDC_LocationReport_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.pucch_Repetition_F0_1_2_3_4_RRC_Config_r17 != nil {
		if err = ie.pucch_Repetition_F0_1_2_3_4_RRC_Config_r17.Encode(w); err != nil {
			return utils.WrapError("Encode pucch_Repetition_F0_1_2_3_4_RRC_Config_r17", err)
		}
	}
	if ie.pucch_Repetition_F0_1_2_3_4_DynamicIndication_r17 != nil {
		if err = ie.pucch_Repetition_F0_1_2_3_4_DynamicIndication_r17.Encode(w); err != nil {
			return utils.WrapError("Encode pucch_Repetition_F0_1_2_3_4_DynamicIndication_r17", err)
		}
	}
	if ie.interSubslotFreqHopping_PUCCH_r17 != nil {
		if err = ie.interSubslotFreqHopping_PUCCH_r17.Encode(w); err != nil {
			return utils.WrapError("Encode interSubslotFreqHopping_PUCCH_r17", err)
		}
	}
	if ie.semiStaticHARQ_ACK_CodebookSub_SlotPUCCH_r17 != nil {
		if err = ie.semiStaticHARQ_ACK_CodebookSub_SlotPUCCH_r17.Encode(w); err != nil {
			return utils.WrapError("Encode semiStaticHARQ_ACK_CodebookSub_SlotPUCCH_r17", err)
		}
	}
	if ie.phy_PrioritizationLowPriorityDG_HighPriorityCG_r17 != nil {
		if err = w.WriteInteger(*ie.phy_PrioritizationLowPriorityDG_HighPriorityCG_r17, &uper.Constraint{Lb: 1, Ub: 16}, false); err != nil {
			return utils.WrapError("Encode phy_PrioritizationLowPriorityDG_HighPriorityCG_r17", err)
		}
	}
	if ie.phy_PrioritizationHighPriorityDG_LowPriorityCG_r17 != nil {
		if err = ie.phy_PrioritizationHighPriorityDG_LowPriorityCG_r17.Encode(w); err != nil {
			return utils.WrapError("Encode phy_PrioritizationHighPriorityDG_LowPriorityCG_r17", err)
		}
	}
	if ie.extendedDC_LocationReport_r17 != nil {
		if err = ie.extendedDC_LocationReport_r17.Encode(w); err != nil {
			return utils.WrapError("Encode extendedDC_LocationReport_r17", err)
		}
	}
	return nil
}

func (ie *FeatureSetUplink_v1720) Decode(r *uper.UperReader) error {
	var err error
	var pucch_Repetition_F0_1_2_3_4_RRC_Config_r17Present bool
	if pucch_Repetition_F0_1_2_3_4_RRC_Config_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var pucch_Repetition_F0_1_2_3_4_DynamicIndication_r17Present bool
	if pucch_Repetition_F0_1_2_3_4_DynamicIndication_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var interSubslotFreqHopping_PUCCH_r17Present bool
	if interSubslotFreqHopping_PUCCH_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var semiStaticHARQ_ACK_CodebookSub_SlotPUCCH_r17Present bool
	if semiStaticHARQ_ACK_CodebookSub_SlotPUCCH_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var phy_PrioritizationLowPriorityDG_HighPriorityCG_r17Present bool
	if phy_PrioritizationLowPriorityDG_HighPriorityCG_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var phy_PrioritizationHighPriorityDG_LowPriorityCG_r17Present bool
	if phy_PrioritizationHighPriorityDG_LowPriorityCG_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var extendedDC_LocationReport_r17Present bool
	if extendedDC_LocationReport_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if pucch_Repetition_F0_1_2_3_4_RRC_Config_r17Present {
		ie.pucch_Repetition_F0_1_2_3_4_RRC_Config_r17 = new(FeatureSetUplink_v1720_pucch_Repetition_F0_1_2_3_4_RRC_Config_r17)
		if err = ie.pucch_Repetition_F0_1_2_3_4_RRC_Config_r17.Decode(r); err != nil {
			return utils.WrapError("Decode pucch_Repetition_F0_1_2_3_4_RRC_Config_r17", err)
		}
	}
	if pucch_Repetition_F0_1_2_3_4_DynamicIndication_r17Present {
		ie.pucch_Repetition_F0_1_2_3_4_DynamicIndication_r17 = new(FeatureSetUplink_v1720_pucch_Repetition_F0_1_2_3_4_DynamicIndication_r17)
		if err = ie.pucch_Repetition_F0_1_2_3_4_DynamicIndication_r17.Decode(r); err != nil {
			return utils.WrapError("Decode pucch_Repetition_F0_1_2_3_4_DynamicIndication_r17", err)
		}
	}
	if interSubslotFreqHopping_PUCCH_r17Present {
		ie.interSubslotFreqHopping_PUCCH_r17 = new(FeatureSetUplink_v1720_interSubslotFreqHopping_PUCCH_r17)
		if err = ie.interSubslotFreqHopping_PUCCH_r17.Decode(r); err != nil {
			return utils.WrapError("Decode interSubslotFreqHopping_PUCCH_r17", err)
		}
	}
	if semiStaticHARQ_ACK_CodebookSub_SlotPUCCH_r17Present {
		ie.semiStaticHARQ_ACK_CodebookSub_SlotPUCCH_r17 = new(FeatureSetUplink_v1720_semiStaticHARQ_ACK_CodebookSub_SlotPUCCH_r17)
		if err = ie.semiStaticHARQ_ACK_CodebookSub_SlotPUCCH_r17.Decode(r); err != nil {
			return utils.WrapError("Decode semiStaticHARQ_ACK_CodebookSub_SlotPUCCH_r17", err)
		}
	}
	if phy_PrioritizationLowPriorityDG_HighPriorityCG_r17Present {
		var tmp_int_phy_PrioritizationLowPriorityDG_HighPriorityCG_r17 int64
		if tmp_int_phy_PrioritizationLowPriorityDG_HighPriorityCG_r17, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 16}, false); err != nil {
			return utils.WrapError("Decode phy_PrioritizationLowPriorityDG_HighPriorityCG_r17", err)
		}
		ie.phy_PrioritizationLowPriorityDG_HighPriorityCG_r17 = &tmp_int_phy_PrioritizationLowPriorityDG_HighPriorityCG_r17
	}
	if phy_PrioritizationHighPriorityDG_LowPriorityCG_r17Present {
		ie.phy_PrioritizationHighPriorityDG_LowPriorityCG_r17 = new(FeatureSetUplink_v1720_phy_PrioritizationHighPriorityDG_LowPriorityCG_r17)
		if err = ie.phy_PrioritizationHighPriorityDG_LowPriorityCG_r17.Decode(r); err != nil {
			return utils.WrapError("Decode phy_PrioritizationHighPriorityDG_LowPriorityCG_r17", err)
		}
	}
	if extendedDC_LocationReport_r17Present {
		ie.extendedDC_LocationReport_r17 = new(FeatureSetUplink_v1720_extendedDC_LocationReport_r17)
		if err = ie.extendedDC_LocationReport_r17.Decode(r); err != nil {
			return utils.WrapError("Decode extendedDC_LocationReport_r17", err)
		}
	}
	return nil
}
