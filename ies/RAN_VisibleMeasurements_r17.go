package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RAN_VisibleMeasurements_r17 struct {
	appLayerBufferLevelList_r17     []AppLayerBufferLevel_r17 `lb:1,ub:8,optional`
	playoutDelayForMediaStartup_r17 *int64                    `lb:0,ub:30000,optional`
	pdu_SessionIdList_r17           []PDU_SessionID           `lb:1,ub:maxNrofPDU_Sessions_r17,optional`
}

func (ie *RAN_VisibleMeasurements_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{len(ie.appLayerBufferLevelList_r17) > 0, ie.playoutDelayForMediaStartup_r17 != nil, len(ie.pdu_SessionIdList_r17) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.appLayerBufferLevelList_r17) > 0 {
		tmp_appLayerBufferLevelList_r17 := utils.NewSequence[*AppLayerBufferLevel_r17]([]*AppLayerBufferLevel_r17{}, uper.Constraint{Lb: 1, Ub: 8}, false)
		for _, i := range ie.appLayerBufferLevelList_r17 {
			tmp_appLayerBufferLevelList_r17.Value = append(tmp_appLayerBufferLevelList_r17.Value, &i)
		}
		if err = tmp_appLayerBufferLevelList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode appLayerBufferLevelList_r17", err)
		}
	}
	if ie.playoutDelayForMediaStartup_r17 != nil {
		if err = w.WriteInteger(*ie.playoutDelayForMediaStartup_r17, &uper.Constraint{Lb: 0, Ub: 30000}, false); err != nil {
			return utils.WrapError("Encode playoutDelayForMediaStartup_r17", err)
		}
	}
	if len(ie.pdu_SessionIdList_r17) > 0 {
		tmp_pdu_SessionIdList_r17 := utils.NewSequence[*PDU_SessionID]([]*PDU_SessionID{}, uper.Constraint{Lb: 1, Ub: maxNrofPDU_Sessions_r17}, false)
		for _, i := range ie.pdu_SessionIdList_r17 {
			tmp_pdu_SessionIdList_r17.Value = append(tmp_pdu_SessionIdList_r17.Value, &i)
		}
		if err = tmp_pdu_SessionIdList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode pdu_SessionIdList_r17", err)
		}
	}
	return nil
}

func (ie *RAN_VisibleMeasurements_r17) Decode(r *uper.UperReader) error {
	var err error
	var appLayerBufferLevelList_r17Present bool
	if appLayerBufferLevelList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var playoutDelayForMediaStartup_r17Present bool
	if playoutDelayForMediaStartup_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var pdu_SessionIdList_r17Present bool
	if pdu_SessionIdList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if appLayerBufferLevelList_r17Present {
		tmp_appLayerBufferLevelList_r17 := utils.NewSequence[*AppLayerBufferLevel_r17]([]*AppLayerBufferLevel_r17{}, uper.Constraint{Lb: 1, Ub: 8}, false)
		fn_appLayerBufferLevelList_r17 := func() *AppLayerBufferLevel_r17 {
			return new(AppLayerBufferLevel_r17)
		}
		if err = tmp_appLayerBufferLevelList_r17.Decode(r, fn_appLayerBufferLevelList_r17); err != nil {
			return utils.WrapError("Decode appLayerBufferLevelList_r17", err)
		}
		ie.appLayerBufferLevelList_r17 = []AppLayerBufferLevel_r17{}
		for _, i := range tmp_appLayerBufferLevelList_r17.Value {
			ie.appLayerBufferLevelList_r17 = append(ie.appLayerBufferLevelList_r17, *i)
		}
	}
	if playoutDelayForMediaStartup_r17Present {
		var tmp_int_playoutDelayForMediaStartup_r17 int64
		if tmp_int_playoutDelayForMediaStartup_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 30000}, false); err != nil {
			return utils.WrapError("Decode playoutDelayForMediaStartup_r17", err)
		}
		ie.playoutDelayForMediaStartup_r17 = &tmp_int_playoutDelayForMediaStartup_r17
	}
	if pdu_SessionIdList_r17Present {
		tmp_pdu_SessionIdList_r17 := utils.NewSequence[*PDU_SessionID]([]*PDU_SessionID{}, uper.Constraint{Lb: 1, Ub: maxNrofPDU_Sessions_r17}, false)
		fn_pdu_SessionIdList_r17 := func() *PDU_SessionID {
			return new(PDU_SessionID)
		}
		if err = tmp_pdu_SessionIdList_r17.Decode(r, fn_pdu_SessionIdList_r17); err != nil {
			return utils.WrapError("Decode pdu_SessionIdList_r17", err)
		}
		ie.pdu_SessionIdList_r17 = []PDU_SessionID{}
		for _, i := range tmp_pdu_SessionIdList_r17.Value {
			ie.pdu_SessionIdList_r17 = append(ie.pdu_SessionIdList_r17, *i)
		}
	}
	return nil
}
