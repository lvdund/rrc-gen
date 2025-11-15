package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	AccessStratumRelease_Enum_rel15  uper.Enumerated = 0
	AccessStratumRelease_Enum_rel16  uper.Enumerated = 1
	AccessStratumRelease_Enum_rel17  uper.Enumerated = 2
	AccessStratumRelease_Enum_spare5 uper.Enumerated = 3
	AccessStratumRelease_Enum_spare4 uper.Enumerated = 4
	AccessStratumRelease_Enum_spare3 uper.Enumerated = 5
	AccessStratumRelease_Enum_spare2 uper.Enumerated = 6
	AccessStratumRelease_Enum_spare1 uper.Enumerated = 7
)

type AccessStratumRelease struct {
	Value uper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *AccessStratumRelease) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode AccessStratumRelease", err)
	}
	return nil
}

func (ie *AccessStratumRelease) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode AccessStratumRelease", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
