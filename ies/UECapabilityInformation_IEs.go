package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UECapabilityInformation_IEs struct {
	ue_CapabilityRAT_ContainerList *UE_CapabilityRAT_ContainerList `optional`
	lateNonCriticalExtension       *[]byte                         `optional`
	nonCriticalExtension           interface{}                     `optional`
}

func (ie *UECapabilityInformation_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.ue_CapabilityRAT_ContainerList != nil, ie.lateNonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.ue_CapabilityRAT_ContainerList != nil {
		if err = ie.ue_CapabilityRAT_ContainerList.Encode(w); err != nil {
			return utils.WrapError("Encode ue_CapabilityRAT_ContainerList", err)
		}
	}
	if ie.lateNonCriticalExtension != nil {
		if err = w.WriteOctetString(*ie.lateNonCriticalExtension, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode lateNonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *UECapabilityInformation_IEs) Decode(r *uper.UperReader) error {
	var err error
	var ue_CapabilityRAT_ContainerListPresent bool
	if ue_CapabilityRAT_ContainerListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var lateNonCriticalExtensionPresent bool
	if lateNonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if ue_CapabilityRAT_ContainerListPresent {
		ie.ue_CapabilityRAT_ContainerList = new(UE_CapabilityRAT_ContainerList)
		if err = ie.ue_CapabilityRAT_ContainerList.Decode(r); err != nil {
			return utils.WrapError("Decode ue_CapabilityRAT_ContainerList", err)
		}
	}
	if lateNonCriticalExtensionPresent {
		var tmp_os_lateNonCriticalExtension []byte
		if tmp_os_lateNonCriticalExtension, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode lateNonCriticalExtension", err)
		}
		ie.lateNonCriticalExtension = &tmp_os_lateNonCriticalExtension
	}
	return nil
}
