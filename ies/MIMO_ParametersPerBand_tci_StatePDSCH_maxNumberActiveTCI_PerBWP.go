package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	MIMO_ParametersPerBand_tci_StatePDSCH_maxNumberActiveTCI_PerBWP_Enum_n1 uper.Enumerated = 0
	MIMO_ParametersPerBand_tci_StatePDSCH_maxNumberActiveTCI_PerBWP_Enum_n2 uper.Enumerated = 1
	MIMO_ParametersPerBand_tci_StatePDSCH_maxNumberActiveTCI_PerBWP_Enum_n4 uper.Enumerated = 2
	MIMO_ParametersPerBand_tci_StatePDSCH_maxNumberActiveTCI_PerBWP_Enum_n8 uper.Enumerated = 3
)

type MIMO_ParametersPerBand_tci_StatePDSCH_maxNumberActiveTCI_PerBWP struct {
	Value uper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *MIMO_ParametersPerBand_tci_StatePDSCH_maxNumberActiveTCI_PerBWP) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode MIMO_ParametersPerBand_tci_StatePDSCH_maxNumberActiveTCI_PerBWP", err)
	}
	return nil
}

func (ie *MIMO_ParametersPerBand_tci_StatePDSCH_maxNumberActiveTCI_PerBWP) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode MIMO_ParametersPerBand_tci_StatePDSCH_maxNumberActiveTCI_PerBWP", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
