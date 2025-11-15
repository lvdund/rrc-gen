package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	RLF_Report_r16_nr_RLF_Report_r16_lastHO_Type_r17_Enum_cho    uper.Enumerated = 0
	RLF_Report_r16_nr_RLF_Report_r16_lastHO_Type_r17_Enum_daps   uper.Enumerated = 1
	RLF_Report_r16_nr_RLF_Report_r16_lastHO_Type_r17_Enum_spare2 uper.Enumerated = 2
	RLF_Report_r16_nr_RLF_Report_r16_lastHO_Type_r17_Enum_spare1 uper.Enumerated = 3
)

type RLF_Report_r16_nr_RLF_Report_r16_lastHO_Type_r17 struct {
	Value uper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *RLF_Report_r16_nr_RLF_Report_r16_lastHO_Type_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode RLF_Report_r16_nr_RLF_Report_r16_lastHO_Type_r17", err)
	}
	return nil
}

func (ie *RLF_Report_r16_nr_RLF_Report_r16_lastHO_Type_r17) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode RLF_Report_r16_nr_RLF_Report_r16_lastHO_Type_r17", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
