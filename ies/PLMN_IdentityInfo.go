package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PLMN_IdentityInfo struct {
	plmn_IdentityList          []PLMN_Identity                              `lb:1,ub:maxPLMN,madatory`
	trackingAreaCode           *TrackingAreaCode                            `optional`
	ranac                      *RAN_AreaCode                                `optional`
	cellIdentity               CellIdentity                                 `madatory`
	cellReservedForOperatorUse PLMN_IdentityInfo_cellReservedForOperatorUse `madatory`
	iab_Support_r16            *PLMN_IdentityInfo_iab_Support_r16           `optional,ext-1`
	trackingAreaList_r17       []TrackingAreaCode                           `lb:1,ub:maxTAC_r17,optional,ext-2`
	gNB_ID_Length_r17          *int64                                       `lb:22,ub:32,optional,ext-2`
}

func (ie *PLMN_IdentityInfo) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.iab_Support_r16 != nil || len(ie.trackingAreaList_r17) > 0 || ie.gNB_ID_Length_r17 != nil
	preambleBits := []bool{hasExtensions, ie.trackingAreaCode != nil, ie.ranac != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	tmp_plmn_IdentityList := utils.NewSequence[*PLMN_Identity]([]*PLMN_Identity{}, uper.Constraint{Lb: 1, Ub: maxPLMN}, false)
	for _, i := range ie.plmn_IdentityList {
		tmp_plmn_IdentityList.Value = append(tmp_plmn_IdentityList.Value, &i)
	}
	if err = tmp_plmn_IdentityList.Encode(w); err != nil {
		return utils.WrapError("Encode plmn_IdentityList", err)
	}
	if ie.trackingAreaCode != nil {
		if err = ie.trackingAreaCode.Encode(w); err != nil {
			return utils.WrapError("Encode trackingAreaCode", err)
		}
	}
	if ie.ranac != nil {
		if err = ie.ranac.Encode(w); err != nil {
			return utils.WrapError("Encode ranac", err)
		}
	}
	if err = ie.cellIdentity.Encode(w); err != nil {
		return utils.WrapError("Encode cellIdentity", err)
	}
	if err = ie.cellReservedForOperatorUse.Encode(w); err != nil {
		return utils.WrapError("Encode cellReservedForOperatorUse", err)
	}
	if hasExtensions {
		// Extension bitmap: 2 bits for 2 extension groups
		extBitmap := []bool{ie.iab_Support_r16 != nil, len(ie.trackingAreaList_r17) > 0 || ie.gNB_ID_Length_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap PLMN_IdentityInfo", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.iab_Support_r16 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode iab_Support_r16 optional
			if ie.iab_Support_r16 != nil {
				if err = ie.iab_Support_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode iab_Support_r16", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 2
		if extBitmap[1] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 2
			optionals_ext_2 := []bool{len(ie.trackingAreaList_r17) > 0, ie.gNB_ID_Length_r17 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode trackingAreaList_r17 optional
			if len(ie.trackingAreaList_r17) > 0 {
				tmp_trackingAreaList_r17 := utils.NewSequence[*TrackingAreaCode]([]*TrackingAreaCode{}, uper.Constraint{Lb: 1, Ub: maxTAC_r17}, false)
				for _, i := range ie.trackingAreaList_r17 {
					tmp_trackingAreaList_r17.Value = append(tmp_trackingAreaList_r17.Value, &i)
				}
				if err = tmp_trackingAreaList_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode trackingAreaList_r17", err)
				}
			}
			// encode gNB_ID_Length_r17 optional
			if ie.gNB_ID_Length_r17 != nil {
				if err = extWriter.WriteInteger(*ie.gNB_ID_Length_r17, &uper.Constraint{Lb: 22, Ub: 32}, false); err != nil {
					return utils.WrapError("Encode gNB_ID_Length_r17", err)
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

func (ie *PLMN_IdentityInfo) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var trackingAreaCodePresent bool
	if trackingAreaCodePresent, err = r.ReadBool(); err != nil {
		return err
	}
	var ranacPresent bool
	if ranacPresent, err = r.ReadBool(); err != nil {
		return err
	}
	tmp_plmn_IdentityList := utils.NewSequence[*PLMN_Identity]([]*PLMN_Identity{}, uper.Constraint{Lb: 1, Ub: maxPLMN}, false)
	fn_plmn_IdentityList := func() *PLMN_Identity {
		return new(PLMN_Identity)
	}
	if err = tmp_plmn_IdentityList.Decode(r, fn_plmn_IdentityList); err != nil {
		return utils.WrapError("Decode plmn_IdentityList", err)
	}
	ie.plmn_IdentityList = []PLMN_Identity{}
	for _, i := range tmp_plmn_IdentityList.Value {
		ie.plmn_IdentityList = append(ie.plmn_IdentityList, *i)
	}
	if trackingAreaCodePresent {
		ie.trackingAreaCode = new(TrackingAreaCode)
		if err = ie.trackingAreaCode.Decode(r); err != nil {
			return utils.WrapError("Decode trackingAreaCode", err)
		}
	}
	if ranacPresent {
		ie.ranac = new(RAN_AreaCode)
		if err = ie.ranac.Decode(r); err != nil {
			return utils.WrapError("Decode ranac", err)
		}
	}
	if err = ie.cellIdentity.Decode(r); err != nil {
		return utils.WrapError("Decode cellIdentity", err)
	}
	if err = ie.cellReservedForOperatorUse.Decode(r); err != nil {
		return utils.WrapError("Decode cellReservedForOperatorUse", err)
	}

	if extensionBit {
		// Read extension bitmap: 2 bits for 2 extension groups
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

			iab_Support_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode iab_Support_r16 optional
			if iab_Support_r16Present {
				ie.iab_Support_r16 = new(PLMN_IdentityInfo_iab_Support_r16)
				if err = ie.iab_Support_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode iab_Support_r16", err)
				}
			}
		}
		// decode extension group 2
		if len(extBitmap) > 1 && extBitmap[1] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			trackingAreaList_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			gNB_ID_Length_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode trackingAreaList_r17 optional
			if trackingAreaList_r17Present {
				tmp_trackingAreaList_r17 := utils.NewSequence[*TrackingAreaCode]([]*TrackingAreaCode{}, uper.Constraint{Lb: 1, Ub: maxTAC_r17}, false)
				fn_trackingAreaList_r17 := func() *TrackingAreaCode {
					return new(TrackingAreaCode)
				}
				if err = tmp_trackingAreaList_r17.Decode(extReader, fn_trackingAreaList_r17); err != nil {
					return utils.WrapError("Decode trackingAreaList_r17", err)
				}
				ie.trackingAreaList_r17 = []TrackingAreaCode{}
				for _, i := range tmp_trackingAreaList_r17.Value {
					ie.trackingAreaList_r17 = append(ie.trackingAreaList_r17, *i)
				}
			}
			// decode gNB_ID_Length_r17 optional
			if gNB_ID_Length_r17Present {
				var tmp_int_gNB_ID_Length_r17 int64
				if tmp_int_gNB_ID_Length_r17, err = extReader.ReadInteger(&uper.Constraint{Lb: 22, Ub: 32}, false); err != nil {
					return utils.WrapError("Decode gNB_ID_Length_r17", err)
				}
				ie.gNB_ID_Length_r17 = &tmp_int_gNB_ID_Length_r17
			}
		}
	}
	return nil
}
