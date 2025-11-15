package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type FeatureSetDownlinkPerCC_v1730 struct {
	intraSlotTDM_UnicastGroupCommonPDSCH_r17 *FeatureSetDownlinkPerCC_v1730_intraSlotTDM_UnicastGroupCommonPDSCH_r17 `optional`
	sps_MulticastSCell_r17                   *FeatureSetDownlinkPerCC_v1730_sps_MulticastSCell_r17                   `optional`
	sps_MulticastSCellMultiConfig_r17        *int64                                                                  `lb:1,ub:8,optional`
	dci_BroadcastWith16Repetitions_r17       *FeatureSetDownlinkPerCC_v1730_dci_BroadcastWith16Repetitions_r17       `optional`
}

func (ie *FeatureSetDownlinkPerCC_v1730) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.intraSlotTDM_UnicastGroupCommonPDSCH_r17 != nil, ie.sps_MulticastSCell_r17 != nil, ie.sps_MulticastSCellMultiConfig_r17 != nil, ie.dci_BroadcastWith16Repetitions_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.intraSlotTDM_UnicastGroupCommonPDSCH_r17 != nil {
		if err = ie.intraSlotTDM_UnicastGroupCommonPDSCH_r17.Encode(w); err != nil {
			return utils.WrapError("Encode intraSlotTDM_UnicastGroupCommonPDSCH_r17", err)
		}
	}
	if ie.sps_MulticastSCell_r17 != nil {
		if err = ie.sps_MulticastSCell_r17.Encode(w); err != nil {
			return utils.WrapError("Encode sps_MulticastSCell_r17", err)
		}
	}
	if ie.sps_MulticastSCellMultiConfig_r17 != nil {
		if err = w.WriteInteger(*ie.sps_MulticastSCellMultiConfig_r17, &uper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
			return utils.WrapError("Encode sps_MulticastSCellMultiConfig_r17", err)
		}
	}
	if ie.dci_BroadcastWith16Repetitions_r17 != nil {
		if err = ie.dci_BroadcastWith16Repetitions_r17.Encode(w); err != nil {
			return utils.WrapError("Encode dci_BroadcastWith16Repetitions_r17", err)
		}
	}
	return nil
}

func (ie *FeatureSetDownlinkPerCC_v1730) Decode(r *uper.UperReader) error {
	var err error
	var intraSlotTDM_UnicastGroupCommonPDSCH_r17Present bool
	if intraSlotTDM_UnicastGroupCommonPDSCH_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sps_MulticastSCell_r17Present bool
	if sps_MulticastSCell_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sps_MulticastSCellMultiConfig_r17Present bool
	if sps_MulticastSCellMultiConfig_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var dci_BroadcastWith16Repetitions_r17Present bool
	if dci_BroadcastWith16Repetitions_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if intraSlotTDM_UnicastGroupCommonPDSCH_r17Present {
		ie.intraSlotTDM_UnicastGroupCommonPDSCH_r17 = new(FeatureSetDownlinkPerCC_v1730_intraSlotTDM_UnicastGroupCommonPDSCH_r17)
		if err = ie.intraSlotTDM_UnicastGroupCommonPDSCH_r17.Decode(r); err != nil {
			return utils.WrapError("Decode intraSlotTDM_UnicastGroupCommonPDSCH_r17", err)
		}
	}
	if sps_MulticastSCell_r17Present {
		ie.sps_MulticastSCell_r17 = new(FeatureSetDownlinkPerCC_v1730_sps_MulticastSCell_r17)
		if err = ie.sps_MulticastSCell_r17.Decode(r); err != nil {
			return utils.WrapError("Decode sps_MulticastSCell_r17", err)
		}
	}
	if sps_MulticastSCellMultiConfig_r17Present {
		var tmp_int_sps_MulticastSCellMultiConfig_r17 int64
		if tmp_int_sps_MulticastSCellMultiConfig_r17, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
			return utils.WrapError("Decode sps_MulticastSCellMultiConfig_r17", err)
		}
		ie.sps_MulticastSCellMultiConfig_r17 = &tmp_int_sps_MulticastSCellMultiConfig_r17
	}
	if dci_BroadcastWith16Repetitions_r17Present {
		ie.dci_BroadcastWith16Repetitions_r17 = new(FeatureSetDownlinkPerCC_v1730_dci_BroadcastWith16Repetitions_r17)
		if err = ie.dci_BroadcastWith16Repetitions_r17.Decode(r); err != nil {
			return utils.WrapError("Decode dci_BroadcastWith16Repetitions_r17", err)
		}
	}
	return nil
}
