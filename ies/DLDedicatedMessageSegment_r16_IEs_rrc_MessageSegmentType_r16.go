package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	DLDedicatedMessageSegment_r16_IEs_rrc_MessageSegmentType_r16_Enum_notLastSegment uper.Enumerated = 0
	DLDedicatedMessageSegment_r16_IEs_rrc_MessageSegmentType_r16_Enum_lastSegment    uper.Enumerated = 1
)

type DLDedicatedMessageSegment_r16_IEs_rrc_MessageSegmentType_r16 struct {
	Value uper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *DLDedicatedMessageSegment_r16_IEs_rrc_MessageSegmentType_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode DLDedicatedMessageSegment_r16_IEs_rrc_MessageSegmentType_r16", err)
	}
	return nil
}

func (ie *DLDedicatedMessageSegment_r16_IEs_rrc_MessageSegmentType_r16) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode DLDedicatedMessageSegment_r16_IEs_rrc_MessageSegmentType_r16", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
