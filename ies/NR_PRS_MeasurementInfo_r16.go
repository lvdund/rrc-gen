package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type NR_PRS_MeasurementInfo_r16 struct {
	dl_PRS_PointA_r16                  ARFCN_ValueNR                                                 `madatory`
	nr_MeasPRS_RepetitionAndOffset_r16 NR_PRS_MeasurementInfo_r16_nr_MeasPRS_RepetitionAndOffset_r16 `lb:0,ub:19,madatory`
	nr_MeasPRS_length_r16              NR_PRS_MeasurementInfo_r16_nr_MeasPRS_length_r16              `madatory,ext`
}

func (ie *NR_PRS_MeasurementInfo_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.dl_PRS_PointA_r16.Encode(w); err != nil {
		return utils.WrapError("Encode dl_PRS_PointA_r16", err)
	}
	if err = ie.nr_MeasPRS_RepetitionAndOffset_r16.Encode(w); err != nil {
		return utils.WrapError("Encode nr_MeasPRS_RepetitionAndOffset_r16", err)
	}
	return nil
}

func (ie *NR_PRS_MeasurementInfo_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.dl_PRS_PointA_r16.Decode(r); err != nil {
		return utils.WrapError("Decode dl_PRS_PointA_r16", err)
	}
	if err = ie.nr_MeasPRS_RepetitionAndOffset_r16.Decode(r); err != nil {
		return utils.WrapError("Decode nr_MeasPRS_RepetitionAndOffset_r16", err)
	}
	return nil
}
