package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasResultsPerCellIdleEUTRA_r16_measIdleResultEUTRA_r16 struct {
	rsrp_ResultEUTRA_r16 *RSRP_RangeEUTRA     `optional`
	rsrq_ResultEUTRA_r16 *RSRQ_RangeEUTRA_r16 `optional`
}

func (ie *MeasResultsPerCellIdleEUTRA_r16_measIdleResultEUTRA_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.rsrp_ResultEUTRA_r16 != nil, ie.rsrq_ResultEUTRA_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.rsrp_ResultEUTRA_r16 != nil {
		if err = ie.rsrp_ResultEUTRA_r16.Encode(w); err != nil {
			return utils.WrapError("Encode rsrp_ResultEUTRA_r16", err)
		}
	}
	if ie.rsrq_ResultEUTRA_r16 != nil {
		if err = ie.rsrq_ResultEUTRA_r16.Encode(w); err != nil {
			return utils.WrapError("Encode rsrq_ResultEUTRA_r16", err)
		}
	}
	return nil
}

func (ie *MeasResultsPerCellIdleEUTRA_r16_measIdleResultEUTRA_r16) Decode(r *uper.UperReader) error {
	var err error
	var rsrp_ResultEUTRA_r16Present bool
	if rsrp_ResultEUTRA_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var rsrq_ResultEUTRA_r16Present bool
	if rsrq_ResultEUTRA_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if rsrp_ResultEUTRA_r16Present {
		ie.rsrp_ResultEUTRA_r16 = new(RSRP_RangeEUTRA)
		if err = ie.rsrp_ResultEUTRA_r16.Decode(r); err != nil {
			return utils.WrapError("Decode rsrp_ResultEUTRA_r16", err)
		}
	}
	if rsrq_ResultEUTRA_r16Present {
		ie.rsrq_ResultEUTRA_r16 = new(RSRQ_RangeEUTRA_r16)
		if err = ie.rsrq_ResultEUTRA_r16.Decode(r); err != nil {
			return utils.WrapError("Decode rsrq_ResultEUTRA_r16", err)
		}
	}
	return nil
}
