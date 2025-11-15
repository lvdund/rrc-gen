package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CSI_RS_CellMobility_csi_rs_MeasurementBW struct {
	nrofPRBs CSI_RS_CellMobility_csi_rs_MeasurementBW_nrofPRBs `madatory`
	startPRB int64                                             `lb:0,ub:2169,madatory`
}

func (ie *CSI_RS_CellMobility_csi_rs_MeasurementBW) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.nrofPRBs.Encode(w); err != nil {
		return utils.WrapError("Encode nrofPRBs", err)
	}
	if err = w.WriteInteger(ie.startPRB, &uper.Constraint{Lb: 0, Ub: 2169}, false); err != nil {
		return utils.WrapError("WriteInteger startPRB", err)
	}
	return nil
}

func (ie *CSI_RS_CellMobility_csi_rs_MeasurementBW) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.nrofPRBs.Decode(r); err != nil {
		return utils.WrapError("Decode nrofPRBs", err)
	}
	var tmp_int_startPRB int64
	if tmp_int_startPRB, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 2169}, false); err != nil {
		return utils.WrapError("ReadInteger startPRB", err)
	}
	ie.startPRB = tmp_int_startPRB
	return nil
}
