package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UE_CapabilityRequestFilterNR_v1540 struct {
	srs_SwitchingTimeRequest *UE_CapabilityRequestFilterNR_v1540_srs_SwitchingTimeRequest `optional`
	nonCriticalExtension     *UE_CapabilityRequestFilterNR_v1710                          `optional`
}

func (ie *UE_CapabilityRequestFilterNR_v1540) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.srs_SwitchingTimeRequest != nil, ie.nonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.srs_SwitchingTimeRequest != nil {
		if err = ie.srs_SwitchingTimeRequest.Encode(w); err != nil {
			return utils.WrapError("Encode srs_SwitchingTimeRequest", err)
		}
	}
	if ie.nonCriticalExtension != nil {
		if err = ie.nonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode nonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *UE_CapabilityRequestFilterNR_v1540) Decode(r *uper.UperReader) error {
	var err error
	var srs_SwitchingTimeRequestPresent bool
	if srs_SwitchingTimeRequestPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var nonCriticalExtensionPresent bool
	if nonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if srs_SwitchingTimeRequestPresent {
		ie.srs_SwitchingTimeRequest = new(UE_CapabilityRequestFilterNR_v1540_srs_SwitchingTimeRequest)
		if err = ie.srs_SwitchingTimeRequest.Decode(r); err != nil {
			return utils.WrapError("Decode srs_SwitchingTimeRequest", err)
		}
	}
	if nonCriticalExtensionPresent {
		ie.nonCriticalExtension = new(UE_CapabilityRequestFilterNR_v1710)
		if err = ie.nonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode nonCriticalExtension", err)
		}
	}
	return nil
}
