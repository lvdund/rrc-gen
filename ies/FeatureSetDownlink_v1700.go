package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type FeatureSetDownlink_v1700 struct {
	scalingFactor_1024QAM_FR1_r17    *FeatureSetDownlink_v1700_scalingFactor_1024QAM_FR1_r17    `optional`
	timeDurationForQCL_v1710         *FeatureSetDownlink_v1700_timeDurationForQCL_v1710         `optional`
	sfn_SchemeA_r17                  *FeatureSetDownlink_v1700_sfn_SchemeA_r17                  `optional`
	sfn_SchemeA_PDCCH_only_r17       *FeatureSetDownlink_v1700_sfn_SchemeA_PDCCH_only_r17       `optional`
	sfn_SchemeA_DynamicSwitching_r17 *FeatureSetDownlink_v1700_sfn_SchemeA_DynamicSwitching_r17 `optional`
	sfn_SchemeA_PDSCH_only_r17       *FeatureSetDownlink_v1700_sfn_SchemeA_PDSCH_only_r17       `optional`
	sfn_SchemeB_r17                  *FeatureSetDownlink_v1700_sfn_SchemeB_r17                  `optional`
	sfn_SchemeB_DynamicSwitching_r17 *FeatureSetDownlink_v1700_sfn_SchemeB_DynamicSwitching_r17 `optional`
	sfn_SchemeB_PDSCH_only_r17       *FeatureSetDownlink_v1700_sfn_SchemeB_PDSCH_only_r17       `optional`
	mTRP_PDCCH_Case2_1SpanGap_r17    *FeatureSetDownlink_v1700_mTRP_PDCCH_Case2_1SpanGap_r17    `optional`
	mTRP_PDCCH_legacyMonitoring_r17  *FeatureSetDownlink_v1700_mTRP_PDCCH_legacyMonitoring_r17  `optional`
	mTRP_PDCCH_multiDCI_multiTRP_r17 *FeatureSetDownlink_v1700_mTRP_PDCCH_multiDCI_multiTRP_r17 `optional`
	dynamicMulticastPCell_r17        *FeatureSetDownlink_v1700_dynamicMulticastPCell_r17        `optional`
	mTRP_PDCCH_Repetition_r17        *FeatureSetDownlink_v1700_mTRP_PDCCH_Repetition_r17        `lb:2,ub:3,optional`
}

func (ie *FeatureSetDownlink_v1700) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.scalingFactor_1024QAM_FR1_r17 != nil, ie.timeDurationForQCL_v1710 != nil, ie.sfn_SchemeA_r17 != nil, ie.sfn_SchemeA_PDCCH_only_r17 != nil, ie.sfn_SchemeA_DynamicSwitching_r17 != nil, ie.sfn_SchemeA_PDSCH_only_r17 != nil, ie.sfn_SchemeB_r17 != nil, ie.sfn_SchemeB_DynamicSwitching_r17 != nil, ie.sfn_SchemeB_PDSCH_only_r17 != nil, ie.mTRP_PDCCH_Case2_1SpanGap_r17 != nil, ie.mTRP_PDCCH_legacyMonitoring_r17 != nil, ie.mTRP_PDCCH_multiDCI_multiTRP_r17 != nil, ie.dynamicMulticastPCell_r17 != nil, ie.mTRP_PDCCH_Repetition_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.scalingFactor_1024QAM_FR1_r17 != nil {
		if err = ie.scalingFactor_1024QAM_FR1_r17.Encode(w); err != nil {
			return utils.WrapError("Encode scalingFactor_1024QAM_FR1_r17", err)
		}
	}
	if ie.timeDurationForQCL_v1710 != nil {
		if err = ie.timeDurationForQCL_v1710.Encode(w); err != nil {
			return utils.WrapError("Encode timeDurationForQCL_v1710", err)
		}
	}
	if ie.sfn_SchemeA_r17 != nil {
		if err = ie.sfn_SchemeA_r17.Encode(w); err != nil {
			return utils.WrapError("Encode sfn_SchemeA_r17", err)
		}
	}
	if ie.sfn_SchemeA_PDCCH_only_r17 != nil {
		if err = ie.sfn_SchemeA_PDCCH_only_r17.Encode(w); err != nil {
			return utils.WrapError("Encode sfn_SchemeA_PDCCH_only_r17", err)
		}
	}
	if ie.sfn_SchemeA_DynamicSwitching_r17 != nil {
		if err = ie.sfn_SchemeA_DynamicSwitching_r17.Encode(w); err != nil {
			return utils.WrapError("Encode sfn_SchemeA_DynamicSwitching_r17", err)
		}
	}
	if ie.sfn_SchemeA_PDSCH_only_r17 != nil {
		if err = ie.sfn_SchemeA_PDSCH_only_r17.Encode(w); err != nil {
			return utils.WrapError("Encode sfn_SchemeA_PDSCH_only_r17", err)
		}
	}
	if ie.sfn_SchemeB_r17 != nil {
		if err = ie.sfn_SchemeB_r17.Encode(w); err != nil {
			return utils.WrapError("Encode sfn_SchemeB_r17", err)
		}
	}
	if ie.sfn_SchemeB_DynamicSwitching_r17 != nil {
		if err = ie.sfn_SchemeB_DynamicSwitching_r17.Encode(w); err != nil {
			return utils.WrapError("Encode sfn_SchemeB_DynamicSwitching_r17", err)
		}
	}
	if ie.sfn_SchemeB_PDSCH_only_r17 != nil {
		if err = ie.sfn_SchemeB_PDSCH_only_r17.Encode(w); err != nil {
			return utils.WrapError("Encode sfn_SchemeB_PDSCH_only_r17", err)
		}
	}
	if ie.mTRP_PDCCH_Case2_1SpanGap_r17 != nil {
		if err = ie.mTRP_PDCCH_Case2_1SpanGap_r17.Encode(w); err != nil {
			return utils.WrapError("Encode mTRP_PDCCH_Case2_1SpanGap_r17", err)
		}
	}
	if ie.mTRP_PDCCH_legacyMonitoring_r17 != nil {
		if err = ie.mTRP_PDCCH_legacyMonitoring_r17.Encode(w); err != nil {
			return utils.WrapError("Encode mTRP_PDCCH_legacyMonitoring_r17", err)
		}
	}
	if ie.mTRP_PDCCH_multiDCI_multiTRP_r17 != nil {
		if err = ie.mTRP_PDCCH_multiDCI_multiTRP_r17.Encode(w); err != nil {
			return utils.WrapError("Encode mTRP_PDCCH_multiDCI_multiTRP_r17", err)
		}
	}
	if ie.dynamicMulticastPCell_r17 != nil {
		if err = ie.dynamicMulticastPCell_r17.Encode(w); err != nil {
			return utils.WrapError("Encode dynamicMulticastPCell_r17", err)
		}
	}
	if ie.mTRP_PDCCH_Repetition_r17 != nil {
		if err = ie.mTRP_PDCCH_Repetition_r17.Encode(w); err != nil {
			return utils.WrapError("Encode mTRP_PDCCH_Repetition_r17", err)
		}
	}
	return nil
}

func (ie *FeatureSetDownlink_v1700) Decode(r *uper.UperReader) error {
	var err error
	var scalingFactor_1024QAM_FR1_r17Present bool
	if scalingFactor_1024QAM_FR1_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var timeDurationForQCL_v1710Present bool
	if timeDurationForQCL_v1710Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sfn_SchemeA_r17Present bool
	if sfn_SchemeA_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sfn_SchemeA_PDCCH_only_r17Present bool
	if sfn_SchemeA_PDCCH_only_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sfn_SchemeA_DynamicSwitching_r17Present bool
	if sfn_SchemeA_DynamicSwitching_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sfn_SchemeA_PDSCH_only_r17Present bool
	if sfn_SchemeA_PDSCH_only_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sfn_SchemeB_r17Present bool
	if sfn_SchemeB_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sfn_SchemeB_DynamicSwitching_r17Present bool
	if sfn_SchemeB_DynamicSwitching_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sfn_SchemeB_PDSCH_only_r17Present bool
	if sfn_SchemeB_PDSCH_only_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var mTRP_PDCCH_Case2_1SpanGap_r17Present bool
	if mTRP_PDCCH_Case2_1SpanGap_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var mTRP_PDCCH_legacyMonitoring_r17Present bool
	if mTRP_PDCCH_legacyMonitoring_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var mTRP_PDCCH_multiDCI_multiTRP_r17Present bool
	if mTRP_PDCCH_multiDCI_multiTRP_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var dynamicMulticastPCell_r17Present bool
	if dynamicMulticastPCell_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var mTRP_PDCCH_Repetition_r17Present bool
	if mTRP_PDCCH_Repetition_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if scalingFactor_1024QAM_FR1_r17Present {
		ie.scalingFactor_1024QAM_FR1_r17 = new(FeatureSetDownlink_v1700_scalingFactor_1024QAM_FR1_r17)
		if err = ie.scalingFactor_1024QAM_FR1_r17.Decode(r); err != nil {
			return utils.WrapError("Decode scalingFactor_1024QAM_FR1_r17", err)
		}
	}
	if timeDurationForQCL_v1710Present {
		ie.timeDurationForQCL_v1710 = new(FeatureSetDownlink_v1700_timeDurationForQCL_v1710)
		if err = ie.timeDurationForQCL_v1710.Decode(r); err != nil {
			return utils.WrapError("Decode timeDurationForQCL_v1710", err)
		}
	}
	if sfn_SchemeA_r17Present {
		ie.sfn_SchemeA_r17 = new(FeatureSetDownlink_v1700_sfn_SchemeA_r17)
		if err = ie.sfn_SchemeA_r17.Decode(r); err != nil {
			return utils.WrapError("Decode sfn_SchemeA_r17", err)
		}
	}
	if sfn_SchemeA_PDCCH_only_r17Present {
		ie.sfn_SchemeA_PDCCH_only_r17 = new(FeatureSetDownlink_v1700_sfn_SchemeA_PDCCH_only_r17)
		if err = ie.sfn_SchemeA_PDCCH_only_r17.Decode(r); err != nil {
			return utils.WrapError("Decode sfn_SchemeA_PDCCH_only_r17", err)
		}
	}
	if sfn_SchemeA_DynamicSwitching_r17Present {
		ie.sfn_SchemeA_DynamicSwitching_r17 = new(FeatureSetDownlink_v1700_sfn_SchemeA_DynamicSwitching_r17)
		if err = ie.sfn_SchemeA_DynamicSwitching_r17.Decode(r); err != nil {
			return utils.WrapError("Decode sfn_SchemeA_DynamicSwitching_r17", err)
		}
	}
	if sfn_SchemeA_PDSCH_only_r17Present {
		ie.sfn_SchemeA_PDSCH_only_r17 = new(FeatureSetDownlink_v1700_sfn_SchemeA_PDSCH_only_r17)
		if err = ie.sfn_SchemeA_PDSCH_only_r17.Decode(r); err != nil {
			return utils.WrapError("Decode sfn_SchemeA_PDSCH_only_r17", err)
		}
	}
	if sfn_SchemeB_r17Present {
		ie.sfn_SchemeB_r17 = new(FeatureSetDownlink_v1700_sfn_SchemeB_r17)
		if err = ie.sfn_SchemeB_r17.Decode(r); err != nil {
			return utils.WrapError("Decode sfn_SchemeB_r17", err)
		}
	}
	if sfn_SchemeB_DynamicSwitching_r17Present {
		ie.sfn_SchemeB_DynamicSwitching_r17 = new(FeatureSetDownlink_v1700_sfn_SchemeB_DynamicSwitching_r17)
		if err = ie.sfn_SchemeB_DynamicSwitching_r17.Decode(r); err != nil {
			return utils.WrapError("Decode sfn_SchemeB_DynamicSwitching_r17", err)
		}
	}
	if sfn_SchemeB_PDSCH_only_r17Present {
		ie.sfn_SchemeB_PDSCH_only_r17 = new(FeatureSetDownlink_v1700_sfn_SchemeB_PDSCH_only_r17)
		if err = ie.sfn_SchemeB_PDSCH_only_r17.Decode(r); err != nil {
			return utils.WrapError("Decode sfn_SchemeB_PDSCH_only_r17", err)
		}
	}
	if mTRP_PDCCH_Case2_1SpanGap_r17Present {
		ie.mTRP_PDCCH_Case2_1SpanGap_r17 = new(FeatureSetDownlink_v1700_mTRP_PDCCH_Case2_1SpanGap_r17)
		if err = ie.mTRP_PDCCH_Case2_1SpanGap_r17.Decode(r); err != nil {
			return utils.WrapError("Decode mTRP_PDCCH_Case2_1SpanGap_r17", err)
		}
	}
	if mTRP_PDCCH_legacyMonitoring_r17Present {
		ie.mTRP_PDCCH_legacyMonitoring_r17 = new(FeatureSetDownlink_v1700_mTRP_PDCCH_legacyMonitoring_r17)
		if err = ie.mTRP_PDCCH_legacyMonitoring_r17.Decode(r); err != nil {
			return utils.WrapError("Decode mTRP_PDCCH_legacyMonitoring_r17", err)
		}
	}
	if mTRP_PDCCH_multiDCI_multiTRP_r17Present {
		ie.mTRP_PDCCH_multiDCI_multiTRP_r17 = new(FeatureSetDownlink_v1700_mTRP_PDCCH_multiDCI_multiTRP_r17)
		if err = ie.mTRP_PDCCH_multiDCI_multiTRP_r17.Decode(r); err != nil {
			return utils.WrapError("Decode mTRP_PDCCH_multiDCI_multiTRP_r17", err)
		}
	}
	if dynamicMulticastPCell_r17Present {
		ie.dynamicMulticastPCell_r17 = new(FeatureSetDownlink_v1700_dynamicMulticastPCell_r17)
		if err = ie.dynamicMulticastPCell_r17.Decode(r); err != nil {
			return utils.WrapError("Decode dynamicMulticastPCell_r17", err)
		}
	}
	if mTRP_PDCCH_Repetition_r17Present {
		ie.mTRP_PDCCH_Repetition_r17 = new(FeatureSetDownlink_v1700_mTRP_PDCCH_Repetition_r17)
		if err = ie.mTRP_PDCCH_Repetition_r17.Decode(r); err != nil {
			return utils.WrapError("Decode mTRP_PDCCH_Repetition_r17", err)
		}
	}
	return nil
}
