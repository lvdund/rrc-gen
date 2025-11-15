package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasResultCellSFTD_NR struct {
	physCellId                PhysCellId  `madatory`
	sfn_OffsetResult          int64       `lb:0,ub:1023,madatory`
	frameBoundaryOffsetResult int64       `lb:-30720,ub:30719,madatory`
	rsrp_Result               *RSRP_Range `optional`
}

func (ie *MeasResultCellSFTD_NR) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.rsrp_Result != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.physCellId.Encode(w); err != nil {
		return utils.WrapError("Encode physCellId", err)
	}
	if err = w.WriteInteger(ie.sfn_OffsetResult, &uper.Constraint{Lb: 0, Ub: 1023}, false); err != nil {
		return utils.WrapError("WriteInteger sfn_OffsetResult", err)
	}
	if err = w.WriteInteger(ie.frameBoundaryOffsetResult, &uper.Constraint{Lb: -30720, Ub: 30719}, false); err != nil {
		return utils.WrapError("WriteInteger frameBoundaryOffsetResult", err)
	}
	if ie.rsrp_Result != nil {
		if err = ie.rsrp_Result.Encode(w); err != nil {
			return utils.WrapError("Encode rsrp_Result", err)
		}
	}
	return nil
}

func (ie *MeasResultCellSFTD_NR) Decode(r *uper.UperReader) error {
	var err error
	var rsrp_ResultPresent bool
	if rsrp_ResultPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.physCellId.Decode(r); err != nil {
		return utils.WrapError("Decode physCellId", err)
	}
	var tmp_int_sfn_OffsetResult int64
	if tmp_int_sfn_OffsetResult, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 1023}, false); err != nil {
		return utils.WrapError("ReadInteger sfn_OffsetResult", err)
	}
	ie.sfn_OffsetResult = tmp_int_sfn_OffsetResult
	var tmp_int_frameBoundaryOffsetResult int64
	if tmp_int_frameBoundaryOffsetResult, err = r.ReadInteger(&uper.Constraint{Lb: -30720, Ub: 30719}, false); err != nil {
		return utils.WrapError("ReadInteger frameBoundaryOffsetResult", err)
	}
	ie.frameBoundaryOffsetResult = tmp_int_frameBoundaryOffsetResult
	if rsrp_ResultPresent {
		ie.rsrp_Result = new(RSRP_Range)
		if err = ie.rsrp_Result.Decode(r); err != nil {
			return utils.WrapError("Decode rsrp_Result", err)
		}
	}
	return nil
}
