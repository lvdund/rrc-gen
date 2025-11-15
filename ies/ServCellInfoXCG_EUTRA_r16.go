package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ServCellInfoXCG_EUTRA_r16 struct {
	dl_CarrierFreq_EUTRA_r16        *ARFCN_ValueEUTRA                `optional`
	ul_CarrierFreq_EUTRA_r16        *ARFCN_ValueEUTRA                `optional`
	transmissionBandwidth_EUTRA_r16 *TransmissionBandwidth_EUTRA_r16 `optional`
}

func (ie *ServCellInfoXCG_EUTRA_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.dl_CarrierFreq_EUTRA_r16 != nil, ie.ul_CarrierFreq_EUTRA_r16 != nil, ie.transmissionBandwidth_EUTRA_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.dl_CarrierFreq_EUTRA_r16 != nil {
		if err = ie.dl_CarrierFreq_EUTRA_r16.Encode(w); err != nil {
			return utils.WrapError("Encode dl_CarrierFreq_EUTRA_r16", err)
		}
	}
	if ie.ul_CarrierFreq_EUTRA_r16 != nil {
		if err = ie.ul_CarrierFreq_EUTRA_r16.Encode(w); err != nil {
			return utils.WrapError("Encode ul_CarrierFreq_EUTRA_r16", err)
		}
	}
	if ie.transmissionBandwidth_EUTRA_r16 != nil {
		if err = ie.transmissionBandwidth_EUTRA_r16.Encode(w); err != nil {
			return utils.WrapError("Encode transmissionBandwidth_EUTRA_r16", err)
		}
	}
	return nil
}

func (ie *ServCellInfoXCG_EUTRA_r16) Decode(r *uper.UperReader) error {
	var err error
	var dl_CarrierFreq_EUTRA_r16Present bool
	if dl_CarrierFreq_EUTRA_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ul_CarrierFreq_EUTRA_r16Present bool
	if ul_CarrierFreq_EUTRA_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var transmissionBandwidth_EUTRA_r16Present bool
	if transmissionBandwidth_EUTRA_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if dl_CarrierFreq_EUTRA_r16Present {
		ie.dl_CarrierFreq_EUTRA_r16 = new(ARFCN_ValueEUTRA)
		if err = ie.dl_CarrierFreq_EUTRA_r16.Decode(r); err != nil {
			return utils.WrapError("Decode dl_CarrierFreq_EUTRA_r16", err)
		}
	}
	if ul_CarrierFreq_EUTRA_r16Present {
		ie.ul_CarrierFreq_EUTRA_r16 = new(ARFCN_ValueEUTRA)
		if err = ie.ul_CarrierFreq_EUTRA_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ul_CarrierFreq_EUTRA_r16", err)
		}
	}
	if transmissionBandwidth_EUTRA_r16Present {
		ie.transmissionBandwidth_EUTRA_r16 = new(TransmissionBandwidth_EUTRA_r16)
		if err = ie.transmissionBandwidth_EUTRA_r16.Decode(r); err != nil {
			return utils.WrapError("Decode transmissionBandwidth_EUTRA_r16", err)
		}
	}
	return nil
}
