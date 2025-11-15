package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	PCI_Range_range_pci_Enum_n4     uper.Enumerated = 0
	PCI_Range_range_pci_Enum_n8     uper.Enumerated = 1
	PCI_Range_range_pci_Enum_n12    uper.Enumerated = 2
	PCI_Range_range_pci_Enum_n16    uper.Enumerated = 3
	PCI_Range_range_pci_Enum_n24    uper.Enumerated = 4
	PCI_Range_range_pci_Enum_n32    uper.Enumerated = 5
	PCI_Range_range_pci_Enum_n48    uper.Enumerated = 6
	PCI_Range_range_pci_Enum_n64    uper.Enumerated = 7
	PCI_Range_range_pci_Enum_n84    uper.Enumerated = 8
	PCI_Range_range_pci_Enum_n96    uper.Enumerated = 9
	PCI_Range_range_pci_Enum_n128   uper.Enumerated = 10
	PCI_Range_range_pci_Enum_n168   uper.Enumerated = 11
	PCI_Range_range_pci_Enum_n252   uper.Enumerated = 12
	PCI_Range_range_pci_Enum_n504   uper.Enumerated = 13
	PCI_Range_range_pci_Enum_n1008  uper.Enumerated = 14
	PCI_Range_range_pci_Enum_spare1 uper.Enumerated = 15
)

type PCI_Range_range_pci struct {
	Value uper.Enumerated `lb:0,ub:15,madatory`
}

func (ie *PCI_Range_range_pci) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Encode PCI_Range_range_pci", err)
	}
	return nil
}

func (ie *PCI_Range_range_pci) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Decode PCI_Range_range_pci", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
