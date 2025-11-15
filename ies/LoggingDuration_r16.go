package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	LoggingDuration_r16_Enum_min10  uper.Enumerated = 0
	LoggingDuration_r16_Enum_min20  uper.Enumerated = 1
	LoggingDuration_r16_Enum_min40  uper.Enumerated = 2
	LoggingDuration_r16_Enum_min60  uper.Enumerated = 3
	LoggingDuration_r16_Enum_min90  uper.Enumerated = 4
	LoggingDuration_r16_Enum_min120 uper.Enumerated = 5
	LoggingDuration_r16_Enum_spare2 uper.Enumerated = 6
	LoggingDuration_r16_Enum_spare1 uper.Enumerated = 7
)

type LoggingDuration_r16 struct {
	Value uper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *LoggingDuration_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode LoggingDuration_r16", err)
	}
	return nil
}

func (ie *LoggingDuration_r16) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode LoggingDuration_r16", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
