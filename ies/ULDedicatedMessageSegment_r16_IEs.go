package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ULDedicatedMessageSegment_r16_IEs struct {
	segmentNumber_r16               int64                                                        `lb:0,ub:15,madatory`
	rrc_MessageSegmentContainer_r16 []byte                                                       `madatory`
	rrc_MessageSegmentType_r16      ULDedicatedMessageSegment_r16_IEs_rrc_MessageSegmentType_r16 `madatory`
	lateNonCriticalExtension        *[]byte                                                      `optional`
	nonCriticalExtension            interface{}                                                  `optional`
}

func (ie *ULDedicatedMessageSegment_r16_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.lateNonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = w.WriteInteger(ie.segmentNumber_r16, &uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("WriteInteger segmentNumber_r16", err)
	}
	if err = w.WriteOctetString(ie.rrc_MessageSegmentContainer_r16, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("WriteOctetString rrc_MessageSegmentContainer_r16", err)
	}
	if ie.lateNonCriticalExtension != nil {
		if err = w.WriteOctetString(*ie.lateNonCriticalExtension, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode lateNonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *ULDedicatedMessageSegment_r16_IEs) Decode(r *uper.UperReader) error {
	var err error
	var lateNonCriticalExtensionPresent bool
	if lateNonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var tmp_int_segmentNumber_r16 int64
	if tmp_int_segmentNumber_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("ReadInteger segmentNumber_r16", err)
	}
	ie.segmentNumber_r16 = tmp_int_segmentNumber_r16
	var tmp_os_rrc_MessageSegmentContainer_r16 []byte
	if tmp_os_rrc_MessageSegmentContainer_r16, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("ReadOctetString rrc_MessageSegmentContainer_r16", err)
	}
	ie.rrc_MessageSegmentContainer_r16 = tmp_os_rrc_MessageSegmentContainer_r16
	if lateNonCriticalExtensionPresent {
		var tmp_os_lateNonCriticalExtension []byte
		if tmp_os_lateNonCriticalExtension, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode lateNonCriticalExtension", err)
		}
		ie.lateNonCriticalExtension = &tmp_os_lateNonCriticalExtension
	}
	return nil
}
