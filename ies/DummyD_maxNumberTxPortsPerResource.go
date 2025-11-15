package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	DummyD_maxNumberTxPortsPerResource_Enum_p4  uper.Enumerated = 0
	DummyD_maxNumberTxPortsPerResource_Enum_p8  uper.Enumerated = 1
	DummyD_maxNumberTxPortsPerResource_Enum_p12 uper.Enumerated = 2
	DummyD_maxNumberTxPortsPerResource_Enum_p16 uper.Enumerated = 3
	DummyD_maxNumberTxPortsPerResource_Enum_p24 uper.Enumerated = 4
	DummyD_maxNumberTxPortsPerResource_Enum_p32 uper.Enumerated = 5
)

type DummyD_maxNumberTxPortsPerResource struct {
	Value uper.Enumerated `lb:0,ub:5,madatory`
}

func (ie *DummyD_maxNumberTxPortsPerResource) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("Encode DummyD_maxNumberTxPortsPerResource", err)
	}
	return nil
}

func (ie *DummyD_maxNumberTxPortsPerResource) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("Decode DummyD_maxNumberTxPortsPerResource", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
