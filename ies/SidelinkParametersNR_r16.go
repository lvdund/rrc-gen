package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SidelinkParametersNR_r16 struct {
	rlc_ParametersSidelink_r16           *RLC_ParametersSidelink_r16                    `optional`
	mac_ParametersSidelink_r16           *MAC_ParametersSidelink_r16                    `optional`
	fdd_Add_UE_Sidelink_Capabilities_r16 *UE_SidelinkCapabilityAddXDD_Mode_r16          `optional`
	tdd_Add_UE_Sidelink_Capabilities_r16 *UE_SidelinkCapabilityAddXDD_Mode_r16          `optional`
	supportedBandListSidelink_r16        []BandSidelink_r16                             `lb:1,ub:maxBands,optional`
	relayParameters_r17                  *RelayParameters_r17                           `optional,ext-1`
	p0_OLPC_Sidelink_r17                 *SidelinkParametersNR_r16_p0_OLPC_Sidelink_r17 `optional,ext-2`
}

func (ie *SidelinkParametersNR_r16) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.relayParameters_r17 != nil || ie.p0_OLPC_Sidelink_r17 != nil
	preambleBits := []bool{hasExtensions, ie.rlc_ParametersSidelink_r16 != nil, ie.mac_ParametersSidelink_r16 != nil, ie.fdd_Add_UE_Sidelink_Capabilities_r16 != nil, ie.tdd_Add_UE_Sidelink_Capabilities_r16 != nil, len(ie.supportedBandListSidelink_r16) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.rlc_ParametersSidelink_r16 != nil {
		if err = ie.rlc_ParametersSidelink_r16.Encode(w); err != nil {
			return utils.WrapError("Encode rlc_ParametersSidelink_r16", err)
		}
	}
	if ie.mac_ParametersSidelink_r16 != nil {
		if err = ie.mac_ParametersSidelink_r16.Encode(w); err != nil {
			return utils.WrapError("Encode mac_ParametersSidelink_r16", err)
		}
	}
	if ie.fdd_Add_UE_Sidelink_Capabilities_r16 != nil {
		if err = ie.fdd_Add_UE_Sidelink_Capabilities_r16.Encode(w); err != nil {
			return utils.WrapError("Encode fdd_Add_UE_Sidelink_Capabilities_r16", err)
		}
	}
	if ie.tdd_Add_UE_Sidelink_Capabilities_r16 != nil {
		if err = ie.tdd_Add_UE_Sidelink_Capabilities_r16.Encode(w); err != nil {
			return utils.WrapError("Encode tdd_Add_UE_Sidelink_Capabilities_r16", err)
		}
	}
	if len(ie.supportedBandListSidelink_r16) > 0 {
		tmp_supportedBandListSidelink_r16 := utils.NewSequence[*BandSidelink_r16]([]*BandSidelink_r16{}, uper.Constraint{Lb: 1, Ub: maxBands}, false)
		for _, i := range ie.supportedBandListSidelink_r16 {
			tmp_supportedBandListSidelink_r16.Value = append(tmp_supportedBandListSidelink_r16.Value, &i)
		}
		if err = tmp_supportedBandListSidelink_r16.Encode(w); err != nil {
			return utils.WrapError("Encode supportedBandListSidelink_r16", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 2 bits for 2 extension groups
		extBitmap := []bool{ie.relayParameters_r17 != nil, ie.p0_OLPC_Sidelink_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap SidelinkParametersNR_r16", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.relayParameters_r17 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode relayParameters_r17 optional
			if ie.relayParameters_r17 != nil {
				if err = ie.relayParameters_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode relayParameters_r17", err)
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
			optionals_ext_2 := []bool{ie.p0_OLPC_Sidelink_r17 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode p0_OLPC_Sidelink_r17 optional
			if ie.p0_OLPC_Sidelink_r17 != nil {
				if err = ie.p0_OLPC_Sidelink_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode p0_OLPC_Sidelink_r17", err)
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

func (ie *SidelinkParametersNR_r16) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var rlc_ParametersSidelink_r16Present bool
	if rlc_ParametersSidelink_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var mac_ParametersSidelink_r16Present bool
	if mac_ParametersSidelink_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var fdd_Add_UE_Sidelink_Capabilities_r16Present bool
	if fdd_Add_UE_Sidelink_Capabilities_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var tdd_Add_UE_Sidelink_Capabilities_r16Present bool
	if tdd_Add_UE_Sidelink_Capabilities_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var supportedBandListSidelink_r16Present bool
	if supportedBandListSidelink_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if rlc_ParametersSidelink_r16Present {
		ie.rlc_ParametersSidelink_r16 = new(RLC_ParametersSidelink_r16)
		if err = ie.rlc_ParametersSidelink_r16.Decode(r); err != nil {
			return utils.WrapError("Decode rlc_ParametersSidelink_r16", err)
		}
	}
	if mac_ParametersSidelink_r16Present {
		ie.mac_ParametersSidelink_r16 = new(MAC_ParametersSidelink_r16)
		if err = ie.mac_ParametersSidelink_r16.Decode(r); err != nil {
			return utils.WrapError("Decode mac_ParametersSidelink_r16", err)
		}
	}
	if fdd_Add_UE_Sidelink_Capabilities_r16Present {
		ie.fdd_Add_UE_Sidelink_Capabilities_r16 = new(UE_SidelinkCapabilityAddXDD_Mode_r16)
		if err = ie.fdd_Add_UE_Sidelink_Capabilities_r16.Decode(r); err != nil {
			return utils.WrapError("Decode fdd_Add_UE_Sidelink_Capabilities_r16", err)
		}
	}
	if tdd_Add_UE_Sidelink_Capabilities_r16Present {
		ie.tdd_Add_UE_Sidelink_Capabilities_r16 = new(UE_SidelinkCapabilityAddXDD_Mode_r16)
		if err = ie.tdd_Add_UE_Sidelink_Capabilities_r16.Decode(r); err != nil {
			return utils.WrapError("Decode tdd_Add_UE_Sidelink_Capabilities_r16", err)
		}
	}
	if supportedBandListSidelink_r16Present {
		tmp_supportedBandListSidelink_r16 := utils.NewSequence[*BandSidelink_r16]([]*BandSidelink_r16{}, uper.Constraint{Lb: 1, Ub: maxBands}, false)
		fn_supportedBandListSidelink_r16 := func() *BandSidelink_r16 {
			return new(BandSidelink_r16)
		}
		if err = tmp_supportedBandListSidelink_r16.Decode(r, fn_supportedBandListSidelink_r16); err != nil {
			return utils.WrapError("Decode supportedBandListSidelink_r16", err)
		}
		ie.supportedBandListSidelink_r16 = []BandSidelink_r16{}
		for _, i := range tmp_supportedBandListSidelink_r16.Value {
			ie.supportedBandListSidelink_r16 = append(ie.supportedBandListSidelink_r16, *i)
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

			relayParameters_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode relayParameters_r17 optional
			if relayParameters_r17Present {
				ie.relayParameters_r17 = new(RelayParameters_r17)
				if err = ie.relayParameters_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode relayParameters_r17", err)
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

			p0_OLPC_Sidelink_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode p0_OLPC_Sidelink_r17 optional
			if p0_OLPC_Sidelink_r17Present {
				ie.p0_OLPC_Sidelink_r17 = new(SidelinkParametersNR_r16_p0_OLPC_Sidelink_r17)
				if err = ie.p0_OLPC_Sidelink_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode p0_OLPC_Sidelink_r17", err)
				}
			}
		}
	}
	return nil
}
