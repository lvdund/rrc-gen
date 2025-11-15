package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	FeatureSetDownlink_v1540_pdcch_MonitoringAnyOccasionsWithSpanGap_scs_15kHz_Enum_set1 uper.Enumerated = 0
	FeatureSetDownlink_v1540_pdcch_MonitoringAnyOccasionsWithSpanGap_scs_15kHz_Enum_set2 uper.Enumerated = 1
	FeatureSetDownlink_v1540_pdcch_MonitoringAnyOccasionsWithSpanGap_scs_15kHz_Enum_set3 uper.Enumerated = 2
)

type FeatureSetDownlink_v1540_pdcch_MonitoringAnyOccasionsWithSpanGap_scs_15kHz struct {
	Value uper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *FeatureSetDownlink_v1540_pdcch_MonitoringAnyOccasionsWithSpanGap_scs_15kHz) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode FeatureSetDownlink_v1540_pdcch_MonitoringAnyOccasionsWithSpanGap_scs_15kHz", err)
	}
	return nil
}

func (ie *FeatureSetDownlink_v1540_pdcch_MonitoringAnyOccasionsWithSpanGap_scs_15kHz) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode FeatureSetDownlink_v1540_pdcch_MonitoringAnyOccasionsWithSpanGap_scs_15kHz", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
