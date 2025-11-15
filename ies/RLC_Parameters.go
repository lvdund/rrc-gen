package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RLC_Parameters struct {
	am_WithShortSN               *RLC_Parameters_am_WithShortSN               `optional`
	um_WithShortSN               *RLC_Parameters_um_WithShortSN               `optional`
	um_WithLongSN                *RLC_Parameters_um_WithLongSN                `optional`
	extendedT_PollRetransmit_r16 *RLC_Parameters_extendedT_PollRetransmit_r16 `optional,ext-1`
	extendedT_StatusProhibit_r16 *RLC_Parameters_extendedT_StatusProhibit_r16 `optional,ext-1`
	am_WithLongSN_RedCap_r17     *RLC_Parameters_am_WithLongSN_RedCap_r17     `optional,ext-2`
}

func (ie *RLC_Parameters) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.extendedT_PollRetransmit_r16 != nil || ie.extendedT_StatusProhibit_r16 != nil || ie.am_WithLongSN_RedCap_r17 != nil
	preambleBits := []bool{hasExtensions, ie.am_WithShortSN != nil, ie.um_WithShortSN != nil, ie.um_WithLongSN != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.am_WithShortSN != nil {
		if err = ie.am_WithShortSN.Encode(w); err != nil {
			return utils.WrapError("Encode am_WithShortSN", err)
		}
	}
	if ie.um_WithShortSN != nil {
		if err = ie.um_WithShortSN.Encode(w); err != nil {
			return utils.WrapError("Encode um_WithShortSN", err)
		}
	}
	if ie.um_WithLongSN != nil {
		if err = ie.um_WithLongSN.Encode(w); err != nil {
			return utils.WrapError("Encode um_WithLongSN", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 2 bits for 2 extension groups
		extBitmap := []bool{ie.extendedT_PollRetransmit_r16 != nil || ie.extendedT_StatusProhibit_r16 != nil, ie.am_WithLongSN_RedCap_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap RLC_Parameters", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.extendedT_PollRetransmit_r16 != nil, ie.extendedT_StatusProhibit_r16 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode extendedT_PollRetransmit_r16 optional
			if ie.extendedT_PollRetransmit_r16 != nil {
				if err = ie.extendedT_PollRetransmit_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode extendedT_PollRetransmit_r16", err)
				}
			}
			// encode extendedT_StatusProhibit_r16 optional
			if ie.extendedT_StatusProhibit_r16 != nil {
				if err = ie.extendedT_StatusProhibit_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode extendedT_StatusProhibit_r16", err)
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
			optionals_ext_2 := []bool{ie.am_WithLongSN_RedCap_r17 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode am_WithLongSN_RedCap_r17 optional
			if ie.am_WithLongSN_RedCap_r17 != nil {
				if err = ie.am_WithLongSN_RedCap_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode am_WithLongSN_RedCap_r17", err)
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

func (ie *RLC_Parameters) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var am_WithShortSNPresent bool
	if am_WithShortSNPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var um_WithShortSNPresent bool
	if um_WithShortSNPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var um_WithLongSNPresent bool
	if um_WithLongSNPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if am_WithShortSNPresent {
		ie.am_WithShortSN = new(RLC_Parameters_am_WithShortSN)
		if err = ie.am_WithShortSN.Decode(r); err != nil {
			return utils.WrapError("Decode am_WithShortSN", err)
		}
	}
	if um_WithShortSNPresent {
		ie.um_WithShortSN = new(RLC_Parameters_um_WithShortSN)
		if err = ie.um_WithShortSN.Decode(r); err != nil {
			return utils.WrapError("Decode um_WithShortSN", err)
		}
	}
	if um_WithLongSNPresent {
		ie.um_WithLongSN = new(RLC_Parameters_um_WithLongSN)
		if err = ie.um_WithLongSN.Decode(r); err != nil {
			return utils.WrapError("Decode um_WithLongSN", err)
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

			extendedT_PollRetransmit_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			extendedT_StatusProhibit_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode extendedT_PollRetransmit_r16 optional
			if extendedT_PollRetransmit_r16Present {
				ie.extendedT_PollRetransmit_r16 = new(RLC_Parameters_extendedT_PollRetransmit_r16)
				if err = ie.extendedT_PollRetransmit_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode extendedT_PollRetransmit_r16", err)
				}
			}
			// decode extendedT_StatusProhibit_r16 optional
			if extendedT_StatusProhibit_r16Present {
				ie.extendedT_StatusProhibit_r16 = new(RLC_Parameters_extendedT_StatusProhibit_r16)
				if err = ie.extendedT_StatusProhibit_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode extendedT_StatusProhibit_r16", err)
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

			am_WithLongSN_RedCap_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode am_WithLongSN_RedCap_r17 optional
			if am_WithLongSN_RedCap_r17Present {
				ie.am_WithLongSN_RedCap_r17 = new(RLC_Parameters_am_WithLongSN_RedCap_r17)
				if err = ie.am_WithLongSN_RedCap_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode am_WithLongSN_RedCap_r17", err)
				}
			}
		}
	}
	return nil
}
