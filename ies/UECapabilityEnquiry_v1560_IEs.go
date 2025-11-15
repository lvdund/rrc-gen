package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UECapabilityEnquiry_v1560_IEs struct {
	capabilityRequestFilterCommon *UE_CapabilityRequestFilterCommon `optional`
	nonCriticalExtension          *UECapabilityEnquiry_v1610_IEs    `optional`
}

func (ie *UECapabilityEnquiry_v1560_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.capabilityRequestFilterCommon != nil, ie.nonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.capabilityRequestFilterCommon != nil {
		if err = ie.capabilityRequestFilterCommon.Encode(w); err != nil {
			return utils.WrapError("Encode capabilityRequestFilterCommon", err)
		}
	}
	if ie.nonCriticalExtension != nil {
		if err = ie.nonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode nonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *UECapabilityEnquiry_v1560_IEs) Decode(r *uper.UperReader) error {
	var err error
	var capabilityRequestFilterCommonPresent bool
	if capabilityRequestFilterCommonPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var nonCriticalExtensionPresent bool
	if nonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if capabilityRequestFilterCommonPresent {
		ie.capabilityRequestFilterCommon = new(UE_CapabilityRequestFilterCommon)
		if err = ie.capabilityRequestFilterCommon.Decode(r); err != nil {
			return utils.WrapError("Decode capabilityRequestFilterCommon", err)
		}
	}
	if nonCriticalExtensionPresent {
		ie.nonCriticalExtension = new(UECapabilityEnquiry_v1610_IEs)
		if err = ie.nonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode nonCriticalExtension", err)
		}
	}
	return nil
}
