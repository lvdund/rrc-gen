package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasResultRLFNR_r16_measResult_r16_rsIndexResults_r16 struct {
	resultsSSB_Indexes_r16    *ResultsPerSSB_IndexList    `optional`
	ssbRLMConfigBitmap_r16    *uper.BitString             `lb:64,ub:64,optional`
	resultsCSI_RS_Indexes_r16 *ResultsPerCSI_RS_IndexList `optional`
	csi_rsRLMConfigBitmap_r16 *uper.BitString             `lb:96,ub:96,optional`
}

func (ie *MeasResultRLFNR_r16_measResult_r16_rsIndexResults_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.resultsSSB_Indexes_r16 != nil, ie.ssbRLMConfigBitmap_r16 != nil, ie.resultsCSI_RS_Indexes_r16 != nil, ie.csi_rsRLMConfigBitmap_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.resultsSSB_Indexes_r16 != nil {
		if err = ie.resultsSSB_Indexes_r16.Encode(w); err != nil {
			return utils.WrapError("Encode resultsSSB_Indexes_r16", err)
		}
	}
	if ie.ssbRLMConfigBitmap_r16 != nil {
		if err = w.WriteBitString(ie.ssbRLMConfigBitmap_r16.Bytes, uint(ie.ssbRLMConfigBitmap_r16.NumBits), &uper.Constraint{Lb: 64, Ub: 64}, false); err != nil {
			return utils.WrapError("Encode ssbRLMConfigBitmap_r16", err)
		}
	}
	if ie.resultsCSI_RS_Indexes_r16 != nil {
		if err = ie.resultsCSI_RS_Indexes_r16.Encode(w); err != nil {
			return utils.WrapError("Encode resultsCSI_RS_Indexes_r16", err)
		}
	}
	if ie.csi_rsRLMConfigBitmap_r16 != nil {
		if err = w.WriteBitString(ie.csi_rsRLMConfigBitmap_r16.Bytes, uint(ie.csi_rsRLMConfigBitmap_r16.NumBits), &uper.Constraint{Lb: 96, Ub: 96}, false); err != nil {
			return utils.WrapError("Encode csi_rsRLMConfigBitmap_r16", err)
		}
	}
	return nil
}

func (ie *MeasResultRLFNR_r16_measResult_r16_rsIndexResults_r16) Decode(r *uper.UperReader) error {
	var err error
	var resultsSSB_Indexes_r16Present bool
	if resultsSSB_Indexes_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ssbRLMConfigBitmap_r16Present bool
	if ssbRLMConfigBitmap_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var resultsCSI_RS_Indexes_r16Present bool
	if resultsCSI_RS_Indexes_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var csi_rsRLMConfigBitmap_r16Present bool
	if csi_rsRLMConfigBitmap_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if resultsSSB_Indexes_r16Present {
		ie.resultsSSB_Indexes_r16 = new(ResultsPerSSB_IndexList)
		if err = ie.resultsSSB_Indexes_r16.Decode(r); err != nil {
			return utils.WrapError("Decode resultsSSB_Indexes_r16", err)
		}
	}
	if ssbRLMConfigBitmap_r16Present {
		var tmp_bs_ssbRLMConfigBitmap_r16 []byte
		var n_ssbRLMConfigBitmap_r16 uint
		if tmp_bs_ssbRLMConfigBitmap_r16, n_ssbRLMConfigBitmap_r16, err = r.ReadBitString(&uper.Constraint{Lb: 64, Ub: 64}, false); err != nil {
			return utils.WrapError("Decode ssbRLMConfigBitmap_r16", err)
		}
		tmp_bitstring := uper.BitString{
			Bytes:   tmp_bs_ssbRLMConfigBitmap_r16,
			NumBits: uint64(n_ssbRLMConfigBitmap_r16),
		}
		ie.ssbRLMConfigBitmap_r16 = &tmp_bitstring
	}
	if resultsCSI_RS_Indexes_r16Present {
		ie.resultsCSI_RS_Indexes_r16 = new(ResultsPerCSI_RS_IndexList)
		if err = ie.resultsCSI_RS_Indexes_r16.Decode(r); err != nil {
			return utils.WrapError("Decode resultsCSI_RS_Indexes_r16", err)
		}
	}
	if csi_rsRLMConfigBitmap_r16Present {
		var tmp_bs_csi_rsRLMConfigBitmap_r16 []byte
		var n_csi_rsRLMConfigBitmap_r16 uint
		if tmp_bs_csi_rsRLMConfigBitmap_r16, n_csi_rsRLMConfigBitmap_r16, err = r.ReadBitString(&uper.Constraint{Lb: 96, Ub: 96}, false); err != nil {
			return utils.WrapError("Decode csi_rsRLMConfigBitmap_r16", err)
		}
		tmp_bitstring := uper.BitString{
			Bytes:   tmp_bs_csi_rsRLMConfigBitmap_r16,
			NumBits: uint64(n_csi_rsRLMConfigBitmap_r16),
		}
		ie.csi_rsRLMConfigBitmap_r16 = &tmp_bitstring
	}
	return nil
}
