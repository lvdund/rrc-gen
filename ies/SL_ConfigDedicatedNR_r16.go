package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_ConfigDedicatedNR_r16 struct {
	sl_PHY_MAC_RLC_Config_r16          *SL_PHY_MAC_RLC_Config_r16         `optional`
	sl_RadioBearerToReleaseList_r16    []SLRB_Uu_ConfigIndex_r16          `lb:1,ub:maxNrofSLRB_r16,optional`
	sl_RadioBearerToAddModList_r16     []SL_RadioBearerConfig_r16         `lb:1,ub:maxNrofSLRB_r16,optional`
	sl_MeasConfigInfoToReleaseList_r16 []SL_DestinationIndex_r16          `lb:1,ub:maxNrofSL_Dest_r16,optional`
	sl_MeasConfigInfoToAddModList_r16  []SL_MeasConfigInfo_r16            `lb:1,ub:maxNrofSL_Dest_r16,optional`
	t400_r16                           *SL_ConfigDedicatedNR_r16_t400_r16 `optional`
	sl_PHY_MAC_RLC_Config_v1700        *SL_PHY_MAC_RLC_Config_v1700       `optional,ext-1,setuprelease`
	sl_DiscConfig_r17                  *SL_DiscConfig_r17                 `optional,ext-1,setuprelease`
}

func (ie *SL_ConfigDedicatedNR_r16) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.sl_PHY_MAC_RLC_Config_v1700 != nil || ie.sl_DiscConfig_r17 != nil
	preambleBits := []bool{hasExtensions, ie.sl_PHY_MAC_RLC_Config_r16 != nil, len(ie.sl_RadioBearerToReleaseList_r16) > 0, len(ie.sl_RadioBearerToAddModList_r16) > 0, len(ie.sl_MeasConfigInfoToReleaseList_r16) > 0, len(ie.sl_MeasConfigInfoToAddModList_r16) > 0, ie.t400_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.sl_PHY_MAC_RLC_Config_r16 != nil {
		if err = ie.sl_PHY_MAC_RLC_Config_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_PHY_MAC_RLC_Config_r16", err)
		}
	}
	if len(ie.sl_RadioBearerToReleaseList_r16) > 0 {
		tmp_sl_RadioBearerToReleaseList_r16 := utils.NewSequence[*SLRB_Uu_ConfigIndex_r16]([]*SLRB_Uu_ConfigIndex_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofSLRB_r16}, false)
		for _, i := range ie.sl_RadioBearerToReleaseList_r16 {
			tmp_sl_RadioBearerToReleaseList_r16.Value = append(tmp_sl_RadioBearerToReleaseList_r16.Value, &i)
		}
		if err = tmp_sl_RadioBearerToReleaseList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_RadioBearerToReleaseList_r16", err)
		}
	}
	if len(ie.sl_RadioBearerToAddModList_r16) > 0 {
		tmp_sl_RadioBearerToAddModList_r16 := utils.NewSequence[*SL_RadioBearerConfig_r16]([]*SL_RadioBearerConfig_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofSLRB_r16}, false)
		for _, i := range ie.sl_RadioBearerToAddModList_r16 {
			tmp_sl_RadioBearerToAddModList_r16.Value = append(tmp_sl_RadioBearerToAddModList_r16.Value, &i)
		}
		if err = tmp_sl_RadioBearerToAddModList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_RadioBearerToAddModList_r16", err)
		}
	}
	if len(ie.sl_MeasConfigInfoToReleaseList_r16) > 0 {
		tmp_sl_MeasConfigInfoToReleaseList_r16 := utils.NewSequence[*SL_DestinationIndex_r16]([]*SL_DestinationIndex_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofSL_Dest_r16}, false)
		for _, i := range ie.sl_MeasConfigInfoToReleaseList_r16 {
			tmp_sl_MeasConfigInfoToReleaseList_r16.Value = append(tmp_sl_MeasConfigInfoToReleaseList_r16.Value, &i)
		}
		if err = tmp_sl_MeasConfigInfoToReleaseList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_MeasConfigInfoToReleaseList_r16", err)
		}
	}
	if len(ie.sl_MeasConfigInfoToAddModList_r16) > 0 {
		tmp_sl_MeasConfigInfoToAddModList_r16 := utils.NewSequence[*SL_MeasConfigInfo_r16]([]*SL_MeasConfigInfo_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofSL_Dest_r16}, false)
		for _, i := range ie.sl_MeasConfigInfoToAddModList_r16 {
			tmp_sl_MeasConfigInfoToAddModList_r16.Value = append(tmp_sl_MeasConfigInfoToAddModList_r16.Value, &i)
		}
		if err = tmp_sl_MeasConfigInfoToAddModList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_MeasConfigInfoToAddModList_r16", err)
		}
	}
	if ie.t400_r16 != nil {
		if err = ie.t400_r16.Encode(w); err != nil {
			return utils.WrapError("Encode t400_r16", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.sl_PHY_MAC_RLC_Config_v1700 != nil || ie.sl_DiscConfig_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap SL_ConfigDedicatedNR_r16", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.sl_PHY_MAC_RLC_Config_v1700 != nil, ie.sl_DiscConfig_r17 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode sl_PHY_MAC_RLC_Config_v1700 optional
			if ie.sl_PHY_MAC_RLC_Config_v1700 != nil {
				tmp_sl_PHY_MAC_RLC_Config_v1700 := utils.SetupRelease[*SL_PHY_MAC_RLC_Config_v1700]{
					Setup: ie.sl_PHY_MAC_RLC_Config_v1700,
				}
				if err = tmp_sl_PHY_MAC_RLC_Config_v1700.Encode(extWriter); err != nil {
					return utils.WrapError("Encode sl_PHY_MAC_RLC_Config_v1700", err)
				}
			}
			// encode sl_DiscConfig_r17 optional
			if ie.sl_DiscConfig_r17 != nil {
				tmp_sl_DiscConfig_r17 := utils.SetupRelease[*SL_DiscConfig_r17]{
					Setup: ie.sl_DiscConfig_r17,
				}
				if err = tmp_sl_DiscConfig_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode sl_DiscConfig_r17", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}
	}
	return nil
}

func (ie *SL_ConfigDedicatedNR_r16) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_PHY_MAC_RLC_Config_r16Present bool
	if sl_PHY_MAC_RLC_Config_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_RadioBearerToReleaseList_r16Present bool
	if sl_RadioBearerToReleaseList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_RadioBearerToAddModList_r16Present bool
	if sl_RadioBearerToAddModList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_MeasConfigInfoToReleaseList_r16Present bool
	if sl_MeasConfigInfoToReleaseList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_MeasConfigInfoToAddModList_r16Present bool
	if sl_MeasConfigInfoToAddModList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var t400_r16Present bool
	if t400_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if sl_PHY_MAC_RLC_Config_r16Present {
		ie.sl_PHY_MAC_RLC_Config_r16 = new(SL_PHY_MAC_RLC_Config_r16)
		if err = ie.sl_PHY_MAC_RLC_Config_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_PHY_MAC_RLC_Config_r16", err)
		}
	}
	if sl_RadioBearerToReleaseList_r16Present {
		tmp_sl_RadioBearerToReleaseList_r16 := utils.NewSequence[*SLRB_Uu_ConfigIndex_r16]([]*SLRB_Uu_ConfigIndex_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofSLRB_r16}, false)
		fn_sl_RadioBearerToReleaseList_r16 := func() *SLRB_Uu_ConfigIndex_r16 {
			return new(SLRB_Uu_ConfigIndex_r16)
		}
		if err = tmp_sl_RadioBearerToReleaseList_r16.Decode(r, fn_sl_RadioBearerToReleaseList_r16); err != nil {
			return utils.WrapError("Decode sl_RadioBearerToReleaseList_r16", err)
		}
		ie.sl_RadioBearerToReleaseList_r16 = []SLRB_Uu_ConfigIndex_r16{}
		for _, i := range tmp_sl_RadioBearerToReleaseList_r16.Value {
			ie.sl_RadioBearerToReleaseList_r16 = append(ie.sl_RadioBearerToReleaseList_r16, *i)
		}
	}
	if sl_RadioBearerToAddModList_r16Present {
		tmp_sl_RadioBearerToAddModList_r16 := utils.NewSequence[*SL_RadioBearerConfig_r16]([]*SL_RadioBearerConfig_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofSLRB_r16}, false)
		fn_sl_RadioBearerToAddModList_r16 := func() *SL_RadioBearerConfig_r16 {
			return new(SL_RadioBearerConfig_r16)
		}
		if err = tmp_sl_RadioBearerToAddModList_r16.Decode(r, fn_sl_RadioBearerToAddModList_r16); err != nil {
			return utils.WrapError("Decode sl_RadioBearerToAddModList_r16", err)
		}
		ie.sl_RadioBearerToAddModList_r16 = []SL_RadioBearerConfig_r16{}
		for _, i := range tmp_sl_RadioBearerToAddModList_r16.Value {
			ie.sl_RadioBearerToAddModList_r16 = append(ie.sl_RadioBearerToAddModList_r16, *i)
		}
	}
	if sl_MeasConfigInfoToReleaseList_r16Present {
		tmp_sl_MeasConfigInfoToReleaseList_r16 := utils.NewSequence[*SL_DestinationIndex_r16]([]*SL_DestinationIndex_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofSL_Dest_r16}, false)
		fn_sl_MeasConfigInfoToReleaseList_r16 := func() *SL_DestinationIndex_r16 {
			return new(SL_DestinationIndex_r16)
		}
		if err = tmp_sl_MeasConfigInfoToReleaseList_r16.Decode(r, fn_sl_MeasConfigInfoToReleaseList_r16); err != nil {
			return utils.WrapError("Decode sl_MeasConfigInfoToReleaseList_r16", err)
		}
		ie.sl_MeasConfigInfoToReleaseList_r16 = []SL_DestinationIndex_r16{}
		for _, i := range tmp_sl_MeasConfigInfoToReleaseList_r16.Value {
			ie.sl_MeasConfigInfoToReleaseList_r16 = append(ie.sl_MeasConfigInfoToReleaseList_r16, *i)
		}
	}
	if sl_MeasConfigInfoToAddModList_r16Present {
		tmp_sl_MeasConfigInfoToAddModList_r16 := utils.NewSequence[*SL_MeasConfigInfo_r16]([]*SL_MeasConfigInfo_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofSL_Dest_r16}, false)
		fn_sl_MeasConfigInfoToAddModList_r16 := func() *SL_MeasConfigInfo_r16 {
			return new(SL_MeasConfigInfo_r16)
		}
		if err = tmp_sl_MeasConfigInfoToAddModList_r16.Decode(r, fn_sl_MeasConfigInfoToAddModList_r16); err != nil {
			return utils.WrapError("Decode sl_MeasConfigInfoToAddModList_r16", err)
		}
		ie.sl_MeasConfigInfoToAddModList_r16 = []SL_MeasConfigInfo_r16{}
		for _, i := range tmp_sl_MeasConfigInfoToAddModList_r16.Value {
			ie.sl_MeasConfigInfoToAddModList_r16 = append(ie.sl_MeasConfigInfoToAddModList_r16, *i)
		}
	}
	if t400_r16Present {
		ie.t400_r16 = new(SL_ConfigDedicatedNR_r16_t400_r16)
		if err = ie.t400_r16.Decode(r); err != nil {
			return utils.WrapError("Decode t400_r16", err)
		}
	}

	if extensionBit {
		// Read extension bitmap: 1 bits for 1 extension groups
		extBitmap, err := r.ReadExtBitMap()
		if err != nil {
			return err
		}

		// decode extension group 1
		if len(extBitmap) > 0 && extBitmap[0] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			sl_PHY_MAC_RLC_Config_v1700Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			sl_DiscConfig_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode sl_PHY_MAC_RLC_Config_v1700 optional
			if sl_PHY_MAC_RLC_Config_v1700Present {
				tmp_sl_PHY_MAC_RLC_Config_v1700 := utils.SetupRelease[*SL_PHY_MAC_RLC_Config_v1700]{}
				if err = tmp_sl_PHY_MAC_RLC_Config_v1700.Decode(extReader); err != nil {
					return utils.WrapError("Decode sl_PHY_MAC_RLC_Config_v1700", err)
				}
				ie.sl_PHY_MAC_RLC_Config_v1700 = tmp_sl_PHY_MAC_RLC_Config_v1700.Setup
			}
			// decode sl_DiscConfig_r17 optional
			if sl_DiscConfig_r17Present {
				tmp_sl_DiscConfig_r17 := utils.SetupRelease[*SL_DiscConfig_r17]{}
				if err = tmp_sl_DiscConfig_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode sl_DiscConfig_r17", err)
				}
				ie.sl_DiscConfig_r17 = tmp_sl_DiscConfig_r17.Setup
			}
		}
	}
	return nil
}
