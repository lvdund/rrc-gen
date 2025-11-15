package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SSB_MTC4_r17 struct {
	pci_List_r17 []PhysCellId `lb:1,ub:maxNrofPCIsPerSMTC,optional`
	offset_r17   int64        `lb:0,ub:159,madatory`
}

func (ie *SSB_MTC4_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{len(ie.pci_List_r17) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.pci_List_r17) > 0 {
		tmp_pci_List_r17 := utils.NewSequence[*PhysCellId]([]*PhysCellId{}, uper.Constraint{Lb: 1, Ub: maxNrofPCIsPerSMTC}, false)
		for _, i := range ie.pci_List_r17 {
			tmp_pci_List_r17.Value = append(tmp_pci_List_r17.Value, &i)
		}
		if err = tmp_pci_List_r17.Encode(w); err != nil {
			return utils.WrapError("Encode pci_List_r17", err)
		}
	}
	if err = w.WriteInteger(ie.offset_r17, &uper.Constraint{Lb: 0, Ub: 159}, false); err != nil {
		return utils.WrapError("WriteInteger offset_r17", err)
	}
	return nil
}

func (ie *SSB_MTC4_r17) Decode(r *uper.UperReader) error {
	var err error
	var pci_List_r17Present bool
	if pci_List_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if pci_List_r17Present {
		tmp_pci_List_r17 := utils.NewSequence[*PhysCellId]([]*PhysCellId{}, uper.Constraint{Lb: 1, Ub: maxNrofPCIsPerSMTC}, false)
		fn_pci_List_r17 := func() *PhysCellId {
			return new(PhysCellId)
		}
		if err = tmp_pci_List_r17.Decode(r, fn_pci_List_r17); err != nil {
			return utils.WrapError("Decode pci_List_r17", err)
		}
		ie.pci_List_r17 = []PhysCellId{}
		for _, i := range tmp_pci_List_r17.Value {
			ie.pci_List_r17 = append(ie.pci_List_r17, *i)
		}
	}
	var tmp_int_offset_r17 int64
	if tmp_int_offset_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 159}, false); err != nil {
		return utils.WrapError("ReadInteger offset_r17", err)
	}
	ie.offset_r17 = tmp_int_offset_r17
	return nil
}
