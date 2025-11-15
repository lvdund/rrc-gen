package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	CSI_RS_CellMobility_csi_rs_MeasurementBW_nrofPRBs_Enum_size24  uper.Enumerated = 0
	CSI_RS_CellMobility_csi_rs_MeasurementBW_nrofPRBs_Enum_size48  uper.Enumerated = 1
	CSI_RS_CellMobility_csi_rs_MeasurementBW_nrofPRBs_Enum_size96  uper.Enumerated = 2
	CSI_RS_CellMobility_csi_rs_MeasurementBW_nrofPRBs_Enum_size192 uper.Enumerated = 3
	CSI_RS_CellMobility_csi_rs_MeasurementBW_nrofPRBs_Enum_size264 uper.Enumerated = 4
)

type CSI_RS_CellMobility_csi_rs_MeasurementBW_nrofPRBs struct {
	Value uper.Enumerated `lb:0,ub:4,madatory`
}

func (ie *CSI_RS_CellMobility_csi_rs_MeasurementBW_nrofPRBs) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
		return utils.WrapError("Encode CSI_RS_CellMobility_csi_rs_MeasurementBW_nrofPRBs", err)
	}
	return nil
}

func (ie *CSI_RS_CellMobility_csi_rs_MeasurementBW_nrofPRBs) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
		return utils.WrapError("Decode CSI_RS_CellMobility_csi_rs_MeasurementBW_nrofPRBs", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
