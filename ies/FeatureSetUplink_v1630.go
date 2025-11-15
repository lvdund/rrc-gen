package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type FeatureSetUplink_v1630 struct {
	offsetSRS_CB_PUSCH_Ant_Switch_fr1_r16                    *FeatureSetUplink_v1630_offsetSRS_CB_PUSCH_Ant_Switch_fr1_r16                    `optional`
	offsetSRS_CB_PUSCH_PDCCH_MonitorSingleOcc_fr1_r16        *FeatureSetUplink_v1630_offsetSRS_CB_PUSCH_PDCCH_MonitorSingleOcc_fr1_r16        `optional`
	offsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithoutGap_fr1_r16 *FeatureSetUplink_v1630_offsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithoutGap_fr1_r16 `optional`
	offsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithGap_fr1_r16    *FeatureSetUplink_v1630_offsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithGap_fr1_r16    `optional`
	dummy                                                    *FeatureSetUplink_v1630_dummy                                                    `optional`
	partialCancellationPUCCH_PUSCH_PRACH_TX_r16              *FeatureSetUplink_v1630_partialCancellationPUCCH_PUSCH_PRACH_TX_r16              `optional`
}

func (ie *FeatureSetUplink_v1630) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.offsetSRS_CB_PUSCH_Ant_Switch_fr1_r16 != nil, ie.offsetSRS_CB_PUSCH_PDCCH_MonitorSingleOcc_fr1_r16 != nil, ie.offsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithoutGap_fr1_r16 != nil, ie.offsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithGap_fr1_r16 != nil, ie.dummy != nil, ie.partialCancellationPUCCH_PUSCH_PRACH_TX_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.offsetSRS_CB_PUSCH_Ant_Switch_fr1_r16 != nil {
		if err = ie.offsetSRS_CB_PUSCH_Ant_Switch_fr1_r16.Encode(w); err != nil {
			return utils.WrapError("Encode offsetSRS_CB_PUSCH_Ant_Switch_fr1_r16", err)
		}
	}
	if ie.offsetSRS_CB_PUSCH_PDCCH_MonitorSingleOcc_fr1_r16 != nil {
		if err = ie.offsetSRS_CB_PUSCH_PDCCH_MonitorSingleOcc_fr1_r16.Encode(w); err != nil {
			return utils.WrapError("Encode offsetSRS_CB_PUSCH_PDCCH_MonitorSingleOcc_fr1_r16", err)
		}
	}
	if ie.offsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithoutGap_fr1_r16 != nil {
		if err = ie.offsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithoutGap_fr1_r16.Encode(w); err != nil {
			return utils.WrapError("Encode offsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithoutGap_fr1_r16", err)
		}
	}
	if ie.offsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithGap_fr1_r16 != nil {
		if err = ie.offsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithGap_fr1_r16.Encode(w); err != nil {
			return utils.WrapError("Encode offsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithGap_fr1_r16", err)
		}
	}
	if ie.dummy != nil {
		if err = ie.dummy.Encode(w); err != nil {
			return utils.WrapError("Encode dummy", err)
		}
	}
	if ie.partialCancellationPUCCH_PUSCH_PRACH_TX_r16 != nil {
		if err = ie.partialCancellationPUCCH_PUSCH_PRACH_TX_r16.Encode(w); err != nil {
			return utils.WrapError("Encode partialCancellationPUCCH_PUSCH_PRACH_TX_r16", err)
		}
	}
	return nil
}

func (ie *FeatureSetUplink_v1630) Decode(r *uper.UperReader) error {
	var err error
	var offsetSRS_CB_PUSCH_Ant_Switch_fr1_r16Present bool
	if offsetSRS_CB_PUSCH_Ant_Switch_fr1_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var offsetSRS_CB_PUSCH_PDCCH_MonitorSingleOcc_fr1_r16Present bool
	if offsetSRS_CB_PUSCH_PDCCH_MonitorSingleOcc_fr1_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var offsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithoutGap_fr1_r16Present bool
	if offsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithoutGap_fr1_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var offsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithGap_fr1_r16Present bool
	if offsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithGap_fr1_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var dummyPresent bool
	if dummyPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var partialCancellationPUCCH_PUSCH_PRACH_TX_r16Present bool
	if partialCancellationPUCCH_PUSCH_PRACH_TX_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if offsetSRS_CB_PUSCH_Ant_Switch_fr1_r16Present {
		ie.offsetSRS_CB_PUSCH_Ant_Switch_fr1_r16 = new(FeatureSetUplink_v1630_offsetSRS_CB_PUSCH_Ant_Switch_fr1_r16)
		if err = ie.offsetSRS_CB_PUSCH_Ant_Switch_fr1_r16.Decode(r); err != nil {
			return utils.WrapError("Decode offsetSRS_CB_PUSCH_Ant_Switch_fr1_r16", err)
		}
	}
	if offsetSRS_CB_PUSCH_PDCCH_MonitorSingleOcc_fr1_r16Present {
		ie.offsetSRS_CB_PUSCH_PDCCH_MonitorSingleOcc_fr1_r16 = new(FeatureSetUplink_v1630_offsetSRS_CB_PUSCH_PDCCH_MonitorSingleOcc_fr1_r16)
		if err = ie.offsetSRS_CB_PUSCH_PDCCH_MonitorSingleOcc_fr1_r16.Decode(r); err != nil {
			return utils.WrapError("Decode offsetSRS_CB_PUSCH_PDCCH_MonitorSingleOcc_fr1_r16", err)
		}
	}
	if offsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithoutGap_fr1_r16Present {
		ie.offsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithoutGap_fr1_r16 = new(FeatureSetUplink_v1630_offsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithoutGap_fr1_r16)
		if err = ie.offsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithoutGap_fr1_r16.Decode(r); err != nil {
			return utils.WrapError("Decode offsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithoutGap_fr1_r16", err)
		}
	}
	if offsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithGap_fr1_r16Present {
		ie.offsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithGap_fr1_r16 = new(FeatureSetUplink_v1630_offsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithGap_fr1_r16)
		if err = ie.offsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithGap_fr1_r16.Decode(r); err != nil {
			return utils.WrapError("Decode offsetSRS_CB_PUSCH_PDCCH_MonitorAnyOccWithGap_fr1_r16", err)
		}
	}
	if dummyPresent {
		ie.dummy = new(FeatureSetUplink_v1630_dummy)
		if err = ie.dummy.Decode(r); err != nil {
			return utils.WrapError("Decode dummy", err)
		}
	}
	if partialCancellationPUCCH_PUSCH_PRACH_TX_r16Present {
		ie.partialCancellationPUCCH_PUSCH_PRACH_TX_r16 = new(FeatureSetUplink_v1630_partialCancellationPUCCH_PUSCH_PRACH_TX_r16)
		if err = ie.partialCancellationPUCCH_PUSCH_PRACH_TX_r16.Decode(r); err != nil {
			return utils.WrapError("Decode partialCancellationPUCCH_PUSCH_PRACH_TX_r16", err)
		}
	}
	return nil
}
