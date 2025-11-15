package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	FeatureSetDownlink_type1_3_CSS_Enum_supported uper.Enumerated = 0
)

type FeatureSetDownlink_type1_3_CSS struct {
	Value uper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *FeatureSetDownlink_type1_3_CSS) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode FeatureSetDownlink_type1_3_CSS", err)
	}
	return nil
}

func (ie *FeatureSetDownlink_type1_3_CSS) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode FeatureSetDownlink_type1_3_CSS", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
