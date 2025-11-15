package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RRCSetup_IEs struct {
	radioBearerConfig        RadioBearerConfig   `madatory`
	masterCellGroup          []byte              `madatory`
	lateNonCriticalExtension *[]byte             `optional`
	nonCriticalExtension     *RRCSetup_v1700_IEs `optional`
}

func (ie *RRCSetup_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.lateNonCriticalExtension != nil, ie.nonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.radioBearerConfig.Encode(w); err != nil {
		return utils.WrapError("Encode radioBearerConfig", err)
	}
	if err = w.WriteOctetString(ie.masterCellGroup, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("WriteOctetString masterCellGroup", err)
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

func (ie *RRCSetup_IEs) Decode(r *uper.UperReader) error {
	var err error
	var lateNonCriticalExtensionPresent bool
	if lateNonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var nonCriticalExtensionPresent bool
	if nonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.radioBearerConfig.Decode(r); err != nil {
		return utils.WrapError("Decode radioBearerConfig", err)
	}
	var tmp_os_masterCellGroup []byte
	if tmp_os_masterCellGroup, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("ReadOctetString masterCellGroup", err)
	}
	ie.masterCellGroup = tmp_os_masterCellGroup
	if lateNonCriticalExtensionPresent {
		var tmp_os_lateNonCriticalExtension []byte
		if tmp_os_lateNonCriticalExtension, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode lateNonCriticalExtension", err)
		}
		ie.lateNonCriticalExtension = &tmp_os_lateNonCriticalExtension
	}
	if nonCriticalExtensionPresent {
		ie.nonCriticalExtension = new(RRCSetup_v1700_IEs)
		if err = ie.nonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode nonCriticalExtension", err)
		}
	}
	return nil
}
