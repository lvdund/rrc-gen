package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	FeatureSetDownlink_timeDurationForQCL_scs_120kHz_Enum_s14 uper.Enumerated = 0
	FeatureSetDownlink_timeDurationForQCL_scs_120kHz_Enum_s28 uper.Enumerated = 1
)

type FeatureSetDownlink_timeDurationForQCL_scs_120kHz struct {
	Value uper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *FeatureSetDownlink_timeDurationForQCL_scs_120kHz) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode FeatureSetDownlink_timeDurationForQCL_scs_120kHz", err)
	}
	return nil
}

func (ie *FeatureSetDownlink_timeDurationForQCL_scs_120kHz) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode FeatureSetDownlink_timeDurationForQCL_scs_120kHz", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
