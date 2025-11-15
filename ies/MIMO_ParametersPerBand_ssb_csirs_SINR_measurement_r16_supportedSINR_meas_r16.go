package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	MIMO_ParametersPerBand_ssb_csirs_SINR_measurement_r16_supportedSINR_meas_r16_Enum_ssbWithCSI_IM    uper.Enumerated = 0
	MIMO_ParametersPerBand_ssb_csirs_SINR_measurement_r16_supportedSINR_meas_r16_Enum_ssbWithNZP_IMR   uper.Enumerated = 1
	MIMO_ParametersPerBand_ssb_csirs_SINR_measurement_r16_supportedSINR_meas_r16_Enum_csirsWithNZP_IMR uper.Enumerated = 2
	MIMO_ParametersPerBand_ssb_csirs_SINR_measurement_r16_supportedSINR_meas_r16_Enum_csi_RSWithoutIMR uper.Enumerated = 3
)

type MIMO_ParametersPerBand_ssb_csirs_SINR_measurement_r16_supportedSINR_meas_r16 struct {
	Value uper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *MIMO_ParametersPerBand_ssb_csirs_SINR_measurement_r16_supportedSINR_meas_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode MIMO_ParametersPerBand_ssb_csirs_SINR_measurement_r16_supportedSINR_meas_r16", err)
	}
	return nil
}

func (ie *MIMO_ParametersPerBand_ssb_csirs_SINR_measurement_r16_supportedSINR_meas_r16) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode MIMO_ParametersPerBand_ssb_csirs_SINR_measurement_r16_supportedSINR_meas_r16", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
