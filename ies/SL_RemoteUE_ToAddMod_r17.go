package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_RemoteUE_ToAddMod_r17 struct {
	sl_L2IdentityRemote_r17 SL_DestinationIdentity_r16 `madatory`
	sl_SRAP_ConfigRelay_r17 *SL_SRAP_Config_r17        `optional`
}

func (ie *SL_RemoteUE_ToAddMod_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.sl_SRAP_ConfigRelay_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.sl_L2IdentityRemote_r17.Encode(w); err != nil {
		return utils.WrapError("Encode sl_L2IdentityRemote_r17", err)
	}
	if ie.sl_SRAP_ConfigRelay_r17 != nil {
		if err = ie.sl_SRAP_ConfigRelay_r17.Encode(w); err != nil {
			return utils.WrapError("Encode sl_SRAP_ConfigRelay_r17", err)
		}
	}
	return nil
}

func (ie *SL_RemoteUE_ToAddMod_r17) Decode(r *uper.UperReader) error {
	var err error
	var sl_SRAP_ConfigRelay_r17Present bool
	if sl_SRAP_ConfigRelay_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.sl_L2IdentityRemote_r17.Decode(r); err != nil {
		return utils.WrapError("Decode sl_L2IdentityRemote_r17", err)
	}
	if sl_SRAP_ConfigRelay_r17Present {
		ie.sl_SRAP_ConfigRelay_r17 = new(SL_SRAP_Config_r17)
		if err = ie.sl_SRAP_ConfigRelay_r17.Decode(r); err != nil {
			return utils.WrapError("Decode sl_SRAP_ConfigRelay_r17", err)
		}
	}
	return nil
}
