package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SSB_InfoNcell_r16 struct {
	physicalCellId_r16    PhysCellId             `madatory`
	ssb_IndexNcell_r16    *SSB_Index             `optional`
	ssb_Configuration_r16 *SSB_Configuration_r16 `optional`
}

func (ie *SSB_InfoNcell_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.ssb_IndexNcell_r16 != nil, ie.ssb_Configuration_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.physicalCellId_r16.Encode(w); err != nil {
		return utils.WrapError("Encode physicalCellId_r16", err)
	}
	if ie.ssb_IndexNcell_r16 != nil {
		if err = ie.ssb_IndexNcell_r16.Encode(w); err != nil {
			return utils.WrapError("Encode ssb_IndexNcell_r16", err)
		}
	}
	if ie.ssb_Configuration_r16 != nil {
		if err = ie.ssb_Configuration_r16.Encode(w); err != nil {
			return utils.WrapError("Encode ssb_Configuration_r16", err)
		}
	}
	return nil
}

func (ie *SSB_InfoNcell_r16) Decode(r *uper.UperReader) error {
	var err error
	var ssb_IndexNcell_r16Present bool
	if ssb_IndexNcell_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ssb_Configuration_r16Present bool
	if ssb_Configuration_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.physicalCellId_r16.Decode(r); err != nil {
		return utils.WrapError("Decode physicalCellId_r16", err)
	}
	if ssb_IndexNcell_r16Present {
		ie.ssb_IndexNcell_r16 = new(SSB_Index)
		if err = ie.ssb_IndexNcell_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ssb_IndexNcell_r16", err)
		}
	}
	if ssb_Configuration_r16Present {
		ie.ssb_Configuration_r16 = new(SSB_Configuration_r16)
		if err = ie.ssb_Configuration_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ssb_Configuration_r16", err)
		}
	}
	return nil
}
