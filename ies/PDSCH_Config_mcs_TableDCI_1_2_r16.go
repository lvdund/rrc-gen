package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	PDSCH_Config_mcs_TableDCI_1_2_r16_Enum_qam256     uper.Enumerated = 0
	PDSCH_Config_mcs_TableDCI_1_2_r16_Enum_qam64LowSE uper.Enumerated = 1
)

type PDSCH_Config_mcs_TableDCI_1_2_r16 struct {
	Value uper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *PDSCH_Config_mcs_TableDCI_1_2_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode PDSCH_Config_mcs_TableDCI_1_2_r16", err)
	}
	return nil
}

func (ie *PDSCH_Config_mcs_TableDCI_1_2_r16) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode PDSCH_Config_mcs_TableDCI_1_2_r16", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
