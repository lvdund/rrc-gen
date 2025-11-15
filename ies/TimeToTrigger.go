package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	TimeToTrigger_Enum_ms0    uper.Enumerated = 0
	TimeToTrigger_Enum_ms40   uper.Enumerated = 1
	TimeToTrigger_Enum_ms64   uper.Enumerated = 2
	TimeToTrigger_Enum_ms80   uper.Enumerated = 3
	TimeToTrigger_Enum_ms100  uper.Enumerated = 4
	TimeToTrigger_Enum_ms128  uper.Enumerated = 5
	TimeToTrigger_Enum_ms160  uper.Enumerated = 6
	TimeToTrigger_Enum_ms256  uper.Enumerated = 7
	TimeToTrigger_Enum_ms320  uper.Enumerated = 8
	TimeToTrigger_Enum_ms480  uper.Enumerated = 9
	TimeToTrigger_Enum_ms512  uper.Enumerated = 10
	TimeToTrigger_Enum_ms640  uper.Enumerated = 11
	TimeToTrigger_Enum_ms1024 uper.Enumerated = 12
	TimeToTrigger_Enum_ms1280 uper.Enumerated = 13
	TimeToTrigger_Enum_ms2560 uper.Enumerated = 14
	TimeToTrigger_Enum_ms5120 uper.Enumerated = 15
)

type TimeToTrigger struct {
	Value uper.Enumerated `lb:0,ub:15,madatory`
}

func (ie *TimeToTrigger) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Encode TimeToTrigger", err)
	}
	return nil
}

func (ie *TimeToTrigger) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Decode TimeToTrigger", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
