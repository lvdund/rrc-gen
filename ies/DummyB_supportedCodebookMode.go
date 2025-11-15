package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	DummyB_supportedCodebookMode_Enum_mode1         uper.Enumerated = 0
	DummyB_supportedCodebookMode_Enum_mode1AndMode2 uper.Enumerated = 1
)

type DummyB_supportedCodebookMode struct {
	Value uper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *DummyB_supportedCodebookMode) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode DummyB_supportedCodebookMode", err)
	}
	return nil
}

func (ie *DummyB_supportedCodebookMode) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode DummyB_supportedCodebookMode", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
