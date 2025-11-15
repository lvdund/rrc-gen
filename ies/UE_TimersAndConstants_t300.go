package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	UE_TimersAndConstants_t300_Enum_ms100  uper.Enumerated = 0
	UE_TimersAndConstants_t300_Enum_ms200  uper.Enumerated = 1
	UE_TimersAndConstants_t300_Enum_ms300  uper.Enumerated = 2
	UE_TimersAndConstants_t300_Enum_ms400  uper.Enumerated = 3
	UE_TimersAndConstants_t300_Enum_ms600  uper.Enumerated = 4
	UE_TimersAndConstants_t300_Enum_ms1000 uper.Enumerated = 5
	UE_TimersAndConstants_t300_Enum_ms1500 uper.Enumerated = 6
	UE_TimersAndConstants_t300_Enum_ms2000 uper.Enumerated = 7
)

type UE_TimersAndConstants_t300 struct {
	Value uper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *UE_TimersAndConstants_t300) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode UE_TimersAndConstants_t300", err)
	}
	return nil
}

func (ie *UE_TimersAndConstants_t300) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode UE_TimersAndConstants_t300", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
