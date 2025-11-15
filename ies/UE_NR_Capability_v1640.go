package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UE_NR_Capability_v1640 struct {
	redirectAtResumeByNAS_r16                *UE_NR_Capability_v1640_redirectAtResumeByNAS_r16 `optional`
	phy_ParametersSharedSpectrumChAccess_r16 *Phy_ParametersSharedSpectrumChAccess_r16         `optional`
	nonCriticalExtension                     *UE_NR_Capability_v1650                           `optional`
}

func (ie *UE_NR_Capability_v1640) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.redirectAtResumeByNAS_r16 != nil, ie.phy_ParametersSharedSpectrumChAccess_r16 != nil, ie.nonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.redirectAtResumeByNAS_r16 != nil {
		if err = ie.redirectAtResumeByNAS_r16.Encode(w); err != nil {
			return utils.WrapError("Encode redirectAtResumeByNAS_r16", err)
		}
	}
	if ie.phy_ParametersSharedSpectrumChAccess_r16 != nil {
		if err = ie.phy_ParametersSharedSpectrumChAccess_r16.Encode(w); err != nil {
			return utils.WrapError("Encode phy_ParametersSharedSpectrumChAccess_r16", err)
		}
	}
	if ie.nonCriticalExtension != nil {
		if err = ie.nonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode nonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *UE_NR_Capability_v1640) Decode(r *uper.UperReader) error {
	var err error
	var redirectAtResumeByNAS_r16Present bool
	if redirectAtResumeByNAS_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var phy_ParametersSharedSpectrumChAccess_r16Present bool
	if phy_ParametersSharedSpectrumChAccess_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var nonCriticalExtensionPresent bool
	if nonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if redirectAtResumeByNAS_r16Present {
		ie.redirectAtResumeByNAS_r16 = new(UE_NR_Capability_v1640_redirectAtResumeByNAS_r16)
		if err = ie.redirectAtResumeByNAS_r16.Decode(r); err != nil {
			return utils.WrapError("Decode redirectAtResumeByNAS_r16", err)
		}
	}
	if phy_ParametersSharedSpectrumChAccess_r16Present {
		ie.phy_ParametersSharedSpectrumChAccess_r16 = new(Phy_ParametersSharedSpectrumChAccess_r16)
		if err = ie.phy_ParametersSharedSpectrumChAccess_r16.Decode(r); err != nil {
			return utils.WrapError("Decode phy_ParametersSharedSpectrumChAccess_r16", err)
		}
	}
	if nonCriticalExtensionPresent {
		ie.nonCriticalExtension = new(UE_NR_Capability_v1650)
		if err = ie.nonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode nonCriticalExtension", err)
		}
	}
	return nil
}
