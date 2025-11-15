package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type FreqBandInformationNR struct {
	bandNR                  FreqBandIndicatorNR  `madatory`
	maxBandwidthRequestedDL *AggregatedBandwidth `optional`
	maxBandwidthRequestedUL *AggregatedBandwidth `optional`
	maxCarriersRequestedDL  *int64               `lb:1,ub:maxNrofServingCells,optional`
	maxCarriersRequestedUL  *int64               `lb:1,ub:maxNrofServingCells,optional`
}

func (ie *FreqBandInformationNR) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.maxBandwidthRequestedDL != nil, ie.maxBandwidthRequestedUL != nil, ie.maxCarriersRequestedDL != nil, ie.maxCarriersRequestedUL != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.bandNR.Encode(w); err != nil {
		return utils.WrapError("Encode bandNR", err)
	}
	if ie.maxBandwidthRequestedDL != nil {
		if err = ie.maxBandwidthRequestedDL.Encode(w); err != nil {
			return utils.WrapError("Encode maxBandwidthRequestedDL", err)
		}
	}
	if ie.maxBandwidthRequestedUL != nil {
		if err = ie.maxBandwidthRequestedUL.Encode(w); err != nil {
			return utils.WrapError("Encode maxBandwidthRequestedUL", err)
		}
	}
	if ie.maxCarriersRequestedDL != nil {
		if err = w.WriteInteger(*ie.maxCarriersRequestedDL, &uper.Constraint{Lb: 1, Ub: maxNrofServingCells}, false); err != nil {
			return utils.WrapError("Encode maxCarriersRequestedDL", err)
		}
	}
	if ie.maxCarriersRequestedUL != nil {
		if err = w.WriteInteger(*ie.maxCarriersRequestedUL, &uper.Constraint{Lb: 1, Ub: maxNrofServingCells}, false); err != nil {
			return utils.WrapError("Encode maxCarriersRequestedUL", err)
		}
	}
	return nil
}

func (ie *FreqBandInformationNR) Decode(r *uper.UperReader) error {
	var err error
	var maxBandwidthRequestedDLPresent bool
	if maxBandwidthRequestedDLPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var maxBandwidthRequestedULPresent bool
	if maxBandwidthRequestedULPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var maxCarriersRequestedDLPresent bool
	if maxCarriersRequestedDLPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var maxCarriersRequestedULPresent bool
	if maxCarriersRequestedULPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.bandNR.Decode(r); err != nil {
		return utils.WrapError("Decode bandNR", err)
	}
	if maxBandwidthRequestedDLPresent {
		ie.maxBandwidthRequestedDL = new(AggregatedBandwidth)
		if err = ie.maxBandwidthRequestedDL.Decode(r); err != nil {
			return utils.WrapError("Decode maxBandwidthRequestedDL", err)
		}
	}
	if maxBandwidthRequestedULPresent {
		ie.maxBandwidthRequestedUL = new(AggregatedBandwidth)
		if err = ie.maxBandwidthRequestedUL.Decode(r); err != nil {
			return utils.WrapError("Decode maxBandwidthRequestedUL", err)
		}
	}
	if maxCarriersRequestedDLPresent {
		var tmp_int_maxCarriersRequestedDL int64
		if tmp_int_maxCarriersRequestedDL, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: maxNrofServingCells}, false); err != nil {
			return utils.WrapError("Decode maxCarriersRequestedDL", err)
		}
		ie.maxCarriersRequestedDL = &tmp_int_maxCarriersRequestedDL
	}
	if maxCarriersRequestedULPresent {
		var tmp_int_maxCarriersRequestedUL int64
		if tmp_int_maxCarriersRequestedUL, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: maxNrofServingCells}, false); err != nil {
			return utils.WrapError("Decode maxCarriersRequestedUL", err)
		}
		ie.maxCarriersRequestedUL = &tmp_int_maxCarriersRequestedUL
	}
	return nil
}
