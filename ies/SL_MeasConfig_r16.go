package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_MeasConfig_r16 struct {
	sl_MeasObjectToRemoveList_r16   *SL_MeasObjectToRemoveList_r16   `optional`
	sl_MeasObjectToAddModList_r16   *SL_MeasObjectList_r16           `optional`
	sl_ReportConfigToRemoveList_r16 *SL_ReportConfigToRemoveList_r16 `optional`
	sl_ReportConfigToAddModList_r16 *SL_ReportConfigList_r16         `optional`
	sl_MeasIdToRemoveList_r16       *SL_MeasIdToRemoveList_r16       `optional`
	sl_MeasIdToAddModList_r16       *SL_MeasIdList_r16               `optional`
	sl_QuantityConfig_r16           *SL_QuantityConfig_r16           `optional`
}

func (ie *SL_MeasConfig_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.sl_MeasObjectToRemoveList_r16 != nil, ie.sl_MeasObjectToAddModList_r16 != nil, ie.sl_ReportConfigToRemoveList_r16 != nil, ie.sl_ReportConfigToAddModList_r16 != nil, ie.sl_MeasIdToRemoveList_r16 != nil, ie.sl_MeasIdToAddModList_r16 != nil, ie.sl_QuantityConfig_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.sl_MeasObjectToRemoveList_r16 != nil {
		if err = ie.sl_MeasObjectToRemoveList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_MeasObjectToRemoveList_r16", err)
		}
	}
	if ie.sl_MeasObjectToAddModList_r16 != nil {
		if err = ie.sl_MeasObjectToAddModList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_MeasObjectToAddModList_r16", err)
		}
	}
	if ie.sl_ReportConfigToRemoveList_r16 != nil {
		if err = ie.sl_ReportConfigToRemoveList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_ReportConfigToRemoveList_r16", err)
		}
	}
	if ie.sl_ReportConfigToAddModList_r16 != nil {
		if err = ie.sl_ReportConfigToAddModList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_ReportConfigToAddModList_r16", err)
		}
	}
	if ie.sl_MeasIdToRemoveList_r16 != nil {
		if err = ie.sl_MeasIdToRemoveList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_MeasIdToRemoveList_r16", err)
		}
	}
	if ie.sl_MeasIdToAddModList_r16 != nil {
		if err = ie.sl_MeasIdToAddModList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_MeasIdToAddModList_r16", err)
		}
	}
	if ie.sl_QuantityConfig_r16 != nil {
		if err = ie.sl_QuantityConfig_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_QuantityConfig_r16", err)
		}
	}
	return nil
}

func (ie *SL_MeasConfig_r16) Decode(r *uper.UperReader) error {
	var err error
	var sl_MeasObjectToRemoveList_r16Present bool
	if sl_MeasObjectToRemoveList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_MeasObjectToAddModList_r16Present bool
	if sl_MeasObjectToAddModList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_ReportConfigToRemoveList_r16Present bool
	if sl_ReportConfigToRemoveList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_ReportConfigToAddModList_r16Present bool
	if sl_ReportConfigToAddModList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_MeasIdToRemoveList_r16Present bool
	if sl_MeasIdToRemoveList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_MeasIdToAddModList_r16Present bool
	if sl_MeasIdToAddModList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_QuantityConfig_r16Present bool
	if sl_QuantityConfig_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if sl_MeasObjectToRemoveList_r16Present {
		ie.sl_MeasObjectToRemoveList_r16 = new(SL_MeasObjectToRemoveList_r16)
		if err = ie.sl_MeasObjectToRemoveList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_MeasObjectToRemoveList_r16", err)
		}
	}
	if sl_MeasObjectToAddModList_r16Present {
		ie.sl_MeasObjectToAddModList_r16 = new(SL_MeasObjectList_r16)
		if err = ie.sl_MeasObjectToAddModList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_MeasObjectToAddModList_r16", err)
		}
	}
	if sl_ReportConfigToRemoveList_r16Present {
		ie.sl_ReportConfigToRemoveList_r16 = new(SL_ReportConfigToRemoveList_r16)
		if err = ie.sl_ReportConfigToRemoveList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_ReportConfigToRemoveList_r16", err)
		}
	}
	if sl_ReportConfigToAddModList_r16Present {
		ie.sl_ReportConfigToAddModList_r16 = new(SL_ReportConfigList_r16)
		if err = ie.sl_ReportConfigToAddModList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_ReportConfigToAddModList_r16", err)
		}
	}
	if sl_MeasIdToRemoveList_r16Present {
		ie.sl_MeasIdToRemoveList_r16 = new(SL_MeasIdToRemoveList_r16)
		if err = ie.sl_MeasIdToRemoveList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_MeasIdToRemoveList_r16", err)
		}
	}
	if sl_MeasIdToAddModList_r16Present {
		ie.sl_MeasIdToAddModList_r16 = new(SL_MeasIdList_r16)
		if err = ie.sl_MeasIdToAddModList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_MeasIdToAddModList_r16", err)
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
