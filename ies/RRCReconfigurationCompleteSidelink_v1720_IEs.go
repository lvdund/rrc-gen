package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RRCReconfigurationCompleteSidelink_v1720_IEs struct {
	sl_DRX_ConfigReject_v1720 *RRCReconfigurationCompleteSidelink_v1720_IEs_sl_DRX_ConfigReject_v1720 `optional`
	nonCriticalExtension      interface{}                                                             `optional`
}

func (ie *RRCReconfigurationCompleteSidelink_v1720_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.sl_DRX_ConfigReject_v1720 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.sl_DRX_ConfigReject_v1720 != nil {
		if err = ie.sl_DRX_ConfigReject_v1720.Encode(w); err != nil {
			return utils.WrapError("Encode sl_DRX_ConfigReject_v1720", err)
		}
	}
	return nil
}

func (ie *RRCReconfigurationCompleteSidelink_v1720_IEs) Decode(r *uper.UperReader) error {
	var err error
	var sl_DRX_ConfigReject_v1720Present bool
	if sl_DRX_ConfigReject_v1720Present, err = r.ReadBool(); err != nil {
		return err
	}
	if sl_DRX_ConfigReject_v1720Present {
		ie.sl_DRX_ConfigReject_v1720 = new(RRCReconfigurationCompleteSidelink_v1720_IEs_sl_DRX_ConfigReject_v1720)
		if err = ie.sl_DRX_ConfigReject_v1720.Decode(r); err != nil {
			return utils.WrapError("Decode sl_DRX_ConfigReject_v1720", err)
		}
	}
	return nil
}
