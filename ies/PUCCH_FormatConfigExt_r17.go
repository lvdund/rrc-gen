package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PUCCH_FormatConfigExt_r17 struct {
	maxCodeRateLP_r17 *PUCCH_MaxCodeRate `optional`
}

func (ie *PUCCH_FormatConfigExt_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.maxCodeRateLP_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.maxCodeRateLP_r17 != nil {
		if err = ie.maxCodeRateLP_r17.Encode(w); err != nil {
			return utils.WrapError("Encode maxCodeRateLP_r17", err)
		}
	}
	return nil
}

func (ie *PUCCH_FormatConfigExt_r17) Decode(r *uper.UperReader) error {
	var err error
	var maxCodeRateLP_r17Present bool
	if maxCodeRateLP_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if maxCodeRateLP_r17Present {
		ie.maxCodeRateLP_r17 = new(PUCCH_MaxCodeRate)
		if err = ie.maxCodeRateLP_r17.Decode(r); err != nil {
			return utils.WrapError("Decode maxCodeRateLP_r17", err)
		}
	}
	return nil
}
