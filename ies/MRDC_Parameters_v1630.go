package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MRDC_Parameters_v1630 struct {
	maxUplinkDutyCycle_interBandENDC_FDD_TDD_PC2_r16 *MRDC_Parameters_v1630_maxUplinkDutyCycle_interBandENDC_FDD_TDD_PC2_r16 `optional`
	interBandMRDC_WithOverlapDL_Bands_r16            *MRDC_Parameters_v1630_interBandMRDC_WithOverlapDL_Bands_r16            `optional`
}

func (ie *MRDC_Parameters_v1630) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.maxUplinkDutyCycle_interBandENDC_FDD_TDD_PC2_r16 != nil, ie.interBandMRDC_WithOverlapDL_Bands_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.maxUplinkDutyCycle_interBandENDC_FDD_TDD_PC2_r16 != nil {
		if err = ie.maxUplinkDutyCycle_interBandENDC_FDD_TDD_PC2_r16.Encode(w); err != nil {
			return utils.WrapError("Encode maxUplinkDutyCycle_interBandENDC_FDD_TDD_PC2_r16", err)
		}
	}
	if ie.interBandMRDC_WithOverlapDL_Bands_r16 != nil {
		if err = ie.interBandMRDC_WithOverlapDL_Bands_r16.Encode(w); err != nil {
			return utils.WrapError("Encode interBandMRDC_WithOverlapDL_Bands_r16", err)
		}
	}
	return nil
}

func (ie *MRDC_Parameters_v1630) Decode(r *uper.UperReader) error {
	var err error
	var maxUplinkDutyCycle_interBandENDC_FDD_TDD_PC2_r16Present bool
	if maxUplinkDutyCycle_interBandENDC_FDD_TDD_PC2_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var interBandMRDC_WithOverlapDL_Bands_r16Present bool
	if interBandMRDC_WithOverlapDL_Bands_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if maxUplinkDutyCycle_interBandENDC_FDD_TDD_PC2_r16Present {
		ie.maxUplinkDutyCycle_interBandENDC_FDD_TDD_PC2_r16 = new(MRDC_Parameters_v1630_maxUplinkDutyCycle_interBandENDC_FDD_TDD_PC2_r16)
		if err = ie.maxUplinkDutyCycle_interBandENDC_FDD_TDD_PC2_r16.Decode(r); err != nil {
			return utils.WrapError("Decode maxUplinkDutyCycle_interBandENDC_FDD_TDD_PC2_r16", err)
		}
	}
	if interBandMRDC_WithOverlapDL_Bands_r16Present {
		ie.interBandMRDC_WithOverlapDL_Bands_r16 = new(MRDC_Parameters_v1630_interBandMRDC_WithOverlapDL_Bands_r16)
		if err = ie.interBandMRDC_WithOverlapDL_Bands_r16.Decode(r); err != nil {
			return utils.WrapError("Decode interBandMRDC_WithOverlapDL_Bands_r16", err)
		}
	}
	return nil
}
