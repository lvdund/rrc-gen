package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	FeatureSetDownlink_v1700_timeDurationForQCL_v1710_scs_480kHz_Enum_s56  uper.Enumerated = 0
	FeatureSetDownlink_v1700_timeDurationForQCL_v1710_scs_480kHz_Enum_s112 uper.Enumerated = 1
)

type FeatureSetDownlink_v1700_timeDurationForQCL_v1710_scs_480kHz struct {
	Value uper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *FeatureSetDownlink_v1700_timeDurationForQCL_v1710_scs_480kHz) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode FeatureSetDownlink_v1700_timeDurationForQCL_v1710_scs_480kHz", err)
	}
	return nil
}

func (ie *FeatureSetDownlink_v1700_timeDurationForQCL_v1710_scs_480kHz) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode FeatureSetDownlink_v1700_timeDurationForQCL_v1710_scs_480kHz", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
