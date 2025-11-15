package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	DummyC_supportedCodebookMode_Enum_mode1 uper.Enumerated = 0
	DummyC_supportedCodebookMode_Enum_mode2 uper.Enumerated = 1
	DummyC_supportedCodebookMode_Enum_both  uper.Enumerated = 2
)

type DummyC_supportedCodebookMode struct {
	Value uper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *DummyC_supportedCodebookMode) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode DummyC_supportedCodebookMode", err)
	}
	return nil
}

func (ie *DummyC_supportedCodebookMode) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode DummyC_supportedCodebookMode", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
