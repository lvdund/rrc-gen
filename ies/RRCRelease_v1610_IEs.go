package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RRCRelease_v1610_IEs struct {
	voiceFallbackIndication_r16 *RRCRelease_v1610_IEs_voiceFallbackIndication_r16 `optional`
	measIdleConfig_r16          *MeasIdleConfigDedicated_r16                      `optional,setuprelease`
	nonCriticalExtension        *RRCRelease_v1650_IEs                             `optional`
}

func (ie *RRCRelease_v1610_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.voiceFallbackIndication_r16 != nil, ie.measIdleConfig_r16 != nil, ie.nonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.voiceFallbackIndication_r16 != nil {
		if err = ie.voiceFallbackIndication_r16.Encode(w); err != nil {
			return utils.WrapError("Encode voiceFallbackIndication_r16", err)
		}
	}
	if ie.measIdleConfig_r16 != nil {
		tmp_measIdleConfig_r16 := utils.SetupRelease[*MeasIdleConfigDedicated_r16]{
			Setup: ie.measIdleConfig_r16,
		}
		if err = tmp_measIdleConfig_r16.Encode(w); err != nil {
			return utils.WrapError("Encode measIdleConfig_r16", err)
		}
	}
	if ie.nonCriticalExtension != nil {
		if err = ie.nonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode nonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *RRCRelease_v1610_IEs) Decode(r *uper.UperReader) error {
	var err error
	var voiceFallbackIndication_r16Present bool
	if voiceFallbackIndication_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var measIdleConfig_r16Present bool
	if measIdleConfig_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var nonCriticalExtensionPresent bool
	if nonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if voiceFallbackIndication_r16Present {
		ie.voiceFallbackIndication_r16 = new(RRCRelease_v1610_IEs_voiceFallbackIndication_r16)
		if err = ie.voiceFallbackIndication_r16.Decode(r); err != nil {
			return utils.WrapError("Decode voiceFallbackIndication_r16", err)
		}
	}
	if measIdleConfig_r16Present {
		tmp_measIdleConfig_r16 := utils.SetupRelease[*MeasIdleConfigDedicated_r16]{}
		if err = tmp_measIdleConfig_r16.Decode(r); err != nil {
			return utils.WrapError("Decode measIdleConfig_r16", err)
		}
		ie.measIdleConfig_r16 = tmp_measIdleConfig_r16.Setup
	}
	if nonCriticalExtensionPresent {
		ie.nonCriticalExtension = new(RRCRelease_v1650_IEs)
		if err = ie.nonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode nonCriticalExtension", err)
		}
	}
	return nil
}
