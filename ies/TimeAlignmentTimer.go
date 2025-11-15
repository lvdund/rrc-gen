package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	TimeAlignmentTimer_Enum_ms500    uper.Enumerated = 0
	TimeAlignmentTimer_Enum_ms750    uper.Enumerated = 1
	TimeAlignmentTimer_Enum_ms1280   uper.Enumerated = 2
	TimeAlignmentTimer_Enum_ms1920   uper.Enumerated = 3
	TimeAlignmentTimer_Enum_ms2560   uper.Enumerated = 4
	TimeAlignmentTimer_Enum_ms5120   uper.Enumerated = 5
	TimeAlignmentTimer_Enum_ms10240  uper.Enumerated = 6
	TimeAlignmentTimer_Enum_infinity uper.Enumerated = 7
)

type TimeAlignmentTimer struct {
	Value uper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *TimeAlignmentTimer) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode TimeAlignmentTimer", err)
	}
	return nil
}

func (ie *TimeAlignmentTimer) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode TimeAlignmentTimer", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
