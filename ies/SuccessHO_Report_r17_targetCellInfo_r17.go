package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SuccessHO_Report_r17_targetCellInfo_r17 struct {
	targetPCellId_r17  CGI_Info_Logging_r16       `madatory`
	targetCellMeas_r17 *MeasResultSuccessHONR_r17 `optional`
}

func (ie *SuccessHO_Report_r17_targetCellInfo_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.targetCellMeas_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.targetPCellId_r17.Encode(w); err != nil {
		return utils.WrapError("Encode targetPCellId_r17", err)
	}
	if ie.targetCellMeas_r17 != nil {
		if err = ie.targetCellMeas_r17.Encode(w); err != nil {
			return utils.WrapError("Encode targetCellMeas_r17", err)
		}
	}
	return nil
}

func (ie *SuccessHO_Report_r17_targetCellInfo_r17) Decode(r *uper.UperReader) error {
	var err error
	var targetCellMeas_r17Present bool
	if targetCellMeas_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.targetPCellId_r17.Decode(r); err != nil {
		return utils.WrapError("Decode targetPCellId_r17", err)
	}
	if targetCellMeas_r17Present {
		ie.targetCellMeas_r17 = new(MeasResultSuccessHONR_r17)
		if err = ie.targetCellMeas_r17.Decode(r); err != nil {
			return utils.WrapError("Decode targetCellMeas_r17", err)
		}
	}
	return nil
}
