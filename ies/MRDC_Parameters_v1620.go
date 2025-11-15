package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MRDC_Parameters_v1620 struct {
	maxUplinkDutyCycle_interBandENDC_TDD_PC2_r16 *MRDC_Parameters_v1620_maxUplinkDutyCycle_interBandENDC_TDD_PC2_r16 `optional`
	tdm_restrictionTDD_endc_r16                  *MRDC_Parameters_v1620_tdm_restrictionTDD_endc_r16                  `optional`
	tdm_restrictionFDD_endc_r16                  *MRDC_Parameters_v1620_tdm_restrictionFDD_endc_r16                  `optional`
	singleUL_HARQ_offsetTDD_PCell_r16            *MRDC_Parameters_v1620_singleUL_HARQ_offsetTDD_PCell_r16            `optional`
	tdm_restrictionDualTX_FDD_endc_r16           *MRDC_Parameters_v1620_tdm_restrictionDualTX_FDD_endc_r16           `optional`
}

func (ie *MRDC_Parameters_v1620) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.maxUplinkDutyCycle_interBandENDC_TDD_PC2_r16 != nil, ie.tdm_restrictionTDD_endc_r16 != nil, ie.tdm_restrictionFDD_endc_r16 != nil, ie.singleUL_HARQ_offsetTDD_PCell_r16 != nil, ie.tdm_restrictionDualTX_FDD_endc_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.maxUplinkDutyCycle_interBandENDC_TDD_PC2_r16 != nil {
		if err = ie.maxUplinkDutyCycle_interBandENDC_TDD_PC2_r16.Encode(w); err != nil {
			return utils.WrapError("Encode maxUplinkDutyCycle_interBandENDC_TDD_PC2_r16", err)
		}
	}
	if ie.tdm_restrictionTDD_endc_r16 != nil {
		if err = ie.tdm_restrictionTDD_endc_r16.Encode(w); err != nil {
			return utils.WrapError("Encode tdm_restrictionTDD_endc_r16", err)
		}
	}
	if ie.tdm_restrictionFDD_endc_r16 != nil {
		if err = ie.tdm_restrictionFDD_endc_r16.Encode(w); err != nil {
			return utils.WrapError("Encode tdm_restrictionFDD_endc_r16", err)
		}
	}
	if ie.singleUL_HARQ_offsetTDD_PCell_r16 != nil {
		if err = ie.singleUL_HARQ_offsetTDD_PCell_r16.Encode(w); err != nil {
			return utils.WrapError("Encode singleUL_HARQ_offsetTDD_PCell_r16", err)
		}
	}
	if ie.tdm_restrictionDualTX_FDD_endc_r16 != nil {
		if err = ie.tdm_restrictionDualTX_FDD_endc_r16.Encode(w); err != nil {
			return utils.WrapError("Encode tdm_restrictionDualTX_FDD_endc_r16", err)
		}
	}
	return nil
}

func (ie *MRDC_Parameters_v1620) Decode(r *uper.UperReader) error {
	var err error
	var maxUplinkDutyCycle_interBandENDC_TDD_PC2_r16Present bool
	if maxUplinkDutyCycle_interBandENDC_TDD_PC2_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var tdm_restrictionTDD_endc_r16Present bool
	if tdm_restrictionTDD_endc_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var tdm_restrictionFDD_endc_r16Present bool
	if tdm_restrictionFDD_endc_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var singleUL_HARQ_offsetTDD_PCell_r16Present bool
	if singleUL_HARQ_offsetTDD_PCell_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var tdm_restrictionDualTX_FDD_endc_r16Present bool
	if tdm_restrictionDualTX_FDD_endc_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if maxUplinkDutyCycle_interBandENDC_TDD_PC2_r16Present {
		ie.maxUplinkDutyCycle_interBandENDC_TDD_PC2_r16 = new(MRDC_Parameters_v1620_maxUplinkDutyCycle_interBandENDC_TDD_PC2_r16)
		if err = ie.maxUplinkDutyCycle_interBandENDC_TDD_PC2_r16.Decode(r); err != nil {
			return utils.WrapError("Decode maxUplinkDutyCycle_interBandENDC_TDD_PC2_r16", err)
		}
	}
	if tdm_restrictionTDD_endc_r16Present {
		ie.tdm_restrictionTDD_endc_r16 = new(MRDC_Parameters_v1620_tdm_restrictionTDD_endc_r16)
		if err = ie.tdm_restrictionTDD_endc_r16.Decode(r); err != nil {
			return utils.WrapError("Decode tdm_restrictionTDD_endc_r16", err)
		}
	}
	if tdm_restrictionFDD_endc_r16Present {
		ie.tdm_restrictionFDD_endc_r16 = new(MRDC_Parameters_v1620_tdm_restrictionFDD_endc_r16)
		if err = ie.tdm_restrictionFDD_endc_r16.Decode(r); err != nil {
			return utils.WrapError("Decode tdm_restrictionFDD_endc_r16", err)
		}
	}
	if singleUL_HARQ_offsetTDD_PCell_r16Present {
		ie.singleUL_HARQ_offsetTDD_PCell_r16 = new(MRDC_Parameters_v1620_singleUL_HARQ_offsetTDD_PCell_r16)
		if err = ie.singleUL_HARQ_offsetTDD_PCell_r16.Decode(r); err != nil {
			return utils.WrapError("Decode singleUL_HARQ_offsetTDD_PCell_r16", err)
		}
	}
	if tdm_restrictionDualTX_FDD_endc_r16Present {
		ie.tdm_restrictionDualTX_FDD_endc_r16 = new(MRDC_Parameters_v1620_tdm_restrictionDualTX_FDD_endc_r16)
		if err = ie.tdm_restrictionDualTX_FDD_endc_r16.Decode(r); err != nil {
			return utils.WrapError("Decode tdm_restrictionDualTX_FDD_endc_r16", err)
		}
	}
	return nil
}
