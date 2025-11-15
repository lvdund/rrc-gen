package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SIB6 struct {
	messageIdentifier        uper.BitString `lb:16,ub:16,madatory`
	serialNumber             uper.BitString `lb:16,ub:16,madatory`
	warningType              []byte         `lb:2,ub:2,madatory`
	lateNonCriticalExtension *[]byte        `optional`
}

func (ie *SIB6) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.lateNonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = w.WriteBitString(ie.messageIdentifier.Bytes, uint(ie.messageIdentifier.NumBits), &uper.Constraint{Lb: 16, Ub: 16}, false); err != nil {
		return utils.WrapError("WriteBitString messageIdentifier", err)
	}
	if err = w.WriteBitString(ie.serialNumber.Bytes, uint(ie.serialNumber.NumBits), &uper.Constraint{Lb: 16, Ub: 16}, false); err != nil {
		return utils.WrapError("WriteBitString serialNumber", err)
	}
	if err = w.WriteOctetString(ie.warningType, &uper.Constraint{Lb: 2, Ub: 2}, false); err != nil {
		return utils.WrapError("WriteOctetString warningType", err)
	}
	if ie.lateNonCriticalExtension != nil {
		if err = w.WriteOctetString(*ie.lateNonCriticalExtension, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode lateNonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *SIB6) Decode(r *uper.UperReader) error {
	var err error
	var lateNonCriticalExtensionPresent bool
	if lateNonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var tmp_bs_messageIdentifier []byte
	var n_messageIdentifier uint
	if tmp_bs_messageIdentifier, n_messageIdentifier, err = r.ReadBitString(&uper.Constraint{Lb: 16, Ub: 16}, false); err != nil {
		return utils.WrapError("ReadBitString messageIdentifier", err)
	}
	ie.messageIdentifier = uper.BitString{
		Bytes:   tmp_bs_messageIdentifier,
		NumBits: uint64(n_messageIdentifier),
	}
	var tmp_bs_serialNumber []byte
	var n_serialNumber uint
	if tmp_bs_serialNumber, n_serialNumber, err = r.ReadBitString(&uper.Constraint{Lb: 16, Ub: 16}, false); err != nil {
		return utils.WrapError("ReadBitString serialNumber", err)
	}
	ie.serialNumber = uper.BitString{
		Bytes:   tmp_bs_serialNumber,
		NumBits: uint64(n_serialNumber),
	}
	var tmp_os_warningType []byte
	if tmp_os_warningType, err = r.ReadOctetString(&uper.Constraint{Lb: 2, Ub: 2}, false); err != nil {
		return utils.WrapError("ReadOctetString warningType", err)
	}
	ie.warningType = tmp_os_warningType
	if lateNonCriticalExtensionPresent {
		var tmp_os_lateNonCriticalExtension []byte
		if tmp_os_lateNonCriticalExtension, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode lateNonCriticalExtension", err)
		}
		ie.lateNonCriticalExtension = &tmp_os_lateNonCriticalExtension
	}
	return nil
}
