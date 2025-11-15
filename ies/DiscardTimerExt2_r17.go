package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	DiscardTimerExt2_r17_Enum_ms2000 uper.Enumerated = 0
	DiscardTimerExt2_r17_Enum_spare3 uper.Enumerated = 1
	DiscardTimerExt2_r17_Enum_spare2 uper.Enumerated = 2
	DiscardTimerExt2_r17_Enum_spare1 uper.Enumerated = 3
)

type DiscardTimerExt2_r17 struct {
	Value uper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *DiscardTimerExt2_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode DiscardTimerExt2_r17", err)
	}
	return nil
}

func (ie *DiscardTimerExt2_r17) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode DiscardTimerExt2_r17", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
