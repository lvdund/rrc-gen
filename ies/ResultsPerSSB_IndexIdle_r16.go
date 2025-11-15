package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ResultsPerSSB_IndexIdle_r16 struct {
	ssb_Index_r16   SSB_Index                                    `madatory`
	ssb_Results_r16 *ResultsPerSSB_IndexIdle_r16_ssb_Results_r16 `optional`
}

func (ie *ResultsPerSSB_IndexIdle_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.ssb_Results_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.ssb_Index_r16.Encode(w); err != nil {
		return utils.WrapError("Encode ssb_Index_r16", err)
	}
	if ie.ssb_Results_r16 != nil {
		if err = ie.ssb_Results_r16.Encode(w); err != nil {
			return utils.WrapError("Encode ssb_Results_r16", err)
		}
	}
	return nil
}

func (ie *ResultsPerSSB_IndexIdle_r16) Decode(r *uper.UperReader) error {
	var err error
	var ssb_Results_r16Present bool
	if ssb_Results_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.ssb_Index_r16.Decode(r); err != nil {
		return utils.WrapError("Decode ssb_Index_r16", err)
	}
	if ssb_Results_r16Present {
		ie.ssb_Results_r16 = new(ResultsPerSSB_IndexIdle_r16_ssb_Results_r16)
		if err = ie.ssb_Results_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ssb_Results_r16", err)
		}
	}
	return nil
}
