package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	IAB_IP_Usage_r16_Enum_f1_C   uper.Enumerated = 0
	IAB_IP_Usage_r16_Enum_f1_U   uper.Enumerated = 1
	IAB_IP_Usage_r16_Enum_non_F1 uper.Enumerated = 2
	IAB_IP_Usage_r16_Enum_spare  uper.Enumerated = 3
)

type IAB_IP_Usage_r16 struct {
	Value uper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *IAB_IP_Usage_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode IAB_IP_Usage_r16", err)
	}
	return nil
}

func (ie *IAB_IP_Usage_r16) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode IAB_IP_Usage_r16", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
