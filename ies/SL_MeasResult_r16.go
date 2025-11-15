package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_MeasResult_r16 struct {
	sl_ResultDMRS_r16 *SL_MeasQuantityResult_r16 `optional`
}

func (ie *SL_MeasResult_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.sl_ResultDMRS_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.sl_ResultDMRS_r16 != nil {
		if err = ie.sl_ResultDMRS_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_ResultDMRS_r16", err)
		}
	}
	return nil
}

func (ie *SL_MeasResult_r16) Decode(r *uper.UperReader) error {
	var err error
	var sl_ResultDMRS_r16Present bool
	if sl_ResultDMRS_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if sl_ResultDMRS_r16Present {
		ie.sl_ResultDMRS_r16 = new(SL_MeasQuantityResult_r16)
		if err = ie.sl_ResultDMRS_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_ResultDMRS_r16", err)
		}
	}
	return nil
}
