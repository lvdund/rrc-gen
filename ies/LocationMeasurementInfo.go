package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	LocationMeasurementInfo_Choice_nothing uint64 = iota
	LocationMeasurementInfo_Choice_eutra_RSTD
	LocationMeasurementInfo_Choice_eutra_FineTimingDetection
	LocationMeasurementInfo_Choice_nr_PRS_Measurement_r16
)

type LocationMeasurementInfo struct {
	Choice                    uint64
	eutra_RSTD                *EUTRA_RSTD_InfoList
	eutra_FineTimingDetection uper.NULL `madatory,ext`
	nr_PRS_Measurement_r16    *NR_PRS_MeasurementInfoList_r16
}

func (ie *LocationMeasurementInfo) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 3, false); err != nil {
		return err
	}
	switch ie.Choice {
	case LocationMeasurementInfo_Choice_eutra_RSTD:
		if err = ie.eutra_RSTD.Encode(w); err != nil {
			err = utils.WrapError("Encode eutra_RSTD", err)
		}
	case LocationMeasurementInfo_Choice_eutra_FineTimingDetection:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode eutra_FineTimingDetection", err)
		}
	case LocationMeasurementInfo_Choice_nr_PRS_Measurement_r16:
		if err = ie.nr_PRS_Measurement_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode nr_PRS_Measurement_r16", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *LocationMeasurementInfo) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(3, false); err != nil {
		return err
	}
	switch ie.Choice {
	case LocationMeasurementInfo_Choice_eutra_RSTD:
		ie.eutra_RSTD = new(EUTRA_RSTD_InfoList)
		if err = ie.eutra_RSTD.Decode(r); err != nil {
			return utils.WrapError("Decode eutra_RSTD", err)
		}
	case LocationMeasurementInfo_Choice_eutra_FineTimingDetection:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode eutra_FineTimingDetection", err)
		}
	case LocationMeasurementInfo_Choice_nr_PRS_Measurement_r16:
		ie.nr_PRS_Measurement_r16 = new(NR_PRS_MeasurementInfoList_r16)
		if err = ie.nr_PRS_Measurement_r16.Decode(r); err != nil {
			return utils.WrapError("Decode nr_PRS_Measurement_r16", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
