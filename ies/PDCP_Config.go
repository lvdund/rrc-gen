package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PDCP_Config struct {
	drb                           *PDCP_Config_drb                          `lb:1,ub:16383,optional`
	moreThanOneRLC                *PDCP_Config_moreThanOneRLC               `optional,ext`
	t_Reordering                  *PDCP_Config_t_Reordering                 `optional,ext`
	cipheringDisabled             *PDCP_Config_cipheringDisabled            `optional,ext-1`
	discardTimerExt_r16           *DiscardTimerExt_r16                      `optional,ext-2,setuprelease`
	moreThanTwoRLC_DRB_r16        *PDCP_Config_moreThanTwoRLC_DRB_r16       `lb:3,ub:3,optional,ext-2`
	ethernetHeaderCompression_r16 *EthernetHeaderCompression_r16            `optional,ext-2,setuprelease`
	survivalTimeStateSupport_r17  *PDCP_Config_survivalTimeStateSupport_r17 `optional,ext-3`
	uplinkDataCompression_r17     *UplinkDataCompression_r17                `optional,ext-3,setuprelease`
	discardTimerExt2_r17          *DiscardTimerExt2_r17                     `optional,ext-3,setuprelease`
	initialRX_DELIV_r17           *uper.BitString                           `lb:32,ub:32,optional,ext-3`
}

func (ie *PDCP_Config) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.cipheringDisabled != nil || ie.discardTimerExt_r16 != nil || ie.moreThanTwoRLC_DRB_r16 != nil || ie.ethernetHeaderCompression_r16 != nil || ie.survivalTimeStateSupport_r17 != nil || ie.uplinkDataCompression_r17 != nil || ie.discardTimerExt2_r17 != nil || ie.initialRX_DELIV_r17 != nil
	preambleBits := []bool{hasExtensions, ie.drb != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.drb != nil {
		if err = ie.drb.Encode(w); err != nil {
			return utils.WrapError("Encode drb", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 3 bits for 3 extension groups
		extBitmap := []bool{ie.cipheringDisabled != nil, ie.discardTimerExt_r16 != nil || ie.moreThanTwoRLC_DRB_r16 != nil || ie.ethernetHeaderCompression_r16 != nil, ie.survivalTimeStateSupport_r17 != nil || ie.uplinkDataCompression_r17 != nil || ie.discardTimerExt2_r17 != nil || ie.initialRX_DELIV_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap PDCP_Config", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.cipheringDisabled != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode cipheringDisabled optional
			if ie.cipheringDisabled != nil {
				if err = ie.cipheringDisabled.Encode(extWriter); err != nil {
					return utils.WrapError("Encode cipheringDisabled", err)
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
			optionals_ext_2 := []bool{ie.discardTimerExt_r16 != nil, ie.moreThanTwoRLC_DRB_r16 != nil, ie.ethernetHeaderCompression_r16 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode discardTimerExt_r16 optional
			if ie.discardTimerExt_r16 != nil {
				tmp_discardTimerExt_r16 := utils.SetupRelease[*DiscardTimerExt_r16]{
					Setup: ie.discardTimerExt_r16,
				}
				if err = tmp_discardTimerExt_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode discardTimerExt_r16", err)
				}
			}
			// encode moreThanTwoRLC_DRB_r16 optional
			if ie.moreThanTwoRLC_DRB_r16 != nil {
				if err = ie.moreThanTwoRLC_DRB_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode moreThanTwoRLC_DRB_r16", err)
				}
			}
			// encode ethernetHeaderCompression_r16 optional
			if ie.ethernetHeaderCompression_r16 != nil {
				tmp_ethernetHeaderCompression_r16 := utils.SetupRelease[*EthernetHeaderCompression_r16]{
					Setup: ie.ethernetHeaderCompression_r16,
				}
				if err = tmp_ethernetHeaderCompression_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ethernetHeaderCompression_r16", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 3
		if extBitmap[2] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 3
			optionals_ext_3 := []bool{ie.survivalTimeStateSupport_r17 != nil, ie.uplinkDataCompression_r17 != nil, ie.discardTimerExt2_r17 != nil, ie.initialRX_DELIV_r17 != nil}
			for _, bit := range optionals_ext_3 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode survivalTimeStateSupport_r17 optional
			if ie.survivalTimeStateSupport_r17 != nil {
				if err = ie.survivalTimeStateSupport_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode survivalTimeStateSupport_r17", err)
				}
			}
			// encode uplinkDataCompression_r17 optional
			if ie.uplinkDataCompression_r17 != nil {
				tmp_uplinkDataCompression_r17 := utils.SetupRelease[*UplinkDataCompression_r17]{
					Setup: ie.uplinkDataCompression_r17,
				}
				if err = tmp_uplinkDataCompression_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode uplinkDataCompression_r17", err)
				}
			}
			// encode discardTimerExt2_r17 optional
			if ie.discardTimerExt2_r17 != nil {
				tmp_discardTimerExt2_r17 := utils.SetupRelease[*DiscardTimerExt2_r17]{
					Setup: ie.discardTimerExt2_r17,
				}
				if err = tmp_discardTimerExt2_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode discardTimerExt2_r17", err)
				}
			}
			// encode initialRX_DELIV_r17 optional
			if ie.initialRX_DELIV_r17 != nil {
				if err = extWriter.WriteBitString(ie.initialRX_DELIV_r17.Bytes, uint(ie.initialRX_DELIV_r17.NumBits), &uper.Constraint{Lb: 32, Ub: 32}, false); err != nil {
					return utils.WrapError("Encode initialRX_DELIV_r17", err)
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

func (ie *PDCP_Config) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var drbPresent bool
	if drbPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if drbPresent {
		ie.drb = new(PDCP_Config_drb)
		if err = ie.drb.Decode(r); err != nil {
			return utils.WrapError("Decode drb", err)
		}
	}

	if extensionBit {
		// Read extension bitmap: 3 bits for 3 extension groups
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

			cipheringDisabledPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode cipheringDisabled optional
			if cipheringDisabledPresent {
				ie.cipheringDisabled = new(PDCP_Config_cipheringDisabled)
				if err = ie.cipheringDisabled.Decode(extReader); err != nil {
					return utils.WrapError("Decode cipheringDisabled", err)
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

			discardTimerExt_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			moreThanTwoRLC_DRB_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			ethernetHeaderCompression_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode discardTimerExt_r16 optional
			if discardTimerExt_r16Present {
				tmp_discardTimerExt_r16 := utils.SetupRelease[*DiscardTimerExt_r16]{}
				if err = tmp_discardTimerExt_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode discardTimerExt_r16", err)
				}
				ie.discardTimerExt_r16 = tmp_discardTimerExt_r16.Setup
			}
			// decode moreThanTwoRLC_DRB_r16 optional
			if moreThanTwoRLC_DRB_r16Present {
				ie.moreThanTwoRLC_DRB_r16 = new(PDCP_Config_moreThanTwoRLC_DRB_r16)
				if err = ie.moreThanTwoRLC_DRB_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode moreThanTwoRLC_DRB_r16", err)
				}
			}
			// decode ethernetHeaderCompression_r16 optional
			if ethernetHeaderCompression_r16Present {
				tmp_ethernetHeaderCompression_r16 := utils.SetupRelease[*EthernetHeaderCompression_r16]{}
				if err = tmp_ethernetHeaderCompression_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode ethernetHeaderCompression_r16", err)
				}
				ie.ethernetHeaderCompression_r16 = tmp_ethernetHeaderCompression_r16.Setup
			}
		}
		// decode extension group 3
		if len(extBitmap) > 2 && extBitmap[2] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			survivalTimeStateSupport_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			uplinkDataCompression_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			discardTimerExt2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			initialRX_DELIV_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode survivalTimeStateSupport_r17 optional
			if survivalTimeStateSupport_r17Present {
				ie.survivalTimeStateSupport_r17 = new(PDCP_Config_survivalTimeStateSupport_r17)
				if err = ie.survivalTimeStateSupport_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode survivalTimeStateSupport_r17", err)
				}
			}
			// decode uplinkDataCompression_r17 optional
			if uplinkDataCompression_r17Present {
				tmp_uplinkDataCompression_r17 := utils.SetupRelease[*UplinkDataCompression_r17]{}
				if err = tmp_uplinkDataCompression_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode uplinkDataCompression_r17", err)
				}
				ie.uplinkDataCompression_r17 = tmp_uplinkDataCompression_r17.Setup
			}
			// decode discardTimerExt2_r17 optional
			if discardTimerExt2_r17Present {
				tmp_discardTimerExt2_r17 := utils.SetupRelease[*DiscardTimerExt2_r17]{}
				if err = tmp_discardTimerExt2_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode discardTimerExt2_r17", err)
				}
				ie.discardTimerExt2_r17 = tmp_discardTimerExt2_r17.Setup
			}
			// decode initialRX_DELIV_r17 optional
			if initialRX_DELIV_r17Present {
				var tmp_bs_initialRX_DELIV_r17 []byte
				var n_initialRX_DELIV_r17 uint
				if tmp_bs_initialRX_DELIV_r17, n_initialRX_DELIV_r17, err = extReader.ReadBitString(&uper.Constraint{Lb: 32, Ub: 32}, false); err != nil {
					return utils.WrapError("Decode initialRX_DELIV_r17", err)
				}
				tmp_bitstring := uper.BitString{
					Bytes:   tmp_bs_initialRX_DELIV_r17,
					NumBits: uint64(n_initialRX_DELIV_r17),
				}
				ie.initialRX_DELIV_r17 = &tmp_bitstring
			}
		}
	}
	return nil
}
