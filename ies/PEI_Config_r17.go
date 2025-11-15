package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PEI_Config_r17 struct {
	po_NumPerPEI_r17       PEI_Config_r17_po_NumPerPEI_r17      `madatory`
	payloadSizeDCI_2_7_r17 int64                                `lb:1,ub:maxDCI_2_7_Size_r17,madatory`
	pei_FrameOffset_r17    int64                                `lb:0,ub:16,madatory`
	subgroupConfig_r17     SubgroupConfig_r17                   `madatory`
	lastUsedCellOnly_r17   *PEI_Config_r17_lastUsedCellOnly_r17 `optional`
}

func (ie *PEI_Config_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.lastUsedCellOnly_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.po_NumPerPEI_r17.Encode(w); err != nil {
		return utils.WrapError("Encode po_NumPerPEI_r17", err)
	}
	if err = w.WriteInteger(ie.payloadSizeDCI_2_7_r17, &uper.Constraint{Lb: 1, Ub: maxDCI_2_7_Size_r17}, false); err != nil {
		return utils.WrapError("WriteInteger payloadSizeDCI_2_7_r17", err)
	}
	if err = w.WriteInteger(ie.pei_FrameOffset_r17, &uper.Constraint{Lb: 0, Ub: 16}, false); err != nil {
		return utils.WrapError("WriteInteger pei_FrameOffset_r17", err)
	}
	if err = ie.subgroupConfig_r17.Encode(w); err != nil {
		return utils.WrapError("Encode subgroupConfig_r17", err)
	}
	if ie.lastUsedCellOnly_r17 != nil {
		if err = ie.lastUsedCellOnly_r17.Encode(w); err != nil {
			return utils.WrapError("Encode lastUsedCellOnly_r17", err)
		}
	}
	return nil
}

func (ie *PEI_Config_r17) Decode(r *uper.UperReader) error {
	var err error
	var lastUsedCellOnly_r17Present bool
	if lastUsedCellOnly_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.po_NumPerPEI_r17.Decode(r); err != nil {
		return utils.WrapError("Decode po_NumPerPEI_r17", err)
	}
	var tmp_int_payloadSizeDCI_2_7_r17 int64
	if tmp_int_payloadSizeDCI_2_7_r17, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: maxDCI_2_7_Size_r17}, false); err != nil {
		return utils.WrapError("ReadInteger payloadSizeDCI_2_7_r17", err)
	}
	ie.payloadSizeDCI_2_7_r17 = tmp_int_payloadSizeDCI_2_7_r17
	var tmp_int_pei_FrameOffset_r17 int64
	if tmp_int_pei_FrameOffset_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 16}, false); err != nil {
		return utils.WrapError("ReadInteger pei_FrameOffset_r17", err)
	}
	ie.pei_FrameOffset_r17 = tmp_int_pei_FrameOffset_r17
	if err = ie.subgroupConfig_r17.Decode(r); err != nil {
		return utils.WrapError("Decode subgroupConfig_r17", err)
	}
	if lastUsedCellOnly_r17Present {
		ie.lastUsedCellOnly_r17 = new(PEI_Config_r17_lastUsedCellOnly_r17)
		if err = ie.lastUsedCellOnly_r17.Decode(r); err != nil {
			return utils.WrapError("Decode lastUsedCellOnly_r17", err)
		}
	}
	return nil
}
