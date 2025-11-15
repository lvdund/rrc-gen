package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MRDC_Parameters_v1630_maxUplinkDutyCycle_interBandENDC_FDD_TDD_PC2_r16 struct {
	maxUplinkDutyCycle_FDD_TDD_EN_DC1_r16 *MRDC_Parameters_v1630_maxUplinkDutyCycle_interBandENDC_FDD_TDD_PC2_r16_maxUplinkDutyCycle_FDD_TDD_EN_DC1_r16 `optional`
	maxUplinkDutyCycle_FDD_TDD_EN_DC2_r16 *MRDC_Parameters_v1630_maxUplinkDutyCycle_interBandENDC_FDD_TDD_PC2_r16_maxUplinkDutyCycle_FDD_TDD_EN_DC2_r16 `optional`
}

func (ie *MRDC_Parameters_v1630_maxUplinkDutyCycle_interBandENDC_FDD_TDD_PC2_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.maxUplinkDutyCycle_FDD_TDD_EN_DC1_r16 != nil, ie.maxUplinkDutyCycle_FDD_TDD_EN_DC2_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.maxUplinkDutyCycle_FDD_TDD_EN_DC1_r16 != nil {
		if err = ie.maxUplinkDutyCycle_FDD_TDD_EN_DC1_r16.Encode(w); err != nil {
			return utils.WrapError("Encode maxUplinkDutyCycle_FDD_TDD_EN_DC1_r16", err)
		}
	}
	if ie.maxUplinkDutyCycle_FDD_TDD_EN_DC2_r16 != nil {
		if err = ie.maxUplinkDutyCycle_FDD_TDD_EN_DC2_r16.Encode(w); err != nil {
			return utils.WrapError("Encode maxUplinkDutyCycle_FDD_TDD_EN_DC2_r16", err)
		}
	}
	return nil
}

func (ie *MRDC_Parameters_v1630_maxUplinkDutyCycle_interBandENDC_FDD_TDD_PC2_r16) Decode(r *uper.UperReader) error {
	var err error
	var maxUplinkDutyCycle_FDD_TDD_EN_DC1_r16Present bool
	if maxUplinkDutyCycle_FDD_TDD_EN_DC1_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var maxUplinkDutyCycle_FDD_TDD_EN_DC2_r16Present bool
	if maxUplinkDutyCycle_FDD_TDD_EN_DC2_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if maxUplinkDutyCycle_FDD_TDD_EN_DC1_r16Present {
		ie.maxUplinkDutyCycle_FDD_TDD_EN_DC1_r16 = new(MRDC_Parameters_v1630_maxUplinkDutyCycle_interBandENDC_FDD_TDD_PC2_r16_maxUplinkDutyCycle_FDD_TDD_EN_DC1_r16)
		if err = ie.maxUplinkDutyCycle_FDD_TDD_EN_DC1_r16.Decode(r); err != nil {
			return utils.WrapError("Decode maxUplinkDutyCycle_FDD_TDD_EN_DC1_r16", err)
		}
	}
	if maxUplinkDutyCycle_FDD_TDD_EN_DC2_r16Present {
		ie.maxUplinkDutyCycle_FDD_TDD_EN_DC2_r16 = new(MRDC_Parameters_v1630_maxUplinkDutyCycle_interBandENDC_FDD_TDD_PC2_r16_maxUplinkDutyCycle_FDD_TDD_EN_DC2_r16)
		if err = ie.maxUplinkDutyCycle_FDD_TDD_EN_DC2_r16.Decode(r); err != nil {
			return utils.WrapError("Decode maxUplinkDutyCycle_FDD_TDD_EN_DC2_r16", err)
		}
	}
	return nil
}
