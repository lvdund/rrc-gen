package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BeamMeasConfigIdle_NR_r16 struct {
	reportQuantityRS_Indexes_r16  BeamMeasConfigIdle_NR_r16_reportQuantityRS_Indexes_r16 `madatory`
	maxNrofRS_IndexesToReport_r16 int64                                                  `lb:1,ub:maxNrofIndexesToReport,madatory`
	includeBeamMeasurements_r16   bool                                                   `madatory`
}

func (ie *BeamMeasConfigIdle_NR_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.reportQuantityRS_Indexes_r16.Encode(w); err != nil {
		return utils.WrapError("Encode reportQuantityRS_Indexes_r16", err)
	}
	if err = w.WriteInteger(ie.maxNrofRS_IndexesToReport_r16, &uper.Constraint{Lb: 1, Ub: maxNrofIndexesToReport}, false); err != nil {
		return utils.WrapError("WriteInteger maxNrofRS_IndexesToReport_r16", err)
	}
	if err = w.WriteBoolean(ie.includeBeamMeasurements_r16); err != nil {
		return utils.WrapError("WriteBoolean includeBeamMeasurements_r16", err)
	}
	return nil
}

func (ie *BeamMeasConfigIdle_NR_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.reportQuantityRS_Indexes_r16.Decode(r); err != nil {
		return utils.WrapError("Decode reportQuantityRS_Indexes_r16", err)
	}
	var tmp_int_maxNrofRS_IndexesToReport_r16 int64
	if tmp_int_maxNrofRS_IndexesToReport_r16, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: maxNrofIndexesToReport}, false); err != nil {
		return utils.WrapError("ReadInteger maxNrofRS_IndexesToReport_r16", err)
	}
	ie.maxNrofRS_IndexesToReport_r16 = tmp_int_maxNrofRS_IndexesToReport_r16
	var tmp_bool_includeBeamMeasurements_r16 bool
	if tmp_bool_includeBeamMeasurements_r16, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean includeBeamMeasurements_r16", err)
	}
	ie.includeBeamMeasurements_r16 = tmp_bool_includeBeamMeasurements_r16
	return nil
}
