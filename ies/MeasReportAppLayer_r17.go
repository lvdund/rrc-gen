package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasReportAppLayer_r17 struct {
	measConfigAppLayerId_r17        MeasConfigAppLayerId_r17                          `madatory`
	measReportAppLayerContainer_r17 *[]byte                                           `optional`
	appLayerSessionStatus_r17       *MeasReportAppLayer_r17_appLayerSessionStatus_r17 `optional`
	ran_VisibleMeasurements_r17     *RAN_VisibleMeasurements_r17                      `optional`
}

func (ie *MeasReportAppLayer_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.measReportAppLayerContainer_r17 != nil, ie.appLayerSessionStatus_r17 != nil, ie.ran_VisibleMeasurements_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.measConfigAppLayerId_r17.Encode(w); err != nil {
		return utils.WrapError("Encode measConfigAppLayerId_r17", err)
	}
	if ie.measReportAppLayerContainer_r17 != nil {
		if err = w.WriteOctetString(*ie.measReportAppLayerContainer_r17, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode measReportAppLayerContainer_r17", err)
		}
	}
	if ie.appLayerSessionStatus_r17 != nil {
		if err = ie.appLayerSessionStatus_r17.Encode(w); err != nil {
			return utils.WrapError("Encode appLayerSessionStatus_r17", err)
		}
	}
	if ie.ran_VisibleMeasurements_r17 != nil {
		if err = ie.ran_VisibleMeasurements_r17.Encode(w); err != nil {
			return utils.WrapError("Encode ran_VisibleMeasurements_r17", err)
		}
	}
	return nil
}

func (ie *MeasReportAppLayer_r17) Decode(r *uper.UperReader) error {
	var err error
	var measReportAppLayerContainer_r17Present bool
	if measReportAppLayerContainer_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var appLayerSessionStatus_r17Present bool
	if appLayerSessionStatus_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ran_VisibleMeasurements_r17Present bool
	if ran_VisibleMeasurements_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.measConfigAppLayerId_r17.Decode(r); err != nil {
		return utils.WrapError("Decode measConfigAppLayerId_r17", err)
	}
	if measReportAppLayerContainer_r17Present {
		var tmp_os_measReportAppLayerContainer_r17 []byte
		if tmp_os_measReportAppLayerContainer_r17, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode measReportAppLayerContainer_r17", err)
		}
		ie.measReportAppLayerContainer_r17 = &tmp_os_measReportAppLayerContainer_r17
	}
	if appLayerSessionStatus_r17Present {
		ie.appLayerSessionStatus_r17 = new(MeasReportAppLayer_r17_appLayerSessionStatus_r17)
		if err = ie.appLayerSessionStatus_r17.Decode(r); err != nil {
			return utils.WrapError("Decode appLayerSessionStatus_r17", err)
		}
	}
	if ran_VisibleMeasurements_r17Present {
		ie.ran_VisibleMeasurements_r17 = new(RAN_VisibleMeasurements_r17)
		if err = ie.ran_VisibleMeasurements_r17.Decode(r); err != nil {
			return utils.WrapError("Decode ran_VisibleMeasurements_r17", err)
		}
	}
	return nil
}
