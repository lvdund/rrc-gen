package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	FreqSeparationClassDL_Only_r16_Enum_mhz200  uper.Enumerated = 0
	FreqSeparationClassDL_Only_r16_Enum_mhz400  uper.Enumerated = 1
	FreqSeparationClassDL_Only_r16_Enum_mhz600  uper.Enumerated = 2
	FreqSeparationClassDL_Only_r16_Enum_mhz800  uper.Enumerated = 3
	FreqSeparationClassDL_Only_r16_Enum_mhz1000 uper.Enumerated = 4
	FreqSeparationClassDL_Only_r16_Enum_mhz1200 uper.Enumerated = 5
)

type FreqSeparationClassDL_Only_r16 struct {
	Value uper.Enumerated `lb:0,ub:5,madatory`
}

func (ie *FreqSeparationClassDL_Only_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("Encode FreqSeparationClassDL_Only_r16", err)
	}
	return nil
}

func (ie *FreqSeparationClassDL_Only_r16) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("Decode FreqSeparationClassDL_Only_r16", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
