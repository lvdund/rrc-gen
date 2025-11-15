package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BWP struct {
	locationAndBandwidth int64             `lb:0,ub:37949,madatory`
	subcarrierSpacing    SubcarrierSpacing `madatory`
	cyclicPrefix         *BWP_cyclicPrefix `optional`
}

func (ie *BWP) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.cyclicPrefix != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = w.WriteInteger(ie.locationAndBandwidth, &uper.Constraint{Lb: 0, Ub: 37949}, false); err != nil {
		return utils.WrapError("WriteInteger locationAndBandwidth", err)
	}
	if err = ie.subcarrierSpacing.Encode(w); err != nil {
		return utils.WrapError("Encode subcarrierSpacing", err)
	}
	if ie.cyclicPrefix != nil {
		if err = ie.cyclicPrefix.Encode(w); err != nil {
			return utils.WrapError("Encode cyclicPrefix", err)
		}
	}
	return nil
}

func (ie *BWP) Decode(r *uper.UperReader) error {
	var err error
	var cyclicPrefixPresent bool
	if cyclicPrefixPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var tmp_int_locationAndBandwidth int64
	if tmp_int_locationAndBandwidth, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 37949}, false); err != nil {
		return utils.WrapError("ReadInteger locationAndBandwidth", err)
	}
	ie.locationAndBandwidth = tmp_int_locationAndBandwidth
	if err = ie.subcarrierSpacing.Decode(r); err != nil {
		return utils.WrapError("Decode subcarrierSpacing", err)
	}
	if cyclicPrefixPresent {
		ie.cyclicPrefix = new(BWP_cyclicPrefix)
		if err = ie.cyclicPrefix.Decode(r); err != nil {
			return utils.WrapError("Decode cyclicPrefix", err)
		}
	}
	return nil
}
