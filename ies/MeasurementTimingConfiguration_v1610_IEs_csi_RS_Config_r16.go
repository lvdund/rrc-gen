package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasurementTimingConfiguration_v1610_IEs_csi_RS_Config_r16 struct {
	csi_RS_SubcarrierSpacing_r16 SubcarrierSpacing   `madatory`
	csi_RS_CellMobility_r16      CSI_RS_CellMobility `madatory`
	refSSBFreq_r16               ARFCN_ValueNR       `madatory`
}

func (ie *MeasurementTimingConfiguration_v1610_IEs_csi_RS_Config_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.csi_RS_SubcarrierSpacing_r16.Encode(w); err != nil {
		return utils.WrapError("Encode csi_RS_SubcarrierSpacing_r16", err)
	}
	if err = ie.csi_RS_CellMobility_r16.Encode(w); err != nil {
		return utils.WrapError("Encode csi_RS_CellMobility_r16", err)
	}
	if err = ie.refSSBFreq_r16.Encode(w); err != nil {
		return utils.WrapError("Encode refSSBFreq_r16", err)
	}
	return nil
}

func (ie *MeasurementTimingConfiguration_v1610_IEs_csi_RS_Config_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.csi_RS_SubcarrierSpacing_r16.Decode(r); err != nil {
		return utils.WrapError("Decode csi_RS_SubcarrierSpacing_r16", err)
	}
	if err = ie.csi_RS_CellMobility_r16.Decode(r); err != nil {
		return utils.WrapError("Decode csi_RS_CellMobility_r16", err)
	}
	if err = ie.refSSBFreq_r16.Decode(r); err != nil {
		return utils.WrapError("Decode refSSBFreq_r16", err)
	}
	return nil
}
