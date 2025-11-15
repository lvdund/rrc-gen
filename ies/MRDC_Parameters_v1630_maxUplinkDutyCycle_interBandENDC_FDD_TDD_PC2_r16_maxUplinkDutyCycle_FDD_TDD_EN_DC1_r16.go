package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	MRDC_Parameters_v1630_maxUplinkDutyCycle_interBandENDC_FDD_TDD_PC2_r16_maxUplinkDutyCycle_FDD_TDD_EN_DC1_r16_Enum_n30  uper.Enumerated = 0
	MRDC_Parameters_v1630_maxUplinkDutyCycle_interBandENDC_FDD_TDD_PC2_r16_maxUplinkDutyCycle_FDD_TDD_EN_DC1_r16_Enum_n40  uper.Enumerated = 1
	MRDC_Parameters_v1630_maxUplinkDutyCycle_interBandENDC_FDD_TDD_PC2_r16_maxUplinkDutyCycle_FDD_TDD_EN_DC1_r16_Enum_n50  uper.Enumerated = 2
	MRDC_Parameters_v1630_maxUplinkDutyCycle_interBandENDC_FDD_TDD_PC2_r16_maxUplinkDutyCycle_FDD_TDD_EN_DC1_r16_Enum_n60  uper.Enumerated = 3
	MRDC_Parameters_v1630_maxUplinkDutyCycle_interBandENDC_FDD_TDD_PC2_r16_maxUplinkDutyCycle_FDD_TDD_EN_DC1_r16_Enum_n70  uper.Enumerated = 4
	MRDC_Parameters_v1630_maxUplinkDutyCycle_interBandENDC_FDD_TDD_PC2_r16_maxUplinkDutyCycle_FDD_TDD_EN_DC1_r16_Enum_n80  uper.Enumerated = 5
	MRDC_Parameters_v1630_maxUplinkDutyCycle_interBandENDC_FDD_TDD_PC2_r16_maxUplinkDutyCycle_FDD_TDD_EN_DC1_r16_Enum_n90  uper.Enumerated = 6
	MRDC_Parameters_v1630_maxUplinkDutyCycle_interBandENDC_FDD_TDD_PC2_r16_maxUplinkDutyCycle_FDD_TDD_EN_DC1_r16_Enum_n100 uper.Enumerated = 7
)

type MRDC_Parameters_v1630_maxUplinkDutyCycle_interBandENDC_FDD_TDD_PC2_r16_maxUplinkDutyCycle_FDD_TDD_EN_DC1_r16 struct {
	Value uper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *MRDC_Parameters_v1630_maxUplinkDutyCycle_interBandENDC_FDD_TDD_PC2_r16_maxUplinkDutyCycle_FDD_TDD_EN_DC1_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode MRDC_Parameters_v1630_maxUplinkDutyCycle_interBandENDC_FDD_TDD_PC2_r16_maxUplinkDutyCycle_FDD_TDD_EN_DC1_r16", err)
	}
	return nil
}

func (ie *MRDC_Parameters_v1630_maxUplinkDutyCycle_interBandENDC_FDD_TDD_PC2_r16_maxUplinkDutyCycle_FDD_TDD_EN_DC1_r16) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode MRDC_Parameters_v1630_maxUplinkDutyCycle_interBandENDC_FDD_TDD_PC2_r16_maxUplinkDutyCycle_FDD_TDD_EN_DC1_r16", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
