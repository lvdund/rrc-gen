package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SIB7_warningMessageSegmentType_Enum_notLastSegment uper.Enumerated = 0
	SIB7_warningMessageSegmentType_Enum_lastSegment    uper.Enumerated = 1
)

type SIB7_warningMessageSegmentType struct {
	Value uper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *SIB7_warningMessageSegmentType) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode SIB7_warningMessageSegmentType", err)
	}
	return nil
}

func (ie *SIB7_warningMessageSegmentType) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode SIB7_warningMessageSegmentType", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
