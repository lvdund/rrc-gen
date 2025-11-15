package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	FrequencyInfoUL_SIB_frequencyShift7p5khz_Enum_true uper.Enumerated = 0
)

type FrequencyInfoUL_SIB_frequencyShift7p5khz struct {
	Value uper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *FrequencyInfoUL_SIB_frequencyShift7p5khz) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode FrequencyInfoUL_SIB_frequencyShift7p5khz", err)
	}
	return nil
}

func (ie *FrequencyInfoUL_SIB_frequencyShift7p5khz) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode FrequencyInfoUL_SIB_frequencyShift7p5khz", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
