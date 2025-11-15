package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CA_ParametersNR_v1630 struct {
	simulTX_SRS_AntSwitchingInterBandUL_CA_r16 *SimulSRS_ForAntennaSwitching_r16                                `optional`
	beamManagementType_r16                     *CA_ParametersNR_v1630_beamManagementType_r16                    `optional`
	intraBandFreqSeparationUL_AggBW_GapBW_r16  *CA_ParametersNR_v1630_intraBandFreqSeparationUL_AggBW_GapBW_r16 `optional`
	interCA_NonAlignedFrame_B_r16              *CA_ParametersNR_v1630_interCA_NonAlignedFrame_B_r16             `optional`
}

func (ie *CA_ParametersNR_v1630) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.simulTX_SRS_AntSwitchingInterBandUL_CA_r16 != nil, ie.beamManagementType_r16 != nil, ie.intraBandFreqSeparationUL_AggBW_GapBW_r16 != nil, ie.interCA_NonAlignedFrame_B_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.simulTX_SRS_AntSwitchingInterBandUL_CA_r16 != nil {
		if err = ie.simulTX_SRS_AntSwitchingInterBandUL_CA_r16.Encode(w); err != nil {
			return utils.WrapError("Encode simulTX_SRS_AntSwitchingInterBandUL_CA_r16", err)
		}
	}
	if ie.beamManagementType_r16 != nil {
		if err = ie.beamManagementType_r16.Encode(w); err != nil {
			return utils.WrapError("Encode beamManagementType_r16", err)
		}
	}
	if ie.intraBandFreqSeparationUL_AggBW_GapBW_r16 != nil {
		if err = ie.intraBandFreqSeparationUL_AggBW_GapBW_r16.Encode(w); err != nil {
			return utils.WrapError("Encode intraBandFreqSeparationUL_AggBW_GapBW_r16", err)
		}
	}
	if ie.interCA_NonAlignedFrame_B_r16 != nil {
		if err = ie.interCA_NonAlignedFrame_B_r16.Encode(w); err != nil {
			return utils.WrapError("Encode interCA_NonAlignedFrame_B_r16", err)
		}
	}
	return nil
}

func (ie *CA_ParametersNR_v1630) Decode(r *uper.UperReader) error {
	var err error
	var simulTX_SRS_AntSwitchingInterBandUL_CA_r16Present bool
	if simulTX_SRS_AntSwitchingInterBandUL_CA_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var beamManagementType_r16Present bool
	if beamManagementType_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var intraBandFreqSeparationUL_AggBW_GapBW_r16Present bool
	if intraBandFreqSeparationUL_AggBW_GapBW_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var interCA_NonAlignedFrame_B_r16Present bool
	if interCA_NonAlignedFrame_B_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if simulTX_SRS_AntSwitchingInterBandUL_CA_r16Present {
		ie.simulTX_SRS_AntSwitchingInterBandUL_CA_r16 = new(SimulSRS_ForAntennaSwitching_r16)
		if err = ie.simulTX_SRS_AntSwitchingInterBandUL_CA_r16.Decode(r); err != nil {
			return utils.WrapError("Decode simulTX_SRS_AntSwitchingInterBandUL_CA_r16", err)
		}
	}
	if beamManagementType_r16Present {
		ie.beamManagementType_r16 = new(CA_ParametersNR_v1630_beamManagementType_r16)
		if err = ie.beamManagementType_r16.Decode(r); err != nil {
			return utils.WrapError("Decode beamManagementType_r16", err)
		}
	}
	if intraBandFreqSeparationUL_AggBW_GapBW_r16Present {
		ie.intraBandFreqSeparationUL_AggBW_GapBW_r16 = new(CA_ParametersNR_v1630_intraBandFreqSeparationUL_AggBW_GapBW_r16)
		if err = ie.intraBandFreqSeparationUL_AggBW_GapBW_r16.Decode(r); err != nil {
			return utils.WrapError("Decode intraBandFreqSeparationUL_AggBW_GapBW_r16", err)
		}
	}
	if interCA_NonAlignedFrame_B_r16Present {
		ie.interCA_NonAlignedFrame_B_r16 = new(CA_ParametersNR_v1630_interCA_NonAlignedFrame_B_r16)
		if err = ie.interCA_NonAlignedFrame_B_r16.Decode(r); err != nil {
			return utils.WrapError("Decode interCA_NonAlignedFrame_B_r16", err)
		}
	}
	return nil
}
