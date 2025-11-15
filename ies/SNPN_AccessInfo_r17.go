package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SNPN_AccessInfo_r17 struct {
	extCH_Supported_r17            *SNPN_AccessInfo_r17_extCH_Supported_r17            `optional`
	extCH_WithoutConfigAllowed_r17 *SNPN_AccessInfo_r17_extCH_WithoutConfigAllowed_r17 `optional`
	onboardingEnabled_r17          *SNPN_AccessInfo_r17_onboardingEnabled_r17          `optional`
	imsEmergencySupportForSNPN_r17 *SNPN_AccessInfo_r17_imsEmergencySupportForSNPN_r17 `optional`
}

func (ie *SNPN_AccessInfo_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.extCH_Supported_r17 != nil, ie.extCH_WithoutConfigAllowed_r17 != nil, ie.onboardingEnabled_r17 != nil, ie.imsEmergencySupportForSNPN_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.extCH_Supported_r17 != nil {
		if err = ie.extCH_Supported_r17.Encode(w); err != nil {
			return utils.WrapError("Encode extCH_Supported_r17", err)
		}
	}
	if ie.extCH_WithoutConfigAllowed_r17 != nil {
		if err = ie.extCH_WithoutConfigAllowed_r17.Encode(w); err != nil {
			return utils.WrapError("Encode extCH_WithoutConfigAllowed_r17", err)
		}
	}
	if ie.onboardingEnabled_r17 != nil {
		if err = ie.onboardingEnabled_r17.Encode(w); err != nil {
			return utils.WrapError("Encode onboardingEnabled_r17", err)
		}
	}
	if ie.imsEmergencySupportForSNPN_r17 != nil {
		if err = ie.imsEmergencySupportForSNPN_r17.Encode(w); err != nil {
			return utils.WrapError("Encode imsEmergencySupportForSNPN_r17", err)
		}
	}
	return nil
}

func (ie *SNPN_AccessInfo_r17) Decode(r *uper.UperReader) error {
	var err error
	var extCH_Supported_r17Present bool
	if extCH_Supported_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var extCH_WithoutConfigAllowed_r17Present bool
	if extCH_WithoutConfigAllowed_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var onboardingEnabled_r17Present bool
	if onboardingEnabled_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var imsEmergencySupportForSNPN_r17Present bool
	if imsEmergencySupportForSNPN_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if extCH_Supported_r17Present {
		ie.extCH_Supported_r17 = new(SNPN_AccessInfo_r17_extCH_Supported_r17)
		if err = ie.extCH_Supported_r17.Decode(r); err != nil {
			return utils.WrapError("Decode extCH_Supported_r17", err)
		}
	}
	if extCH_WithoutConfigAllowed_r17Present {
		ie.extCH_WithoutConfigAllowed_r17 = new(SNPN_AccessInfo_r17_extCH_WithoutConfigAllowed_r17)
		if err = ie.extCH_WithoutConfigAllowed_r17.Decode(r); err != nil {
			return utils.WrapError("Decode extCH_WithoutConfigAllowed_r17", err)
		}
	}
	if onboardingEnabled_r17Present {
		ie.onboardingEnabled_r17 = new(SNPN_AccessInfo_r17_onboardingEnabled_r17)
		if err = ie.onboardingEnabled_r17.Decode(r); err != nil {
			return utils.WrapError("Decode onboardingEnabled_r17", err)
		}
	}
	if imsEmergencySupportForSNPN_r17Present {
		ie.imsEmergencySupportForSNPN_r17 = new(SNPN_AccessInfo_r17_imsEmergencySupportForSNPN_r17)
		if err = ie.imsEmergencySupportForSNPN_r17.Decode(r); err != nil {
			return utils.WrapError("Decode imsEmergencySupportForSNPN_r17", err)
		}
	}
	return nil
}
