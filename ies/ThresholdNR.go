package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ThresholdNR struct {
	thresholdRSRP *RSRP_Range `optional`
	thresholdRSRQ *RSRQ_Range `optional`
	thresholdSINR *SINR_Range `optional`
}

func (ie *ThresholdNR) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.thresholdRSRP != nil, ie.thresholdRSRQ != nil, ie.thresholdSINR != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.thresholdRSRP != nil {
		if err = ie.thresholdRSRP.Encode(w); err != nil {
			return utils.WrapError("Encode thresholdRSRP", err)
		}
	}
	if ie.thresholdRSRQ != nil {
		if err = ie.thresholdRSRQ.Encode(w); err != nil {
			return utils.WrapError("Encode thresholdRSRQ", err)
		}
	}
	if ie.thresholdSINR != nil {
		if err = ie.thresholdSINR.Encode(w); err != nil {
			return utils.WrapError("Encode thresholdSINR", err)
		}
	}
	return nil
}

func (ie *ThresholdNR) Decode(r *uper.UperReader) error {
	var err error
	var thresholdRSRPPresent bool
	if thresholdRSRPPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var thresholdRSRQPresent bool
	if thresholdRSRQPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var thresholdSINRPresent bool
	if thresholdSINRPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if thresholdRSRPPresent {
		ie.thresholdRSRP = new(RSRP_Range)
		if err = ie.thresholdRSRP.Decode(r); err != nil {
			return utils.WrapError("Decode thresholdRSRP", err)
		}
	}
	if thresholdRSRQPresent {
		ie.thresholdRSRQ = new(RSRQ_Range)
		if err = ie.thresholdRSRQ.Decode(r); err != nil {
			return utils.WrapError("Decode thresholdRSRQ", err)
		}
	}
	if thresholdSINRPresent {
		ie.thresholdSINR = new(SINR_Range)
		if err = ie.thresholdSINR.Decode(r); err != nil {
			return utils.WrapError("Decode thresholdSINR", err)
		}
	}
	return nil
}
