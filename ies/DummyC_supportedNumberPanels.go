package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	DummyC_supportedNumberPanels_Enum_n2 uper.Enumerated = 0
	DummyC_supportedNumberPanels_Enum_n4 uper.Enumerated = 1
)

type DummyC_supportedNumberPanels struct {
	Value uper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *DummyC_supportedNumberPanels) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode DummyC_supportedNumberPanels", err)
	}
	return nil
}

func (ie *DummyC_supportedNumberPanels) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode DummyC_supportedNumberPanels", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
