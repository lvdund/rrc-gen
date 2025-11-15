package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	DummyE_maxNumberTxPortsPerResource_Enum_p4  uper.Enumerated = 0
	DummyE_maxNumberTxPortsPerResource_Enum_p8  uper.Enumerated = 1
	DummyE_maxNumberTxPortsPerResource_Enum_p12 uper.Enumerated = 2
	DummyE_maxNumberTxPortsPerResource_Enum_p16 uper.Enumerated = 3
	DummyE_maxNumberTxPortsPerResource_Enum_p24 uper.Enumerated = 4
	DummyE_maxNumberTxPortsPerResource_Enum_p32 uper.Enumerated = 5
)

type DummyE_maxNumberTxPortsPerResource struct {
	Value uper.Enumerated `lb:0,ub:5,madatory`
}

func (ie *DummyE_maxNumberTxPortsPerResource) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("Encode DummyE_maxNumberTxPortsPerResource", err)
	}
	return nil
}

func (ie *DummyE_maxNumberTxPortsPerResource) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("Decode DummyE_maxNumberTxPortsPerResource", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
