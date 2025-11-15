package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_SDAP_Config_r16 struct {
	sl_SDAP_Header_r16     SL_SDAP_Config_r16_sl_SDAP_Header_r16      `madatory`
	sl_DefaultRB_r16       bool                                       `madatory`
	sl_MappedQoS_Flows_r16 *SL_SDAP_Config_r16_sl_MappedQoS_Flows_r16 `lb:1,ub:maxNrofSL_QFIs_r16,optional`
	sl_CastType_r16        *SL_SDAP_Config_r16_sl_CastType_r16        `optional`
}

func (ie *SL_SDAP_Config_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.sl_MappedQoS_Flows_r16 != nil, ie.sl_CastType_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.sl_SDAP_Header_r16.Encode(w); err != nil {
		return utils.WrapError("Encode sl_SDAP_Header_r16", err)
	}
	if err = w.WriteBoolean(ie.sl_DefaultRB_r16); err != nil {
		return utils.WrapError("WriteBoolean sl_DefaultRB_r16", err)
	}
	if ie.sl_MappedQoS_Flows_r16 != nil {
		if err = ie.sl_MappedQoS_Flows_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_MappedQoS_Flows_r16", err)
		}
	}
	if ie.sl_CastType_r16 != nil {
		if err = ie.sl_CastType_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_CastType_r16", err)
		}
	}
	return nil
}

func (ie *SL_SDAP_Config_r16) Decode(r *uper.UperReader) error {
	var err error
	var sl_MappedQoS_Flows_r16Present bool
	if sl_MappedQoS_Flows_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_CastType_r16Present bool
	if sl_CastType_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.sl_SDAP_Header_r16.Decode(r); err != nil {
		return utils.WrapError("Decode sl_SDAP_Header_r16", err)
	}
	var tmp_bool_sl_DefaultRB_r16 bool
	if tmp_bool_sl_DefaultRB_r16, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean sl_DefaultRB_r16", err)
	}
	ie.sl_DefaultRB_r16 = tmp_bool_sl_DefaultRB_r16
	if sl_MappedQoS_Flows_r16Present {
		ie.sl_MappedQoS_Flows_r16 = new(SL_SDAP_Config_r16_sl_MappedQoS_Flows_r16)
		if err = ie.sl_MappedQoS_Flows_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_MappedQoS_Flows_r16", err)
		}
	}
	if sl_CastType_r16Present {
		ie.sl_CastType_r16 = new(SL_SDAP_Config_r16_sl_CastType_r16)
		if err = ie.sl_CastType_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_CastType_r16", err)
		}
	}
	return nil
}
