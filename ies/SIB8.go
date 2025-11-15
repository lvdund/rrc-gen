package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SIB8 struct {
	messageIdentifier             uper.BitString                 `lb:16,ub:16,madatory`
	serialNumber                  uper.BitString                 `lb:16,ub:16,madatory`
	warningMessageSegmentType     SIB8_warningMessageSegmentType `madatory`
	warningMessageSegmentNumber   int64                          `lb:0,ub:63,madatory`
	warningMessageSegment         []byte                         `madatory`
	dataCodingScheme              *[]byte                        `lb:1,ub:1,optional`
	warningAreaCoordinatesSegment *[]byte                        `optional`
	lateNonCriticalExtension      *[]byte                        `optional`
}

func (ie *SIB8) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.dataCodingScheme != nil, ie.warningAreaCoordinatesSegment != nil, ie.lateNonCriticalExtension != nil}
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
	if err = ie.warningMessageSegmentType.Encode(w); err != nil {
		return utils.WrapError("Encode warningMessageSegmentType", err)
	}
	if err = w.WriteInteger(ie.warningMessageSegmentNumber, &uper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
		return utils.WrapError("WriteInteger warningMessageSegmentNumber", err)
	}
	if err = w.WriteOctetString(ie.warningMessageSegment, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("WriteOctetString warningMessageSegment", err)
	}
	if ie.dataCodingScheme != nil {
		if err = w.WriteOctetString(*ie.dataCodingScheme, &uper.Constraint{Lb: 1, Ub: 1}, false); err != nil {
			return utils.WrapError("Encode dataCodingScheme", err)
		}
	}
	if ie.warningAreaCoordinatesSegment != nil {
		if err = w.WriteOctetString(*ie.warningAreaCoordinatesSegment, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode warningAreaCoordinatesSegment", err)
		}
	}
	if ie.lateNonCriticalExtension != nil {
		if err = w.WriteOctetString(*ie.lateNonCriticalExtension, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode lateNonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *SIB8) Decode(r *uper.UperReader) error {
	var err error
	var dataCodingSchemePresent bool
	if dataCodingSchemePresent, err = r.ReadBool(); err != nil {
		return err
	}
	var warningAreaCoordinatesSegmentPresent bool
	if warningAreaCoordinatesSegmentPresent, err = r.ReadBool(); err != nil {
		return err
	}
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
	if err = ie.warningMessageSegmentType.Decode(r); err != nil {
		return utils.WrapError("Decode warningMessageSegmentType", err)
	}
	var tmp_int_warningMessageSegmentNumber int64
	if tmp_int_warningMessageSegmentNumber, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
		return utils.WrapError("ReadInteger warningMessageSegmentNumber", err)
	}
	ie.warningMessageSegmentNumber = tmp_int_warningMessageSegmentNumber
	var tmp_os_warningMessageSegment []byte
	if tmp_os_warningMessageSegment, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("ReadOctetString warningMessageSegment", err)
	}
	ie.warningMessageSegment = tmp_os_warningMessageSegment
	if dataCodingSchemePresent {
		var tmp_os_dataCodingScheme []byte
		if tmp_os_dataCodingScheme, err = r.ReadOctetString(&uper.Constraint{Lb: 1, Ub: 1}, false); err != nil {
			return utils.WrapError("Decode dataCodingScheme", err)
		}
		ie.dataCodingScheme = &tmp_os_dataCodingScheme
	}
	if warningAreaCoordinatesSegmentPresent {
		var tmp_os_warningAreaCoordinatesSegment []byte
		if tmp_os_warningAreaCoordinatesSegment, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode warningAreaCoordinatesSegment", err)
		}
		ie.warningAreaCoordinatesSegment = &tmp_os_warningAreaCoordinatesSegment
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
