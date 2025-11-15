package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SDT_ConfigCommonSIB_r17 struct {
	sdt_RSRP_Threshold_r17              *RSRP_Range                                                  `optional`
	sdt_LogicalChannelSR_DelayTimer_r17 *SDT_ConfigCommonSIB_r17_sdt_LogicalChannelSR_DelayTimer_r17 `optional`
	sdt_DataVolumeThreshold_r17         SDT_ConfigCommonSIB_r17_sdt_DataVolumeThreshold_r17          `madatory`
	t319a_r17                           SDT_ConfigCommonSIB_r17_t319a_r17                            `madatory`
}

func (ie *SDT_ConfigCommonSIB_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.sdt_RSRP_Threshold_r17 != nil, ie.sdt_LogicalChannelSR_DelayTimer_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.sdt_RSRP_Threshold_r17 != nil {
		if err = ie.sdt_RSRP_Threshold_r17.Encode(w); err != nil {
			return utils.WrapError("Encode sdt_RSRP_Threshold_r17", err)
		}
	}
	if ie.sdt_LogicalChannelSR_DelayTimer_r17 != nil {
		if err = ie.sdt_LogicalChannelSR_DelayTimer_r17.Encode(w); err != nil {
			return utils.WrapError("Encode sdt_LogicalChannelSR_DelayTimer_r17", err)
		}
	}
	if err = ie.sdt_DataVolumeThreshold_r17.Encode(w); err != nil {
		return utils.WrapError("Encode sdt_DataVolumeThreshold_r17", err)
	}
	if err = ie.t319a_r17.Encode(w); err != nil {
		return utils.WrapError("Encode t319a_r17", err)
	}
	return nil
}

func (ie *SDT_ConfigCommonSIB_r17) Decode(r *uper.UperReader) error {
	var err error
	var sdt_RSRP_Threshold_r17Present bool
	if sdt_RSRP_Threshold_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sdt_LogicalChannelSR_DelayTimer_r17Present bool
	if sdt_LogicalChannelSR_DelayTimer_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if sdt_RSRP_Threshold_r17Present {
		ie.sdt_RSRP_Threshold_r17 = new(RSRP_Range)
		if err = ie.sdt_RSRP_Threshold_r17.Decode(r); err != nil {
			return utils.WrapError("Decode sdt_RSRP_Threshold_r17", err)
		}
	}
	if sdt_LogicalChannelSR_DelayTimer_r17Present {
		ie.sdt_LogicalChannelSR_DelayTimer_r17 = new(SDT_ConfigCommonSIB_r17_sdt_LogicalChannelSR_DelayTimer_r17)
		if err = ie.sdt_LogicalChannelSR_DelayTimer_r17.Decode(r); err != nil {
			return utils.WrapError("Decode sdt_LogicalChannelSR_DelayTimer_r17", err)
		}
	}
	if err = ie.sdt_DataVolumeThreshold_r17.Decode(r); err != nil {
		return utils.WrapError("Decode sdt_DataVolumeThreshold_r17", err)
	}
	if err = ie.t319a_r17.Decode(r); err != nil {
		return utils.WrapError("Decode t319a_r17", err)
	}
	return nil
}
