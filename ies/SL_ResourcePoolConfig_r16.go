package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_ResourcePoolConfig_r16 struct {
	sl_ResourcePoolID_r16 SL_ResourcePoolID_r16 `madatory`
	sl_ResourcePool_r16   *SL_ResourcePool_r16  `optional`
}

func (ie *SL_ResourcePoolConfig_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.sl_ResourcePool_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.sl_ResourcePoolID_r16.Encode(w); err != nil {
		return utils.WrapError("Encode sl_ResourcePoolID_r16", err)
	}
	if ie.sl_ResourcePool_r16 != nil {
		if err = ie.sl_ResourcePool_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_ResourcePool_r16", err)
		}
	}
	return nil
}

func (ie *SL_ResourcePoolConfig_r16) Decode(r *uper.UperReader) error {
	var err error
	var sl_ResourcePool_r16Present bool
	if sl_ResourcePool_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.sl_ResourcePoolID_r16.Decode(r); err != nil {
		return utils.WrapError("Decode sl_ResourcePoolID_r16", err)
	}
	if sl_ResourcePool_r16Present {
		ie.sl_ResourcePool_r16 = new(SL_ResourcePool_r16)
		if err = ie.sl_ResourcePool_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_ResourcePool_r16", err)
		}
	}
	return nil
}
