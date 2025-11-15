package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type AppLayerMeasParameters_r17 struct {
	qoe_Streaming_MeasReport_r17            *AppLayerMeasParameters_r17_qoe_Streaming_MeasReport_r17            `optional`
	qoe_MTSI_MeasReport_r17                 *AppLayerMeasParameters_r17_qoe_MTSI_MeasReport_r17                 `optional`
	qoe_VR_MeasReport_r17                   *AppLayerMeasParameters_r17_qoe_VR_MeasReport_r17                   `optional`
	ran_VisibleQoE_Streaming_MeasReport_r17 *AppLayerMeasParameters_r17_ran_VisibleQoE_Streaming_MeasReport_r17 `optional`
	ran_VisibleQoE_VR_MeasReport_r17        *AppLayerMeasParameters_r17_ran_VisibleQoE_VR_MeasReport_r17        `optional`
	ul_MeasurementReportAppLayer_Seg_r17    *AppLayerMeasParameters_r17_ul_MeasurementReportAppLayer_Seg_r17    `optional`
}

func (ie *AppLayerMeasParameters_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.qoe_Streaming_MeasReport_r17 != nil, ie.qoe_MTSI_MeasReport_r17 != nil, ie.qoe_VR_MeasReport_r17 != nil, ie.ran_VisibleQoE_Streaming_MeasReport_r17 != nil, ie.ran_VisibleQoE_VR_MeasReport_r17 != nil, ie.ul_MeasurementReportAppLayer_Seg_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.qoe_Streaming_MeasReport_r17 != nil {
		if err = ie.qoe_Streaming_MeasReport_r17.Encode(w); err != nil {
			return utils.WrapError("Encode qoe_Streaming_MeasReport_r17", err)
		}
	}
	if ie.qoe_MTSI_MeasReport_r17 != nil {
		if err = ie.qoe_MTSI_MeasReport_r17.Encode(w); err != nil {
			return utils.WrapError("Encode qoe_MTSI_MeasReport_r17", err)
		}
	}
	if ie.qoe_VR_MeasReport_r17 != nil {
		if err = ie.qoe_VR_MeasReport_r17.Encode(w); err != nil {
			return utils.WrapError("Encode qoe_VR_MeasReport_r17", err)
		}
	}
	if ie.ran_VisibleQoE_Streaming_MeasReport_r17 != nil {
		if err = ie.ran_VisibleQoE_Streaming_MeasReport_r17.Encode(w); err != nil {
			return utils.WrapError("Encode ran_VisibleQoE_Streaming_MeasReport_r17", err)
		}
	}
	if ie.ran_VisibleQoE_VR_MeasReport_r17 != nil {
		if err = ie.ran_VisibleQoE_VR_MeasReport_r17.Encode(w); err != nil {
			return utils.WrapError("Encode ran_VisibleQoE_VR_MeasReport_r17", err)
		}
	}
	if ie.ul_MeasurementReportAppLayer_Seg_r17 != nil {
		if err = ie.ul_MeasurementReportAppLayer_Seg_r17.Encode(w); err != nil {
			return utils.WrapError("Encode ul_MeasurementReportAppLayer_Seg_r17", err)
		}
	}
	return nil
}

func (ie *AppLayerMeasParameters_r17) Decode(r *uper.UperReader) error {
	var err error
	var qoe_Streaming_MeasReport_r17Present bool
	if qoe_Streaming_MeasReport_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var qoe_MTSI_MeasReport_r17Present bool
	if qoe_MTSI_MeasReport_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var qoe_VR_MeasReport_r17Present bool
	if qoe_VR_MeasReport_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ran_VisibleQoE_Streaming_MeasReport_r17Present bool
	if ran_VisibleQoE_Streaming_MeasReport_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ran_VisibleQoE_VR_MeasReport_r17Present bool
	if ran_VisibleQoE_VR_MeasReport_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ul_MeasurementReportAppLayer_Seg_r17Present bool
	if ul_MeasurementReportAppLayer_Seg_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if qoe_Streaming_MeasReport_r17Present {
		ie.qoe_Streaming_MeasReport_r17 = new(AppLayerMeasParameters_r17_qoe_Streaming_MeasReport_r17)
		if err = ie.qoe_Streaming_MeasReport_r17.Decode(r); err != nil {
			return utils.WrapError("Decode qoe_Streaming_MeasReport_r17", err)
		}
	}
	if qoe_MTSI_MeasReport_r17Present {
		ie.qoe_MTSI_MeasReport_r17 = new(AppLayerMeasParameters_r17_qoe_MTSI_MeasReport_r17)
		if err = ie.qoe_MTSI_MeasReport_r17.Decode(r); err != nil {
			return utils.WrapError("Decode qoe_MTSI_MeasReport_r17", err)
		}
	}
	if qoe_VR_MeasReport_r17Present {
		ie.qoe_VR_MeasReport_r17 = new(AppLayerMeasParameters_r17_qoe_VR_MeasReport_r17)
		if err = ie.qoe_VR_MeasReport_r17.Decode(r); err != nil {
			return utils.WrapError("Decode qoe_VR_MeasReport_r17", err)
		}
	}
	if ran_VisibleQoE_Streaming_MeasReport_r17Present {
		ie.ran_VisibleQoE_Streaming_MeasReport_r17 = new(AppLayerMeasParameters_r17_ran_VisibleQoE_Streaming_MeasReport_r17)
		if err = ie.ran_VisibleQoE_Streaming_MeasReport_r17.Decode(r); err != nil {
			return utils.WrapError("Decode ran_VisibleQoE_Streaming_MeasReport_r17", err)
		}
	}
	if ran_VisibleQoE_VR_MeasReport_r17Present {
		ie.ran_VisibleQoE_VR_MeasReport_r17 = new(AppLayerMeasParameters_r17_ran_VisibleQoE_VR_MeasReport_r17)
		if err = ie.ran_VisibleQoE_VR_MeasReport_r17.Decode(r); err != nil {
			return utils.WrapError("Decode ran_VisibleQoE_VR_MeasReport_r17", err)
		}
	}
	if ul_MeasurementReportAppLayer_Seg_r17Present {
		ie.ul_MeasurementReportAppLayer_Seg_r17 = new(AppLayerMeasParameters_r17_ul_MeasurementReportAppLayer_Seg_r17)
		if err = ie.ul_MeasurementReportAppLayer_Seg_r17.Decode(r); err != nil {
			return utils.WrapError("Decode ul_MeasurementReportAppLayer_Seg_r17", err)
		}
	}
	return nil
}
