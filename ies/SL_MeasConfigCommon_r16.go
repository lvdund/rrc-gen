package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_MeasConfigCommon_r16 struct {
	sl_MeasObjectListCommon_r16   *SL_MeasObjectList_r16   `optional`
	sl_ReportConfigListCommon_r16 *SL_ReportConfigList_r16 `optional`
	sl_MeasIdListCommon_r16       *SL_MeasIdList_r16       `optional`
	sl_QuantityConfigCommon_r16   *SL_QuantityConfig_r16   `optional`
}

func (ie *SL_MeasConfigCommon_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.sl_MeasObjectListCommon_r16 != nil, ie.sl_ReportConfigListCommon_r16 != nil, ie.sl_MeasIdListCommon_r16 != nil, ie.sl_QuantityConfigCommon_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.sl_MeasObjectListCommon_r16 != nil {
		if err = ie.sl_MeasObjectListCommon_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_MeasObjectListCommon_r16", err)
		}
	}
	if ie.sl_ReportConfigListCommon_r16 != nil {
		if err = ie.sl_ReportConfigListCommon_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_ReportConfigListCommon_r16", err)
		}
	}
	if ie.sl_MeasIdListCommon_r16 != nil {
		if err = ie.sl_MeasIdListCommon_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_MeasIdListCommon_r16", err)
		}
	}
	if ie.sl_QuantityConfigCommon_r16 != nil {
		if err = ie.sl_QuantityConfigCommon_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_QuantityConfigCommon_r16", err)
		}
	}
	return nil
}

func (ie *SL_MeasConfigCommon_r16) Decode(r *uper.UperReader) error {
	var err error
	var sl_MeasObjectListCommon_r16Present bool
	if sl_MeasObjectListCommon_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_ReportConfigListCommon_r16Present bool
	if sl_ReportConfigListCommon_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_MeasIdListCommon_r16Present bool
	if sl_MeasIdListCommon_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_QuantityConfigCommon_r16Present bool
	if sl_QuantityConfigCommon_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if sl_MeasObjectListCommon_r16Present {
		ie.sl_MeasObjectListCommon_r16 = new(SL_MeasObjectList_r16)
		if err = ie.sl_MeasObjectListCommon_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_MeasObjectListCommon_r16", err)
		}
	}
	if sl_ReportConfigListCommon_r16Present {
		ie.sl_ReportConfigListCommon_r16 = new(SL_ReportConfigList_r16)
		if err = ie.sl_ReportConfigListCommon_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_ReportConfigListCommon_r16", err)
		}
	}
	if sl_MeasIdListCommon_r16Present {
		ie.sl_MeasIdListCommon_r16 = new(SL_MeasIdList_r16)
		if err = ie.sl_MeasIdListCommon_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_MeasIdListCommon_r16", err)
		}
	}
	if sl_QuantityConfigCommon_r16Present {
		ie.sl_QuantityConfigCommon_r16 = new(SL_QuantityConfig_r16)
		if err = ie.sl_QuantityConfigCommon_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_QuantityConfigCommon_r16", err)
		}
	}
	return nil
}
