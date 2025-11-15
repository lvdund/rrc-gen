package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UECapabilityEnquiry_IEs struct {
	ue_CapabilityRAT_RequestList UE_CapabilityRAT_RequestList `madatory`
	lateNonCriticalExtension     *[]byte                      `optional`
	ue_CapabilityEnquiryExt      *[]byte                      `optional`
}

func (ie *UECapabilityEnquiry_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.lateNonCriticalExtension != nil, ie.ue_CapabilityEnquiryExt != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.ue_CapabilityRAT_RequestList.Encode(w); err != nil {
		return utils.WrapError("Encode ue_CapabilityRAT_RequestList", err)
	}
	if ie.lateNonCriticalExtension != nil {
		if err = w.WriteOctetString(*ie.lateNonCriticalExtension, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode lateNonCriticalExtension", err)
		}
	}
	if ie.ue_CapabilityEnquiryExt != nil {
		if err = w.WriteOctetString(*ie.ue_CapabilityEnquiryExt, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode ue_CapabilityEnquiryExt", err)
		}
	}
	return nil
}

func (ie *UECapabilityEnquiry_IEs) Decode(r *uper.UperReader) error {
	var err error
	var lateNonCriticalExtensionPresent bool
	if lateNonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var ue_CapabilityEnquiryExtPresent bool
	if ue_CapabilityEnquiryExtPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.ue_CapabilityRAT_RequestList.Decode(r); err != nil {
		return utils.WrapError("Decode ue_CapabilityRAT_RequestList", err)
	}
	if lateNonCriticalExtensionPresent {
		var tmp_os_lateNonCriticalExtension []byte
		if tmp_os_lateNonCriticalExtension, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode lateNonCriticalExtension", err)
		}
		ie.lateNonCriticalExtension = &tmp_os_lateNonCriticalExtension
	}
	if ue_CapabilityEnquiryExtPresent {
		var tmp_os_ue_CapabilityEnquiryExt []byte
		if tmp_os_ue_CapabilityEnquiryExt, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode ue_CapabilityEnquiryExt", err)
		}
		ie.ue_CapabilityEnquiryExt = &tmp_os_ue_CapabilityEnquiryExt
	}
	return nil
}
