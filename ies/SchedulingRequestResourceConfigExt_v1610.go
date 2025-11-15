package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SchedulingRequestResourceConfigExt_v1610 struct {
	phy_PriorityIndex_r16 *SchedulingRequestResourceConfigExt_v1610_phy_PriorityIndex_r16 `optional`
}

func (ie *SchedulingRequestResourceConfigExt_v1610) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.phy_PriorityIndex_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.phy_PriorityIndex_r16 != nil {
		if err = ie.phy_PriorityIndex_r16.Encode(w); err != nil {
			return utils.WrapError("Encode phy_PriorityIndex_r16", err)
		}
	}
	return nil
}

func (ie *SchedulingRequestResourceConfigExt_v1610) Decode(r *uper.UperReader) error {
	var err error
	var phy_PriorityIndex_r16Present bool
	if phy_PriorityIndex_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if phy_PriorityIndex_r16Present {
		ie.phy_PriorityIndex_r16 = new(SchedulingRequestResourceConfigExt_v1610_phy_PriorityIndex_r16)
		if err = ie.phy_PriorityIndex_r16.Decode(r); err != nil {
			return utils.WrapError("Decode phy_PriorityIndex_r16", err)
		}
	}
	return nil
}
