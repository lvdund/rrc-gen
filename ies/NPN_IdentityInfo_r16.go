package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type NPN_IdentityInfo_r16 struct {
	npn_IdentityList_r16           []NPN_Identity_r16                                  `lb:1,ub:maxNPN_r16,madatory`
	trackingAreaCode_r16           TrackingAreaCode                                    `madatory`
	ranac_r16                      *RAN_AreaCode                                       `optional`
	cellIdentity_r16               CellIdentity                                        `madatory`
	cellReservedForOperatorUse_r16 NPN_IdentityInfo_r16_cellReservedForOperatorUse_r16 `madatory`
	iab_Support_r16                *NPN_IdentityInfo_r16_iab_Support_r16               `optional`
	gNB_ID_Length_r17              *int64                                              `lb:22,ub:32,optional,ext-1`
}

func (ie *NPN_IdentityInfo_r16) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.gNB_ID_Length_r17 != nil
	preambleBits := []bool{hasExtensions, ie.ranac_r16 != nil, ie.iab_Support_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	tmp_npn_IdentityList_r16 := utils.NewSequence[*NPN_Identity_r16]([]*NPN_Identity_r16{}, uper.Constraint{Lb: 1, Ub: maxNPN_r16}, false)
	for _, i := range ie.npn_IdentityList_r16 {
		tmp_npn_IdentityList_r16.Value = append(tmp_npn_IdentityList_r16.Value, &i)
	}
	if err = tmp_npn_IdentityList_r16.Encode(w); err != nil {
		return utils.WrapError("Encode npn_IdentityList_r16", err)
	}
	if err = ie.trackingAreaCode_r16.Encode(w); err != nil {
		return utils.WrapError("Encode trackingAreaCode_r16", err)
	}
	if ie.ranac_r16 != nil {
		if err = ie.ranac_r16.Encode(w); err != nil {
			return utils.WrapError("Encode ranac_r16", err)
		}
	}
	if err = ie.cellIdentity_r16.Encode(w); err != nil {
		return utils.WrapError("Encode cellIdentity_r16", err)
	}
	if err = ie.cellReservedForOperatorUse_r16.Encode(w); err != nil {
		return utils.WrapError("Encode cellReservedForOperatorUse_r16", err)
	}
	if ie.iab_Support_r16 != nil {
		if err = ie.iab_Support_r16.Encode(w); err != nil {
			return utils.WrapError("Encode iab_Support_r16", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.gNB_ID_Length_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap NPN_IdentityInfo_r16", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.gNB_ID_Length_r17 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
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

func (ie *NPN_IdentityInfo_r16) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var ranac_r16Present bool
	if ranac_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var iab_Support_r16Present bool
	if iab_Support_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	tmp_npn_IdentityList_r16 := utils.NewSequence[*NPN_Identity_r16]([]*NPN_Identity_r16{}, uper.Constraint{Lb: 1, Ub: maxNPN_r16}, false)
	fn_npn_IdentityList_r16 := func() *NPN_Identity_r16 {
		return new(NPN_Identity_r16)
	}
	if err = tmp_npn_IdentityList_r16.Decode(r, fn_npn_IdentityList_r16); err != nil {
		return utils.WrapError("Decode npn_IdentityList_r16", err)
	}
	ie.npn_IdentityList_r16 = []NPN_Identity_r16{}
	for _, i := range tmp_npn_IdentityList_r16.Value {
		ie.npn_IdentityList_r16 = append(ie.npn_IdentityList_r16, *i)
	}
	if err = ie.trackingAreaCode_r16.Decode(r); err != nil {
		return utils.WrapError("Decode trackingAreaCode_r16", err)
	}
	if ranac_r16Present {
		ie.ranac_r16 = new(RAN_AreaCode)
		if err = ie.ranac_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ranac_r16", err)
		}
	}
	if err = ie.cellIdentity_r16.Decode(r); err != nil {
		return utils.WrapError("Decode cellIdentity_r16", err)
	}
	if err = ie.cellReservedForOperatorUse_r16.Decode(r); err != nil {
		return utils.WrapError("Decode cellReservedForOperatorUse_r16", err)
	}
	if iab_Support_r16Present {
		ie.iab_Support_r16 = new(NPN_IdentityInfo_r16_iab_Support_r16)
		if err = ie.iab_Support_r16.Decode(r); err != nil {
			return utils.WrapError("Decode iab_Support_r16", err)
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

			gNB_ID_Length_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
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
