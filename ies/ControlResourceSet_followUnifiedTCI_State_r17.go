package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	ControlResourceSet_followUnifiedTCI_State_r17_Enum_enabled uper.Enumerated = 0
)

type ControlResourceSet_followUnifiedTCI_State_r17 struct {
	Value uper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *ControlResourceSet_followUnifiedTCI_State_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode ControlResourceSet_followUnifiedTCI_State_r17", err)
	}
	return nil
}

func (ie *ControlResourceSet_followUnifiedTCI_State_r17) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode ControlResourceSet_followUnifiedTCI_State_r17", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
