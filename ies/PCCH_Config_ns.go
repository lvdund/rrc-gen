package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	PCCH_Config_ns_Enum_four uper.Enumerated = 0
	PCCH_Config_ns_Enum_two  uper.Enumerated = 1
	PCCH_Config_ns_Enum_one  uper.Enumerated = 2
)

type PCCH_Config_ns struct {
	Value uper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *PCCH_Config_ns) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode PCCH_Config_ns", err)
	}
	return nil
}

func (ie *PCCH_Config_ns) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode PCCH_Config_ns", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
