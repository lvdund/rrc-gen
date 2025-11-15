package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type AdditionalRACH_Config_r17 struct {
	rach_ConfigCommon_r17 *RACH_ConfigCommon     `optional`
	msgA_ConfigCommon_r17 *MsgA_ConfigCommon_r16 `optional`
}

func (ie *AdditionalRACH_Config_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.rach_ConfigCommon_r17 != nil, ie.msgA_ConfigCommon_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.rach_ConfigCommon_r17 != nil {
		if err = ie.rach_ConfigCommon_r17.Encode(w); err != nil {
			return utils.WrapError("Encode rach_ConfigCommon_r17", err)
		}
	}
	if ie.msgA_ConfigCommon_r17 != nil {
		if err = ie.msgA_ConfigCommon_r17.Encode(w); err != nil {
			return utils.WrapError("Encode msgA_ConfigCommon_r17", err)
		}
	}
	return nil
}

func (ie *AdditionalRACH_Config_r17) Decode(r *uper.UperReader) error {
	var err error
	var rach_ConfigCommon_r17Present bool
	if rach_ConfigCommon_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var msgA_ConfigCommon_r17Present bool
	if msgA_ConfigCommon_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if rach_ConfigCommon_r17Present {
		ie.rach_ConfigCommon_r17 = new(RACH_ConfigCommon)
		if err = ie.rach_ConfigCommon_r17.Decode(r); err != nil {
			return utils.WrapError("Decode rach_ConfigCommon_r17", err)
		}
	}
	if msgA_ConfigCommon_r17Present {
		ie.msgA_ConfigCommon_r17 = new(MsgA_ConfigCommon_r16)
		if err = ie.msgA_ConfigCommon_r17.Decode(r); err != nil {
			return utils.WrapError("Decode msgA_ConfigCommon_r17", err)
		}
	}
	return nil
}
