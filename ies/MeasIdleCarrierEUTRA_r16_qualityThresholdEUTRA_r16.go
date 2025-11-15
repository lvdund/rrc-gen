package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasIdleCarrierEUTRA_r16_qualityThresholdEUTRA_r16 struct {
	idleRSRP_Threshold_EUTRA_r16 *RSRP_RangeEUTRA     `optional`
	idleRSRQ_Threshold_EUTRA_r16 *RSRQ_RangeEUTRA_r16 `optional`
}

func (ie *MeasIdleCarrierEUTRA_r16_qualityThresholdEUTRA_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.idleRSRP_Threshold_EUTRA_r16 != nil, ie.idleRSRQ_Threshold_EUTRA_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.idleRSRP_Threshold_EUTRA_r16 != nil {
		if err = ie.idleRSRP_Threshold_EUTRA_r16.Encode(w); err != nil {
			return utils.WrapError("Encode idleRSRP_Threshold_EUTRA_r16", err)
		}
	}
	if ie.idleRSRQ_Threshold_EUTRA_r16 != nil {
		if err = ie.idleRSRQ_Threshold_EUTRA_r16.Encode(w); err != nil {
			return utils.WrapError("Encode idleRSRQ_Threshold_EUTRA_r16", err)
		}
	}
	return nil
}

func (ie *MeasIdleCarrierEUTRA_r16_qualityThresholdEUTRA_r16) Decode(r *uper.UperReader) error {
	var err error
	var idleRSRP_Threshold_EUTRA_r16Present bool
	if idleRSRP_Threshold_EUTRA_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var idleRSRQ_Threshold_EUTRA_r16Present bool
	if idleRSRQ_Threshold_EUTRA_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if idleRSRP_Threshold_EUTRA_r16Present {
		ie.idleRSRP_Threshold_EUTRA_r16 = new(RSRP_RangeEUTRA)
		if err = ie.idleRSRP_Threshold_EUTRA_r16.Decode(r); err != nil {
			return utils.WrapError("Decode idleRSRP_Threshold_EUTRA_r16", err)
		}
	}
	if idleRSRQ_Threshold_EUTRA_r16Present {
		ie.idleRSRQ_Threshold_EUTRA_r16 = new(RSRQ_RangeEUTRA_r16)
		if err = ie.idleRSRQ_Threshold_EUTRA_r16.Decode(r); err != nil {
			return utils.WrapError("Decode idleRSRQ_Threshold_EUTRA_r16", err)
		}
	}
	return nil
}
