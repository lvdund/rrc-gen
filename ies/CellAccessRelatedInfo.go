package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CellAccessRelatedInfo struct {
	plmn_IdentityInfoList        PLMN_IdentityInfoList                               `madatory`
	cellReservedForOtherUse      *CellAccessRelatedInfo_cellReservedForOtherUse      `optional`
	cellReservedForFutureUse_r16 *CellAccessRelatedInfo_cellReservedForFutureUse_r16 `optional,ext-1`
	npn_IdentityInfoList_r16     *NPN_IdentityInfoList_r16                           `optional,ext-1`
	snpn_AccessInfoList_r17      []SNPN_AccessInfo_r17                               `lb:1,ub:maxNPN_r16,optional,ext-2`
}

func (ie *CellAccessRelatedInfo) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.cellReservedForFutureUse_r16 != nil || ie.npn_IdentityInfoList_r16 != nil || len(ie.snpn_AccessInfoList_r17) > 0
	preambleBits := []bool{hasExtensions, ie.cellReservedForOtherUse != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.plmn_IdentityInfoList.Encode(w); err != nil {
		return utils.WrapError("Encode plmn_IdentityInfoList", err)
	}
	if ie.cellReservedForOtherUse != nil {
		if err = ie.cellReservedForOtherUse.Encode(w); err != nil {
			return utils.WrapError("Encode cellReservedForOtherUse", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 2 bits for 2 extension groups
		extBitmap := []bool{ie.cellReservedForFutureUse_r16 != nil || ie.npn_IdentityInfoList_r16 != nil, len(ie.snpn_AccessInfoList_r17) > 0}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap CellAccessRelatedInfo", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.cellReservedForFutureUse_r16 != nil, ie.npn_IdentityInfoList_r16 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode cellReservedForFutureUse_r16 optional
			if ie.cellReservedForFutureUse_r16 != nil {
				if err = ie.cellReservedForFutureUse_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode cellReservedForFutureUse_r16", err)
				}
			}
			// encode npn_IdentityInfoList_r16 optional
			if ie.npn_IdentityInfoList_r16 != nil {
				if err = ie.npn_IdentityInfoList_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode npn_IdentityInfoList_r16", err)
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
			optionals_ext_2 := []bool{len(ie.snpn_AccessInfoList_r17) > 0}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode snpn_AccessInfoList_r17 optional
			if len(ie.snpn_AccessInfoList_r17) > 0 {
				tmp_snpn_AccessInfoList_r17 := utils.NewSequence[*SNPN_AccessInfo_r17]([]*SNPN_AccessInfo_r17{}, uper.Constraint{Lb: 1, Ub: maxNPN_r16}, false)
				for _, i := range ie.snpn_AccessInfoList_r17 {
					tmp_snpn_AccessInfoList_r17.Value = append(tmp_snpn_AccessInfoList_r17.Value, &i)
				}
				if err = tmp_snpn_AccessInfoList_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode snpn_AccessInfoList_r17", err)
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

func (ie *CellAccessRelatedInfo) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var cellReservedForOtherUsePresent bool
	if cellReservedForOtherUsePresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.plmn_IdentityInfoList.Decode(r); err != nil {
		return utils.WrapError("Decode plmn_IdentityInfoList", err)
	}
	if cellReservedForOtherUsePresent {
		ie.cellReservedForOtherUse = new(CellAccessRelatedInfo_cellReservedForOtherUse)
		if err = ie.cellReservedForOtherUse.Decode(r); err != nil {
			return utils.WrapError("Decode cellReservedForOtherUse", err)
		}
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

			cellReservedForFutureUse_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			npn_IdentityInfoList_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode cellReservedForFutureUse_r16 optional
			if cellReservedForFutureUse_r16Present {
				ie.cellReservedForFutureUse_r16 = new(CellAccessRelatedInfo_cellReservedForFutureUse_r16)
				if err = ie.cellReservedForFutureUse_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode cellReservedForFutureUse_r16", err)
				}
			}
			// decode npn_IdentityInfoList_r16 optional
			if npn_IdentityInfoList_r16Present {
				ie.npn_IdentityInfoList_r16 = new(NPN_IdentityInfoList_r16)
				if err = ie.npn_IdentityInfoList_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode npn_IdentityInfoList_r16", err)
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

			snpn_AccessInfoList_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode snpn_AccessInfoList_r17 optional
			if snpn_AccessInfoList_r17Present {
				tmp_snpn_AccessInfoList_r17 := utils.NewSequence[*SNPN_AccessInfo_r17]([]*SNPN_AccessInfo_r17{}, uper.Constraint{Lb: 1, Ub: maxNPN_r16}, false)
				fn_snpn_AccessInfoList_r17 := func() *SNPN_AccessInfo_r17 {
					return new(SNPN_AccessInfo_r17)
				}
				if err = tmp_snpn_AccessInfoList_r17.Decode(extReader, fn_snpn_AccessInfoList_r17); err != nil {
					return utils.WrapError("Decode snpn_AccessInfoList_r17", err)
				}
				ie.snpn_AccessInfoList_r17 = []SNPN_AccessInfo_r17{}
				for _, i := range tmp_snpn_AccessInfoList_r17.Value {
					ie.snpn_AccessInfoList_r17 = append(ie.snpn_AccessInfoList_r17, *i)
				}
			}
		}
	}
	return nil
}
