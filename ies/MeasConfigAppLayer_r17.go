package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasConfigAppLayer_r17 struct {
	measConfigAppLayerId_r17           MeasConfigAppLayerId_r17                `madatory`
	measConfigAppLayerContainer_r17    *[]byte                                 `lb:1,ub:8000,optional`
	serviceType_r17                    *MeasConfigAppLayer_r17_serviceType_r17 `optional`
	pauseReporting_r17                 *bool                                   `optional`
	transmissionOfSessionStartStop_r17 *bool                                   `optional`
	ran_VisibleParameters_r17          *RAN_VisibleParameters_r17              `optional,setuprelease`
}

func (ie *MeasConfigAppLayer_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.measConfigAppLayerContainer_r17 != nil, ie.serviceType_r17 != nil, ie.pauseReporting_r17 != nil, ie.transmissionOfSessionStartStop_r17 != nil, ie.ran_VisibleParameters_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.measConfigAppLayerId_r17.Encode(w); err != nil {
		return utils.WrapError("Encode measConfigAppLayerId_r17", err)
	}
	if ie.measConfigAppLayerContainer_r17 != nil {
		if err = w.WriteOctetString(*ie.measConfigAppLayerContainer_r17, &uper.Constraint{Lb: 1, Ub: 8000}, false); err != nil {
			return utils.WrapError("Encode measConfigAppLayerContainer_r17", err)
		}
	}
	if ie.serviceType_r17 != nil {
		if err = ie.serviceType_r17.Encode(w); err != nil {
			return utils.WrapError("Encode serviceType_r17", err)
		}
	}
	if ie.pauseReporting_r17 != nil {
		if err = w.WriteBoolean(*ie.pauseReporting_r17); err != nil {
			return utils.WrapError("Encode pauseReporting_r17", err)
		}
	}
	if ie.transmissionOfSessionStartStop_r17 != nil {
		if err = w.WriteBoolean(*ie.transmissionOfSessionStartStop_r17); err != nil {
			return utils.WrapError("Encode transmissionOfSessionStartStop_r17", err)
		}
	}
	if ie.ran_VisibleParameters_r17 != nil {
		tmp_ran_VisibleParameters_r17 := utils.SetupRelease[*RAN_VisibleParameters_r17]{
			Setup: ie.ran_VisibleParameters_r17,
		}
		if err = tmp_ran_VisibleParameters_r17.Encode(w); err != nil {
			return utils.WrapError("Encode ran_VisibleParameters_r17", err)
		}
	}
	return nil
}

func (ie *MeasConfigAppLayer_r17) Decode(r *uper.UperReader) error {
	var err error
	var measConfigAppLayerContainer_r17Present bool
	if measConfigAppLayerContainer_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var serviceType_r17Present bool
	if serviceType_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var pauseReporting_r17Present bool
	if pauseReporting_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var transmissionOfSessionStartStop_r17Present bool
	if transmissionOfSessionStartStop_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ran_VisibleParameters_r17Present bool
	if ran_VisibleParameters_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.measConfigAppLayerId_r17.Decode(r); err != nil {
		return utils.WrapError("Decode measConfigAppLayerId_r17", err)
	}
	if measConfigAppLayerContainer_r17Present {
		var tmp_os_measConfigAppLayerContainer_r17 []byte
		if tmp_os_measConfigAppLayerContainer_r17, err = r.ReadOctetString(&uper.Constraint{Lb: 1, Ub: 8000}, false); err != nil {
			return utils.WrapError("Decode measConfigAppLayerContainer_r17", err)
		}
		ie.measConfigAppLayerContainer_r17 = &tmp_os_measConfigAppLayerContainer_r17
	}
	if serviceType_r17Present {
		ie.serviceType_r17 = new(MeasConfigAppLayer_r17_serviceType_r17)
		if err = ie.serviceType_r17.Decode(r); err != nil {
			return utils.WrapError("Decode serviceType_r17", err)
		}
	}
	if pauseReporting_r17Present {
		var tmp_bool_pauseReporting_r17 bool
		if tmp_bool_pauseReporting_r17, err = r.ReadBoolean(); err != nil {
			return utils.WrapError("Decode pauseReporting_r17", err)
		}
		ie.pauseReporting_r17 = &tmp_bool_pauseReporting_r17
	}
	if transmissionOfSessionStartStop_r17Present {
		var tmp_bool_transmissionOfSessionStartStop_r17 bool
		if tmp_bool_transmissionOfSessionStartStop_r17, err = r.ReadBoolean(); err != nil {
			return utils.WrapError("Decode transmissionOfSessionStartStop_r17", err)
		}
		ie.transmissionOfSessionStartStop_r17 = &tmp_bool_transmissionOfSessionStartStop_r17
	}
	if ran_VisibleParameters_r17Present {
		tmp_ran_VisibleParameters_r17 := utils.SetupRelease[*RAN_VisibleParameters_r17]{}
		if err = tmp_ran_VisibleParameters_r17.Decode(r); err != nil {
			return utils.WrapError("Decode ran_VisibleParameters_r17", err)
		}
		ie.ran_VisibleParameters_r17 = tmp_ran_VisibleParameters_r17.Setup
	}
	return nil
}
