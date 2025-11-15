package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	PDSCH_ServingCellConfig_nrofHARQ_ProcessesForPDSCH_Enum_n2  uper.Enumerated = 0
	PDSCH_ServingCellConfig_nrofHARQ_ProcessesForPDSCH_Enum_n4  uper.Enumerated = 1
	PDSCH_ServingCellConfig_nrofHARQ_ProcessesForPDSCH_Enum_n6  uper.Enumerated = 2
	PDSCH_ServingCellConfig_nrofHARQ_ProcessesForPDSCH_Enum_n10 uper.Enumerated = 3
	PDSCH_ServingCellConfig_nrofHARQ_ProcessesForPDSCH_Enum_n12 uper.Enumerated = 4
	PDSCH_ServingCellConfig_nrofHARQ_ProcessesForPDSCH_Enum_n16 uper.Enumerated = 5
)

type PDSCH_ServingCellConfig_nrofHARQ_ProcessesForPDSCH struct {
	Value uper.Enumerated `lb:0,ub:5,madatory`
}

func (ie *PDSCH_ServingCellConfig_nrofHARQ_ProcessesForPDSCH) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("Encode PDSCH_ServingCellConfig_nrofHARQ_ProcessesForPDSCH", err)
	}
	return nil
}

func (ie *PDSCH_ServingCellConfig_nrofHARQ_ProcessesForPDSCH) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("Decode PDSCH_ServingCellConfig_nrofHARQ_ProcessesForPDSCH", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
