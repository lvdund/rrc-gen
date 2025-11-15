package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	LoggingInterval_r16_Enum_ms320    uper.Enumerated = 0
	LoggingInterval_r16_Enum_ms640    uper.Enumerated = 1
	LoggingInterval_r16_Enum_ms1280   uper.Enumerated = 2
	LoggingInterval_r16_Enum_ms2560   uper.Enumerated = 3
	LoggingInterval_r16_Enum_ms5120   uper.Enumerated = 4
	LoggingInterval_r16_Enum_ms10240  uper.Enumerated = 5
	LoggingInterval_r16_Enum_ms20480  uper.Enumerated = 6
	LoggingInterval_r16_Enum_ms30720  uper.Enumerated = 7
	LoggingInterval_r16_Enum_ms40960  uper.Enumerated = 8
	LoggingInterval_r16_Enum_ms61440  uper.Enumerated = 9
	LoggingInterval_r16_Enum_infinity uper.Enumerated = 10
)

type LoggingInterval_r16 struct {
	Value uper.Enumerated `lb:0,ub:10,madatory`
}

func (ie *LoggingInterval_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 10}, false); err != nil {
		return utils.WrapError("Encode LoggingInterval_r16", err)
	}
	return nil
}

func (ie *LoggingInterval_r16) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 10}, false); err != nil {
		return utils.WrapError("Decode LoggingInterval_r16", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
