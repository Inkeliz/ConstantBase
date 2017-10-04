ConstantBase
===================

> This is experimental!

A constant time implementation of Base64.

----------

Download
=======

    go get github.com/Inkeliz/ConstantBase

Usage
=======

For now we only support a Base64, to encode use:

    Encoded := ConstantBase64.EncodeWithPad(Text)

If you don't care about padding use `Encode` instead.

To decode the base64 use:

    Decoded := ConstantBase64.Decode(Encoded)

    