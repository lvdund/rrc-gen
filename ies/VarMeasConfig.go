package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type VarMeasConfig struct {
	measIdList       *MeasIdToAddModList            `optional`
	measObjectList   *MeasObjectToAddModList        `optional`
	reportConfigList *ReportConfigToAddModList      `optional`
	quantityConfig   *QuantityConfig                `optional`
	s_MeasureConfig  *VarMeasConfig_s_MeasureConfig `optional`
}

func (ie *VarMeasConfig) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.measIdList != nil, ie.measObjectList != nil, ie.reportConfigList != nil, ie.quantityConfig != nil, ie.s_MeasureConfig != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.measIdList != nil {
		if err = ie.measIdList.Encode(w); err != nil {
			return utils.WrapError("Encode measIdList", err)
		}
	}
	if ie.measObjectList != nil {
		if err = ie.measObjectList.Encode(w); err != nil {
			return utils.WrapError("Encode measObjectList", err)
		}
	}
	if ie.reportConfigList != nil {
		if err = ie.reportConfigList.Encode(w); err != nil {
			return utils.WrapError("Encode reportConfigList", err)
		}
	}
	if ie.quantityConfig != nil {
		if err = ie.quantityConfig.Encode(w); err != nil {
			return utils.WrapError("Encode quantityConfig", err)
		}
	}
	if ie.s_MeasureConfig != nil {
		if err = ie.s_MeasureConfig.Encode(w); err != nil {
			return utils.WrapError("Encode s_MeasureConfig", err)
		}
	}
	return nil
}

func (ie *VarMeasConfig) Decode(r *uper.UperReader) error {
	var err error
	var measIdListPresent bool
	if measIdListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var measObjectListPresent bool
	if measObjectListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var reportConfigListPresent bool
	if reportConfigListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var quantityConfigPresent bool
	if quantityConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var s_MeasureConfigPresent bool
	if s_MeasureConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if measIdListPresent {
		ie.measIdList = new(MeasIdToAddModList)
		if err = ie.measIdList.Decode(r); err != nil {
			return utils.WrapError("Decode measIdList", err)
		}
	}
	if measObjectListPresent {
		ie.measObjectList = new(MeasObjectToAddModList)
		if err = ie.measObjectList.Decode(r); err != nil {
			return utils.WrapError("Decode measObjectList", err)
		}
	}
	if reportConfigListPresent {
		ie.reportConfigList = new(ReportConfigToAddModList)
		if err = ie.reportConfigList.Decode(r); err != nil {
			return utils.WrapError("Decode reportConfigList", err)
		}
	}
	if quantityConfigPresent {
		ie.quantityConfig = new(QuantityConfig)
		if err = ie.quantityConfig.Decode(r); err != nil {
			return utils.WrapError("Decode quantityConfig", err)
		}
	}
	if s_MeasureConfigPresent {
		ie.s_MeasureConfig = new(VarMeasConfig_s_MeasureConfig)
		if err = ie.s_MeasureConfig.Decode(r); err != nil {
			return utils.WrapError("Decode s_MeasureConfig", err)
		}
	}
	return nil
}
