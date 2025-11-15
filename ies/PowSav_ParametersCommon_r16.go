package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PowSav_ParametersCommon_r16 struct {
	drx_Preference_r16                *PowSav_ParametersCommon_r16_drx_Preference_r16                `optional`
	maxCC_Preference_r16              *PowSav_ParametersCommon_r16_maxCC_Preference_r16              `optional`
	releasePreference_r16             *PowSav_ParametersCommon_r16_releasePreference_r16             `optional`
	minSchedulingOffsetPreference_r16 *PowSav_ParametersCommon_r16_minSchedulingOffsetPreference_r16 `optional`
}

func (ie *PowSav_ParametersCommon_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.drx_Preference_r16 != nil, ie.maxCC_Preference_r16 != nil, ie.releasePreference_r16 != nil, ie.minSchedulingOffsetPreference_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.drx_Preference_r16 != nil {
		if err = ie.drx_Preference_r16.Encode(w); err != nil {
			return utils.WrapError("Encode drx_Preference_r16", err)
		}
	}
	if ie.maxCC_Preference_r16 != nil {
		if err = ie.maxCC_Preference_r16.Encode(w); err != nil {
			return utils.WrapError("Encode maxCC_Preference_r16", err)
		}
	}
	if ie.releasePreference_r16 != nil {
		if err = ie.releasePreference_r16.Encode(w); err != nil {
			return utils.WrapError("Encode releasePreference_r16", err)
		}
	}
	if ie.minSchedulingOffsetPreference_r16 != nil {
		if err = ie.minSchedulingOffsetPreference_r16.Encode(w); err != nil {
			return utils.WrapError("Encode minSchedulingOffsetPreference_r16", err)
		}
	}
	return nil
}

func (ie *PowSav_ParametersCommon_r16) Decode(r *uper.UperReader) error {
	var err error
	var drx_Preference_r16Present bool
	if drx_Preference_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var maxCC_Preference_r16Present bool
	if maxCC_Preference_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var releasePreference_r16Present bool
	if releasePreference_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var minSchedulingOffsetPreference_r16Present bool
	if minSchedulingOffsetPreference_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if drx_Preference_r16Present {
		ie.drx_Preference_r16 = new(PowSav_ParametersCommon_r16_drx_Preference_r16)
		if err = ie.drx_Preference_r16.Decode(r); err != nil {
			return utils.WrapError("Decode drx_Preference_r16", err)
		}
	}
	if maxCC_Preference_r16Present {
		ie.maxCC_Preference_r16 = new(PowSav_ParametersCommon_r16_maxCC_Preference_r16)
		if err = ie.maxCC_Preference_r16.Decode(r); err != nil {
			return utils.WrapError("Decode maxCC_Preference_r16", err)
		}
	}
	if releasePreference_r16Present {
		ie.releasePreference_r16 = new(PowSav_ParametersCommon_r16_releasePreference_r16)
		if err = ie.releasePreference_r16.Decode(r); err != nil {
			return utils.WrapError("Decode releasePreference_r16", err)
		}
	}
	if minSchedulingOffsetPreference_r16Present {
		ie.minSchedulingOffsetPreference_r16 = new(PowSav_ParametersCommon_r16_minSchedulingOffsetPreference_r16)
		if err = ie.minSchedulingOffsetPreference_r16.Decode(r); err != nil {
			return utils.WrapError("Decode minSchedulingOffsetPreference_r16", err)
		}
	}
	return nil
}
