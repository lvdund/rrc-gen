package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MIMO_ParametersPerBand_ssb_csirs_SINR_measurement_r16 struct {
	maxNumberSSB_CSIRS_OneTx_CMR_r16    MIMO_ParametersPerBand_ssb_csirs_SINR_measurement_r16_maxNumberSSB_CSIRS_OneTx_CMR_r16    `madatory`
	maxNumberCSI_IM_NZP_IMR_res_r16     MIMO_ParametersPerBand_ssb_csirs_SINR_measurement_r16_maxNumberCSI_IM_NZP_IMR_res_r16     `madatory`
	maxNumberCSIRS_2Tx_res_r16          MIMO_ParametersPerBand_ssb_csirs_SINR_measurement_r16_maxNumberCSIRS_2Tx_res_r16          `madatory`
	maxNumberSSB_CSIRS_res_r16          MIMO_ParametersPerBand_ssb_csirs_SINR_measurement_r16_maxNumberSSB_CSIRS_res_r16          `madatory`
	maxNumberCSI_IM_NZP_IMR_res_mem_r16 MIMO_ParametersPerBand_ssb_csirs_SINR_measurement_r16_maxNumberCSI_IM_NZP_IMR_res_mem_r16 `madatory`
	supportedCSI_RS_Density_CMR_r16     MIMO_ParametersPerBand_ssb_csirs_SINR_measurement_r16_supportedCSI_RS_Density_CMR_r16     `madatory`
	maxNumberAperiodicCSI_RS_Res_r16    MIMO_ParametersPerBand_ssb_csirs_SINR_measurement_r16_maxNumberAperiodicCSI_RS_Res_r16    `madatory`
	supportedSINR_meas_r16              *MIMO_ParametersPerBand_ssb_csirs_SINR_measurement_r16_supportedSINR_meas_r16             `optional`
}

func (ie *MIMO_ParametersPerBand_ssb_csirs_SINR_measurement_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.supportedSINR_meas_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.maxNumberSSB_CSIRS_OneTx_CMR_r16.Encode(w); err != nil {
		return utils.WrapError("Encode maxNumberSSB_CSIRS_OneTx_CMR_r16", err)
	}
	if err = ie.maxNumberCSI_IM_NZP_IMR_res_r16.Encode(w); err != nil {
		return utils.WrapError("Encode maxNumberCSI_IM_NZP_IMR_res_r16", err)
	}
	if err = ie.maxNumberCSIRS_2Tx_res_r16.Encode(w); err != nil {
		return utils.WrapError("Encode maxNumberCSIRS_2Tx_res_r16", err)
	}
	if err = ie.maxNumberSSB_CSIRS_res_r16.Encode(w); err != nil {
		return utils.WrapError("Encode maxNumberSSB_CSIRS_res_r16", err)
	}
	if err = ie.maxNumberCSI_IM_NZP_IMR_res_mem_r16.Encode(w); err != nil {
		return utils.WrapError("Encode maxNumberCSI_IM_NZP_IMR_res_mem_r16", err)
	}
	if err = ie.supportedCSI_RS_Density_CMR_r16.Encode(w); err != nil {
		return utils.WrapError("Encode supportedCSI_RS_Density_CMR_r16", err)
	}
	if err = ie.maxNumberAperiodicCSI_RS_Res_r16.Encode(w); err != nil {
		return utils.WrapError("Encode maxNumberAperiodicCSI_RS_Res_r16", err)
	}
	if ie.supportedSINR_meas_r16 != nil {
		if err = ie.supportedSINR_meas_r16.Encode(w); err != nil {
			return utils.WrapError("Encode supportedSINR_meas_r16", err)
		}
	}
	return nil
}

func (ie *MIMO_ParametersPerBand_ssb_csirs_SINR_measurement_r16) Decode(r *uper.UperReader) error {
	var err error
	var supportedSINR_meas_r16Present bool
	if supportedSINR_meas_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.maxNumberSSB_CSIRS_OneTx_CMR_r16.Decode(r); err != nil {
		return utils.WrapError("Decode maxNumberSSB_CSIRS_OneTx_CMR_r16", err)
	}
	if err = ie.maxNumberCSI_IM_NZP_IMR_res_r16.Decode(r); err != nil {
		return utils.WrapError("Decode maxNumberCSI_IM_NZP_IMR_res_r16", err)
	}
	if err = ie.maxNumberCSIRS_2Tx_res_r16.Decode(r); err != nil {
		return utils.WrapError("Decode maxNumberCSIRS_2Tx_res_r16", err)
	}
	if err = ie.maxNumberSSB_CSIRS_res_r16.Decode(r); err != nil {
		return utils.WrapError("Decode maxNumberSSB_CSIRS_res_r16", err)
	}
	if err = ie.maxNumberCSI_IM_NZP_IMR_res_mem_r16.Decode(r); err != nil {
		return utils.WrapError("Decode maxNumberCSI_IM_NZP_IMR_res_mem_r16", err)
	}
	if err = ie.supportedCSI_RS_Density_CMR_r16.Decode(r); err != nil {
		return utils.WrapError("Decode supportedCSI_RS_Density_CMR_r16", err)
	}
	if err = ie.maxNumberAperiodicCSI_RS_Res_r16.Decode(r); err != nil {
		return utils.WrapError("Decode maxNumberAperiodicCSI_RS_Res_r16", err)
	}
	if supportedSINR_meas_r16Present {
		ie.supportedSINR_meas_r16 = new(MIMO_ParametersPerBand_ssb_csirs_SINR_measurement_r16_supportedSINR_meas_r16)
		if err = ie.supportedSINR_meas_r16.Decode(r); err != nil {
			return utils.WrapError("Decode supportedSINR_meas_r16", err)
		}
	}
	return nil
}
