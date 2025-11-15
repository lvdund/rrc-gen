package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MRDC_Parameters_v1620_maxUplinkDutyCycle_interBandENDC_TDD_PC2_r16 struct {
	eutra_TDD_Config0_r16 *MRDC_Parameters_v1620_maxUplinkDutyCycle_interBandENDC_TDD_PC2_r16_eutra_TDD_Config0_r16 `optional`
	eutra_TDD_Config1_r16 *MRDC_Parameters_v1620_maxUplinkDutyCycle_interBandENDC_TDD_PC2_r16_eutra_TDD_Config1_r16 `optional`
	eutra_TDD_Config2_r16 *MRDC_Parameters_v1620_maxUplinkDutyCycle_interBandENDC_TDD_PC2_r16_eutra_TDD_Config2_r16 `optional`
	eutra_TDD_Config3_r16 *MRDC_Parameters_v1620_maxUplinkDutyCycle_interBandENDC_TDD_PC2_r16_eutra_TDD_Config3_r16 `optional`
	eutra_TDD_Config4_r16 *MRDC_Parameters_v1620_maxUplinkDutyCycle_interBandENDC_TDD_PC2_r16_eutra_TDD_Config4_r16 `optional`
	eutra_TDD_Config5_r16 *MRDC_Parameters_v1620_maxUplinkDutyCycle_interBandENDC_TDD_PC2_r16_eutra_TDD_Config5_r16 `optional`
	eutra_TDD_Config6_r16 *MRDC_Parameters_v1620_maxUplinkDutyCycle_interBandENDC_TDD_PC2_r16_eutra_TDD_Config6_r16 `optional`
}

func (ie *MRDC_Parameters_v1620_maxUplinkDutyCycle_interBandENDC_TDD_PC2_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.eutra_TDD_Config0_r16 != nil, ie.eutra_TDD_Config1_r16 != nil, ie.eutra_TDD_Config2_r16 != nil, ie.eutra_TDD_Config3_r16 != nil, ie.eutra_TDD_Config4_r16 != nil, ie.eutra_TDD_Config5_r16 != nil, ie.eutra_TDD_Config6_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.eutra_TDD_Config0_r16 != nil {
		if err = ie.eutra_TDD_Config0_r16.Encode(w); err != nil {
			return utils.WrapError("Encode eutra_TDD_Config0_r16", err)
		}
	}
	if ie.eutra_TDD_Config1_r16 != nil {
		if err = ie.eutra_TDD_Config1_r16.Encode(w); err != nil {
			return utils.WrapError("Encode eutra_TDD_Config1_r16", err)
		}
	}
	if ie.eutra_TDD_Config2_r16 != nil {
		if err = ie.eutra_TDD_Config2_r16.Encode(w); err != nil {
			return utils.WrapError("Encode eutra_TDD_Config2_r16", err)
		}
	}
	if ie.eutra_TDD_Config3_r16 != nil {
		if err = ie.eutra_TDD_Config3_r16.Encode(w); err != nil {
			return utils.WrapError("Encode eutra_TDD_Config3_r16", err)
		}
	}
	if ie.eutra_TDD_Config4_r16 != nil {
		if err = ie.eutra_TDD_Config4_r16.Encode(w); err != nil {
			return utils.WrapError("Encode eutra_TDD_Config4_r16", err)
		}
	}
	if ie.eutra_TDD_Config5_r16 != nil {
		if err = ie.eutra_TDD_Config5_r16.Encode(w); err != nil {
			return utils.WrapError("Encode eutra_TDD_Config5_r16", err)
		}
	}
	if ie.eutra_TDD_Config6_r16 != nil {
		if err = ie.eutra_TDD_Config6_r16.Encode(w); err != nil {
			return utils.WrapError("Encode eutra_TDD_Config6_r16", err)
		}
	}
	return nil
}

func (ie *MRDC_Parameters_v1620_maxUplinkDutyCycle_interBandENDC_TDD_PC2_r16) Decode(r *uper.UperReader) error {
	var err error
	var eutra_TDD_Config0_r16Present bool
	if eutra_TDD_Config0_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var eutra_TDD_Config1_r16Present bool
	if eutra_TDD_Config1_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var eutra_TDD_Config2_r16Present bool
	if eutra_TDD_Config2_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var eutra_TDD_Config3_r16Present bool
	if eutra_TDD_Config3_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var eutra_TDD_Config4_r16Present bool
	if eutra_TDD_Config4_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var eutra_TDD_Config5_r16Present bool
	if eutra_TDD_Config5_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var eutra_TDD_Config6_r16Present bool
	if eutra_TDD_Config6_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if eutra_TDD_Config0_r16Present {
		ie.eutra_TDD_Config0_r16 = new(MRDC_Parameters_v1620_maxUplinkDutyCycle_interBandENDC_TDD_PC2_r16_eutra_TDD_Config0_r16)
		if err = ie.eutra_TDD_Config0_r16.Decode(r); err != nil {
			return utils.WrapError("Decode eutra_TDD_Config0_r16", err)
		}
	}
	if eutra_TDD_Config1_r16Present {
		ie.eutra_TDD_Config1_r16 = new(MRDC_Parameters_v1620_maxUplinkDutyCycle_interBandENDC_TDD_PC2_r16_eutra_TDD_Config1_r16)
		if err = ie.eutra_TDD_Config1_r16.Decode(r); err != nil {
			return utils.WrapError("Decode eutra_TDD_Config1_r16", err)
		}
	}
	if eutra_TDD_Config2_r16Present {
		ie.eutra_TDD_Config2_r16 = new(MRDC_Parameters_v1620_maxUplinkDutyCycle_interBandENDC_TDD_PC2_r16_eutra_TDD_Config2_r16)
		if err = ie.eutra_TDD_Config2_r16.Decode(r); err != nil {
			return utils.WrapError("Decode eutra_TDD_Config2_r16", err)
		}
	}
	if eutra_TDD_Config3_r16Present {
		ie.eutra_TDD_Config3_r16 = new(MRDC_Parameters_v1620_maxUplinkDutyCycle_interBandENDC_TDD_PC2_r16_eutra_TDD_Config3_r16)
		if err = ie.eutra_TDD_Config3_r16.Decode(r); err != nil {
			return utils.WrapError("Decode eutra_TDD_Config3_r16", err)
		}
	}
	if eutra_TDD_Config4_r16Present {
		ie.eutra_TDD_Config4_r16 = new(MRDC_Parameters_v1620_maxUplinkDutyCycle_interBandENDC_TDD_PC2_r16_eutra_TDD_Config4_r16)
		if err = ie.eutra_TDD_Config4_r16.Decode(r); err != nil {
			return utils.WrapError("Decode eutra_TDD_Config4_r16", err)
		}
	}
	if eutra_TDD_Config5_r16Present {
		ie.eutra_TDD_Config5_r16 = new(MRDC_Parameters_v1620_maxUplinkDutyCycle_interBandENDC_TDD_PC2_r16_eutra_TDD_Config5_r16)
		if err = ie.eutra_TDD_Config5_r16.Decode(r); err != nil {
			return utils.WrapError("Decode eutra_TDD_Config5_r16", err)
		}
	}
	if eutra_TDD_Config6_r16Present {
		ie.eutra_TDD_Config6_r16 = new(MRDC_Parameters_v1620_maxUplinkDutyCycle_interBandENDC_TDD_PC2_r16_eutra_TDD_Config6_r16)
		if err = ie.eutra_TDD_Config6_r16.Decode(r); err != nil {
			return utils.WrapError("Decode eutra_TDD_Config6_r16", err)
		}
	}
	return nil
}
