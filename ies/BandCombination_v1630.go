package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BandCombination_v1630 struct {
	ca_ParametersNR_v1630                     *CA_ParametersNR_v1630      `optional`
	ca_ParametersNRDC_v1630                   *CA_ParametersNRDC_v1630    `optional`
	mrdc_Parameters_v1630                     *MRDC_Parameters_v1630      `optional`
	supportedTxBandCombListPerBC_Sidelink_r16 *uper.BitString             `lb:1,ub:maxBandComb,optional`
	supportedRxBandCombListPerBC_Sidelink_r16 *uper.BitString             `lb:1,ub:maxBandComb,optional`
	scalingFactorTxSidelink_r16               []ScalingFactorSidelink_r16 `lb:1,ub:maxBandComb,optional`
	scalingFactorRxSidelink_r16               []ScalingFactorSidelink_r16 `lb:1,ub:maxBandComb,optional`
}

func (ie *BandCombination_v1630) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.ca_ParametersNR_v1630 != nil, ie.ca_ParametersNRDC_v1630 != nil, ie.mrdc_Parameters_v1630 != nil, ie.supportedTxBandCombListPerBC_Sidelink_r16 != nil, ie.supportedRxBandCombListPerBC_Sidelink_r16 != nil, len(ie.scalingFactorTxSidelink_r16) > 0, len(ie.scalingFactorRxSidelink_r16) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.ca_ParametersNR_v1630 != nil {
		if err = ie.ca_ParametersNR_v1630.Encode(w); err != nil {
			return utils.WrapError("Encode ca_ParametersNR_v1630", err)
		}
	}
	if ie.ca_ParametersNRDC_v1630 != nil {
		if err = ie.ca_ParametersNRDC_v1630.Encode(w); err != nil {
			return utils.WrapError("Encode ca_ParametersNRDC_v1630", err)
		}
	}
	if ie.mrdc_Parameters_v1630 != nil {
		if err = ie.mrdc_Parameters_v1630.Encode(w); err != nil {
			return utils.WrapError("Encode mrdc_Parameters_v1630", err)
		}
	}
	if ie.supportedTxBandCombListPerBC_Sidelink_r16 != nil {
		if err = w.WriteBitString(ie.supportedTxBandCombListPerBC_Sidelink_r16.Bytes, uint(ie.supportedTxBandCombListPerBC_Sidelink_r16.NumBits), &uper.Constraint{Lb: 1, Ub: maxBandComb}, false); err != nil {
			return utils.WrapError("Encode supportedTxBandCombListPerBC_Sidelink_r16", err)
		}
	}
	if ie.supportedRxBandCombListPerBC_Sidelink_r16 != nil {
		if err = w.WriteBitString(ie.supportedRxBandCombListPerBC_Sidelink_r16.Bytes, uint(ie.supportedRxBandCombListPerBC_Sidelink_r16.NumBits), &uper.Constraint{Lb: 1, Ub: maxBandComb}, false); err != nil {
			return utils.WrapError("Encode supportedRxBandCombListPerBC_Sidelink_r16", err)
		}
	}
	if len(ie.scalingFactorTxSidelink_r16) > 0 {
		tmp_scalingFactorTxSidelink_r16 := utils.NewSequence[*ScalingFactorSidelink_r16]([]*ScalingFactorSidelink_r16{}, uper.Constraint{Lb: 1, Ub: maxBandComb}, false)
		for _, i := range ie.scalingFactorTxSidelink_r16 {
			tmp_scalingFactorTxSidelink_r16.Value = append(tmp_scalingFactorTxSidelink_r16.Value, &i)
		}
		if err = tmp_scalingFactorTxSidelink_r16.Encode(w); err != nil {
			return utils.WrapError("Encode scalingFactorTxSidelink_r16", err)
		}
	}
	if len(ie.scalingFactorRxSidelink_r16) > 0 {
		tmp_scalingFactorRxSidelink_r16 := utils.NewSequence[*ScalingFactorSidelink_r16]([]*ScalingFactorSidelink_r16{}, uper.Constraint{Lb: 1, Ub: maxBandComb}, false)
		for _, i := range ie.scalingFactorRxSidelink_r16 {
			tmp_scalingFactorRxSidelink_r16.Value = append(tmp_scalingFactorRxSidelink_r16.Value, &i)
		}
		if err = tmp_scalingFactorRxSidelink_r16.Encode(w); err != nil {
			return utils.WrapError("Encode scalingFactorRxSidelink_r16", err)
		}
	}
	return nil
}

func (ie *BandCombination_v1630) Decode(r *uper.UperReader) error {
	var err error
	var ca_ParametersNR_v1630Present bool
	if ca_ParametersNR_v1630Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ca_ParametersNRDC_v1630Present bool
	if ca_ParametersNRDC_v1630Present, err = r.ReadBool(); err != nil {
		return err
	}
	var mrdc_Parameters_v1630Present bool
	if mrdc_Parameters_v1630Present, err = r.ReadBool(); err != nil {
		return err
	}
	var supportedTxBandCombListPerBC_Sidelink_r16Present bool
	if supportedTxBandCombListPerBC_Sidelink_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var supportedRxBandCombListPerBC_Sidelink_r16Present bool
	if supportedRxBandCombListPerBC_Sidelink_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var scalingFactorTxSidelink_r16Present bool
	if scalingFactorTxSidelink_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var scalingFactorRxSidelink_r16Present bool
	if scalingFactorRxSidelink_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if ca_ParametersNR_v1630Present {
		ie.ca_ParametersNR_v1630 = new(CA_ParametersNR_v1630)
		if err = ie.ca_ParametersNR_v1630.Decode(r); err != nil {
			return utils.WrapError("Decode ca_ParametersNR_v1630", err)
		}
	}
	if ca_ParametersNRDC_v1630Present {
		ie.ca_ParametersNRDC_v1630 = new(CA_ParametersNRDC_v1630)
		if err = ie.ca_ParametersNRDC_v1630.Decode(r); err != nil {
			return utils.WrapError("Decode ca_ParametersNRDC_v1630", err)
		}
	}
	if mrdc_Parameters_v1630Present {
		ie.mrdc_Parameters_v1630 = new(MRDC_Parameters_v1630)
		if err = ie.mrdc_Parameters_v1630.Decode(r); err != nil {
			return utils.WrapError("Decode mrdc_Parameters_v1630", err)
		}
	}
	if supportedTxBandCombListPerBC_Sidelink_r16Present {
		var tmp_bs_supportedTxBandCombListPerBC_Sidelink_r16 []byte
		var n_supportedTxBandCombListPerBC_Sidelink_r16 uint
		if tmp_bs_supportedTxBandCombListPerBC_Sidelink_r16, n_supportedTxBandCombListPerBC_Sidelink_r16, err = r.ReadBitString(&uper.Constraint{Lb: 1, Ub: maxBandComb}, false); err != nil {
			return utils.WrapError("Decode supportedTxBandCombListPerBC_Sidelink_r16", err)
		}
		tmp_bitstring := uper.BitString{
			Bytes:   tmp_bs_supportedTxBandCombListPerBC_Sidelink_r16,
			NumBits: uint64(n_supportedTxBandCombListPerBC_Sidelink_r16),
		}
		ie.supportedTxBandCombListPerBC_Sidelink_r16 = &tmp_bitstring
	}
	if supportedRxBandCombListPerBC_Sidelink_r16Present {
		var tmp_bs_supportedRxBandCombListPerBC_Sidelink_r16 []byte
		var n_supportedRxBandCombListPerBC_Sidelink_r16 uint
		if tmp_bs_supportedRxBandCombListPerBC_Sidelink_r16, n_supportedRxBandCombListPerBC_Sidelink_r16, err = r.ReadBitString(&uper.Constraint{Lb: 1, Ub: maxBandComb}, false); err != nil {
			return utils.WrapError("Decode supportedRxBandCombListPerBC_Sidelink_r16", err)
		}
		tmp_bitstring := uper.BitString{
			Bytes:   tmp_bs_supportedRxBandCombListPerBC_Sidelink_r16,
			NumBits: uint64(n_supportedRxBandCombListPerBC_Sidelink_r16),
		}
		ie.supportedRxBandCombListPerBC_Sidelink_r16 = &tmp_bitstring
	}
	if scalingFactorTxSidelink_r16Present {
		tmp_scalingFactorTxSidelink_r16 := utils.NewSequence[*ScalingFactorSidelink_r16]([]*ScalingFactorSidelink_r16{}, uper.Constraint{Lb: 1, Ub: maxBandComb}, false)
		fn_scalingFactorTxSidelink_r16 := func() *ScalingFactorSidelink_r16 {
			return new(ScalingFactorSidelink_r16)
		}
		if err = tmp_scalingFactorTxSidelink_r16.Decode(r, fn_scalingFactorTxSidelink_r16); err != nil {
			return utils.WrapError("Decode scalingFactorTxSidelink_r16", err)
		}
		ie.scalingFactorTxSidelink_r16 = []ScalingFactorSidelink_r16{}
		for _, i := range tmp_scalingFactorTxSidelink_r16.Value {
			ie.scalingFactorTxSidelink_r16 = append(ie.scalingFactorTxSidelink_r16, *i)
		}
	}
	if scalingFactorRxSidelink_r16Present {
		tmp_scalingFactorRxSidelink_r16 := utils.NewSequence[*ScalingFactorSidelink_r16]([]*ScalingFactorSidelink_r16{}, uper.Constraint{Lb: 1, Ub: maxBandComb}, false)
		fn_scalingFactorRxSidelink_r16 := func() *ScalingFactorSidelink_r16 {
			return new(ScalingFactorSidelink_r16)
		}
		if err = tmp_scalingFactorRxSidelink_r16.Decode(r, fn_scalingFactorRxSidelink_r16); err != nil {
			return utils.WrapError("Decode scalingFactorRxSidelink_r16", err)
		}
		ie.scalingFactorRxSidelink_r16 = []ScalingFactorSidelink_r16{}
		for _, i := range tmp_scalingFactorRxSidelink_r16.Value {
			ie.scalingFactorRxSidelink_r16 = append(ie.scalingFactorRxSidelink_r16, *i)
		}
	}
	return nil
}
