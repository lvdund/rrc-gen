package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_MeasQuantityResult_r16 struct {
	sl_RSRP_r16 *RSRP_Range `optional`
}

func (ie *SL_MeasQuantityResult_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.sl_RSRP_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.sl_RSRP_r16 != nil {
		if err = ie.sl_RSRP_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_RSRP_r16", err)
		}
	}
	return nil
}

func (ie *SL_MeasQuantityResult_r16) Decode(r *uper.UperReader) error {
	var err error
	var sl_RSRP_r16Present bool
	if sl_RSRP_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if sl_RSRP_r16Present {
		ie.sl_RSRP_r16 = new(RSRP_Range)
		if err = ie.sl_RSRP_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_RSRP_r16", err)
		}
	}
	return nil
}
