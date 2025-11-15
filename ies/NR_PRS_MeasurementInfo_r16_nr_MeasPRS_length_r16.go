package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	NR_PRS_MeasurementInfo_r16_nr_MeasPRS_length_r16_Enum_ms1dot5 uper.Enumerated = 0
	NR_PRS_MeasurementInfo_r16_nr_MeasPRS_length_r16_Enum_ms3     uper.Enumerated = 1
	NR_PRS_MeasurementInfo_r16_nr_MeasPRS_length_r16_Enum_ms3dot5 uper.Enumerated = 2
	NR_PRS_MeasurementInfo_r16_nr_MeasPRS_length_r16_Enum_ms4     uper.Enumerated = 3
	NR_PRS_MeasurementInfo_r16_nr_MeasPRS_length_r16_Enum_ms5dot5 uper.Enumerated = 4
	NR_PRS_MeasurementInfo_r16_nr_MeasPRS_length_r16_Enum_ms6     uper.Enumerated = 5
	NR_PRS_MeasurementInfo_r16_nr_MeasPRS_length_r16_Enum_ms10    uper.Enumerated = 6
	NR_PRS_MeasurementInfo_r16_nr_MeasPRS_length_r16_Enum_ms20    uper.Enumerated = 7
)

type NR_PRS_MeasurementInfo_r16_nr_MeasPRS_length_r16 struct {
	Value uper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *NR_PRS_MeasurementInfo_r16_nr_MeasPRS_length_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode NR_PRS_MeasurementInfo_r16_nr_MeasPRS_length_r16", err)
	}
	return nil
}

func (ie *NR_PRS_MeasurementInfo_r16_nr_MeasPRS_length_r16) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode NR_PRS_MeasurementInfo_r16_nr_MeasPRS_length_r16", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
