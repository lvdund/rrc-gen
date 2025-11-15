package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type FeatureSetDownlink_v1720 struct {
	rtt_BasedPDC_CSI_RS_ForTracking_r17 *FeatureSetDownlink_v1720_rtt_BasedPDC_CSI_RS_ForTracking_r17 `optional`
	rtt_BasedPDC_PRS_r17                *FeatureSetDownlink_v1720_rtt_BasedPDC_PRS_r17                `optional`
	sps_Multicast_r17                   *FeatureSetDownlink_v1720_sps_Multicast_r17                   `optional`
}

func (ie *FeatureSetDownlink_v1720) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.rtt_BasedPDC_CSI_RS_ForTracking_r17 != nil, ie.rtt_BasedPDC_PRS_r17 != nil, ie.sps_Multicast_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.rtt_BasedPDC_CSI_RS_ForTracking_r17 != nil {
		if err = ie.rtt_BasedPDC_CSI_RS_ForTracking_r17.Encode(w); err != nil {
			return utils.WrapError("Encode rtt_BasedPDC_CSI_RS_ForTracking_r17", err)
		}
	}
	if ie.rtt_BasedPDC_PRS_r17 != nil {
		if err = ie.rtt_BasedPDC_PRS_r17.Encode(w); err != nil {
			return utils.WrapError("Encode rtt_BasedPDC_PRS_r17", err)
		}
	}
	if ie.sps_Multicast_r17 != nil {
		if err = ie.sps_Multicast_r17.Encode(w); err != nil {
			return utils.WrapError("Encode sps_Multicast_r17", err)
		}
	}
	return nil
}

func (ie *FeatureSetDownlink_v1720) Decode(r *uper.UperReader) error {
	var err error
	var rtt_BasedPDC_CSI_RS_ForTracking_r17Present bool
	if rtt_BasedPDC_CSI_RS_ForTracking_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var rtt_BasedPDC_PRS_r17Present bool
	if rtt_BasedPDC_PRS_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sps_Multicast_r17Present bool
	if sps_Multicast_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if rtt_BasedPDC_CSI_RS_ForTracking_r17Present {
		ie.rtt_BasedPDC_CSI_RS_ForTracking_r17 = new(FeatureSetDownlink_v1720_rtt_BasedPDC_CSI_RS_ForTracking_r17)
		if err = ie.rtt_BasedPDC_CSI_RS_ForTracking_r17.Decode(r); err != nil {
			return utils.WrapError("Decode rtt_BasedPDC_CSI_RS_ForTracking_r17", err)
		}
	}
	if rtt_BasedPDC_PRS_r17Present {
		ie.rtt_BasedPDC_PRS_r17 = new(FeatureSetDownlink_v1720_rtt_BasedPDC_PRS_r17)
		if err = ie.rtt_BasedPDC_PRS_r17.Decode(r); err != nil {
			return utils.WrapError("Decode rtt_BasedPDC_PRS_r17", err)
		}
	}
	if sps_Multicast_r17Present {
		ie.sps_Multicast_r17 = new(FeatureSetDownlink_v1720_sps_Multicast_r17)
		if err = ie.sps_Multicast_r17.Decode(r); err != nil {
			return utils.WrapError("Decode sps_Multicast_r17", err)
		}
	}
	return nil
}
