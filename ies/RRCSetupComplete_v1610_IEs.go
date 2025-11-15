package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RRCSetupComplete_v1610_IEs struct {
	iab_NodeIndication_r16       *RRCSetupComplete_v1610_IEs_iab_NodeIndication_r16   `optional`
	idleMeasAvailable_r16        *RRCSetupComplete_v1610_IEs_idleMeasAvailable_r16    `optional`
	ue_MeasurementsAvailable_r16 *UE_MeasurementsAvailable_r16                        `optional`
	mobilityHistoryAvail_r16     *RRCSetupComplete_v1610_IEs_mobilityHistoryAvail_r16 `optional`
	mobilityState_r16            *RRCSetupComplete_v1610_IEs_mobilityState_r16        `optional`
	nonCriticalExtension         *RRCSetupComplete_v1690_IEs                          `optional`
}

func (ie *RRCSetupComplete_v1610_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.iab_NodeIndication_r16 != nil, ie.idleMeasAvailable_r16 != nil, ie.ue_MeasurementsAvailable_r16 != nil, ie.mobilityHistoryAvail_r16 != nil, ie.mobilityState_r16 != nil, ie.nonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.iab_NodeIndication_r16 != nil {
		if err = ie.iab_NodeIndication_r16.Encode(w); err != nil {
			return utils.WrapError("Encode iab_NodeIndication_r16", err)
		}
	}
	if ie.idleMeasAvailable_r16 != nil {
		if err = ie.idleMeasAvailable_r16.Encode(w); err != nil {
			return utils.WrapError("Encode idleMeasAvailable_r16", err)
		}
	}
	if ie.ue_MeasurementsAvailable_r16 != nil {
		if err = ie.ue_MeasurementsAvailable_r16.Encode(w); err != nil {
			return utils.WrapError("Encode ue_MeasurementsAvailable_r16", err)
		}
	}
	if ie.mobilityHistoryAvail_r16 != nil {
		if err = ie.mobilityHistoryAvail_r16.Encode(w); err != nil {
			return utils.WrapError("Encode mobilityHistoryAvail_r16", err)
		}
	}
	if ie.mobilityState_r16 != nil {
		if err = ie.mobilityState_r16.Encode(w); err != nil {
			return utils.WrapError("Encode mobilityState_r16", err)
		}
	}
	if ie.nonCriticalExtension != nil {
		if err = ie.nonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode nonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *RRCSetupComplete_v1610_IEs) Decode(r *uper.UperReader) error {
	var err error
	var iab_NodeIndication_r16Present bool
	if iab_NodeIndication_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var idleMeasAvailable_r16Present bool
	if idleMeasAvailable_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ue_MeasurementsAvailable_r16Present bool
	if ue_MeasurementsAvailable_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var mobilityHistoryAvail_r16Present bool
	if mobilityHistoryAvail_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var mobilityState_r16Present bool
	if mobilityState_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var nonCriticalExtensionPresent bool
	if nonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if iab_NodeIndication_r16Present {
		ie.iab_NodeIndication_r16 = new(RRCSetupComplete_v1610_IEs_iab_NodeIndication_r16)
		if err = ie.iab_NodeIndication_r16.Decode(r); err != nil {
			return utils.WrapError("Decode iab_NodeIndication_r16", err)
		}
	}
	if idleMeasAvailable_r16Present {
		ie.idleMeasAvailable_r16 = new(RRCSetupComplete_v1610_IEs_idleMeasAvailable_r16)
		if err = ie.idleMeasAvailable_r16.Decode(r); err != nil {
			return utils.WrapError("Decode idleMeasAvailable_r16", err)
		}
	}
	if ue_MeasurementsAvailable_r16Present {
		ie.ue_MeasurementsAvailable_r16 = new(UE_MeasurementsAvailable_r16)
		if err = ie.ue_MeasurementsAvailable_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ue_MeasurementsAvailable_r16", err)
		}
	}
	if mobilityHistoryAvail_r16Present {
		ie.mobilityHistoryAvail_r16 = new(RRCSetupComplete_v1610_IEs_mobilityHistoryAvail_r16)
		if err = ie.mobilityHistoryAvail_r16.Decode(r); err != nil {
			return utils.WrapError("Decode mobilityHistoryAvail_r16", err)
		}
	}
	if mobilityState_r16Present {
		ie.mobilityState_r16 = new(RRCSetupComplete_v1610_IEs_mobilityState_r16)
		if err = ie.mobilityState_r16.Decode(r); err != nil {
			return utils.WrapError("Decode mobilityState_r16", err)
		}
	}
	if nonCriticalExtensionPresent {
		ie.nonCriticalExtension = new(RRCSetupComplete_v1690_IEs)
		if err = ie.nonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode nonCriticalExtension", err)
		}
	}
	return nil
}
