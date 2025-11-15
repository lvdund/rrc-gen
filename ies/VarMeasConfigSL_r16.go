package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type VarMeasConfigSL_r16 struct {
	sl_MeasIdList_r16       *SL_MeasIdList_r16       `optional`
	sl_MeasObjectList_r16   *SL_MeasObjectList_r16   `optional`
	sl_reportConfigList_r16 *SL_ReportConfigList_r16 `optional`
	sl_QuantityConfig_r16   *SL_QuantityConfig_r16   `optional`
}

func (ie *VarMeasConfigSL_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.sl_MeasIdList_r16 != nil, ie.sl_MeasObjectList_r16 != nil, ie.sl_reportConfigList_r16 != nil, ie.sl_QuantityConfig_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.sl_MeasIdList_r16 != nil {
		if err = ie.sl_MeasIdList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_MeasIdList_r16", err)
		}
	}
	if ie.sl_MeasObjectList_r16 != nil {
		if err = ie.sl_MeasObjectList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_MeasObjectList_r16", err)
		}
	}
	if ie.sl_reportConfigList_r16 != nil {
		if err = ie.sl_reportConfigList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_reportConfigList_r16", err)
		}
	}
	if ie.sl_QuantityConfig_r16 != nil {
		if err = ie.sl_QuantityConfig_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_QuantityConfig_r16", err)
		}
	}
	return nil
}

func (ie *VarMeasConfigSL_r16) Decode(r *uper.UperReader) error {
	var err error
	var sl_MeasIdList_r16Present bool
	if sl_MeasIdList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_MeasObjectList_r16Present bool
	if sl_MeasObjectList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_reportConfigList_r16Present bool
	if sl_reportConfigList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_QuantityConfig_r16Present bool
	if sl_QuantityConfig_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if sl_MeasIdList_r16Present {
		ie.sl_MeasIdList_r16 = new(SL_MeasIdList_r16)
		if err = ie.sl_MeasIdList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_MeasIdList_r16", err)
		}
	}
	if sl_MeasObjectList_r16Present {
		ie.sl_MeasObjectList_r16 = new(SL_MeasObjectList_r16)
		if err = ie.sl_MeasObjectList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_MeasObjectList_r16", err)
		}
	}
	if sl_reportConfigList_r16Present {
		ie.sl_reportConfigList_r16 = new(SL_ReportConfigList_r16)
		if err = ie.sl_reportConfigList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_reportConfigList_r16", err)
		}
	}
	if sl_QuantityConfig_r16Present {
		ie.sl_QuantityConfig_r16 = new(SL_QuantityConfig_r16)
		if err = ie.sl_QuantityConfig_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_QuantityConfig_r16", err)
		}
	}
	return nil
}
