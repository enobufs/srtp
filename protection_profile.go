// SPDX-FileCopyrightText: 2023 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package srtp

import "fmt"

// ProtectionProfile specifies Cipher and AuthTag details, similar to TLS cipher suite
type ProtectionProfile uint16

// Supported protection profiles
// See https://www.iana.org/assignments/srtp-protection/srtp-protection.xhtml
const (
	ProtectionProfileAes128CmHmacSha1_80 ProtectionProfile = 0x0001
	ProtectionProfileAes128CmHmacSha1_32 ProtectionProfile = 0x0002
	ProtectionProfileAeadAes128Gcm       ProtectionProfile = 0x0007
	ProtectionProfileAeadAes256Gcm       ProtectionProfile = 0x0008
)

// KeyLen returns length of encryption key in bytes.
func (p ProtectionProfile) KeyLen() (int, error) {
	switch p {
	case ProtectionProfileAes128CmHmacSha1_32, ProtectionProfileAes128CmHmacSha1_80, ProtectionProfileAeadAes128Gcm:
		return 16, nil
	case ProtectionProfileAeadAes256Gcm:
		return 32, nil
	default:
		return 0, fmt.Errorf("%w: %#v", errNoSuchSRTPProfile, p)
	}
}

// SaltLen returns length of salt key in bytes.
func (p ProtectionProfile) SaltLen() (int, error) {
	switch p {
	case ProtectionProfileAes128CmHmacSha1_32, ProtectionProfileAes128CmHmacSha1_80:
		return 14, nil
	case ProtectionProfileAeadAes128Gcm, ProtectionProfileAeadAes256Gcm:
		return 12, nil
	default:
		return 0, fmt.Errorf("%w: %#v", errNoSuchSRTPProfile, p)
	}
}

// AuthTagRTPLen returns length of RTP authentication tag in bytes for AES protection profiles. For AEAD ones it returns zero.
func (p ProtectionProfile) AuthTagRTPLen() (int, error) {
	switch p {
	case ProtectionProfileAes128CmHmacSha1_80:
		return 10, nil
	case ProtectionProfileAes128CmHmacSha1_32:
		return 4, nil
	case ProtectionProfileAeadAes128Gcm, ProtectionProfileAeadAes256Gcm:
		return 0, nil
	default:
		return 0, fmt.Errorf("%w: %#v", errNoSuchSRTPProfile, p)
	}
}

// AuthTagRTCPLen returns length of RTCP authentication tag in bytes for AES protection profiles. For AEAD ones it returns zero.
func (p ProtectionProfile) AuthTagRTCPLen() (int, error) {
	switch p {
	case ProtectionProfileAes128CmHmacSha1_32, ProtectionProfileAes128CmHmacSha1_80:
		return 10, nil
	case ProtectionProfileAeadAes128Gcm, ProtectionProfileAeadAes256Gcm:
		return 0, nil
	default:
		return 0, fmt.Errorf("%w: %#v", errNoSuchSRTPProfile, p)
	}
}

// AEADAuthTagLen returns length of authentication tag in bytes for AEAD protection profiles. For AES ones it returns zero.
func (p ProtectionProfile) AEADAuthTagLen() (int, error) {
	switch p {
	case ProtectionProfileAes128CmHmacSha1_32, ProtectionProfileAes128CmHmacSha1_80:
		return 0, nil
	case ProtectionProfileAeadAes128Gcm, ProtectionProfileAeadAes256Gcm:
		return 16, nil
	default:
		return 0, fmt.Errorf("%w: %#v", errNoSuchSRTPProfile, p)
	}
}

// AuthKeyLen returns length of authentication key in bytes for AES protection profiles. For AEAD ones it returns zero.
func (p ProtectionProfile) AuthKeyLen() (int, error) {
	switch p {
	case ProtectionProfileAes128CmHmacSha1_32, ProtectionProfileAes128CmHmacSha1_80:
		return 20, nil
	case ProtectionProfileAeadAes128Gcm, ProtectionProfileAeadAes256Gcm:
		return 0, nil
	default:
		return 0, fmt.Errorf("%w: %#v", errNoSuchSRTPProfile, p)
	}
}

// String returns the name of the protection profile.
func (p ProtectionProfile) String() string {
	switch p {
	case ProtectionProfileAes128CmHmacSha1_80:
		return "SRTP_AES128_CM_HMAC_SHA1_80"
	case ProtectionProfileAes128CmHmacSha1_32:
		return "SRTP_AES128_CM_HMAC_SHA1_32"
	case ProtectionProfileAeadAes128Gcm:
		return "SRTP_AEAD_AES_128_GCM"
	case ProtectionProfileAeadAes256Gcm:
		return "SRTP_AEAD_AES_256_GCM"
	default:
		return fmt.Sprintf("Unknown SRTP profile: %#v", p)
	}
}
