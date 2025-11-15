package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type VarMeasIdleReport_r16 struct {
	measReportIdleNR_r16    *MeasResultIdleNR_r16    `optional`
	measReportIdleEUTRA_r16 *MeasResultIdleEUTRA_r16 `optional`
}

func (ie *VarMeasIdleReport_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.measReportIdleNR_r16 != nil, ie.measReportIdleEUTRA_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.measReportIdleNR_r16 != nil {
		if err = ie.measReportIdleNR_r16.Encode(w); err != nil {
			return utils.WrapError("Encode measReportIdleNR_r16", err)
		}
	}
	if ie.measReportIdleEUTRA_r16 != nil {
		if err = ie.measReportIdleEUTRA_r16.Encode(w); err != nil {
			return utils.WrapError("Encode measReportIdleEUTRA_r16", err)
		}
	}
	return nil
}

func (ie *VarMeasIdleReport_r16) Decode(r *uper.UperReader) error {
	var err error
	var measReportIdleNR_r16Present bool
	if measReportIdleNR_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var measReportIdleEUTRA_r16Present bool
	if measReportIdleEUTRA_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if measReportIdleNR_r16Present {
		ie.measReportIdleNR_r16 = new(MeasResultIdleNR_r16)
		if err = ie.measReportIdleNR_r16.Decode(r); err != nil {
			return utils.WrapError("Decode measReportIdleNR_r16", err)
		}
	}
	if measReportIdleEUTRA_r16Present {
		ie.measReportIdleEUTRA_r16 = new(MeasResultIdleEUTRA_r16)
		if err = ie.measReportIdleEUTRA_r16.Decode(r); err != nil {
			return utils.WrapError("Decode measReportIdleEUTRA_r16", err)
		}
	}
	return nil
}
