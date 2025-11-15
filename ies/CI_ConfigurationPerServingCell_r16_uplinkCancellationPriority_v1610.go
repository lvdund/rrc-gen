package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	CI_ConfigurationPerServingCell_r16_uplinkCancellationPriority_v1610_Enum_enabled uper.Enumerated = 0
)

type CI_ConfigurationPerServingCell_r16_uplinkCancellationPriority_v1610 struct {
	Value uper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *CI_ConfigurationPerServingCell_r16_uplinkCancellationPriority_v1610) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode CI_ConfigurationPerServingCell_r16_uplinkCancellationPriority_v1610", err)
	}
	return nil
}

func (ie *CI_ConfigurationPerServingCell_r16_uplinkCancellationPriority_v1610) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode CI_ConfigurationPerServingCell_r16_uplinkCancellationPriority_v1610", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
