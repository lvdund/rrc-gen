package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PCI_RangeElement struct {
	pci_RangeIndex PCI_RangeIndex `madatory`
	pci_Range      PCI_Range      `madatory`
}

func (ie *PCI_RangeElement) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.pci_RangeIndex.Encode(w); err != nil {
		return utils.WrapError("Encode pci_RangeIndex", err)
	}
	if err = ie.pci_Range.Encode(w); err != nil {
		return utils.WrapError("Encode pci_Range", err)
	}
	return nil
}

func (ie *PCI_RangeElement) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.pci_RangeIndex.Decode(r); err != nil {
		return utils.WrapError("Decode pci_RangeIndex", err)
	}
	if err = ie.pci_Range.Decode(r); err != nil {
		return utils.WrapError("Decode pci_Range", err)
	}
	return nil
}
