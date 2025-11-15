package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	RLF_TimersAndConstants_n311_Enum_n1  uper.Enumerated = 0
	RLF_TimersAndConstants_n311_Enum_n2  uper.Enumerated = 1
	RLF_TimersAndConstants_n311_Enum_n3  uper.Enumerated = 2
	RLF_TimersAndConstants_n311_Enum_n4  uper.Enumerated = 3
	RLF_TimersAndConstants_n311_Enum_n5  uper.Enumerated = 4
	RLF_TimersAndConstants_n311_Enum_n6  uper.Enumerated = 5
	RLF_TimersAndConstants_n311_Enum_n8  uper.Enumerated = 6
	RLF_TimersAndConstants_n311_Enum_n10 uper.Enumerated = 7
)

type RLF_TimersAndConstants_n311 struct {
	Value uper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *RLF_TimersAndConstants_n311) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode RLF_TimersAndConstants_n311", err)
	}
	return nil
}

func (ie *RLF_TimersAndConstants_n311) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode RLF_TimersAndConstants_n311", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
