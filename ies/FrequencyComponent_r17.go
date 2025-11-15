package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	FrequencyComponent_r17_Enum_activeCarrier     uper.Enumerated = 0
	FrequencyComponent_r17_Enum_configuredCarrier uper.Enumerated = 1
	FrequencyComponent_r17_Enum_activeBWP         uper.Enumerated = 2
	FrequencyComponent_r17_Enum_configuredBWP     uper.Enumerated = 3
)

type FrequencyComponent_r17 struct {
	Value uper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *FrequencyComponent_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode FrequencyComponent_r17", err)
	}
	return nil
}

func (ie *FrequencyComponent_r17) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode FrequencyComponent_r17", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
