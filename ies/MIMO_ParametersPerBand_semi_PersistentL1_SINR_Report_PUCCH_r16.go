package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MIMO_ParametersPerBand_semi_PersistentL1_SINR_Report_PUCCH_r16 struct {
	supportReportFormat1_2OFDM_syms_r16  *MIMO_ParametersPerBand_semi_PersistentL1_SINR_Report_PUCCH_r16_supportReportFormat1_2OFDM_syms_r16  `optional`
	supportReportFormat4_14OFDM_syms_r16 *MIMO_ParametersPerBand_semi_PersistentL1_SINR_Report_PUCCH_r16_supportReportFormat4_14OFDM_syms_r16 `optional`
}

func (ie *MIMO_ParametersPerBand_semi_PersistentL1_SINR_Report_PUCCH_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.supportReportFormat1_2OFDM_syms_r16 != nil, ie.supportReportFormat4_14OFDM_syms_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.supportReportFormat1_2OFDM_syms_r16 != nil {
		if err = ie.supportReportFormat1_2OFDM_syms_r16.Encode(w); err != nil {
			return utils.WrapError("Encode supportReportFormat1_2OFDM_syms_r16", err)
		}
	}
	if ie.supportReportFormat4_14OFDM_syms_r16 != nil {
		if err = ie.supportReportFormat4_14OFDM_syms_r16.Encode(w); err != nil {
			return utils.WrapError("Encode supportReportFormat4_14OFDM_syms_r16", err)
		}
	}
	return nil
}

func (ie *MIMO_ParametersPerBand_semi_PersistentL1_SINR_Report_PUCCH_r16) Decode(r *uper.UperReader) error {
	var err error
	var supportReportFormat1_2OFDM_syms_r16Present bool
	if supportReportFormat1_2OFDM_syms_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var supportReportFormat4_14OFDM_syms_r16Present bool
	if supportReportFormat4_14OFDM_syms_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if supportReportFormat1_2OFDM_syms_r16Present {
		ie.supportReportFormat1_2OFDM_syms_r16 = new(MIMO_ParametersPerBand_semi_PersistentL1_SINR_Report_PUCCH_r16_supportReportFormat1_2OFDM_syms_r16)
		if err = ie.supportReportFormat1_2OFDM_syms_r16.Decode(r); err != nil {
			return utils.WrapError("Decode supportReportFormat1_2OFDM_syms_r16", err)
		}
	}
	if supportReportFormat4_14OFDM_syms_r16Present {
		ie.supportReportFormat4_14OFDM_syms_r16 = new(MIMO_ParametersPerBand_semi_PersistentL1_SINR_Report_PUCCH_r16_supportReportFormat4_14OFDM_syms_r16)
		if err = ie.supportReportFormat4_14OFDM_syms_r16.Decode(r); err != nil {
			return utils.WrapError("Decode supportReportFormat4_14OFDM_syms_r16", err)
		}
	}
	return nil
}
