package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	UE_TimersAndConstants_t311_Enum_ms1000  uper.Enumerated = 0
	UE_TimersAndConstants_t311_Enum_ms3000  uper.Enumerated = 1
	UE_TimersAndConstants_t311_Enum_ms5000  uper.Enumerated = 2
	UE_TimersAndConstants_t311_Enum_ms10000 uper.Enumerated = 3
	UE_TimersAndConstants_t311_Enum_ms15000 uper.Enumerated = 4
	UE_TimersAndConstants_t311_Enum_ms20000 uper.Enumerated = 5
	UE_TimersAndConstants_t311_Enum_ms30000 uper.Enumerated = 6
)

type UE_TimersAndConstants_t311 struct {
	Value uper.Enumerated `lb:0,ub:6,madatory`
}

func (ie *UE_TimersAndConstants_t311) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 6}, false); err != nil {
		return utils.WrapError("Encode UE_TimersAndConstants_t311", err)
	}
	return nil
}

func (ie *UE_TimersAndConstants_t311) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 6}, false); err != nil {
		return utils.WrapError("Decode UE_TimersAndConstants_t311", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
