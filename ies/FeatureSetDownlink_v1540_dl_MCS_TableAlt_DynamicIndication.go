package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	FeatureSetDownlink_v1540_dl_MCS_TableAlt_DynamicIndication_Enum_supported uper.Enumerated = 0
)

type FeatureSetDownlink_v1540_dl_MCS_TableAlt_DynamicIndication struct {
	Value uper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *FeatureSetDownlink_v1540_dl_MCS_TableAlt_DynamicIndication) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode FeatureSetDownlink_v1540_dl_MCS_TableAlt_DynamicIndication", err)
	}
	return nil
}

func (ie *FeatureSetDownlink_v1540_dl_MCS_TableAlt_DynamicIndication) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode FeatureSetDownlink_v1540_dl_MCS_TableAlt_DynamicIndication", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
