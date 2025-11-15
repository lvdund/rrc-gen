package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SSB_MTC3_r16 struct {
	periodicityAndOffset_r16 SSB_MTC3_r16_periodicityAndOffset_r16 `lb:0,ub:4,madatory`
	duration_r16             SSB_MTC3_r16_duration_r16             `madatory`
	pci_List_r16             []PhysCellId                          `lb:1,ub:maxNrofPCIsPerSMTC,optional`
	ssb_ToMeasure_r16        *SSB_ToMeasure                        `optional,setuprelease`
}

func (ie *SSB_MTC3_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{len(ie.pci_List_r16) > 0, ie.ssb_ToMeasure_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.periodicityAndOffset_r16.Encode(w); err != nil {
		return utils.WrapError("Encode periodicityAndOffset_r16", err)
	}
	if err = ie.duration_r16.Encode(w); err != nil {
		return utils.WrapError("Encode duration_r16", err)
	}
	if len(ie.pci_List_r16) > 0 {
		tmp_pci_List_r16 := utils.NewSequence[*PhysCellId]([]*PhysCellId{}, uper.Constraint{Lb: 1, Ub: maxNrofPCIsPerSMTC}, false)
		for _, i := range ie.pci_List_r16 {
			tmp_pci_List_r16.Value = append(tmp_pci_List_r16.Value, &i)
		}
		if err = tmp_pci_List_r16.Encode(w); err != nil {
			return utils.WrapError("Encode pci_List_r16", err)
		}
	}
	if ie.ssb_ToMeasure_r16 != nil {
		tmp_ssb_ToMeasure_r16 := utils.SetupRelease[*SSB_ToMeasure]{
			Setup: ie.ssb_ToMeasure_r16,
		}
		if err = tmp_ssb_ToMeasure_r16.Encode(w); err != nil {
			return utils.WrapError("Encode ssb_ToMeasure_r16", err)
		}
	}
	return nil
}

func (ie *SSB_MTC3_r16) Decode(r *uper.UperReader) error {
	var err error
	var pci_List_r16Present bool
	if pci_List_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ssb_ToMeasure_r16Present bool
	if ssb_ToMeasure_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.periodicityAndOffset_r16.Decode(r); err != nil {
		return utils.WrapError("Decode periodicityAndOffset_r16", err)
	}
	if err = ie.duration_r16.Decode(r); err != nil {
		return utils.WrapError("Decode duration_r16", err)
	}
	if pci_List_r16Present {
		tmp_pci_List_r16 := utils.NewSequence[*PhysCellId]([]*PhysCellId{}, uper.Constraint{Lb: 1, Ub: maxNrofPCIsPerSMTC}, false)
		fn_pci_List_r16 := func() *PhysCellId {
			return new(PhysCellId)
		}
		if err = tmp_pci_List_r16.Decode(r, fn_pci_List_r16); err != nil {
			return utils.WrapError("Decode pci_List_r16", err)
		}
		ie.pci_List_r16 = []PhysCellId{}
		for _, i := range tmp_pci_List_r16.Value {
			ie.pci_List_r16 = append(ie.pci_List_r16, *i)
		}
	}
	if ssb_ToMeasure_r16Present {
		tmp_ssb_ToMeasure_r16 := utils.SetupRelease[*SSB_ToMeasure]{}
		if err = tmp_ssb_ToMeasure_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ssb_ToMeasure_r16", err)
		}
		ie.ssb_ToMeasure_r16 = tmp_ssb_ToMeasure_r16.Setup
	}
	return nil
}
