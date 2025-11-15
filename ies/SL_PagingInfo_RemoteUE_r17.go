package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_PagingInfo_RemoteUE_r17 struct {
	sl_PagingIdentityRemoteUE_r17 SL_PagingIdentityRemoteUE_r17 `madatory`
	sl_PagingCycleRemoteUE_r17    *PagingCycle                  `optional`
}

func (ie *SL_PagingInfo_RemoteUE_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.sl_PagingCycleRemoteUE_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.sl_PagingIdentityRemoteUE_r17.Encode(w); err != nil {
		return utils.WrapError("Encode sl_PagingIdentityRemoteUE_r17", err)
	}
	if ie.sl_PagingCycleRemoteUE_r17 != nil {
		if err = ie.sl_PagingCycleRemoteUE_r17.Encode(w); err != nil {
			return utils.WrapError("Encode sl_PagingCycleRemoteUE_r17", err)
		}
	}
	return nil
}

func (ie *SL_PagingInfo_RemoteUE_r17) Decode(r *uper.UperReader) error {
	var err error
	var sl_PagingCycleRemoteUE_r17Present bool
	if sl_PagingCycleRemoteUE_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.sl_PagingIdentityRemoteUE_r17.Decode(r); err != nil {
		return utils.WrapError("Decode sl_PagingIdentityRemoteUE_r17", err)
	}
	if sl_PagingCycleRemoteUE_r17Present {
		ie.sl_PagingCycleRemoteUE_r17 = new(PagingCycle)
		if err = ie.sl_PagingCycleRemoteUE_r17.Decode(r); err != nil {
			return utils.WrapError("Decode sl_PagingCycleRemoteUE_r17", err)
		}
	}
	return nil
}
