package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	PDSCH_ConfigPTM_r17_pdsch_AggregationFactor_r17_Enum_n2 uper.Enumerated = 0
	PDSCH_ConfigPTM_r17_pdsch_AggregationFactor_r17_Enum_n4 uper.Enumerated = 1
	PDSCH_ConfigPTM_r17_pdsch_AggregationFactor_r17_Enum_n8 uper.Enumerated = 2
)

type PDSCH_ConfigPTM_r17_pdsch_AggregationFactor_r17 struct {
	Value uper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *PDSCH_ConfigPTM_r17_pdsch_AggregationFactor_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode PDSCH_ConfigPTM_r17_pdsch_AggregationFactor_r17", err)
	}
	return nil
}

func (ie *PDSCH_ConfigPTM_r17_pdsch_AggregationFactor_r17) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode PDSCH_ConfigPTM_r17_pdsch_AggregationFactor_r17", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
