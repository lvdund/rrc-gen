package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type DLInformationTransfer_IEs struct {
	dedicatedNAS_Message     *DedicatedNAS_Message            `optional`
	lateNonCriticalExtension *[]byte                          `optional`
	nonCriticalExtension     *DLInformationTransfer_v1610_IEs `optional`
}

func (ie *DLInformationTransfer_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.dedicatedNAS_Message != nil, ie.lateNonCriticalExtension != nil, ie.nonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.dedicatedNAS_Message != nil {
		if err = ie.dedicatedNAS_Message.Encode(w); err != nil {
			return utils.WrapError("Encode dedicatedNAS_Message", err)
		}
	}
	if ie.lateNonCriticalExtension != nil {
		if err = w.WriteOctetString(*ie.lateNonCriticalExtension, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode lateNonCriticalExtension", err)
		}
	}
	if ie.nonCriticalExtension != nil {
		if err = ie.nonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode nonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *DLInformationTransfer_IEs) Decode(r *uper.UperReader) error {
	var err error
	var dedicatedNAS_MessagePresent bool
	if dedicatedNAS_MessagePresent, err = r.ReadBool(); err != nil {
		return err
	}
	var lateNonCriticalExtensionPresent bool
	if lateNonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var nonCriticalExtensionPresent bool
	if nonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if dedicatedNAS_MessagePresent {
		ie.dedicatedNAS_Message = new(DedicatedNAS_Message)
		if err = ie.dedicatedNAS_Message.Decode(r); err != nil {
			return utils.WrapError("Decode dedicatedNAS_Message", err)
		}
	}
	if lateNonCriticalExtensionPresent {
		var tmp_os_lateNonCriticalExtension []byte
		if tmp_os_lateNonCriticalExtension, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode lateNonCriticalExtension", err)
		}
		ie.lateNonCriticalExtension = &tmp_os_lateNonCriticalExtension
	}
	if nonCriticalExtensionPresent {
		ie.nonCriticalExtension = new(DLInformationTransfer_v1610_IEs)
		if err = ie.nonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode nonCriticalExtension", err)
		}
	}
	return nil
}
