package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	RLF_Report_r16_nr_RLF_Report_r16_choCellId_r17_Choice_nothing uint64 = iota
	RLF_Report_r16_nr_RLF_Report_r16_choCellId_r17_Choice_cellGlobalId_r17
	RLF_Report_r16_nr_RLF_Report_r16_choCellId_r17_Choice_pci_arfcn_r17
)

type RLF_Report_r16_nr_RLF_Report_r16_choCellId_r17 struct {
	Choice           uint64
	cellGlobalId_r17 *CGI_Info_Logging_r16
	pci_arfcn_r17    *PCI_ARFCN_NR_r16
}

func (ie *RLF_Report_r16_nr_RLF_Report_r16_choCellId_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RLF_Report_r16_nr_RLF_Report_r16_choCellId_r17_Choice_cellGlobalId_r17:
		if err = ie.cellGlobalId_r17.Encode(w); err != nil {
			err = utils.WrapError("Encode cellGlobalId_r17", err)
		}
	case RLF_Report_r16_nr_RLF_Report_r16_choCellId_r17_Choice_pci_arfcn_r17:
		if err = ie.pci_arfcn_r17.Encode(w); err != nil {
			err = utils.WrapError("Encode pci_arfcn_r17", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *RLF_Report_r16_nr_RLF_Report_r16_choCellId_r17) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RLF_Report_r16_nr_RLF_Report_r16_choCellId_r17_Choice_cellGlobalId_r17:
		ie.cellGlobalId_r17 = new(CGI_Info_Logging_r16)
		if err = ie.cellGlobalId_r17.Decode(r); err != nil {
			return utils.WrapError("Decode cellGlobalId_r17", err)
		}
	case RLF_Report_r16_nr_RLF_Report_r16_choCellId_r17_Choice_pci_arfcn_r17:
		ie.pci_arfcn_r17 = new(PCI_ARFCN_NR_r16)
		if err = ie.pci_arfcn_r17.Decode(r); err != nil {
			return utils.WrapError("Decode pci_arfcn_r17", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
