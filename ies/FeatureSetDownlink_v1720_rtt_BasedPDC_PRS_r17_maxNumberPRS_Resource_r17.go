package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	FeatureSetDownlink_v1720_rtt_BasedPDC_PRS_r17_maxNumberPRS_Resource_r17_Enum_n1  uper.Enumerated = 0
	FeatureSetDownlink_v1720_rtt_BasedPDC_PRS_r17_maxNumberPRS_Resource_r17_Enum_n2  uper.Enumerated = 1
	FeatureSetDownlink_v1720_rtt_BasedPDC_PRS_r17_maxNumberPRS_Resource_r17_Enum_n4  uper.Enumerated = 2
	FeatureSetDownlink_v1720_rtt_BasedPDC_PRS_r17_maxNumberPRS_Resource_r17_Enum_n8  uper.Enumerated = 3
	FeatureSetDownlink_v1720_rtt_BasedPDC_PRS_r17_maxNumberPRS_Resource_r17_Enum_n16 uper.Enumerated = 4
	FeatureSetDownlink_v1720_rtt_BasedPDC_PRS_r17_maxNumberPRS_Resource_r17_Enum_n32 uper.Enumerated = 5
	FeatureSetDownlink_v1720_rtt_BasedPDC_PRS_r17_maxNumberPRS_Resource_r17_Enum_n64 uper.Enumerated = 6
)

type FeatureSetDownlink_v1720_rtt_BasedPDC_PRS_r17_maxNumberPRS_Resource_r17 struct {
	Value uper.Enumerated `lb:0,ub:6,madatory`
}

func (ie *FeatureSetDownlink_v1720_rtt_BasedPDC_PRS_r17_maxNumberPRS_Resource_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 6}, false); err != nil {
		return utils.WrapError("Encode FeatureSetDownlink_v1720_rtt_BasedPDC_PRS_r17_maxNumberPRS_Resource_r17", err)
	}
	return nil
}

func (ie *FeatureSetDownlink_v1720_rtt_BasedPDC_PRS_r17_maxNumberPRS_Resource_r17) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 6}, false); err != nil {
		return utils.WrapError("Decode FeatureSetDownlink_v1720_rtt_BasedPDC_PRS_r17_maxNumberPRS_Resource_r17", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
