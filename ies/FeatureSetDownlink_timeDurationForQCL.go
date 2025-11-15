package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type FeatureSetDownlink_timeDurationForQCL struct {
	scs_60kHz  *FeatureSetDownlink_timeDurationForQCL_scs_60kHz  `optional`
	scs_120kHz *FeatureSetDownlink_timeDurationForQCL_scs_120kHz `optional`
}

func (ie *FeatureSetDownlink_timeDurationForQCL) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.scs_60kHz != nil, ie.scs_120kHz != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.scs_60kHz != nil {
		if err = ie.scs_60kHz.Encode(w); err != nil {
			return utils.WrapError("Encode scs_60kHz", err)
		}
	}
	if ie.scs_120kHz != nil {
		if err = ie.scs_120kHz.Encode(w); err != nil {
			return utils.WrapError("Encode scs_120kHz", err)
		}
	}
	return nil
}

func (ie *FeatureSetDownlink_timeDurationForQCL) Decode(r *uper.UperReader) error {
	var err error
	var scs_60kHzPresent bool
	if scs_60kHzPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var scs_120kHzPresent bool
	if scs_120kHzPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if scs_60kHzPresent {
		ie.scs_60kHz = new(FeatureSetDownlink_timeDurationForQCL_scs_60kHz)
		if err = ie.scs_60kHz.Decode(r); err != nil {
			return utils.WrapError("Decode scs_60kHz", err)
		}
	}
	if scs_120kHzPresent {
		ie.scs_120kHz = new(FeatureSetDownlink_timeDurationForQCL_scs_120kHz)
		if err = ie.scs_120kHz.Decode(r); err != nil {
			return utils.WrapError("Decode scs_120kHz", err)
		}
	}
	return nil
}
