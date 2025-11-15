package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	DummyB_maxNumberTxPortsPerResource_Enum_p2  uper.Enumerated = 0
	DummyB_maxNumberTxPortsPerResource_Enum_p4  uper.Enumerated = 1
	DummyB_maxNumberTxPortsPerResource_Enum_p8  uper.Enumerated = 2
	DummyB_maxNumberTxPortsPerResource_Enum_p12 uper.Enumerated = 3
	DummyB_maxNumberTxPortsPerResource_Enum_p16 uper.Enumerated = 4
	DummyB_maxNumberTxPortsPerResource_Enum_p24 uper.Enumerated = 5
	DummyB_maxNumberTxPortsPerResource_Enum_p32 uper.Enumerated = 6
)

type DummyB_maxNumberTxPortsPerResource struct {
	Value uper.Enumerated `lb:0,ub:6,madatory`
}

func (ie *DummyB_maxNumberTxPortsPerResource) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 6}, false); err != nil {
		return utils.WrapError("Encode DummyB_maxNumberTxPortsPerResource", err)
	}
	return nil
}

func (ie *DummyB_maxNumberTxPortsPerResource) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 6}, false); err != nil {
		return utils.WrapError("Decode DummyB_maxNumberTxPortsPerResource", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
