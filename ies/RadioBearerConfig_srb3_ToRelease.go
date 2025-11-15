package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	RadioBearerConfig_srb3_ToRelease_Enum_true uper.Enumerated = 0
)

type RadioBearerConfig_srb3_ToRelease struct {
	Value uper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *RadioBearerConfig_srb3_ToRelease) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode RadioBearerConfig_srb3_ToRelease", err)
	}
	return nil
}

func (ie *RadioBearerConfig_srb3_ToRelease) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode RadioBearerConfig_srb3_ToRelease", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
