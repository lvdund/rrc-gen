package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasurementReportAppLayerList_r17 struct {
	Value []MeasReportAppLayer_r17 `lb:1,ub:maxNrofAppLayerMeas_r17,madatory`
}

func (ie *MeasurementReportAppLayerList_r17) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*MeasReportAppLayer_r17]([]*MeasReportAppLayer_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofAppLayerMeas_r17}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode MeasurementReportAppLayerList_r17", err)
	}
	return nil
}

func (ie *MeasurementReportAppLayerList_r17) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*MeasReportAppLayer_r17]([]*MeasReportAppLayer_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofAppLayerMeas_r17}, false)
	fn := func() *MeasReportAppLayer_r17 {
		return new(MeasReportAppLayer_r17)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode MeasurementReportAppLayerList_r17", err)
	}
	ie.Value = []MeasReportAppLayer_r17{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
