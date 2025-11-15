package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UAC_BarringInfoSet_v1700 struct {
	uac_BarringFactorForAI3_r17 *UAC_BarringInfoSet_v1700_uac_BarringFactorForAI3_r17 `optional`
}

func (ie *UAC_BarringInfoSet_v1700) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.uac_BarringFactorForAI3_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.uac_BarringFactorForAI3_r17 != nil {
		if err = ie.uac_BarringFactorForAI3_r17.Encode(w); err != nil {
			return utils.WrapError("Encode uac_BarringFactorForAI3_r17", err)
		}
	}
	return nil
}

func (ie *UAC_BarringInfoSet_v1700) Decode(r *uper.UperReader) error {
	var err error
	var uac_BarringFactorForAI3_r17Present bool
	if uac_BarringFactorForAI3_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if uac_BarringFactorForAI3_r17Present {
		ie.uac_BarringFactorForAI3_r17 = new(UAC_BarringInfoSet_v1700_uac_BarringFactorForAI3_r17)
		if err = ie.uac_BarringFactorForAI3_r17.Decode(r); err != nil {
			return utils.WrapError("Decode uac_BarringFactorForAI3_r17", err)
		}
	}
	return nil
}
