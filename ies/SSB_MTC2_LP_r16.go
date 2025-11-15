package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SSB_MTC2_LP_r16 struct {
	pci_List    []PhysCellId                `lb:1,ub:maxNrofPCIsPerSMTC,optional`
	periodicity SSB_MTC2_LP_r16_periodicity `madatory`
}

func (ie *SSB_MTC2_LP_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{len(ie.pci_List) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.pci_List) > 0 {
		tmp_pci_List := utils.NewSequence[*PhysCellId]([]*PhysCellId{}, uper.Constraint{Lb: 1, Ub: maxNrofPCIsPerSMTC}, false)
		for _, i := range ie.pci_List {
			tmp_pci_List.Value = append(tmp_pci_List.Value, &i)
		}
		if err = tmp_pci_List.Encode(w); err != nil {
			return utils.WrapError("Encode pci_List", err)
		}
	}
	if err = ie.periodicity.Encode(w); err != nil {
		return utils.WrapError("Encode periodicity", err)
	}
	return nil
}

func (ie *SSB_MTC2_LP_r16) Decode(r *uper.UperReader) error {
	var err error
	var pci_ListPresent bool
	if pci_ListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if pci_ListPresent {
		tmp_pci_List := utils.NewSequence[*PhysCellId]([]*PhysCellId{}, uper.Constraint{Lb: 1, Ub: maxNrofPCIsPerSMTC}, false)
		fn_pci_List := func() *PhysCellId {
			return new(PhysCellId)
		}
		if err = tmp_pci_List.Decode(r, fn_pci_List); err != nil {
			return utils.WrapError("Decode pci_List", err)
		}
		ie.pci_List = []PhysCellId{}
		for _, i := range tmp_pci_List.Value {
			ie.pci_List = append(ie.pci_List, *i)
		}
	}
	if err = ie.periodicity.Decode(r); err != nil {
		return utils.WrapError("Decode periodicity", err)
	}
	return nil
}
