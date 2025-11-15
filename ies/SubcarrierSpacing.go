package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SubcarrierSpacing_Enum_kHz15        uper.Enumerated = 0
	SubcarrierSpacing_Enum_kHz30        uper.Enumerated = 1
	SubcarrierSpacing_Enum_kHz60        uper.Enumerated = 2
	SubcarrierSpacing_Enum_kHz120       uper.Enumerated = 3
	SubcarrierSpacing_Enum_kHz240       uper.Enumerated = 4
	SubcarrierSpacing_Enum_kHz480_v1700 uper.Enumerated = 5
	SubcarrierSpacing_Enum_kHz960_v1700 uper.Enumerated = 6
	SubcarrierSpacing_Enum_spare1       uper.Enumerated = 7
)

type SubcarrierSpacing struct {
	Value uper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *SubcarrierSpacing) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode SubcarrierSpacing", err)
	}
	return nil
}

func (ie *SubcarrierSpacing) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode SubcarrierSpacing", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
