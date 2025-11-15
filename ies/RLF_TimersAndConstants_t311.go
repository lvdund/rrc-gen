package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	RLF_TimersAndConstants_t311_Enum_ms1000  uper.Enumerated = 0
	RLF_TimersAndConstants_t311_Enum_ms3000  uper.Enumerated = 1
	RLF_TimersAndConstants_t311_Enum_ms5000  uper.Enumerated = 2
	RLF_TimersAndConstants_t311_Enum_ms10000 uper.Enumerated = 3
	RLF_TimersAndConstants_t311_Enum_ms15000 uper.Enumerated = 4
	RLF_TimersAndConstants_t311_Enum_ms20000 uper.Enumerated = 5
	RLF_TimersAndConstants_t311_Enum_ms30000 uper.Enumerated = 6
)

type RLF_TimersAndConstants_t311 struct {
	Value uper.Enumerated `lb:0,ub:6,madatory`
}

func (ie *RLF_TimersAndConstants_t311) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 6}, false); err != nil {
		return utils.WrapError("Encode RLF_TimersAndConstants_t311", err)
	}
	return nil
}

func (ie *RLF_TimersAndConstants_t311) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 6}, false); err != nil {
		return utils.WrapError("Decode RLF_TimersAndConstants_t311", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
