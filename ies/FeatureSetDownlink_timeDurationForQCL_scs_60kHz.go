package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	FeatureSetDownlink_timeDurationForQCL_scs_60kHz_Enum_s7  uper.Enumerated = 0
	FeatureSetDownlink_timeDurationForQCL_scs_60kHz_Enum_s14 uper.Enumerated = 1
	FeatureSetDownlink_timeDurationForQCL_scs_60kHz_Enum_s28 uper.Enumerated = 2
)

type FeatureSetDownlink_timeDurationForQCL_scs_60kHz struct {
	Value uper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *FeatureSetDownlink_timeDurationForQCL_scs_60kHz) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode FeatureSetDownlink_timeDurationForQCL_scs_60kHz", err)
	}
	return nil
}

func (ie *FeatureSetDownlink_timeDurationForQCL_scs_60kHz) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode FeatureSetDownlink_timeDurationForQCL_scs_60kHz", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
