package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UE_NR_Capability_v1540 struct {
	sdap_Parameters                  *SDAP_Parameters                       `optional`
	overheatingInd                   *UE_NR_Capability_v1540_overheatingInd `optional`
	ims_Parameters                   *IMS_Parameters                        `optional`
	fr1_Add_UE_NR_Capabilities_v1540 *UE_NR_CapabilityAddFRX_Mode_v1540     `optional`
	fr2_Add_UE_NR_Capabilities_v1540 *UE_NR_CapabilityAddFRX_Mode_v1540     `optional`
	fr1_fr2_Add_UE_NR_Capabilities   *UE_NR_CapabilityAddFRX_Mode           `optional`
	nonCriticalExtension             *UE_NR_Capability_v1550                `optional`
}

func (ie *UE_NR_Capability_v1540) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.sdap_Parameters != nil, ie.overheatingInd != nil, ie.ims_Parameters != nil, ie.fr1_Add_UE_NR_Capabilities_v1540 != nil, ie.fr2_Add_UE_NR_Capabilities_v1540 != nil, ie.fr1_fr2_Add_UE_NR_Capabilities != nil, ie.nonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.sdap_Parameters != nil {
		if err = ie.sdap_Parameters.Encode(w); err != nil {
			return utils.WrapError("Encode sdap_Parameters", err)
		}
	}
	if ie.overheatingInd != nil {
		if err = ie.overheatingInd.Encode(w); err != nil {
			return utils.WrapError("Encode overheatingInd", err)
		}
	}
	if ie.ims_Parameters != nil {
		if err = ie.ims_Parameters.Encode(w); err != nil {
			return utils.WrapError("Encode ims_Parameters", err)
		}
	}
	if ie.fr1_Add_UE_NR_Capabilities_v1540 != nil {
		if err = ie.fr1_Add_UE_NR_Capabilities_v1540.Encode(w); err != nil {
			return utils.WrapError("Encode fr1_Add_UE_NR_Capabilities_v1540", err)
		}
	}
	if ie.fr2_Add_UE_NR_Capabilities_v1540 != nil {
		if err = ie.fr2_Add_UE_NR_Capabilities_v1540.Encode(w); err != nil {
			return utils.WrapError("Encode fr2_Add_UE_NR_Capabilities_v1540", err)
		}
	}
	if ie.fr1_fr2_Add_UE_NR_Capabilities != nil {
		if err = ie.fr1_fr2_Add_UE_NR_Capabilities.Encode(w); err != nil {
			return utils.WrapError("Encode fr1_fr2_Add_UE_NR_Capabilities", err)
		}
	}
	if ie.nonCriticalExtension != nil {
		if err = ie.nonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode nonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *UE_NR_Capability_v1540) Decode(r *uper.UperReader) error {
	var err error
	var sdap_ParametersPresent bool
	if sdap_ParametersPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var overheatingIndPresent bool
	if overheatingIndPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var ims_ParametersPresent bool
	if ims_ParametersPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var fr1_Add_UE_NR_Capabilities_v1540Present bool
	if fr1_Add_UE_NR_Capabilities_v1540Present, err = r.ReadBool(); err != nil {
		return err
	}
	var fr2_Add_UE_NR_Capabilities_v1540Present bool
	if fr2_Add_UE_NR_Capabilities_v1540Present, err = r.ReadBool(); err != nil {
		return err
	}
	var fr1_fr2_Add_UE_NR_CapabilitiesPresent bool
	if fr1_fr2_Add_UE_NR_CapabilitiesPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var nonCriticalExtensionPresent bool
	if nonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if sdap_ParametersPresent {
		ie.sdap_Parameters = new(SDAP_Parameters)
		if err = ie.sdap_Parameters.Decode(r); err != nil {
			return utils.WrapError("Decode sdap_Parameters", err)
		}
	}
	if overheatingIndPresent {
		ie.overheatingInd = new(UE_NR_Capability_v1540_overheatingInd)
		if err = ie.overheatingInd.Decode(r); err != nil {
			return utils.WrapError("Decode overheatingInd", err)
		}
	}
	if ims_ParametersPresent {
		ie.ims_Parameters = new(IMS_Parameters)
		if err = ie.ims_Parameters.Decode(r); err != nil {
			return utils.WrapError("Decode ims_Parameters", err)
		}
	}
	if fr1_Add_UE_NR_Capabilities_v1540Present {
		ie.fr1_Add_UE_NR_Capabilities_v1540 = new(UE_NR_CapabilityAddFRX_Mode_v1540)
		if err = ie.fr1_Add_UE_NR_Capabilities_v1540.Decode(r); err != nil {
			return utils.WrapError("Decode fr1_Add_UE_NR_Capabilities_v1540", err)
		}
	}
	if fr2_Add_UE_NR_Capabilities_v1540Present {
		ie.fr2_Add_UE_NR_Capabilities_v1540 = new(UE_NR_CapabilityAddFRX_Mode_v1540)
		if err = ie.fr2_Add_UE_NR_Capabilities_v1540.Decode(r); err != nil {
			return utils.WrapError("Decode fr2_Add_UE_NR_Capabilities_v1540", err)
		}
	}
	if fr1_fr2_Add_UE_NR_CapabilitiesPresent {
		ie.fr1_fr2_Add_UE_NR_Capabilities = new(UE_NR_CapabilityAddFRX_Mode)
		if err = ie.fr1_fr2_Add_UE_NR_Capabilities.Decode(r); err != nil {
			return utils.WrapError("Decode fr1_fr2_Add_UE_NR_Capabilities", err)
		}
	}
	if nonCriticalExtensionPresent {
		ie.nonCriticalExtension = new(UE_NR_Capability_v1550)
		if err = ie.nonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode nonCriticalExtension", err)
		}
	}
	return nil
}
