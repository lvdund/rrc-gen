package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CGI_InfoNR struct {
	plmn_IdentityInfoList       *PLMN_IdentityInfoList                  `optional`
	frequencyBandList           *MultiFrequencyBandListNR               `optional`
	noSIB1                      *CGI_InfoNR_noSIB1                      `lb:0,ub:15,optional`
	npn_IdentityInfoList_r16    *NPN_IdentityInfoList_r16               `optional,ext-1`
	cellReservedForOtherUse_r16 *CGI_InfoNR_cellReservedForOtherUse_r16 `optional,ext-2`
}

func (ie *CGI_InfoNR) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.npn_IdentityInfoList_r16 != nil || ie.cellReservedForOtherUse_r16 != nil
	preambleBits := []bool{hasExtensions, ie.plmn_IdentityInfoList != nil, ie.frequencyBandList != nil, ie.noSIB1 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.plmn_IdentityInfoList != nil {
		if err = ie.plmn_IdentityInfoList.Encode(w); err != nil {
			return utils.WrapError("Encode plmn_IdentityInfoList", err)
		}
	}
	if ie.frequencyBandList != nil {
		if err = ie.frequencyBandList.Encode(w); err != nil {
			return utils.WrapError("Encode frequencyBandList", err)
		}
	}
	if ie.noSIB1 != nil {
		if err = ie.noSIB1.Encode(w); err != nil {
			return utils.WrapError("Encode noSIB1", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 2 bits for 2 extension groups
		extBitmap := []bool{ie.npn_IdentityInfoList_r16 != nil, ie.cellReservedForOtherUse_r16 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap CGI_InfoNR", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.npn_IdentityInfoList_r16 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
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
			optionals_ext_2 := []bool{ie.cellReservedForOtherUse_r16 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode cellReservedForOtherUse_r16 optional
			if ie.cellReservedForOtherUse_r16 != nil {
				if err = ie.cellReservedForOtherUse_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode cellReservedForOtherUse_r16", err)
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

func (ie *CGI_InfoNR) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var plmn_IdentityInfoListPresent bool
	if plmn_IdentityInfoListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var frequencyBandListPresent bool
	if frequencyBandListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var noSIB1Present bool
	if noSIB1Present, err = r.ReadBool(); err != nil {
		return err
	}
	if plmn_IdentityInfoListPresent {
		ie.plmn_IdentityInfoList = new(PLMN_IdentityInfoList)
		if err = ie.plmn_IdentityInfoList.Decode(r); err != nil {
			return utils.WrapError("Decode plmn_IdentityInfoList", err)
		}
	}
	if frequencyBandListPresent {
		ie.frequencyBandList = new(MultiFrequencyBandListNR)
		if err = ie.frequencyBandList.Decode(r); err != nil {
			return utils.WrapError("Decode frequencyBandList", err)
		}
	}
	if noSIB1Present {
		ie.noSIB1 = new(CGI_InfoNR_noSIB1)
		if err = ie.noSIB1.Decode(r); err != nil {
			return utils.WrapError("Decode noSIB1", err)
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

			npn_IdentityInfoList_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
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

			cellReservedForOtherUse_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode cellReservedForOtherUse_r16 optional
			if cellReservedForOtherUse_r16Present {
				ie.cellReservedForOtherUse_r16 = new(CGI_InfoNR_cellReservedForOtherUse_r16)
				if err = ie.cellReservedForOtherUse_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode cellReservedForOtherUse_r16", err)
				}
			}
		}
	}
	return nil
}
