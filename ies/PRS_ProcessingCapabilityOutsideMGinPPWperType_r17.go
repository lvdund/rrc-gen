package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PRS_ProcessingCapabilityOutsideMGinPPWperType_r17 struct {
	prsProcessingType_r17                      PRS_ProcessingCapabilityOutsideMGinPPWperType_r17_prsProcessingType_r17                       `madatory`
	ppw_dl_PRS_BufferType_r17                  PRS_ProcessingCapabilityOutsideMGinPPWperType_r17_ppw_dl_PRS_BufferType_r17                   `madatory`
	ppw_durationOfPRS_Processing_r17           *PRS_ProcessingCapabilityOutsideMGinPPWperType_r17_ppw_durationOfPRS_Processing_r17           `optional`
	ppw_maxNumOfDL_PRS_ResProcessedPerSlot_r17 *PRS_ProcessingCapabilityOutsideMGinPPWperType_r17_ppw_maxNumOfDL_PRS_ResProcessedPerSlot_r17 `optional`
	ppw_maxNumOfDL_Bandwidth_r17               *PRS_ProcessingCapabilityOutsideMGinPPWperType_r17_ppw_maxNumOfDL_Bandwidth_r17               `optional,ext`
}

func (ie *PRS_ProcessingCapabilityOutsideMGinPPWperType_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.ppw_durationOfPRS_Processing_r17 != nil, ie.ppw_maxNumOfDL_PRS_ResProcessedPerSlot_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.prsProcessingType_r17.Encode(w); err != nil {
		return utils.WrapError("Encode prsProcessingType_r17", err)
	}
	if err = ie.ppw_dl_PRS_BufferType_r17.Encode(w); err != nil {
		return utils.WrapError("Encode ppw_dl_PRS_BufferType_r17", err)
	}
	if ie.ppw_durationOfPRS_Processing_r17 != nil {
		if err = ie.ppw_durationOfPRS_Processing_r17.Encode(w); err != nil {
			return utils.WrapError("Encode ppw_durationOfPRS_Processing_r17", err)
		}
	}
	if ie.ppw_maxNumOfDL_PRS_ResProcessedPerSlot_r17 != nil {
		if err = ie.ppw_maxNumOfDL_PRS_ResProcessedPerSlot_r17.Encode(w); err != nil {
			return utils.WrapError("Encode ppw_maxNumOfDL_PRS_ResProcessedPerSlot_r17", err)
		}
	}
	return nil
}

func (ie *PRS_ProcessingCapabilityOutsideMGinPPWperType_r17) Decode(r *uper.UperReader) error {
	var err error
	var ppw_durationOfPRS_Processing_r17Present bool
	if ppw_durationOfPRS_Processing_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ppw_maxNumOfDL_PRS_ResProcessedPerSlot_r17Present bool
	if ppw_maxNumOfDL_PRS_ResProcessedPerSlot_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.prsProcessingType_r17.Decode(r); err != nil {
		return utils.WrapError("Decode prsProcessingType_r17", err)
	}
	if err = ie.ppw_dl_PRS_BufferType_r17.Decode(r); err != nil {
		return utils.WrapError("Decode ppw_dl_PRS_BufferType_r17", err)
	}
	if ppw_durationOfPRS_Processing_r17Present {
		ie.ppw_durationOfPRS_Processing_r17 = new(PRS_ProcessingCapabilityOutsideMGinPPWperType_r17_ppw_durationOfPRS_Processing_r17)
		if err = ie.ppw_durationOfPRS_Processing_r17.Decode(r); err != nil {
			return utils.WrapError("Decode ppw_durationOfPRS_Processing_r17", err)
		}
	}
	if ppw_maxNumOfDL_PRS_ResProcessedPerSlot_r17Present {
		ie.ppw_maxNumOfDL_PRS_ResProcessedPerSlot_r17 = new(PRS_ProcessingCapabilityOutsideMGinPPWperType_r17_ppw_maxNumOfDL_PRS_ResProcessedPerSlot_r17)
		if err = ie.ppw_maxNumOfDL_PRS_ResProcessedPerSlot_r17.Decode(r); err != nil {
			return utils.WrapError("Decode ppw_maxNumOfDL_PRS_ResProcessedPerSlot_r17", err)
		}
	}
	return nil
}
