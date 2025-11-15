package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	ModulationOrder_Enum_bpsk_halfpi uper.Enumerated = 0
	ModulationOrder_Enum_bpsk        uper.Enumerated = 1
	ModulationOrder_Enum_qpsk        uper.Enumerated = 2
	ModulationOrder_Enum_qam16       uper.Enumerated = 3
	ModulationOrder_Enum_qam64       uper.Enumerated = 4
	ModulationOrder_Enum_qam256      uper.Enumerated = 5
)

type ModulationOrder struct {
	Value uper.Enumerated `lb:0,ub:5,madatory`
}

func (ie *ModulationOrder) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("Encode ModulationOrder", err)
	}
	return nil
}

func (ie *ModulationOrder) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("Decode ModulationOrder", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
