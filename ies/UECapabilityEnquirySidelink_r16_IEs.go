package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UECapabilityEnquirySidelink_r16_IEs struct {
	frequencyBandListFilterSidelink_r16  *FreqBandList `optional`
	ue_CapabilityInformationSidelink_r16 *[]byte       `optional`
	lateNonCriticalExtension             *[]byte       `optional`
	nonCriticalExtension                 interface{}   `optional`
}

func (ie *UECapabilityEnquirySidelink_r16_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.frequencyBandListFilterSidelink_r16 != nil, ie.ue_CapabilityInformationSidelink_r16 != nil, ie.lateNonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.frequencyBandListFilterSidelink_r16 != nil {
		if err = ie.frequencyBandListFilterSidelink_r16.Encode(w); err != nil {
			return utils.WrapError("Encode frequencyBandListFilterSidelink_r16", err)
		}
	}
	if ie.ue_CapabilityInformationSidelink_r16 != nil {
		if err = w.WriteOctetString(*ie.ue_CapabilityInformationSidelink_r16, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode ue_CapabilityInformationSidelink_r16", err)
		}
	}
	if ie.lateNonCriticalExtension != nil {
		if err = w.WriteOctetString(*ie.lateNonCriticalExtension, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode lateNonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *UECapabilityEnquirySidelink_r16_IEs) Decode(r *uper.UperReader) error {
	var err error
	var frequencyBandListFilterSidelink_r16Present bool
	if frequencyBandListFilterSidelink_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ue_CapabilityInformationSidelink_r16Present bool
	if ue_CapabilityInformationSidelink_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var lateNonCriticalExtensionPresent bool
	if lateNonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if frequencyBandListFilterSidelink_r16Present {
		ie.frequencyBandListFilterSidelink_r16 = new(FreqBandList)
		if err = ie.frequencyBandListFilterSidelink_r16.Decode(r); err != nil {
			return utils.WrapError("Decode frequencyBandListFilterSidelink_r16", err)
		}
	}
	if ue_CapabilityInformationSidelink_r16Present {
		var tmp_os_ue_CapabilityInformationSidelink_r16 []byte
		if tmp_os_ue_CapabilityInformationSidelink_r16, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode ue_CapabilityInformationSidelink_r16", err)
		}
		ie.ue_CapabilityInformationSidelink_r16 = &tmp_os_ue_CapabilityInformationSidelink_r16
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
