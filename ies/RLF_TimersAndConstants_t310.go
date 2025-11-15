package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	RLF_TimersAndConstants_t310_Enum_ms0    uper.Enumerated = 0
	RLF_TimersAndConstants_t310_Enum_ms50   uper.Enumerated = 1
	RLF_TimersAndConstants_t310_Enum_ms100  uper.Enumerated = 2
	RLF_TimersAndConstants_t310_Enum_ms200  uper.Enumerated = 3
	RLF_TimersAndConstants_t310_Enum_ms500  uper.Enumerated = 4
	RLF_TimersAndConstants_t310_Enum_ms1000 uper.Enumerated = 5
	RLF_TimersAndConstants_t310_Enum_ms2000 uper.Enumerated = 6
	RLF_TimersAndConstants_t310_Enum_ms4000 uper.Enumerated = 7
	RLF_TimersAndConstants_t310_Enum_ms6000 uper.Enumerated = 8
)

type RLF_TimersAndConstants_t310 struct {
	Value uper.Enumerated `lb:0,ub:8,madatory`
}

func (ie *RLF_TimersAndConstants_t310) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 8}, false); err != nil {
		return utils.WrapError("Encode RLF_TimersAndConstants_t310", err)
	}
	return nil
}

func (ie *RLF_TimersAndConstants_t310) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 8}, false); err != nil {
		return utils.WrapError("Decode RLF_TimersAndConstants_t310", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
