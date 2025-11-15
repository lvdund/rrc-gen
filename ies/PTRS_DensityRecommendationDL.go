package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PTRS_DensityRecommendationDL struct {
	frequencyDensity1 int64 `lb:1,ub:276,madatory`
	frequencyDensity2 int64 `lb:1,ub:276,madatory`
	timeDensity1      int64 `lb:0,ub:29,madatory`
	timeDensity2      int64 `lb:0,ub:29,madatory`
	timeDensity3      int64 `lb:0,ub:29,madatory`
}

func (ie *PTRS_DensityRecommendationDL) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(ie.frequencyDensity1, &uper.Constraint{Lb: 1, Ub: 276}, false); err != nil {
		return utils.WrapError("WriteInteger frequencyDensity1", err)
	}
	if err = w.WriteInteger(ie.frequencyDensity2, &uper.Constraint{Lb: 1, Ub: 276}, false); err != nil {
		return utils.WrapError("WriteInteger frequencyDensity2", err)
	}
	if err = w.WriteInteger(ie.timeDensity1, &uper.Constraint{Lb: 0, Ub: 29}, false); err != nil {
		return utils.WrapError("WriteInteger timeDensity1", err)
	}
	if err = w.WriteInteger(ie.timeDensity2, &uper.Constraint{Lb: 0, Ub: 29}, false); err != nil {
		return utils.WrapError("WriteInteger timeDensity2", err)
	}
	if err = w.WriteInteger(ie.timeDensity3, &uper.Constraint{Lb: 0, Ub: 29}, false); err != nil {
		return utils.WrapError("WriteInteger timeDensity3", err)
	}
	return nil
}

func (ie *PTRS_DensityRecommendationDL) Decode(r *uper.UperReader) error {
	var err error
	var tmp_int_frequencyDensity1 int64
	if tmp_int_frequencyDensity1, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 276}, false); err != nil {
		return utils.WrapError("ReadInteger frequencyDensity1", err)
	}
	ie.frequencyDensity1 = tmp_int_frequencyDensity1
	var tmp_int_frequencyDensity2 int64
	if tmp_int_frequencyDensity2, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 276}, false); err != nil {
		return utils.WrapError("ReadInteger frequencyDensity2", err)
	}
	ie.frequencyDensity2 = tmp_int_frequencyDensity2
	var tmp_int_timeDensity1 int64
	if tmp_int_timeDensity1, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 29}, false); err != nil {
		return utils.WrapError("ReadInteger timeDensity1", err)
	}
	ie.timeDensity1 = tmp_int_timeDensity1
	var tmp_int_timeDensity2 int64
	if tmp_int_timeDensity2, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 29}, false); err != nil {
		return utils.WrapError("ReadInteger timeDensity2", err)
	}
	ie.timeDensity2 = tmp_int_timeDensity2
	var tmp_int_timeDensity3 int64
	if tmp_int_timeDensity3, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 29}, false); err != nil {
		return utils.WrapError("ReadInteger timeDensity3", err)
	}
	ie.timeDensity3 = tmp_int_timeDensity3
	return nil
}
