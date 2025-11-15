package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SuccessHO_Report_r17_sourceCellInfo_r17 struct {
	sourcePCellId_r17    CGI_Info_Logging_r16                                          `madatory`
	sourceCellMeas_r17   *MeasResultSuccessHONR_r17                                    `optional`
	rlf_InSourceDAPS_r17 *SuccessHO_Report_r17_sourceCellInfo_r17_rlf_InSourceDAPS_r17 `optional`
}

func (ie *SuccessHO_Report_r17_sourceCellInfo_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.sourceCellMeas_r17 != nil, ie.rlf_InSourceDAPS_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.sourcePCellId_r17.Encode(w); err != nil {
		return utils.WrapError("Encode sourcePCellId_r17", err)
	}
	if ie.sourceCellMeas_r17 != nil {
		if err = ie.sourceCellMeas_r17.Encode(w); err != nil {
			return utils.WrapError("Encode sourceCellMeas_r17", err)
		}
	}
	if ie.rlf_InSourceDAPS_r17 != nil {
		if err = ie.rlf_InSourceDAPS_r17.Encode(w); err != nil {
			return utils.WrapError("Encode rlf_InSourceDAPS_r17", err)
		}
	}
	return nil
}

func (ie *SuccessHO_Report_r17_sourceCellInfo_r17) Decode(r *uper.UperReader) error {
	var err error
	var sourceCellMeas_r17Present bool
	if sourceCellMeas_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var rlf_InSourceDAPS_r17Present bool
	if rlf_InSourceDAPS_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.sourcePCellId_r17.Decode(r); err != nil {
		return utils.WrapError("Decode sourcePCellId_r17", err)
	}
	if sourceCellMeas_r17Present {
		ie.sourceCellMeas_r17 = new(MeasResultSuccessHONR_r17)
		if err = ie.sourceCellMeas_r17.Decode(r); err != nil {
			return utils.WrapError("Decode sourceCellMeas_r17", err)
		}
	}
	if rlf_InSourceDAPS_r17Present {
		ie.rlf_InSourceDAPS_r17 = new(SuccessHO_Report_r17_sourceCellInfo_r17_rlf_InSourceDAPS_r17)
		if err = ie.rlf_InSourceDAPS_r17.Decode(r); err != nil {
			return utils.WrapError("Decode rlf_InSourceDAPS_r17", err)
		}
	}
	return nil
}
