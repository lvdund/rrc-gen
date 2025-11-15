package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CA_ParametersEUTRA struct {
	multipleTimingAdvance                       *CA_ParametersEUTRA_multipleTimingAdvance          `optional`
	simultaneousRx_Tx                           *CA_ParametersEUTRA_simultaneousRx_Tx              `optional`
	supportedNAICS_2CRS_AP                      *uper.BitString                                    `lb:1,ub:8,optional`
	additionalRx_Tx_PerformanceReq              *CA_ParametersEUTRA_additionalRx_Tx_PerformanceReq `optional`
	ue_CA_PowerClass_N                          *CA_ParametersEUTRA_ue_CA_PowerClass_N             `optional`
	supportedBandwidthCombinationSetEUTRA_v1530 *uper.BitString                                    `lb:1,ub:32,optional`
}

func (ie *CA_ParametersEUTRA) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.multipleTimingAdvance != nil, ie.simultaneousRx_Tx != nil, ie.supportedNAICS_2CRS_AP != nil, ie.additionalRx_Tx_PerformanceReq != nil, ie.ue_CA_PowerClass_N != nil, ie.supportedBandwidthCombinationSetEUTRA_v1530 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.multipleTimingAdvance != nil {
		if err = ie.multipleTimingAdvance.Encode(w); err != nil {
			return utils.WrapError("Encode multipleTimingAdvance", err)
		}
	}
	if ie.simultaneousRx_Tx != nil {
		if err = ie.simultaneousRx_Tx.Encode(w); err != nil {
			return utils.WrapError("Encode simultaneousRx_Tx", err)
		}
	}
	if ie.supportedNAICS_2CRS_AP != nil {
		if err = w.WriteBitString(ie.supportedNAICS_2CRS_AP.Bytes, uint(ie.supportedNAICS_2CRS_AP.NumBits), &uper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
			return utils.WrapError("Encode supportedNAICS_2CRS_AP", err)
		}
	}
	if ie.additionalRx_Tx_PerformanceReq != nil {
		if err = ie.additionalRx_Tx_PerformanceReq.Encode(w); err != nil {
			return utils.WrapError("Encode additionalRx_Tx_PerformanceReq", err)
		}
	}
	if ie.ue_CA_PowerClass_N != nil {
		if err = ie.ue_CA_PowerClass_N.Encode(w); err != nil {
			return utils.WrapError("Encode ue_CA_PowerClass_N", err)
		}
	}
	if ie.supportedBandwidthCombinationSetEUTRA_v1530 != nil {
		if err = w.WriteBitString(ie.supportedBandwidthCombinationSetEUTRA_v1530.Bytes, uint(ie.supportedBandwidthCombinationSetEUTRA_v1530.NumBits), &uper.Constraint{Lb: 1, Ub: 32}, false); err != nil {
			return utils.WrapError("Encode supportedBandwidthCombinationSetEUTRA_v1530", err)
		}
	}
	return nil
}

func (ie *CA_ParametersEUTRA) Decode(r *uper.UperReader) error {
	var err error
	var multipleTimingAdvancePresent bool
	if multipleTimingAdvancePresent, err = r.ReadBool(); err != nil {
		return err
	}
	var simultaneousRx_TxPresent bool
	if simultaneousRx_TxPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var supportedNAICS_2CRS_APPresent bool
	if supportedNAICS_2CRS_APPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var additionalRx_Tx_PerformanceReqPresent bool
	if additionalRx_Tx_PerformanceReqPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var ue_CA_PowerClass_NPresent bool
	if ue_CA_PowerClass_NPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var supportedBandwidthCombinationSetEUTRA_v1530Present bool
	if supportedBandwidthCombinationSetEUTRA_v1530Present, err = r.ReadBool(); err != nil {
		return err
	}
	if multipleTimingAdvancePresent {
		ie.multipleTimingAdvance = new(CA_ParametersEUTRA_multipleTimingAdvance)
		if err = ie.multipleTimingAdvance.Decode(r); err != nil {
			return utils.WrapError("Decode multipleTimingAdvance", err)
		}
	}
	if simultaneousRx_TxPresent {
		ie.simultaneousRx_Tx = new(CA_ParametersEUTRA_simultaneousRx_Tx)
		if err = ie.simultaneousRx_Tx.Decode(r); err != nil {
			return utils.WrapError("Decode simultaneousRx_Tx", err)
		}
	}
	if supportedNAICS_2CRS_APPresent {
		var tmp_bs_supportedNAICS_2CRS_AP []byte
		var n_supportedNAICS_2CRS_AP uint
		if tmp_bs_supportedNAICS_2CRS_AP, n_supportedNAICS_2CRS_AP, err = r.ReadBitString(&uper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
			return utils.WrapError("Decode supportedNAICS_2CRS_AP", err)
		}
		tmp_bitstring := uper.BitString{
			Bytes:   tmp_bs_supportedNAICS_2CRS_AP,
			NumBits: uint64(n_supportedNAICS_2CRS_AP),
		}
		ie.supportedNAICS_2CRS_AP = &tmp_bitstring
	}
	if additionalRx_Tx_PerformanceReqPresent {
		ie.additionalRx_Tx_PerformanceReq = new(CA_ParametersEUTRA_additionalRx_Tx_PerformanceReq)
		if err = ie.additionalRx_Tx_PerformanceReq.Decode(r); err != nil {
			return utils.WrapError("Decode additionalRx_Tx_PerformanceReq", err)
		}
	}
	if ue_CA_PowerClass_NPresent {
		ie.ue_CA_PowerClass_N = new(CA_ParametersEUTRA_ue_CA_PowerClass_N)
		if err = ie.ue_CA_PowerClass_N.Decode(r); err != nil {
			return utils.WrapError("Decode ue_CA_PowerClass_N", err)
		}
	}
	if supportedBandwidthCombinationSetEUTRA_v1530Present {
		var tmp_bs_supportedBandwidthCombinationSetEUTRA_v1530 []byte
		var n_supportedBandwidthCombinationSetEUTRA_v1530 uint
		if tmp_bs_supportedBandwidthCombinationSetEUTRA_v1530, n_supportedBandwidthCombinationSetEUTRA_v1530, err = r.ReadBitString(&uper.Constraint{Lb: 1, Ub: 32}, false); err != nil {
			return utils.WrapError("Decode supportedBandwidthCombinationSetEUTRA_v1530", err)
		}
		tmp_bitstring := uper.BitString{
			Bytes:   tmp_bs_supportedBandwidthCombinationSetEUTRA_v1530,
			NumBits: uint64(n_supportedBandwidthCombinationSetEUTRA_v1530),
		}
		ie.supportedBandwidthCombinationSetEUTRA_v1530 = &tmp_bitstring
	}
	return nil
}
