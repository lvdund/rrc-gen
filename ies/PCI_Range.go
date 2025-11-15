package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PCI_Range struct {
	start     PhysCellId           `madatory`
	range_pci *PCI_Range_range_pci `optional`
}

func (ie *PCI_Range) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.range_pci != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.start.Encode(w); err != nil {
		return utils.WrapError("Encode start", err)
	}
	if ie.range_pci != nil {
		if err = ie.range_pci.Encode(w); err != nil {
			return utils.WrapError("Encode range_pci", err)
		}
	}
	return nil
}

func (ie *PCI_Range) Decode(r *uper.UperReader) error {
	var err error
	var range_pciPresent bool
	if range_pciPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.start.Decode(r); err != nil {
		return utils.WrapError("Decode start", err)
	}
	if range_pciPresent {
		ie.range_pci = new(PCI_Range_range_pci)
		if err = ie.range_pci.Decode(r); err != nil {
			return utils.WrapError("Decode range_pci", err)
		}
	}
	return nil
}
