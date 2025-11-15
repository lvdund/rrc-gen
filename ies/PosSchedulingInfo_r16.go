package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PosSchedulingInfo_r16 struct {
	offsetToSI_Used_r16       *PosSchedulingInfo_r16_offsetToSI_Used_r16      `optional`
	posSI_Periodicity_r16     PosSchedulingInfo_r16_posSI_Periodicity_r16     `madatory`
	posSI_BroadcastStatus_r16 PosSchedulingInfo_r16_posSI_BroadcastStatus_r16 `madatory`
	posSIB_MappingInfo_r16    PosSIB_MappingInfo_r16                          `madatory`
}

func (ie *PosSchedulingInfo_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.offsetToSI_Used_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.offsetToSI_Used_r16 != nil {
		if err = ie.offsetToSI_Used_r16.Encode(w); err != nil {
			return utils.WrapError("Encode offsetToSI_Used_r16", err)
		}
	}
	if err = ie.posSI_Periodicity_r16.Encode(w); err != nil {
		return utils.WrapError("Encode posSI_Periodicity_r16", err)
	}
	if err = ie.posSI_BroadcastStatus_r16.Encode(w); err != nil {
		return utils.WrapError("Encode posSI_BroadcastStatus_r16", err)
	}
	if err = ie.posSIB_MappingInfo_r16.Encode(w); err != nil {
		return utils.WrapError("Encode posSIB_MappingInfo_r16", err)
	}
	return nil
}

func (ie *PosSchedulingInfo_r16) Decode(r *uper.UperReader) error {
	var err error
	var offsetToSI_Used_r16Present bool
	if offsetToSI_Used_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if offsetToSI_Used_r16Present {
		ie.offsetToSI_Used_r16 = new(PosSchedulingInfo_r16_offsetToSI_Used_r16)
		if err = ie.offsetToSI_Used_r16.Decode(r); err != nil {
			return utils.WrapError("Decode offsetToSI_Used_r16", err)
		}
	}
	if err = ie.posSI_Periodicity_r16.Decode(r); err != nil {
		return utils.WrapError("Decode posSI_Periodicity_r16", err)
	}
	if err = ie.posSI_BroadcastStatus_r16.Decode(r); err != nil {
		return utils.WrapError("Decode posSI_BroadcastStatus_r16", err)
	}
	if err = ie.posSIB_MappingInfo_r16.Decode(r); err != nil {
		return utils.WrapError("Decode posSIB_MappingInfo_r16", err)
	}
	return nil
}
