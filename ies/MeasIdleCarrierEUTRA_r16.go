package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasIdleCarrierEUTRA_r16 struct {
	carrierFreqEUTRA_r16      ARFCN_ValueEUTRA                                    `madatory`
	allowedMeasBandwidth_r16  EUTRA_AllowedMeasBandwidth                          `madatory`
	measCellListEUTRA_r16     *CellListEUTRA_r16                                  `optional`
	reportQuantitiesEUTRA_r16 MeasIdleCarrierEUTRA_r16_reportQuantitiesEUTRA_r16  `madatory`
	qualityThresholdEUTRA_r16 *MeasIdleCarrierEUTRA_r16_qualityThresholdEUTRA_r16 `optional`
}

func (ie *MeasIdleCarrierEUTRA_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.measCellListEUTRA_r16 != nil, ie.qualityThresholdEUTRA_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.carrierFreqEUTRA_r16.Encode(w); err != nil {
		return utils.WrapError("Encode carrierFreqEUTRA_r16", err)
	}
	if err = ie.allowedMeasBandwidth_r16.Encode(w); err != nil {
		return utils.WrapError("Encode allowedMeasBandwidth_r16", err)
	}
	if ie.measCellListEUTRA_r16 != nil {
		if err = ie.measCellListEUTRA_r16.Encode(w); err != nil {
			return utils.WrapError("Encode measCellListEUTRA_r16", err)
		}
	}
	if err = ie.reportQuantitiesEUTRA_r16.Encode(w); err != nil {
		return utils.WrapError("Encode reportQuantitiesEUTRA_r16", err)
	}
	if ie.qualityThresholdEUTRA_r16 != nil {
		if err = ie.qualityThresholdEUTRA_r16.Encode(w); err != nil {
			return utils.WrapError("Encode qualityThresholdEUTRA_r16", err)
		}
	}
	return nil
}

func (ie *MeasIdleCarrierEUTRA_r16) Decode(r *uper.UperReader) error {
	var err error
	var measCellListEUTRA_r16Present bool
	if measCellListEUTRA_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var qualityThresholdEUTRA_r16Present bool
	if qualityThresholdEUTRA_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.carrierFreqEUTRA_r16.Decode(r); err != nil {
		return utils.WrapError("Decode carrierFreqEUTRA_r16", err)
	}
	if err = ie.allowedMeasBandwidth_r16.Decode(r); err != nil {
		return utils.WrapError("Decode allowedMeasBandwidth_r16", err)
	}
	if measCellListEUTRA_r16Present {
		ie.measCellListEUTRA_r16 = new(CellListEUTRA_r16)
		if err = ie.measCellListEUTRA_r16.Decode(r); err != nil {
			return utils.WrapError("Decode measCellListEUTRA_r16", err)
		}
	}
	if err = ie.reportQuantitiesEUTRA_r16.Decode(r); err != nil {
		return utils.WrapError("Decode reportQuantitiesEUTRA_r16", err)
	}
	if qualityThresholdEUTRA_r16Present {
		ie.qualityThresholdEUTRA_r16 = new(MeasIdleCarrierEUTRA_r16_qualityThresholdEUTRA_r16)
		if err = ie.qualityThresholdEUTRA_r16.Decode(r); err != nil {
			return utils.WrapError("Decode qualityThresholdEUTRA_r16", err)
		}
	}
	return nil
}
