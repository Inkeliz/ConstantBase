ConstantBase
===================

> This is experimental!

A constant time implementation of Base64, this is based on work of [Paragon Initiative Enterprises](https://github.com/paragonie/constant_time_encoding) and [Steve "Sc00bz" Thomas](https://github.com/Sc00bz/ConstTimeEncoding), both for PHP.

----------

Download
=======

    go get github.com/Inkeliz/ConstantBase/Base64

Usage
=======

For now we only support a Base64, but all bases will have the following functions:

    encodedBytes := ConstantBase64.EncodeWithPad(text)
    encodedString := ConstantBase64.EncodeWithPadToString(text)

If you don't care about padding use `Encode` instead:

    encodedBytes := ConstantBase64.Encode(text)
    encodedString := ConstantBase64.EncodeToString(text)

To decode the data we have two functions:

    decodedBytes, err := ConstantBase64.Decode(encodedBytes)
    decodedString, err := ConstantBase64.DecodeToString(encodedString)