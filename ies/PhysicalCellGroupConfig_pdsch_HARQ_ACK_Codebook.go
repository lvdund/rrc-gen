package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	PhysicalCellGroupConfig_pdsch_HARQ_ACK_Codebook_Enum_semiStatic uper.Enumerated = 0
	PhysicalCellGroupConfig_pdsch_HARQ_ACK_Codebook_Enum_dynamic    uper.Enumerated = 1
)

type PhysicalCellGroupConfig_pdsch_HARQ_ACK_Codebook struct {
	Value uper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *PhysicalCellGroupConfig_pdsch_HARQ_ACK_Codebook) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode PhysicalCellGroupConfig_pdsch_HARQ_ACK_Codebook", err)
	}
	return nil
}

func (ie *PhysicalCellGroupConfig_pdsch_HARQ_ACK_Codebook) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode PhysicalCellGroupConfig_pdsch_HARQ_ACK_Codebook", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
