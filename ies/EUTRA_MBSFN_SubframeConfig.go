package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type EUTRA_MBSFN_SubframeConfig struct {
	radioframeAllocationPeriod EUTRA_MBSFN_SubframeConfig_radioframeAllocationPeriod `madatory`
	radioframeAllocationOffset int64                                                 `lb:0,ub:7,madatory`
	subframeAllocation1        EUTRA_MBSFN_SubframeConfig_subframeAllocation1        `lb:6,ub:6,madatory`
	subframeAllocation2        *EUTRA_MBSFN_SubframeConfig_subframeAllocation2       `lb:2,ub:2,optional`
}

func (ie *EUTRA_MBSFN_SubframeConfig) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.subframeAllocation2 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.radioframeAllocationPeriod.Encode(w); err != nil {
		return utils.WrapError("Encode radioframeAllocationPeriod", err)
	}
	if err = w.WriteInteger(ie.radioframeAllocationOffset, &uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("WriteInteger radioframeAllocationOffset", err)
	}
	if err = ie.subframeAllocation1.Encode(w); err != nil {
		return utils.WrapError("Encode subframeAllocation1", err)
	}
	if ie.subframeAllocation2 != nil {
		if err = ie.subframeAllocation2.Encode(w); err != nil {
			return utils.WrapError("Encode subframeAllocation2", err)
		}
	}
	return nil
}

func (ie *EUTRA_MBSFN_SubframeConfig) Decode(r *uper.UperReader) error {
	var err error
	var subframeAllocation2Present bool
	if subframeAllocation2Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.radioframeAllocationPeriod.Decode(r); err != nil {
		return utils.WrapError("Decode radioframeAllocationPeriod", err)
	}
	var tmp_int_radioframeAllocationOffset int64
	if tmp_int_radioframeAllocationOffset, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("ReadInteger radioframeAllocationOffset", err)
	}
	ie.radioframeAllocationOffset = tmp_int_radioframeAllocationOffset
	if err = ie.subframeAllocation1.Decode(r); err != nil {
		return utils.WrapError("Decode subframeAllocation1", err)
	}
	if subframeAllocation2Present {
		ie.subframeAllocation2 = new(EUTRA_MBSFN_SubframeConfig_subframeAllocation2)
		if err = ie.subframeAllocation2.Decode(r); err != nil {
			return utils.WrapError("Decode subframeAllocation2", err)
		}
	}
	return nil
}
