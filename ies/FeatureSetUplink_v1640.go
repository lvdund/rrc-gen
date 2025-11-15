package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type FeatureSetUplink_v1640 struct {
	twoHARQ_ACK_Codebook_type1_r16                            *SubSlot_Config_r16                                                               `optional`
	twoHARQ_ACK_Codebook_type2_r16                            *SubSlot_Config_r16                                                               `optional`
	offsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithSpanGap_fr1_r16 *FeatureSetUplink_v1640_offsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithSpanGap_fr1_r16 `optional`
}

func (ie *FeatureSetUplink_v1640) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.twoHARQ_ACK_Codebook_type1_r16 != nil, ie.twoHARQ_ACK_Codebook_type2_r16 != nil, ie.offsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithSpanGap_fr1_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.twoHARQ_ACK_Codebook_type1_r16 != nil {
		if err = ie.twoHARQ_ACK_Codebook_type1_r16.Encode(w); err != nil {
			return utils.WrapError("Encode twoHARQ_ACK_Codebook_type1_r16", err)
		}
	}
	if ie.twoHARQ_ACK_Codebook_type2_r16 != nil {
		if err = ie.twoHARQ_ACK_Codebook_type2_r16.Encode(w); err != nil {
			return utils.WrapError("Encode twoHARQ_ACK_Codebook_type2_r16", err)
		}
	}
	if ie.offsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithSpanGap_fr1_r16 != nil {
		if err = ie.offsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithSpanGap_fr1_r16.Encode(w); err != nil {
			return utils.WrapError("Encode offsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithSpanGap_fr1_r16", err)
		}
	}
	return nil
}

func (ie *FeatureSetUplink_v1640) Decode(r *uper.UperReader) error {
	var err error
	var twoHARQ_ACK_Codebook_type1_r16Present bool
	if twoHARQ_ACK_Codebook_type1_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var twoHARQ_ACK_Codebook_type2_r16Present bool
	if twoHARQ_ACK_Codebook_type2_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var offsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithSpanGap_fr1_r16Present bool
	if offsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithSpanGap_fr1_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if twoHARQ_ACK_Codebook_type1_r16Present {
		ie.twoHARQ_ACK_Codebook_type1_r16 = new(SubSlot_Config_r16)
		if err = ie.twoHARQ_ACK_Codebook_type1_r16.Decode(r); err != nil {
			return utils.WrapError("Decode twoHARQ_ACK_Codebook_type1_r16", err)
		}
	}
	if twoHARQ_ACK_Codebook_type2_r16Present {
		ie.twoHARQ_ACK_Codebook_type2_r16 = new(SubSlot_Config_r16)
		if err = ie.twoHARQ_ACK_Codebook_type2_r16.Decode(r); err != nil {
			return utils.WrapError("Decode twoHARQ_ACK_Codebook_type2_r16", err)
		}
	}
	if offsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithSpanGap_fr1_r16Present {
		ie.offsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithSpanGap_fr1_r16 = new(FeatureSetUplink_v1640_offsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithSpanGap_fr1_r16)
		if err = ie.offsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithSpanGap_fr1_r16.Decode(r); err != nil {
			return utils.WrapError("Decode offsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithSpanGap_fr1_r16", err)
		}
	}
	return nil
}
