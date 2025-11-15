package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	AccessStratumReleaseSidelink_r16_Enum_rel16  uper.Enumerated = 0
	AccessStratumReleaseSidelink_r16_Enum_rel17  uper.Enumerated = 1
	AccessStratumReleaseSidelink_r16_Enum_spare6 uper.Enumerated = 2
	AccessStratumReleaseSidelink_r16_Enum_spare5 uper.Enumerated = 3
	AccessStratumReleaseSidelink_r16_Enum_spare4 uper.Enumerated = 4
	AccessStratumReleaseSidelink_r16_Enum_spare3 uper.Enumerated = 5
	AccessStratumReleaseSidelink_r16_Enum_spare2 uper.Enumerated = 6
	AccessStratumReleaseSidelink_r16_Enum_spare1 uper.Enumerated = 7
)

type AccessStratumReleaseSidelink_r16 struct {
	Value uper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *AccessStratumReleaseSidelink_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode AccessStratumReleaseSidelink_r16", err)
	}
	return nil
}

func (ie *AccessStratumReleaseSidelink_r16) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode AccessStratumReleaseSidelink_r16", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
