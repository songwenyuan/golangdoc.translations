// Copyright The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ingore

// Package rc4 implements RC4 encryption, as defined in Bruce Schneier's Applied
// Cryptography.

// rc4包实现了RC4加密算法，参见Bruce Schneier's Applied Cryptography。
package rc4

// A Cipher is an instance of RC4 using a particular key.

// Cipher是一个使用特定密钥的RC4实例，本类型实现了cipher.Stream接口。
type Cipher struct {
	// contains filtered or unexported fields
}

// NewCipher creates and returns a new Cipher. The key argument should be the RC4
// key, at least 1 byte and at most 256 bytes.

// NewCipher创建并返回一个新的Cipher。参数key是RC4密钥，至少1字节，最多256字节。
func NewCipher(key []byte) (*Cipher, error)

// Reset zeros the key data so that it will no longer appear in the process's
// memory.

// Reset方法会清空密钥数据，以便将其数据从程序内存中清除（以免被破解）
func (c *Cipher) Reset()

// XORKeyStream sets dst to the result of XORing src with the key stream. Dst and
// src may be the same slice but otherwise should not overlap.

// XORKeyStream方法将src的数据与秘钥生成的伪随机位流取XOR并写入dst。dst和src可指向同一内存地址；但如果指向不同则其底层内存不可重叠。
func (c *Cipher) XORKeyStream(dst, src []byte)

type KeySizeError int

func (k KeySizeError) Error() string
