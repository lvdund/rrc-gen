package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BandCombination_v1590 struct {
	supportedBandwidthCombinationSetIntraENDC *uper.BitString       `lb:1,ub:32,optional`
	mrdc_Parameters_v1590                     MRDC_Parameters_v1590 `madatory`
}

func (ie *BandCombination_v1590) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.supportedBandwidthCombinationSetIntraENDC != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.supportedBandwidthCombinationSetIntraENDC != nil {
		if err = w.WriteBitString(ie.supportedBandwidthCombinationSetIntraENDC.Bytes, uint(ie.supportedBandwidthCombinationSetIntraENDC.NumBits), &uper.Constraint{Lb: 1, Ub: 32}, false); err != nil {
			return utils.WrapError("Encode supportedBandwidthCombinationSetIntraENDC", err)
		}
	}
	if err = ie.mrdc_Parameters_v1590.Encode(w); err != nil {
		return utils.WrapError("Encode mrdc_Parameters_v1590", err)
	}
	return nil
}

func (ie *BandCombination_v1590) Decode(r *uper.UperReader) error {
	var err error
	var supportedBandwidthCombinationSetIntraENDCPresent bool
	if supportedBandwidthCombinationSetIntraENDCPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if supportedBandwidthCombinationSetIntraENDCPresent {
		var tmp_bs_supportedBandwidthCombinationSetIntraENDC []byte
		var n_supportedBandwidthCombinationSetIntraENDC uint
		if tmp_bs_supportedBandwidthCombinationSetIntraENDC, n_supportedBandwidthCombinationSetIntraENDC, err = r.ReadBitString(&uper.Constraint{Lb: 1, Ub: 32}, false); err != nil {
			return utils.WrapError("Decode supportedBandwidthCombinationSetIntraENDC", err)
		}
		tmp_bitstring := uper.BitString{
			Bytes:   tmp_bs_supportedBandwidthCombinationSetIntraENDC,
			NumBits: uint64(n_supportedBandwidthCombinationSetIntraENDC),
		}
		ie.supportedBandwidthCombinationSetIntraENDC = &tmp_bitstring
	}
	if err = ie.mrdc_Parameters_v1590.Decode(r); err != nil {
		return utils.WrapError("Decode mrdc_Parameters_v1590", err)
	}
	return nil
}
