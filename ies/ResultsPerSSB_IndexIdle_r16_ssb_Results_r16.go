package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ResultsPerSSB_IndexIdle_r16_ssb_Results_r16 struct {
	ssb_RSRP_Result_r16 *RSRP_Range `optional`
	ssb_RSRQ_Result_r16 *RSRQ_Range `optional`
}

func (ie *ResultsPerSSB_IndexIdle_r16_ssb_Results_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.ssb_RSRP_Result_r16 != nil, ie.ssb_RSRQ_Result_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.ssb_RSRP_Result_r16 != nil {
		if err = ie.ssb_RSRP_Result_r16.Encode(w); err != nil {
			return utils.WrapError("Encode ssb_RSRP_Result_r16", err)
		}
	}
	if ie.ssb_RSRQ_Result_r16 != nil {
		if err = ie.ssb_RSRQ_Result_r16.Encode(w); err != nil {
			return utils.WrapError("Encode ssb_RSRQ_Result_r16", err)
		}
	}
	return nil
}

func (ie *ResultsPerSSB_IndexIdle_r16_ssb_Results_r16) Decode(r *uper.UperReader) error {
	var err error
	var ssb_RSRP_Result_r16Present bool
	if ssb_RSRP_Result_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ssb_RSRQ_Result_r16Present bool
	if ssb_RSRQ_Result_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if ssb_RSRP_Result_r16Present {
		ie.ssb_RSRP_Result_r16 = new(RSRP_Range)
		if err = ie.ssb_RSRP_Result_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ssb_RSRP_Result_r16", err)
		}
	}
	if ssb_RSRQ_Result_r16Present {
		ie.ssb_RSRQ_Result_r16 = new(RSRQ_Range)
		if err = ie.ssb_RSRQ_Result_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ssb_RSRQ_Result_r16", err)
		}
	}
	return nil
}
