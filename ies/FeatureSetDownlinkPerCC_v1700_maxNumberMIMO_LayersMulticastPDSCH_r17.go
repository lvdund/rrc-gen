package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	FeatureSetDownlinkPerCC_v1700_maxNumberMIMO_LayersMulticastPDSCH_r17_Enum_n2 uper.Enumerated = 0
	FeatureSetDownlinkPerCC_v1700_maxNumberMIMO_LayersMulticastPDSCH_r17_Enum_n4 uper.Enumerated = 1
	FeatureSetDownlinkPerCC_v1700_maxNumberMIMO_LayersMulticastPDSCH_r17_Enum_n8 uper.Enumerated = 2
)

type FeatureSetDownlinkPerCC_v1700_maxNumberMIMO_LayersMulticastPDSCH_r17 struct {
	Value uper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *FeatureSetDownlinkPerCC_v1700_maxNumberMIMO_LayersMulticastPDSCH_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode FeatureSetDownlinkPerCC_v1700_maxNumberMIMO_LayersMulticastPDSCH_r17", err)
	}
	return nil
}

func (ie *FeatureSetDownlinkPerCC_v1700_maxNumberMIMO_LayersMulticastPDSCH_r17) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode FeatureSetDownlinkPerCC_v1700_maxNumberMIMO_LayersMulticastPDSCH_r17", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
