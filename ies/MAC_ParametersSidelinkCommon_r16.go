package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MAC_ParametersSidelinkCommon_r16 struct {
	lcp_RestrictionSidelink_r16          *MAC_ParametersSidelinkCommon_r16_lcp_RestrictionSidelink_r16          `optional`
	multipleConfiguredGrantsSidelink_r16 *MAC_ParametersSidelinkCommon_r16_multipleConfiguredGrantsSidelink_r16 `optional`
	drx_OnSidelink_r17                   *MAC_ParametersSidelinkCommon_r16_drx_OnSidelink_r17                   `optional,ext-1`
}

func (ie *MAC_ParametersSidelinkCommon_r16) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.drx_OnSidelink_r17 != nil
	preambleBits := []bool{hasExtensions, ie.lcp_RestrictionSidelink_r16 != nil, ie.multipleConfiguredGrantsSidelink_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.lcp_RestrictionSidelink_r16 != nil {
		if err = ie.lcp_RestrictionSidelink_r16.Encode(w); err != nil {
			return utils.WrapError("Encode lcp_RestrictionSidelink_r16", err)
		}
	}
	if ie.multipleConfiguredGrantsSidelink_r16 != nil {
		if err = ie.multipleConfiguredGrantsSidelink_r16.Encode(w); err != nil {
			return utils.WrapError("Encode multipleConfiguredGrantsSidelink_r16", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.drx_OnSidelink_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap MAC_ParametersSidelinkCommon_r16", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.drx_OnSidelink_r17 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode drx_OnSidelink_r17 optional
			if ie.drx_OnSidelink_r17 != nil {
				if err = ie.drx_OnSidelink_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode drx_OnSidelink_r17", err)
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

func (ie *MAC_ParametersSidelinkCommon_r16) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var lcp_RestrictionSidelink_r16Present bool
	if lcp_RestrictionSidelink_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var multipleConfiguredGrantsSidelink_r16Present bool
	if multipleConfiguredGrantsSidelink_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if lcp_RestrictionSidelink_r16Present {
		ie.lcp_RestrictionSidelink_r16 = new(MAC_ParametersSidelinkCommon_r16_lcp_RestrictionSidelink_r16)
		if err = ie.lcp_RestrictionSidelink_r16.Decode(r); err != nil {
			return utils.WrapError("Decode lcp_RestrictionSidelink_r16", err)
		}
	}
	if multipleConfiguredGrantsSidelink_r16Present {
		ie.multipleConfiguredGrantsSidelink_r16 = new(MAC_ParametersSidelinkCommon_r16_multipleConfiguredGrantsSidelink_r16)
		if err = ie.multipleConfiguredGrantsSidelink_r16.Decode(r); err != nil {
			return utils.WrapError("Decode multipleConfiguredGrantsSidelink_r16", err)
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

			drx_OnSidelink_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode drx_OnSidelink_r17 optional
			if drx_OnSidelink_r17Present {
				ie.drx_OnSidelink_r17 = new(MAC_ParametersSidelinkCommon_r16_drx_OnSidelink_r17)
				if err = ie.drx_OnSidelink_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode drx_OnSidelink_r17", err)
				}
			}
		}
	}
	return nil
}
