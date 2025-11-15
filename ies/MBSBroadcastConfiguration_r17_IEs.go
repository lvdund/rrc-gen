package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MBSBroadcastConfiguration_r17_IEs struct {
	mbs_SessionInfoList_r17        *MBS_SessionInfoList_r17        `optional`
	mbs_NeighbourCellList_r17      *MBS_NeighbourCellList_r17      `optional`
	drx_ConfigPTM_List_r17         []DRX_ConfigPTM_r17             `lb:1,ub:maxNrofDRX_ConfigPTM_r17,optional`
	pdsch_ConfigMTCH_r17           *PDSCH_ConfigBroadcast_r17      `optional`
	mtch_SSB_MappingWindowList_r17 *MTCH_SSB_MappingWindowList_r17 `optional`
	lateNonCriticalExtension       *[]byte                         `optional`
	nonCriticalExtension           interface{}                     `optional`
}

func (ie *MBSBroadcastConfiguration_r17_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.mbs_SessionInfoList_r17 != nil, ie.mbs_NeighbourCellList_r17 != nil, len(ie.drx_ConfigPTM_List_r17) > 0, ie.pdsch_ConfigMTCH_r17 != nil, ie.mtch_SSB_MappingWindowList_r17 != nil, ie.lateNonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.mbs_SessionInfoList_r17 != nil {
		if err = ie.mbs_SessionInfoList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode mbs_SessionInfoList_r17", err)
		}
	}
	if ie.mbs_NeighbourCellList_r17 != nil {
		if err = ie.mbs_NeighbourCellList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode mbs_NeighbourCellList_r17", err)
		}
	}
	if len(ie.drx_ConfigPTM_List_r17) > 0 {
		tmp_drx_ConfigPTM_List_r17 := utils.NewSequence[*DRX_ConfigPTM_r17]([]*DRX_ConfigPTM_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofDRX_ConfigPTM_r17}, false)
		for _, i := range ie.drx_ConfigPTM_List_r17 {
			tmp_drx_ConfigPTM_List_r17.Value = append(tmp_drx_ConfigPTM_List_r17.Value, &i)
		}
		if err = tmp_drx_ConfigPTM_List_r17.Encode(w); err != nil {
			return utils.WrapError("Encode drx_ConfigPTM_List_r17", err)
		}
	}
	if ie.pdsch_ConfigMTCH_r17 != nil {
		if err = ie.pdsch_ConfigMTCH_r17.Encode(w); err != nil {
			return utils.WrapError("Encode pdsch_ConfigMTCH_r17", err)
		}
	}
	if ie.mtch_SSB_MappingWindowList_r17 != nil {
		if err = ie.mtch_SSB_MappingWindowList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode mtch_SSB_MappingWindowList_r17", err)
		}
	}
	if ie.lateNonCriticalExtension != nil {
		if err = w.WriteOctetString(*ie.lateNonCriticalExtension, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode lateNonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *MBSBroadcastConfiguration_r17_IEs) Decode(r *uper.UperReader) error {
	var err error
	var mbs_SessionInfoList_r17Present bool
	if mbs_SessionInfoList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var mbs_NeighbourCellList_r17Present bool
	if mbs_NeighbourCellList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var drx_ConfigPTM_List_r17Present bool
	if drx_ConfigPTM_List_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var pdsch_ConfigMTCH_r17Present bool
	if pdsch_ConfigMTCH_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var mtch_SSB_MappingWindowList_r17Present bool
	if mtch_SSB_MappingWindowList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var lateNonCriticalExtensionPresent bool
	if lateNonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if mbs_SessionInfoList_r17Present {
		ie.mbs_SessionInfoList_r17 = new(MBS_SessionInfoList_r17)
		if err = ie.mbs_SessionInfoList_r17.Decode(r); err != nil {
			return utils.WrapError("Decode mbs_SessionInfoList_r17", err)
		}
	}
	if mbs_NeighbourCellList_r17Present {
		ie.mbs_NeighbourCellList_r17 = new(MBS_NeighbourCellList_r17)
		if err = ie.mbs_NeighbourCellList_r17.Decode(r); err != nil {
			return utils.WrapError("Decode mbs_NeighbourCellList_r17", err)
		}
	}
	if drx_ConfigPTM_List_r17Present {
		tmp_drx_ConfigPTM_List_r17 := utils.NewSequence[*DRX_ConfigPTM_r17]([]*DRX_ConfigPTM_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofDRX_ConfigPTM_r17}, false)
		fn_drx_ConfigPTM_List_r17 := func() *DRX_ConfigPTM_r17 {
			return new(DRX_ConfigPTM_r17)
		}
		if err = tmp_drx_ConfigPTM_List_r17.Decode(r, fn_drx_ConfigPTM_List_r17); err != nil {
			return utils.WrapError("Decode drx_ConfigPTM_List_r17", err)
		}
		ie.drx_ConfigPTM_List_r17 = []DRX_ConfigPTM_r17{}
		for _, i := range tmp_drx_ConfigPTM_List_r17.Value {
			ie.drx_ConfigPTM_List_r17 = append(ie.drx_ConfigPTM_List_r17, *i)
		}
	}
	if pdsch_ConfigMTCH_r17Present {
		ie.pdsch_ConfigMTCH_r17 = new(PDSCH_ConfigBroadcast_r17)
		if err = ie.pdsch_ConfigMTCH_r17.Decode(r); err != nil {
			return utils.WrapError("Decode pdsch_ConfigMTCH_r17", err)
		}
	}
	if mtch_SSB_MappingWindowList_r17Present {
		ie.mtch_SSB_MappingWindowList_r17 = new(MTCH_SSB_MappingWindowList_r17)
		if err = ie.mtch_SSB_MappingWindowList_r17.Decode(r); err != nil {
			return utils.WrapError("Decode mtch_SSB_MappingWindowList_r17", err)
		}
	}
	if lateNonCriticalExtensionPresent {
		var tmp_os_lateNonCriticalExtension []byte
		if tmp_os_lateNonCriticalExtension, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode lateNonCriticalExtension", err)
		}
		ie.lateNonCriticalExtension = &tmp_os_lateNonCriticalExtension
	}
	return nil
}
