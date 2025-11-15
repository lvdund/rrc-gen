package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	RACH_ConfigGeneric_preambleTransMax_Enum_n3   uper.Enumerated = 0
	RACH_ConfigGeneric_preambleTransMax_Enum_n4   uper.Enumerated = 1
	RACH_ConfigGeneric_preambleTransMax_Enum_n5   uper.Enumerated = 2
	RACH_ConfigGeneric_preambleTransMax_Enum_n6   uper.Enumerated = 3
	RACH_ConfigGeneric_preambleTransMax_Enum_n7   uper.Enumerated = 4
	RACH_ConfigGeneric_preambleTransMax_Enum_n8   uper.Enumerated = 5
	RACH_ConfigGeneric_preambleTransMax_Enum_n10  uper.Enumerated = 6
	RACH_ConfigGeneric_preambleTransMax_Enum_n20  uper.Enumerated = 7
	RACH_ConfigGeneric_preambleTransMax_Enum_n50  uper.Enumerated = 8
	RACH_ConfigGeneric_preambleTransMax_Enum_n100 uper.Enumerated = 9
	RACH_ConfigGeneric_preambleTransMax_Enum_n200 uper.Enumerated = 10
)

type RACH_ConfigGeneric_preambleTransMax struct {
	Value uper.Enumerated `lb:0,ub:10,madatory`
}

func (ie *RACH_ConfigGeneric_preambleTransMax) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 10}, false); err != nil {
		return utils.WrapError("Encode RACH_ConfigGeneric_preambleTransMax", err)
	}
	return nil
}

func (ie *RACH_ConfigGeneric_preambleTransMax) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 10}, false); err != nil {
		return utils.WrapError("Decode RACH_ConfigGeneric_preambleTransMax", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
