package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SharedSpectrumChAccessParamsPerBand_v1640 struct {
	csi_RSRP_AndRSRQ_MeasWithSSB_r16    *SharedSpectrumChAccessParamsPerBand_v1640_csi_RSRP_AndRSRQ_MeasWithSSB_r16    `optional`
	csi_RSRP_AndRSRQ_MeasWithoutSSB_r16 *SharedSpectrumChAccessParamsPerBand_v1640_csi_RSRP_AndRSRQ_MeasWithoutSSB_r16 `optional`
	csi_SINR_Meas_r16                   *SharedSpectrumChAccessParamsPerBand_v1640_csi_SINR_Meas_r16                   `optional`
	ssb_AndCSI_RS_RLM_r16               *SharedSpectrumChAccessParamsPerBand_v1640_ssb_AndCSI_RS_RLM_r16               `optional`
	csi_RS_CFRA_ForHO_r16               *SharedSpectrumChAccessParamsPerBand_v1640_csi_RS_CFRA_ForHO_r16               `optional`
}

func (ie *SharedSpectrumChAccessParamsPerBand_v1640) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.csi_RSRP_AndRSRQ_MeasWithSSB_r16 != nil, ie.csi_RSRP_AndRSRQ_MeasWithoutSSB_r16 != nil, ie.csi_SINR_Meas_r16 != nil, ie.ssb_AndCSI_RS_RLM_r16 != nil, ie.csi_RS_CFRA_ForHO_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.csi_RSRP_AndRSRQ_MeasWithSSB_r16 != nil {
		if err = ie.csi_RSRP_AndRSRQ_MeasWithSSB_r16.Encode(w); err != nil {
			return utils.WrapError("Encode csi_RSRP_AndRSRQ_MeasWithSSB_r16", err)
		}
	}
	if ie.csi_RSRP_AndRSRQ_MeasWithoutSSB_r16 != nil {
		if err = ie.csi_RSRP_AndRSRQ_MeasWithoutSSB_r16.Encode(w); err != nil {
			return utils.WrapError("Encode csi_RSRP_AndRSRQ_MeasWithoutSSB_r16", err)
		}
	}
	if ie.csi_SINR_Meas_r16 != nil {
		if err = ie.csi_SINR_Meas_r16.Encode(w); err != nil {
			return utils.WrapError("Encode csi_SINR_Meas_r16", err)
		}
	}
	if ie.ssb_AndCSI_RS_RLM_r16 != nil {
		if err = ie.ssb_AndCSI_RS_RLM_r16.Encode(w); err != nil {
			return utils.WrapError("Encode ssb_AndCSI_RS_RLM_r16", err)
		}
	}
	if ie.csi_RS_CFRA_ForHO_r16 != nil {
		if err = ie.csi_RS_CFRA_ForHO_r16.Encode(w); err != nil {
			return utils.WrapError("Encode csi_RS_CFRA_ForHO_r16", err)
		}
	}
	return nil
}

func (ie *SharedSpectrumChAccessParamsPerBand_v1640) Decode(r *uper.UperReader) error {
	var err error
	var csi_RSRP_AndRSRQ_MeasWithSSB_r16Present bool
	if csi_RSRP_AndRSRQ_MeasWithSSB_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var csi_RSRP_AndRSRQ_MeasWithoutSSB_r16Present bool
	if csi_RSRP_AndRSRQ_MeasWithoutSSB_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var csi_SINR_Meas_r16Present bool
	if csi_SINR_Meas_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ssb_AndCSI_RS_RLM_r16Present bool
	if ssb_AndCSI_RS_RLM_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var csi_RS_CFRA_ForHO_r16Present bool
	if csi_RS_CFRA_ForHO_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if csi_RSRP_AndRSRQ_MeasWithSSB_r16Present {
		ie.csi_RSRP_AndRSRQ_MeasWithSSB_r16 = new(SharedSpectrumChAccessParamsPerBand_v1640_csi_RSRP_AndRSRQ_MeasWithSSB_r16)
		if err = ie.csi_RSRP_AndRSRQ_MeasWithSSB_r16.Decode(r); err != nil {
			return utils.WrapError("Decode csi_RSRP_AndRSRQ_MeasWithSSB_r16", err)
		}
	}
	if csi_RSRP_AndRSRQ_MeasWithoutSSB_r16Present {
		ie.csi_RSRP_AndRSRQ_MeasWithoutSSB_r16 = new(SharedSpectrumChAccessParamsPerBand_v1640_csi_RSRP_AndRSRQ_MeasWithoutSSB_r16)
		if err = ie.csi_RSRP_AndRSRQ_MeasWithoutSSB_r16.Decode(r); err != nil {
			return utils.WrapError("Decode csi_RSRP_AndRSRQ_MeasWithoutSSB_r16", err)
		}
	}
	if csi_SINR_Meas_r16Present {
		ie.csi_SINR_Meas_r16 = new(SharedSpectrumChAccessParamsPerBand_v1640_csi_SINR_Meas_r16)
		if err = ie.csi_SINR_Meas_r16.Decode(r); err != nil {
			return utils.WrapError("Decode csi_SINR_Meas_r16", err)
		}
	}
	if ssb_AndCSI_RS_RLM_r16Present {
		ie.ssb_AndCSI_RS_RLM_r16 = new(SharedSpectrumChAccessParamsPerBand_v1640_ssb_AndCSI_RS_RLM_r16)
		if err = ie.ssb_AndCSI_RS_RLM_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ssb_AndCSI_RS_RLM_r16", err)
		}
	}
	if csi_RS_CFRA_ForHO_r16Present {
		ie.csi_RS_CFRA_ForHO_r16 = new(SharedSpectrumChAccessParamsPerBand_v1640_csi_RS_CFRA_ForHO_r16)
		if err = ie.csi_RS_CFRA_ForHO_r16.Decode(r); err != nil {
			return utils.WrapError("Decode csi_RS_CFRA_ForHO_r16", err)
		}
	}
	return nil
}
