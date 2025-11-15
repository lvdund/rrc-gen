package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UE_TimersAndConstantsRemoteUE_r17 struct {
	t300_RemoteUE_r17 *UE_TimersAndConstantsRemoteUE_r17_t300_RemoteUE_r17 `optional`
	t301_RemoteUE_r17 *UE_TimersAndConstantsRemoteUE_r17_t301_RemoteUE_r17 `optional`
	t319_RemoteUE_r17 *UE_TimersAndConstantsRemoteUE_r17_t319_RemoteUE_r17 `optional`
}

func (ie *UE_TimersAndConstantsRemoteUE_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.t300_RemoteUE_r17 != nil, ie.t301_RemoteUE_r17 != nil, ie.t319_RemoteUE_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.t300_RemoteUE_r17 != nil {
		if err = ie.t300_RemoteUE_r17.Encode(w); err != nil {
			return utils.WrapError("Encode t300_RemoteUE_r17", err)
		}
	}
	if ie.t301_RemoteUE_r17 != nil {
		if err = ie.t301_RemoteUE_r17.Encode(w); err != nil {
			return utils.WrapError("Encode t301_RemoteUE_r17", err)
		}
	}
	if ie.t319_RemoteUE_r17 != nil {
		if err = ie.t319_RemoteUE_r17.Encode(w); err != nil {
			return utils.WrapError("Encode t319_RemoteUE_r17", err)
		}
	}
	return nil
}

func (ie *UE_TimersAndConstantsRemoteUE_r17) Decode(r *uper.UperReader) error {
	var err error
	var t300_RemoteUE_r17Present bool
	if t300_RemoteUE_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var t301_RemoteUE_r17Present bool
	if t301_RemoteUE_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var t319_RemoteUE_r17Present bool
	if t319_RemoteUE_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if t300_RemoteUE_r17Present {
		ie.t300_RemoteUE_r17 = new(UE_TimersAndConstantsRemoteUE_r17_t300_RemoteUE_r17)
		if err = ie.t300_RemoteUE_r17.Decode(r); err != nil {
			return utils.WrapError("Decode t300_RemoteUE_r17", err)
		}
	}
	if t301_RemoteUE_r17Present {
		ie.t301_RemoteUE_r17 = new(UE_TimersAndConstantsRemoteUE_r17_t301_RemoteUE_r17)
		if err = ie.t301_RemoteUE_r17.Decode(r); err != nil {
			return utils.WrapError("Decode t301_RemoteUE_r17", err)
		}
	}
	if t319_RemoteUE_r17Present {
		ie.t319_RemoteUE_r17 = new(UE_TimersAndConstantsRemoteUE_r17_t319_RemoteUE_r17)
		if err = ie.t319_RemoteUE_r17.Decode(r); err != nil {
			return utils.WrapError("Decode t319_RemoteUE_r17", err)
		}
	}
	return nil
}
