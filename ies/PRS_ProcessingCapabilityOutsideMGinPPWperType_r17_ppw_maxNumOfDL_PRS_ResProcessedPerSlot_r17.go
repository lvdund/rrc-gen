package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PRS_ProcessingCapabilityOutsideMGinPPWperType_r17_ppw_maxNumOfDL_PRS_ResProcessedPerSlot_r17 struct {
	scs15_r17  *PRS_ProcessingCapabilityOutsideMGinPPWperType_r17_ppw_maxNumOfDL_PRS_ResProcessedPerSlot_r17_scs15_r17  `optional`
	scs30_r17  *PRS_ProcessingCapabilityOutsideMGinPPWperType_r17_ppw_maxNumOfDL_PRS_ResProcessedPerSlot_r17_scs30_r17  `optional`
	scs60_r17  *PRS_ProcessingCapabilityOutsideMGinPPWperType_r17_ppw_maxNumOfDL_PRS_ResProcessedPerSlot_r17_scs60_r17  `optional`
	scs120_r17 *PRS_ProcessingCapabilityOutsideMGinPPWperType_r17_ppw_maxNumOfDL_PRS_ResProcessedPerSlot_r17_scs120_r17 `optional`
}

func (ie *PRS_ProcessingCapabilityOutsideMGinPPWperType_r17_ppw_maxNumOfDL_PRS_ResProcessedPerSlot_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.scs15_r17 != nil, ie.scs30_r17 != nil, ie.scs60_r17 != nil, ie.scs120_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.scs15_r17 != nil {
		if err = ie.scs15_r17.Encode(w); err != nil {
			return utils.WrapError("Encode scs15_r17", err)
		}
	}
	if ie.scs30_r17 != nil {
		if err = ie.scs30_r17.Encode(w); err != nil {
			return utils.WrapError("Encode scs30_r17", err)
		}
	}
	if ie.scs60_r17 != nil {
		if err = ie.scs60_r17.Encode(w); err != nil {
			return utils.WrapError("Encode scs60_r17", err)
		}
	}
	if ie.scs120_r17 != nil {
		if err = ie.scs120_r17.Encode(w); err != nil {
			return utils.WrapError("Encode scs120_r17", err)
		}
	}
	return nil
}

func (ie *PRS_ProcessingCapabilityOutsideMGinPPWperType_r17_ppw_maxNumOfDL_PRS_ResProcessedPerSlot_r17) Decode(r *uper.UperReader) error {
	var err error
	var scs15_r17Present bool
	if scs15_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var scs30_r17Present bool
	if scs30_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var scs60_r17Present bool
	if scs60_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var scs120_r17Present bool
	if scs120_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if scs15_r17Present {
		ie.scs15_r17 = new(PRS_ProcessingCapabilityOutsideMGinPPWperType_r17_ppw_maxNumOfDL_PRS_ResProcessedPerSlot_r17_scs15_r17)
		if err = ie.scs15_r17.Decode(r); err != nil {
			return utils.WrapError("Decode scs15_r17", err)
		}
	}
	if scs30_r17Present {
		ie.scs30_r17 = new(PRS_ProcessingCapabilityOutsideMGinPPWperType_r17_ppw_maxNumOfDL_PRS_ResProcessedPerSlot_r17_scs30_r17)
		if err = ie.scs30_r17.Decode(r); err != nil {
			return utils.WrapError("Decode scs30_r17", err)
		}
	}
	if scs60_r17Present {
		ie.scs60_r17 = new(PRS_ProcessingCapabilityOutsideMGinPPWperType_r17_ppw_maxNumOfDL_PRS_ResProcessedPerSlot_r17_scs60_r17)
		if err = ie.scs60_r17.Decode(r); err != nil {
			return utils.WrapError("Decode scs60_r17", err)
		}
	}
	if scs120_r17Present {
		ie.scs120_r17 = new(PRS_ProcessingCapabilityOutsideMGinPPWperType_r17_ppw_maxNumOfDL_PRS_ResProcessedPerSlot_r17_scs120_r17)
		if err = ie.scs120_r17.Decode(r); err != nil {
			return utils.WrapError("Decode scs120_r17", err)
		}
	}
	return nil
}
