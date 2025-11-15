package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type NR_PRS_MeasurementInfoList_r16 struct {
	Value []NR_PRS_MeasurementInfo_r16 `lb:1,ub:maxFreqLayers,madatory`
}

func (ie *NR_PRS_MeasurementInfoList_r16) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*NR_PRS_MeasurementInfo_r16]([]*NR_PRS_MeasurementInfo_r16{}, uper.Constraint{Lb: 1, Ub: maxFreqLayers}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode NR_PRS_MeasurementInfoList_r16", err)
	}
	return nil
}

func (ie *NR_PRS_MeasurementInfoList_r16) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*NR_PRS_MeasurementInfo_r16]([]*NR_PRS_MeasurementInfo_r16{}, uper.Constraint{Lb: 1, Ub: maxFreqLayers}, false)
	fn := func() *NR_PRS_MeasurementInfo_r16 {
		return new(NR_PRS_MeasurementInfo_r16)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode NR_PRS_MeasurementInfoList_r16", err)
	}
	ie.Value = []NR_PRS_MeasurementInfo_r16{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
