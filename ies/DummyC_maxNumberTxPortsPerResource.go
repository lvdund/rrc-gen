package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	DummyC_maxNumberTxPortsPerResource_Enum_p8  uper.Enumerated = 0
	DummyC_maxNumberTxPortsPerResource_Enum_p16 uper.Enumerated = 1
	DummyC_maxNumberTxPortsPerResource_Enum_p32 uper.Enumerated = 2
)

type DummyC_maxNumberTxPortsPerResource struct {
	Value uper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *DummyC_maxNumberTxPortsPerResource) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode DummyC_maxNumberTxPortsPerResource", err)
	}
	return nil
}

func (ie *DummyC_maxNumberTxPortsPerResource) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode DummyC_maxNumberTxPortsPerResource", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
