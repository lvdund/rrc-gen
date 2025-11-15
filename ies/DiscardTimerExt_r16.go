package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	DiscardTimerExt_r16_Enum_ms0dot5 uper.Enumerated = 0
	DiscardTimerExt_r16_Enum_ms1     uper.Enumerated = 1
	DiscardTimerExt_r16_Enum_ms2     uper.Enumerated = 2
	DiscardTimerExt_r16_Enum_ms4     uper.Enumerated = 3
	DiscardTimerExt_r16_Enum_ms6     uper.Enumerated = 4
	DiscardTimerExt_r16_Enum_ms8     uper.Enumerated = 5
	DiscardTimerExt_r16_Enum_spare2  uper.Enumerated = 6
	DiscardTimerExt_r16_Enum_spare1  uper.Enumerated = 7
)

type DiscardTimerExt_r16 struct {
	Value uper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *DiscardTimerExt_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode DiscardTimerExt_r16", err)
	}
	return nil
}

func (ie *DiscardTimerExt_r16) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode DiscardTimerExt_r16", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
