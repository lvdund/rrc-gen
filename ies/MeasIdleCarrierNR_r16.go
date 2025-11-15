package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasIdleCarrierNR_r16 struct {
	carrierFreq_r16          ARFCN_ValueNR                               `madatory`
	ssbSubcarrierSpacing_r16 SubcarrierSpacing                           `madatory`
	frequencyBandList        *MultiFrequencyBandListNR                   `optional`
	measCellListNR_r16       *CellListNR_r16                             `optional`
	reportQuantities_r16     MeasIdleCarrierNR_r16_reportQuantities_r16  `madatory`
	qualityThreshold_r16     *MeasIdleCarrierNR_r16_qualityThreshold_r16 `optional`
	ssb_MeasConfig_r16       *MeasIdleCarrierNR_r16_ssb_MeasConfig_r16   `lb:2,ub:maxNrofSS_BlocksToAverage,optional`
	beamMeasConfigIdle_r16   *BeamMeasConfigIdle_NR_r16                  `optional`
}

func (ie *MeasIdleCarrierNR_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.frequencyBandList != nil, ie.measCellListNR_r16 != nil, ie.qualityThreshold_r16 != nil, ie.ssb_MeasConfig_r16 != nil, ie.beamMeasConfigIdle_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.carrierFreq_r16.Encode(w); err != nil {
		return utils.WrapError("Encode carrierFreq_r16", err)
	}
	if err = ie.ssbSubcarrierSpacing_r16.Encode(w); err != nil {
		return utils.WrapError("Encode ssbSubcarrierSpacing_r16", err)
	}
	if ie.frequencyBandList != nil {
		if err = ie.frequencyBandList.Encode(w); err != nil {
			return utils.WrapError("Encode frequencyBandList", err)
		}
	}
	if ie.measCellListNR_r16 != nil {
		if err = ie.measCellListNR_r16.Encode(w); err != nil {
			return utils.WrapError("Encode measCellListNR_r16", err)
		}
	}
	if err = ie.reportQuantities_r16.Encode(w); err != nil {
		return utils.WrapError("Encode reportQuantities_r16", err)
	}
	if ie.qualityThreshold_r16 != nil {
		if err = ie.qualityThreshold_r16.Encode(w); err != nil {
			return utils.WrapError("Encode qualityThreshold_r16", err)
		}
	}
	if ie.ssb_MeasConfig_r16 != nil {
		if err = ie.ssb_MeasConfig_r16.Encode(w); err != nil {
			return utils.WrapError("Encode ssb_MeasConfig_r16", err)
		}
	}
	if ie.beamMeasConfigIdle_r16 != nil {
		if err = ie.beamMeasConfigIdle_r16.Encode(w); err != nil {
			return utils.WrapError("Encode beamMeasConfigIdle_r16", err)
		}
	}
	return nil
}

func (ie *MeasIdleCarrierNR_r16) Decode(r *uper.UperReader) error {
	var err error
	var frequencyBandListPresent bool
	if frequencyBandListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var measCellListNR_r16Present bool
	if measCellListNR_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var qualityThreshold_r16Present bool
	if qualityThreshold_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ssb_MeasConfig_r16Present bool
	if ssb_MeasConfig_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var beamMeasConfigIdle_r16Present bool
	if beamMeasConfigIdle_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.carrierFreq_r16.Decode(r); err != nil {
		return utils.WrapError("Decode carrierFreq_r16", err)
	}
	if err = ie.ssbSubcarrierSpacing_r16.Decode(r); err != nil {
		return utils.WrapError("Decode ssbSubcarrierSpacing_r16", err)
	}
	if frequencyBandListPresent {
		ie.frequencyBandList = new(MultiFrequencyBandListNR)
		if err = ie.frequencyBandList.Decode(r); err != nil {
			return utils.WrapError("Decode frequencyBandList", err)
		}
	}
	if measCellListNR_r16Present {
		ie.measCellListNR_r16 = new(CellListNR_r16)
		if err = ie.measCellListNR_r16.Decode(r); err != nil {
			return utils.WrapError("Decode measCellListNR_r16", err)
		}
	}
	if err = ie.reportQuantities_r16.Decode(r); err != nil {
		return utils.WrapError("Decode reportQuantities_r16", err)
	}
	if qualityThreshold_r16Present {
		ie.qualityThreshold_r16 = new(MeasIdleCarrierNR_r16_qualityThreshold_r16)
		if err = ie.qualityThreshold_r16.Decode(r); err != nil {
			return utils.WrapError("Decode qualityThreshold_r16", err)
		}
	}
	if ssb_MeasConfig_r16Present {
		ie.ssb_MeasConfig_r16 = new(MeasIdleCarrierNR_r16_ssb_MeasConfig_r16)
		if err = ie.ssb_MeasConfig_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ssb_MeasConfig_r16", err)
		}
	}
	if beamMeasConfigIdle_r16Present {
		ie.beamMeasConfigIdle_r16 = new(BeamMeasConfigIdle_NR_r16)
		if err = ie.beamMeasConfigIdle_r16.Decode(r); err != nil {
			return utils.WrapError("Decode beamMeasConfigIdle_r16", err)
		}
	}
	return nil
}
