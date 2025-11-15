package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BandCombination_v1700 struct {
	ca_ParametersNR_v1700                               *CA_ParametersNR_v1700   `optional`
	ca_ParametersNRDC_v1700                             *CA_ParametersNRDC_v1700 `optional`
	mrdc_Parameters_v1700                               *MRDC_Parameters_v1700   `optional`
	bandList_v1710                                      []BandParameters_v1710   `lb:1,ub:maxSimultaneousBands,optional`
	supportedBandCombListPerBC_SL_RelayDiscovery_r17    *uper.BitString          `lb:1,ub:maxBandComb,optional`
	supportedBandCombListPerBC_SL_NonRelayDiscovery_r17 *uper.BitString          `lb:1,ub:maxBandComb,optional`
}

func (ie *BandCombination_v1700) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.ca_ParametersNR_v1700 != nil, ie.ca_ParametersNRDC_v1700 != nil, ie.mrdc_Parameters_v1700 != nil, len(ie.bandList_v1710) > 0, ie.supportedBandCombListPerBC_SL_RelayDiscovery_r17 != nil, ie.supportedBandCombListPerBC_SL_NonRelayDiscovery_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.ca_ParametersNR_v1700 != nil {
		if err = ie.ca_ParametersNR_v1700.Encode(w); err != nil {
			return utils.WrapError("Encode ca_ParametersNR_v1700", err)
		}
	}
	if ie.ca_ParametersNRDC_v1700 != nil {
		if err = ie.ca_ParametersNRDC_v1700.Encode(w); err != nil {
			return utils.WrapError("Encode ca_ParametersNRDC_v1700", err)
		}
	}
	if ie.mrdc_Parameters_v1700 != nil {
		if err = ie.mrdc_Parameters_v1700.Encode(w); err != nil {
			return utils.WrapError("Encode mrdc_Parameters_v1700", err)
		}
	}
	if len(ie.bandList_v1710) > 0 {
		tmp_bandList_v1710 := utils.NewSequence[*BandParameters_v1710]([]*BandParameters_v1710{}, uper.Constraint{Lb: 1, Ub: maxSimultaneousBands}, false)
		for _, i := range ie.bandList_v1710 {
			tmp_bandList_v1710.Value = append(tmp_bandList_v1710.Value, &i)
		}
		if err = tmp_bandList_v1710.Encode(w); err != nil {
			return utils.WrapError("Encode bandList_v1710", err)
		}
	}
	if ie.supportedBandCombListPerBC_SL_RelayDiscovery_r17 != nil {
		if err = w.WriteBitString(ie.supportedBandCombListPerBC_SL_RelayDiscovery_r17.Bytes, uint(ie.supportedBandCombListPerBC_SL_RelayDiscovery_r17.NumBits), &uper.Constraint{Lb: 1, Ub: maxBandComb}, false); err != nil {
			return utils.WrapError("Encode supportedBandCombListPerBC_SL_RelayDiscovery_r17", err)
		}
	}
	if ie.supportedBandCombListPerBC_SL_NonRelayDiscovery_r17 != nil {
		if err = w.WriteBitString(ie.supportedBandCombListPerBC_SL_NonRelayDiscovery_r17.Bytes, uint(ie.supportedBandCombListPerBC_SL_NonRelayDiscovery_r17.NumBits), &uper.Constraint{Lb: 1, Ub: maxBandComb}, false); err != nil {
			return utils.WrapError("Encode supportedBandCombListPerBC_SL_NonRelayDiscovery_r17", err)
		}
	}
	return nil
}

func (ie *BandCombination_v1700) Decode(r *uper.UperReader) error {
	var err error
	var ca_ParametersNR_v1700Present bool
	if ca_ParametersNR_v1700Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ca_ParametersNRDC_v1700Present bool
	if ca_ParametersNRDC_v1700Present, err = r.ReadBool(); err != nil {
		return err
	}
	var mrdc_Parameters_v1700Present bool
	if mrdc_Parameters_v1700Present, err = r.ReadBool(); err != nil {
		return err
	}
	var bandList_v1710Present bool
	if bandList_v1710Present, err = r.ReadBool(); err != nil {
		return err
	}
	var supportedBandCombListPerBC_SL_RelayDiscovery_r17Present bool
	if supportedBandCombListPerBC_SL_RelayDiscovery_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var supportedBandCombListPerBC_SL_NonRelayDiscovery_r17Present bool
	if supportedBandCombListPerBC_SL_NonRelayDiscovery_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if ca_ParametersNR_v1700Present {
		ie.ca_ParametersNR_v1700 = new(CA_ParametersNR_v1700)
		if err = ie.ca_ParametersNR_v1700.Decode(r); err != nil {
			return utils.WrapError("Decode ca_ParametersNR_v1700", err)
		}
	}
	if ca_ParametersNRDC_v1700Present {
		ie.ca_ParametersNRDC_v1700 = new(CA_ParametersNRDC_v1700)
		if err = ie.ca_ParametersNRDC_v1700.Decode(r); err != nil {
			return utils.WrapError("Decode ca_ParametersNRDC_v1700", err)
		}
	}
	if mrdc_Parameters_v1700Present {
		ie.mrdc_Parameters_v1700 = new(MRDC_Parameters_v1700)
		if err = ie.mrdc_Parameters_v1700.Decode(r); err != nil {
			return utils.WrapError("Decode mrdc_Parameters_v1700", err)
		}
	}
	if bandList_v1710Present {
		tmp_bandList_v1710 := utils.NewSequence[*BandParameters_v1710]([]*BandParameters_v1710{}, uper.Constraint{Lb: 1, Ub: maxSimultaneousBands}, false)
		fn_bandList_v1710 := func() *BandParameters_v1710 {
			return new(BandParameters_v1710)
		}
		if err = tmp_bandList_v1710.Decode(r, fn_bandList_v1710); err != nil {
			return utils.WrapError("Decode bandList_v1710", err)
		}
		ie.bandList_v1710 = []BandParameters_v1710{}
		for _, i := range tmp_bandList_v1710.Value {
			ie.bandList_v1710 = append(ie.bandList_v1710, *i)
		}
	}
	if supportedBandCombListPerBC_SL_RelayDiscovery_r17Present {
		var tmp_bs_supportedBandCombListPerBC_SL_RelayDiscovery_r17 []byte
		var n_supportedBandCombListPerBC_SL_RelayDiscovery_r17 uint
		if tmp_bs_supportedBandCombListPerBC_SL_RelayDiscovery_r17, n_supportedBandCombListPerBC_SL_RelayDiscovery_r17, err = r.ReadBitString(&uper.Constraint{Lb: 1, Ub: maxBandComb}, false); err != nil {
			return utils.WrapError("Decode supportedBandCombListPerBC_SL_RelayDiscovery_r17", err)
		}
		tmp_bitstring := uper.BitString{
			Bytes:   tmp_bs_supportedBandCombListPerBC_SL_RelayDiscovery_r17,
			NumBits: uint64(n_supportedBandCombListPerBC_SL_RelayDiscovery_r17),
		}
		ie.supportedBandCombListPerBC_SL_RelayDiscovery_r17 = &tmp_bitstring
	}
	if supportedBandCombListPerBC_SL_NonRelayDiscovery_r17Present {
		var tmp_bs_supportedBandCombListPerBC_SL_NonRelayDiscovery_r17 []byte
		var n_supportedBandCombListPerBC_SL_NonRelayDiscovery_r17 uint
		if tmp_bs_supportedBandCombListPerBC_SL_NonRelayDiscovery_r17, n_supportedBandCombListPerBC_SL_NonRelayDiscovery_r17, err = r.ReadBitString(&uper.Constraint{Lb: 1, Ub: maxBandComb}, false); err != nil {
			return utils.WrapError("Decode supportedBandCombListPerBC_SL_NonRelayDiscovery_r17", err)
		}
		tmp_bitstring := uper.BitString{
			Bytes:   tmp_bs_supportedBandCombListPerBC_SL_NonRelayDiscovery_r17,
			NumBits: uint64(n_supportedBandCombListPerBC_SL_NonRelayDiscovery_r17),
		}
		ie.supportedBandCombListPerBC_SL_NonRelayDiscovery_r17 = &tmp_bitstring
	}
	return nil
}
