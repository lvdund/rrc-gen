package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MBS_RNTI_SpecificConfig_r17 struct {
	mbs_RNTI_SpecificConfigId_r17     MBS_RNTI_SpecificConfigId_r17                                  `madatory`
	groupCommon_RNTI_r17              MBS_RNTI_SpecificConfig_r17_groupCommon_RNTI_r17               `madatory`
	drx_ConfigPTM_r17                 *DRX_ConfigPTM_r17                                             `optional,setuprelease`
	harq_FeedbackEnablerMulticast_r17 *MBS_RNTI_SpecificConfig_r17_harq_FeedbackEnablerMulticast_r17 `optional`
	harq_FeedbackOptionMulticast_r17  *MBS_RNTI_SpecificConfig_r17_harq_FeedbackOptionMulticast_r17  `optional`
	pdsch_AggregationFactor_r17       *MBS_RNTI_SpecificConfig_r17_pdsch_AggregationFactor_r17       `optional`
}

func (ie *MBS_RNTI_SpecificConfig_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.drx_ConfigPTM_r17 != nil, ie.harq_FeedbackEnablerMulticast_r17 != nil, ie.harq_FeedbackOptionMulticast_r17 != nil, ie.pdsch_AggregationFactor_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.mbs_RNTI_SpecificConfigId_r17.Encode(w); err != nil {
		return utils.WrapError("Encode mbs_RNTI_SpecificConfigId_r17", err)
	}
	if err = ie.groupCommon_RNTI_r17.Encode(w); err != nil {
		return utils.WrapError("Encode groupCommon_RNTI_r17", err)
	}
	if ie.drx_ConfigPTM_r17 != nil {
		tmp_drx_ConfigPTM_r17 := utils.SetupRelease[*DRX_ConfigPTM_r17]{
			Setup: ie.drx_ConfigPTM_r17,
		}
		if err = tmp_drx_ConfigPTM_r17.Encode(w); err != nil {
			return utils.WrapError("Encode drx_ConfigPTM_r17", err)
		}
	}
	if ie.harq_FeedbackEnablerMulticast_r17 != nil {
		if err = ie.harq_FeedbackEnablerMulticast_r17.Encode(w); err != nil {
			return utils.WrapError("Encode harq_FeedbackEnablerMulticast_r17", err)
		}
	}
	if ie.harq_FeedbackOptionMulticast_r17 != nil {
		if err = ie.harq_FeedbackOptionMulticast_r17.Encode(w); err != nil {
			return utils.WrapError("Encode harq_FeedbackOptionMulticast_r17", err)
		}
	}
	if ie.pdsch_AggregationFactor_r17 != nil {
		if err = ie.pdsch_AggregationFactor_r17.Encode(w); err != nil {
			return utils.WrapError("Encode pdsch_AggregationFactor_r17", err)
		}
	}
	return nil
}

func (ie *MBS_RNTI_SpecificConfig_r17) Decode(r *uper.UperReader) error {
	var err error
	var drx_ConfigPTM_r17Present bool
	if drx_ConfigPTM_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var harq_FeedbackEnablerMulticast_r17Present bool
	if harq_FeedbackEnablerMulticast_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var harq_FeedbackOptionMulticast_r17Present bool
	if harq_FeedbackOptionMulticast_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var pdsch_AggregationFactor_r17Present bool
	if pdsch_AggregationFactor_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.mbs_RNTI_SpecificConfigId_r17.Decode(r); err != nil {
		return utils.WrapError("Decode mbs_RNTI_SpecificConfigId_r17", err)
	}
	if err = ie.groupCommon_RNTI_r17.Decode(r); err != nil {
		return utils.WrapError("Decode groupCommon_RNTI_r17", err)
	}
	if drx_ConfigPTM_r17Present {
		tmp_drx_ConfigPTM_r17 := utils.SetupRelease[*DRX_ConfigPTM_r17]{}
		if err = tmp_drx_ConfigPTM_r17.Decode(r); err != nil {
			return utils.WrapError("Decode drx_ConfigPTM_r17", err)
		}
		ie.drx_ConfigPTM_r17 = tmp_drx_ConfigPTM_r17.Setup
	}
	if harq_FeedbackEnablerMulticast_r17Present {
		ie.harq_FeedbackEnablerMulticast_r17 = new(MBS_RNTI_SpecificConfig_r17_harq_FeedbackEnablerMulticast_r17)
		if err = ie.harq_FeedbackEnablerMulticast_r17.Decode(r); err != nil {
			return utils.WrapError("Decode harq_FeedbackEnablerMulticast_r17", err)
		}
	}
	if harq_FeedbackOptionMulticast_r17Present {
		ie.harq_FeedbackOptionMulticast_r17 = new(MBS_RNTI_SpecificConfig_r17_harq_FeedbackOptionMulticast_r17)
		if err = ie.harq_FeedbackOptionMulticast_r17.Decode(r); err != nil {
			return utils.WrapError("Decode harq_FeedbackOptionMulticast_r17", err)
		}
	}
	if pdsch_AggregationFactor_r17Present {
		ie.pdsch_AggregationFactor_r17 = new(MBS_RNTI_SpecificConfig_r17_pdsch_AggregationFactor_r17)
		if err = ie.pdsch_AggregationFactor_r17.Decode(r); err != nil {
			return utils.WrapError("Decode pdsch_AggregationFactor_r17", err)
		}
	}
	return nil
}
