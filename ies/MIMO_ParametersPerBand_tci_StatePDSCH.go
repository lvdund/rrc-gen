package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MIMO_ParametersPerBand_tci_StatePDSCH struct {
	maxNumberConfiguredTCI_StatesPerCC *MIMO_ParametersPerBand_tci_StatePDSCH_maxNumberConfiguredTCI_StatesPerCC `optional`
	maxNumberActiveTCI_PerBWP          *MIMO_ParametersPerBand_tci_StatePDSCH_maxNumberActiveTCI_PerBWP          `optional`
}

func (ie *MIMO_ParametersPerBand_tci_StatePDSCH) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.maxNumberConfiguredTCI_StatesPerCC != nil, ie.maxNumberActiveTCI_PerBWP != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.maxNumberConfiguredTCI_StatesPerCC != nil {
		if err = ie.maxNumberConfiguredTCI_StatesPerCC.Encode(w); err != nil {
			return utils.WrapError("Encode maxNumberConfiguredTCI_StatesPerCC", err)
		}
	}
	if ie.maxNumberActiveTCI_PerBWP != nil {
		if err = ie.maxNumberActiveTCI_PerBWP.Encode(w); err != nil {
			return utils.WrapError("Encode maxNumberActiveTCI_PerBWP", err)
		}
	}
	return nil
}

func (ie *MIMO_ParametersPerBand_tci_StatePDSCH) Decode(r *uper.UperReader) error {
	var err error
	var maxNumberConfiguredTCI_StatesPerCCPresent bool
	if maxNumberConfiguredTCI_StatesPerCCPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var maxNumberActiveTCI_PerBWPPresent bool
	if maxNumberActiveTCI_PerBWPPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if maxNumberConfiguredTCI_StatesPerCCPresent {
		ie.maxNumberConfiguredTCI_StatesPerCC = new(MIMO_ParametersPerBand_tci_StatePDSCH_maxNumberConfiguredTCI_StatesPerCC)
		if err = ie.maxNumberConfiguredTCI_StatesPerCC.Decode(r); err != nil {
			return utils.WrapError("Decode maxNumberConfiguredTCI_StatesPerCC", err)
		}
	}
	if maxNumberActiveTCI_PerBWPPresent {
		ie.maxNumberActiveTCI_PerBWP = new(MIMO_ParametersPerBand_tci_StatePDSCH_maxNumberActiveTCI_PerBWP)
		if err = ie.maxNumberActiveTCI_PerBWP.Decode(r); err != nil {
			return utils.WrapError("Decode maxNumberActiveTCI_PerBWP", err)
		}
	}
	return nil
}
